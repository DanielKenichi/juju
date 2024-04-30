// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secrets_test

import (
	"context"
	"time"

	"github.com/juju/loggo/v2"
	"github.com/juju/names/v5"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/common/secrets"
	"github.com/juju/juju/apiserver/common/secrets/mocks"
	facademocks "github.com/juju/juju/apiserver/facade/mocks"
	"github.com/juju/juju/core/model"
	coresecrets "github.com/juju/juju/core/secrets"
	secretservice "github.com/juju/juju/domain/secret/service"
	backendservice "github.com/juju/juju/domain/secretbackend/service"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/rpc/params"
	coretesting "github.com/juju/juju/testing"
)

type secretsDrainSuite struct {
	testing.IsolationSuite

	authorizer      *facademocks.MockAuthorizer
	watcherRegistry *facademocks.MockWatcherRegistry

	leadership                *mocks.MockChecker
	token                     *mocks.MockToken
	secretService             *mocks.MockSecretService
	secretBackendService      *mocks.MockSecretBackendService
	model                     *mocks.MockModelConfig
	modelConfigChangesWatcher *mocks.MockNotifyWatcher

	authTag names.Tag

	facade *secrets.SecretsDrainAPI
}

var _ = gc.Suite(&secretsDrainSuite{})

func (s *secretsDrainSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)

	s.authTag = names.NewUnitTag("mariadb/0")
}

func (s *secretsDrainSuite) setup(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.authorizer = facademocks.NewMockAuthorizer(ctrl)
	s.watcherRegistry = facademocks.NewMockWatcherRegistry(ctrl)

	s.leadership = mocks.NewMockChecker(ctrl)
	s.token = mocks.NewMockToken(ctrl)
	s.secretService = mocks.NewMockSecretService(ctrl)
	s.secretBackendService = mocks.NewMockSecretBackendService(ctrl)
	s.model = mocks.NewMockModelConfig(ctrl)
	s.modelConfigChangesWatcher = mocks.NewMockNotifyWatcher(ctrl)
	s.expectAuthUnitAgent()

	var err error
	s.facade, err = secrets.NewSecretsDrainAPI(
		s.authTag,
		s.authorizer,
		loggo.GetLogger("juju.apiserver.secretsdrain"),
		s.leadership,
		model.UUID(coretesting.ModelTag.Id()),
		s.model,
		s.secretService,
		s.secretBackendService,
		s.watcherRegistry,
	)
	c.Assert(err, jc.ErrorIsNil)
	return ctrl
}

func (s *secretsDrainSuite) expectAuthUnitAgent() {
	s.authorizer.EXPECT().AuthUnitAgent().Return(true)
}

func (s *secretsDrainSuite) assertGetSecretsToDrain(c *gc.C, expectedRevions ...params.SecretRevision) {
	defer s.setup(c).Finish()

	s.leadership.EXPECT().LeadershipCheck("mariadb", "mariadb/0").Return(s.token)
	s.token.EXPECT().Check().Return(nil)

	now := time.Now()
	uri := coresecrets.NewURI()
	revisions := []*coresecrets.SecretRevisionMetadata{
		{
			// External backend.
			Revision: 666,
			ValueRef: &coresecrets.ValueRef{
				BackendID:  "backend-id",
				RevisionID: "rev-666",
			},
		}, {
			// Internal backend.
			Revision: 667,
		},
		{
			// k8s backend.
			Revision: 668,
			ValueRef: &coresecrets.ValueRef{
				BackendID:  coretesting.ModelTag.Id(),
				RevisionID: "rev-668",
			},
		},
	}
	s.secretService.EXPECT().ListCharmSecrets(
		gomock.Any(),
		[]secretservice.CharmSecretOwner{{
			Kind: secretservice.UnitOwner,
			ID:   "mariadb/0",
		}, {
			Kind: secretservice.ApplicationOwner,
			ID:   "mariadb",
		}}).Return([]*coresecrets.SecretMetadata{{
		URI:              uri,
		Owner:            coresecrets.Owner{Kind: coresecrets.ApplicationOwner, ID: "mariadb"},
		Label:            "label",
		RotatePolicy:     coresecrets.RotateHourly,
		LatestRevision:   666,
		LatestExpireTime: &now,
		NextRotateTime:   &now,
	}}, [][]*coresecrets.SecretRevisionMetadata{revisions}, nil)
	revInfo := make([]backendservice.RevisionInfo, len(expectedRevions))
	for i, r := range expectedRevions {
		revInfo[i] = backendservice.RevisionInfo{
			Revision: r.Revision,
		}
		if r.ValueRef != nil {
			revInfo[i].ValueRef = &coresecrets.ValueRef{
				BackendID:  r.ValueRef.BackendID,
				RevisionID: r.ValueRef.RevisionID,
			}
		}
	}
	s.secretBackendService.EXPECT().GetRevisionsToDrain(gomock.Any(), model.UUID(coretesting.ModelTag.Id()), revisions).
		Return(revInfo, nil)

	results, err := s.facade.GetSecretsToDrain(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.ListSecretResults{
		Results: []params.ListSecretResult{{
			URI:              uri.String(),
			OwnerTag:         "application-mariadb",
			Label:            "label",
			RotatePolicy:     coresecrets.RotateHourly.String(),
			LatestRevision:   666,
			LatestExpireTime: &now,
			NextRotateTime:   &now,
			Revisions:        expectedRevions,
		}},
	})
}

