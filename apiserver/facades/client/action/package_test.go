// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package action

import (
	"github.com/juju/names/v4"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	facademocks "github.com/juju/juju/apiserver/facade/mocks"
	"github.com/juju/juju/state"
)

//go:generate go run github.com/golang/mock/mockgen -package action -destination package_mock_test.go github.com/juju/juju/apiserver/facades/client/action State,Model
//go:generate go run github.com/golang/mock/mockgen -package action -destination state_mock_test.go github.com/juju/juju/state Action,ActionReceiver
//go:generate go run github.com/golang/mock/mockgen -package action -destination leader_mock_test.go github.com/juju/juju/core/leadership Reader

type MockBaseSuite struct {
	State          *MockState
	Authorizer     *facademocks.MockAuthorizer
	ActionReceiver *MockActionReceiver
	Leadership     *MockReader
}

func (s *MockBaseSuite) NewActionAPI(c *gc.C) *ActionAPI {
	api, err := newActionAPI(s.State, nil, s.Authorizer, s.Leadership)
	c.Assert(err, jc.ErrorIsNil)

	api.tagToActionReceiverFn = s.tagToActionReceiverFn
	return api
}

func (s *MockBaseSuite) tagToActionReceiverFn(
	_ func(names.Tag) (state.Entity, error),
) func(tag string) (state.ActionReceiver, error) {
	return func(tag string) (state.ActionReceiver, error) {
		return s.ActionReceiver, nil
	}
}
