// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package bootstrap

import (
	"testing"

	"github.com/juju/charm/v12"
	jujutesting "github.com/juju/testing"
	"github.com/juju/utils/v3"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/facades/client/charms/interfaces"
	services "github.com/juju/juju/apiserver/facades/client/charms/services"
	"github.com/juju/juju/controller"
	corecharm "github.com/juju/juju/core/charm"
	"github.com/juju/juju/core/constraints"
	state "github.com/juju/juju/state"
	jujujujutesting "github.com/juju/juju/testing"
)

//go:generate go run go.uber.org/mock/mockgen -package bootstrap -destination bootstrap_mock_test.go github.com/juju/juju/internal/bootstrap AgentBinaryStorage,ControllerCharmDeployer,ControllerUnit,HTTPClient,LoggerFactory,CloudService,CloudServiceGetter,OperationApplier,MachineGetter,Machine
//go:generate go run go.uber.org/mock/mockgen -package bootstrap -destination objectstore_mock_test.go github.com/juju/juju/core/objectstore ObjectStore

func Test(t *testing.T) {
	gc.TestingT(t)
}

type baseSuite struct {
	jujutesting.IsolationSuite

	storage        *MockAgentBinaryStorage
	deployer       *MockControllerCharmDeployer
	controllerUnit *MockControllerUnit
	httpClient     *MockHTTPClient
	loggerFactory  *MockLoggerFactory
	objectStore    *MockObjectStore

	logger Logger
}

func (s *baseSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.storage = NewMockAgentBinaryStorage(ctrl)
	s.deployer = NewMockControllerCharmDeployer(ctrl)
	s.controllerUnit = NewMockControllerUnit(ctrl)
	s.httpClient = NewMockHTTPClient(ctrl)

	s.logger = jujujujutesting.NewCheckLogger(c)

	s.loggerFactory = NewMockLoggerFactory(ctrl)
	s.loggerFactory.EXPECT().Child(gomock.Any()).Return(s.logger).AnyTimes()

	return ctrl
}

func (s *baseSuite) newConfig(c *gc.C) BaseDeployerConfig {
	controllerUUID := utils.MustNewUUID()

	return BaseDeployerConfig{
		DataDir:     c.MkDir(),
		State:       &state.State{},
		ObjectStore: s.objectStore,
		Constraints: constraints.Value{},
		ControllerConfig: controller.Config{
			controller.ControllerUUIDKey: controllerUUID.String(),
		},
		NewCharmRepo: func(services.CharmRepoFactoryConfig) (corecharm.Repository, error) {
			return nil, nil
		},
		NewCharmDownloader: func(services.CharmDownloaderConfig) (interfaces.Downloader, error) {
			return nil, nil
		},
		CharmhubHTTPClient: s.httpClient,
		Channel:            charm.Channel{},
		LoggerFactory:      s.loggerFactory,
	}
}

func ptr[T any](v T) *T {
	return &v
}