func (s *secretsDrainSuite) TestGetSecretsToDrainInternal(c *gc.C) {
	s.assertGetSecretsToDrain(c,
		// External backend.
		params.SecretRevision{
			Revision: 666,
			ValueRef: &params.SecretValueRef{
				BackendID:  "backend-id",
				RevisionID: "rev-666",
			},
		},
		// k8s backend.
		params.SecretRevision{
			Revision: 668,
			ValueRef: &params.SecretValueRef{
				BackendID:  coretesting.ModelTag.Id(),
				RevisionID: "rev-668",
			},
		},
	)
}

func (s *secretsDrainSuite) TestGetSecretsToDrainExternal(c *gc.C) {
	s.assertGetSecretsToDrain(c,
		// Internal backend.
		params.SecretRevision{
			Revision: 667,
		},
		// k8s backend.
		params.SecretRevision{
			Revision: 668,
			ValueRef: &params.SecretValueRef{
				BackendID:  coretesting.ModelTag.Id(),
				RevisionID: "rev-668",
			},
		},
	)
}

func (s *secretsDrainSuite) TestChangeSecretBackend(c *gc.C) {
	defer s.setup(c).Finish()

	uri1 := coresecrets.NewURI()
	uri2 := coresecrets.NewURI()
	s.secretService.EXPECT().ChangeSecretBackend(
		gomock.Any(),
		uri1, 666,
		secretservice.ChangeSecretBackendParams{
			Accessor: secretservice.SecretAccessor{
				Kind: secretservice.UnitAccessor,
				ID:   s.authTag.Id(),
			},
			LeaderToken: s.token,
			ValueRef: &coresecrets.ValueRef{
				BackendID:  "backend-id",
				RevisionID: "rev-666",
			},
		},
	).Return(nil)
	s.secretService.EXPECT().ChangeSecretBackend(
		gomock.Any(),
		uri2, 888,
		secretservice.ChangeSecretBackendParams{
			Accessor: secretservice.SecretAccessor{
				Kind: secretservice.UnitAccessor,
				ID:   s.authTag.Id(),
			},
			LeaderToken: s.token,
			Data:        map[string]string{"foo": "bar"},
		},
	).Return(nil)
	s.leadership.EXPECT().LeadershipCheck("mariadb", "mariadb/0").Return(s.token).Times(2)

	result, err := s.facade.ChangeSecretBackend(context.Background(), params.ChangeSecretBackendArgs{
		Args: []params.ChangeSecretBackendArg{
			{
				URI:      uri1.String(),
				Revision: 666,
				Content: params.SecretContentParams{
					// Change to external backend.
					ValueRef: &params.SecretValueRef{
						BackendID:  "backend-id",
						RevisionID: "rev-666",
					},
				},
			},
			{
				URI:      uri2.String(),
				Revision: 888,
				Content: params.SecretContentParams{
					// Change to internal backend.
					Data: map[string]string{"foo": "bar"},
				},
			},
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, jc.DeepEquals, params.ErrorResults{
		Results: []params.ErrorResult{{Error: nil}, {Error: nil}},
	})
}

func (s *secretsDrainSuite) TestWatchSecretBackendChanged(c *gc.C) {
	defer s.setup(c).Finish()

	done := make(chan struct{})
	changeChan := make(chan struct{}, 1)
	changeChan <- struct{}{}
	s.modelConfigChangesWatcher.EXPECT().Wait().DoAndReturn(func() error {
		close(done)
		return nil
	})
	s.modelConfigChangesWatcher.EXPECT().Changes().Return(changeChan).AnyTimes()

	s.model.EXPECT().WatchForModelConfigChanges().Return(s.modelConfigChangesWatcher)
	configAttrs := map[string]interface{}{
		"name":           "some-name",
		"type":           "some-type",
		"uuid":           coretesting.ModelTag.Id(),
		"secret-backend": "backend-id",
	}
	cfg, err := config.New(config.NoDefaults, configAttrs)
	c.Assert(err, jc.ErrorIsNil)
	s.model.EXPECT().ModelConfig(gomock.Any()).Return(cfg, nil).Times(2)

	s.watcherRegistry.EXPECT().Register(gomock.Any()).Return("11", nil)

	result, err := s.facade.WatchSecretBackendChanged(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, jc.DeepEquals, params.NotifyWatchResult{
		NotifyWatcherId: "11",
	})

	select {
	case <-done:
		// We need to wait for the watcher to fully start to ensure that all expected methods are called.
	case <-time.After(coretesting.LongWait):
		c.Fatalf("timed out waiting for 2nd loop")
	}
}
