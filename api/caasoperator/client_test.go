// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package caasoperator_test

import (
	"github.com/juju/errors"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	basetesting "github.com/juju/juju/api/base/testing"
	"github.com/juju/juju/api/caasoperator"
	"github.com/juju/juju/apiserver/params"
)

type operatorSuite struct {
	testing.IsolationSuite
}

var _ = gc.Suite(&operatorSuite{})

func newClient(f basetesting.APICallerFunc) *caasoperator.Client {
	return caasoperator.NewClient(basetesting.BestVersionCaller{f, 1})
}

func (s *operatorSuite) TestSetStatus(c *gc.C) {
	apiCaller := basetesting.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		c.Check(objType, gc.Equals, "CAASOperator")
		c.Check(version, gc.Equals, 0)
		c.Check(id, gc.Equals, "")
		c.Check(request, gc.Equals, "SetStatus")
		c.Check(arg, jc.DeepEquals, params.SetStatus{
			Entities: []params.EntityStatusArgs{{
				Tag:    "application-gitlab",
				Status: "foo",
				Info:   "bar",
				Data: map[string]interface{}{
					"baz": "qux",
				},
			}},
		})
		c.Assert(result, gc.FitsTypeOf, &params.ErrorResults{})
		*(result.(*params.ErrorResults)) = params.ErrorResults{
			Results: []params.ErrorResult{{Error: &params.Error{Message: "bletch"}}},
		}
		return nil
	})

	client := caasoperator.NewClient(apiCaller)
	err := client.SetStatus("gitlab", "foo", "bar", map[string]interface{}{
		"baz": "qux",
	})
	c.Assert(err, gc.ErrorMatches, "bletch")
}

func (s *operatorSuite) TestSetStatusInvalidApplicationName(c *gc.C) {
	client := caasoperator.NewClient(basetesting.APICallerFunc(func(_ string, _ int, _, _ string, _, _ interface{}) error {
		return errors.New("should not be called")
	}))
	err := client.SetStatus("", "foo", "bar", nil)
	c.Assert(err, gc.ErrorMatches, `application name "" not valid`)
}

func (s *operatorSuite) TestCharm(c *gc.C) {
	apiCaller := basetesting.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		c.Check(objType, gc.Equals, "CAASOperator")
		c.Check(version, gc.Equals, 0)
		c.Check(id, gc.Equals, "")
		c.Check(request, gc.Equals, "Charm")
		c.Check(arg, jc.DeepEquals, params.Entities{
			Entities: []params.Entity{{
				Tag: "application-gitlab",
			}},
		})
		c.Assert(result, gc.FitsTypeOf, &params.ApplicationCharmResults{})
		*(result.(*params.ApplicationCharmResults)) = params.ApplicationCharmResults{
			Results: []params.ApplicationCharmResult{{
				Result: &params.ApplicationCharm{
					URL:          "cs:foo/bar-1",
					ForceUpgrade: true,
					SHA256:       "fake-sha256",
				},
			}},
		}
		return nil
	})

	client := caasoperator.NewClient(apiCaller)
	curl, sha256, err := client.Charm("gitlab")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(curl, gc.NotNil)
	c.Assert(curl.String(), gc.Equals, "cs:foo/bar-1")
	c.Assert(sha256, gc.Equals, "fake-sha256")
}

func (s *operatorSuite) TestCharmError(c *gc.C) {
	apiCaller := basetesting.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		*(result.(*params.ApplicationCharmResults)) = params.ApplicationCharmResults{
			Results: []params.ApplicationCharmResult{{Error: &params.Error{Message: "bletch"}}},
		}
		return nil
	})
	client := caasoperator.NewClient(apiCaller)
	_, _, err := client.Charm("gitlab")
	c.Assert(err, gc.ErrorMatches, "bletch")
}

func (s *operatorSuite) TestCharmInvalidApplicationName(c *gc.C) {
	client := caasoperator.NewClient(basetesting.APICallerFunc(func(_ string, _ int, _, _ string, _, _ interface{}) error {
		return errors.New("should not be called")
	}))
	_, _, err := client.Charm("")
	c.Assert(err, gc.ErrorMatches, `application name "" not valid`)
}
