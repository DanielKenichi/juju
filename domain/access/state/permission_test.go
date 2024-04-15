// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"
	"database/sql"
	"time"

	"github.com/juju/names/v5"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	corepermission "github.com/juju/juju/core/permission"
	"github.com/juju/juju/core/user"
	"github.com/juju/juju/domain/access"
	accesserrors "github.com/juju/juju/domain/access/errors"
	schematesting "github.com/juju/juju/domain/schema/testing"
	"github.com/juju/juju/internal/uuid"
	jujutesting "github.com/juju/juju/testing"
)

type permissionStateSuite struct {
	schematesting.ControllerSuite

	debug bool
}

var _ = gc.Suite(&permissionStateSuite{})

func (s *permissionStateSuite) SetUpTest(c *gc.C) {
	s.ControllerSuite.SetUpTest(c)
	s.ensureUser(c, "42", "admin", "42") // model owner
	s.ensureUser(c, "123", "bob", "42")
	s.ensureUser(c, "456", "sue", "42")
	s.ensureCloud(c, "987", "test-cloud")
	s.ensureCloud(c, "654", "another-cloud")
	s.ensureModel(c, "default-model")
	s.ensureModel(c, "model-uuid")
}

func (s *permissionStateSuite) TestCreatePermissionModel(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	userAccess, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "model-uuid",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.WriteAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Check(userAccess.UserID, gc.Equals, "123")
	c.Check(userAccess.UserTag, gc.Equals, names.NewUserTag("bob"))
	c.Check(userAccess.Object.Id(), gc.Equals, "model-uuid")
	c.Check(userAccess.Access, gc.Equals, corepermission.WriteAccess)
	c.Check(userAccess.DisplayName, gc.Equals, "bob")
	c.Check(userAccess.UserName, gc.Equals, "bob")
	c.Check(userAccess.CreatedBy, gc.Equals, names.NewUserTag("admin"))

	s.checkPermissionRow(c, corepermission.WriteAccess, "123", "model-uuid")
}

func (s *permissionStateSuite) TestCreatePermissionCloud(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	userAccess, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "test-cloud",
				ObjectType: corepermission.Cloud,
			},
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Check(userAccess.UserID, gc.Equals, "123")
	c.Check(userAccess.UserTag, gc.Equals, names.NewUserTag("bob"))
	c.Check(userAccess.Object.Id(), gc.Equals, "test-cloud")
	c.Check(userAccess.Access, gc.Equals, corepermission.AddModelAccess)
	c.Check(userAccess.DisplayName, gc.Equals, "bob")
	c.Check(userAccess.UserName, gc.Equals, "bob")
	c.Check(userAccess.CreatedBy, gc.Equals, names.NewUserTag("admin"))

	s.checkPermissionRow(c, corepermission.AddModelAccess, "123", "test-cloud")
}

func (s *permissionStateSuite) TestCreatePermissionController(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	userAccess, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "controller",
				ObjectType: corepermission.Controller,
			},
			Access: corepermission.SuperuserAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Check(userAccess.UserID, gc.Equals, "123")
	c.Check(userAccess.UserTag, gc.Equals, names.NewUserTag("bob"))
	c.Check(userAccess.Object.Id(), gc.Equals, "controller")
	c.Check(userAccess.Access, gc.Equals, corepermission.SuperuserAccess)
	c.Check(userAccess.DisplayName, gc.Equals, "bob")
	c.Check(userAccess.UserName, gc.Equals, "bob")
	c.Check(userAccess.CreatedBy, gc.Equals, names.NewUserTag("admin"))

	s.checkPermissionRow(c, corepermission.SuperuserAccess, "123", "controller")
}

func (s *permissionStateSuite) TestCreatePermissionForModelWithBadInfo(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "foo-bar",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.ReadAccess,
		},
	})
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionTargetInvalid)
}

func (s *permissionStateSuite) TestCreatePermissionForControllerWithBadInfo(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "foo-bar",
				ObjectType: corepermission.Controller,
			},
			Access: corepermission.SuperuserAccess,
		},
	})
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionTargetInvalid)
}

func (s *permissionStateSuite) checkPermissionRow(c *gc.C, access corepermission.Access, expectedGrantTo, expectedGrantON string) {
	db := s.DB()

	// Find the permission
	row := db.QueryRow(`
SELECT uuid, access_type, grant_to, grant_on
FROM v_permission
`)
	c.Assert(row.Err(), jc.ErrorIsNil)
	var (
		accessType, userUuid, grantTo, grantOn string
	)
	err := row.Scan(&userUuid, &accessType, &grantTo, &grantOn)
	c.Assert(err, jc.ErrorIsNil)

	// Verify the permission as expected.
	c.Check(userUuid, gc.Not(gc.Equals), "")
	c.Check(accessType, gc.Equals, string(access))
	c.Check(grantTo, gc.Equals, expectedGrantTo)
	c.Check(grantOn, gc.Equals, expectedGrantON)
}

func (s *permissionStateSuite) TestCreatePermissionErrorNoUser(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "testme",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "model-uuid",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.WriteAccess,
		},
	})
	c.Assert(err, jc.ErrorIs, accesserrors.UserNotFound)
}

func (s *permissionStateSuite) TestCreatePermissionErrorDuplicate(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	spec := corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "model-uuid",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.ReadAccess,
		},
	}
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), spec)
	c.Assert(err, jc.ErrorIsNil)

	// Find the permission
	row := s.DB().QueryRow(`
SELECT uuid, permission_type_id, grant_to, grant_on 
FROM permission
WHERE permission_type_id = 0
`)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var (
		userUuid, grantTo, grantOn string
		permissionTypeId           int
	)
	err = row.Scan(&userUuid, &permissionTypeId, &grantTo, &grantOn)
	c.Assert(err, jc.ErrorIsNil)

	// Ensure each combination of grant_on and grant_two
	// is unique
	spec.Access = corepermission.WriteAccess
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), spec)
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionAlreadyExists)
	row2 := s.DB().QueryRow(`
SELECT uuid, permission_type_id, grant_to, grant_on 
FROM permission
WHERE permission_type_id = 1
`)
	c.Assert(row2.Err(), jc.ErrorIsNil)
	err = row2.Scan(&userUuid, &permissionTypeId, &grantTo, &grantOn)
	c.Assert(err, jc.ErrorIs, sql.ErrNoRows)
}

func (s *permissionStateSuite) TestDeletePermission(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	target := corepermission.ID{
		Key:        "model-uuid",
		ObjectType: corepermission.Model,
	}
	spec := corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.ReadAccess,
		},
	}
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), spec)
	c.Assert(err, jc.ErrorIsNil)

	err = st.DeletePermission(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIsNil)

	db := s.DB()

	var num int
	err = db.QueryRowContext(context.Background(), "SELECT count(*) FROM permission").Scan(&num)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(num, gc.Equals, 0)
}

func (s *permissionStateSuite) TestDeletePermissionDoesNotExist(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	target := corepermission.ID{
		Key:        "model-uuid",
		ObjectType: corepermission.Model,
	}

	// Don't fail if the permission does not exist.
	err := st.DeletePermission(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *permissionStateSuite) TestReadUserAccessForTarget(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	target := corepermission.ID{
		Key:        "controller",
		ObjectType: corepermission.Controller,
	}
	createUserAccess, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.SuperuserAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	var (
		userUuid, grantTo, grantOn string
		permissionTypeId           int
	)

	row2 := s.DB().QueryRow(`
SELECT uuid, permission_type_id, grant_to, grant_on 
FROM permission
WHERE grant_to = 123
`)
	c.Assert(row2.Err(), jc.ErrorIsNil)
	err = row2.Scan(&userUuid, &permissionTypeId, &grantTo, &grantOn)
	c.Assert(err, jc.ErrorIsNil)
	c.Logf("%q, %d, to %q, on %q", userUuid, permissionTypeId, grantTo, grantOn)

	readUserAccess, err := st.ReadUserAccessForTarget(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(createUserAccess, gc.DeepEquals, readUserAccess)
}

func (s *permissionStateSuite) TestReadUserAccessLevelForTarget(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	target := corepermission.ID{
		Key:        "test-cloud",
		ObjectType: corepermission.Cloud,
	}
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	readUserAccessType, err := st.ReadUserAccessLevelForTarget(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(readUserAccessType, gc.Equals, corepermission.AddModelAccess)
}

func (s *permissionStateSuite) TestReadAllUserAccessForUser(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)

	userAccesses, err := st.ReadAllUserAccessForUser(context.Background(), "bob")
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(userAccesses, gc.HasLen, 4)
	for _, access := range userAccesses {
		c.Assert(access.UserName, gc.Equals, "bob")
		c.Assert(access.CreatedBy.Id(), gc.Equals, "admin")
	}
	accessOne := userAccesses[0]
	c.Assert(accessOne.Access, gc.Equals, corepermission.AddModelAccess)
}

func (s *permissionStateSuite) TestReadAllUserAccessForTarget(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)
	targetCloud := corepermission.ID{
		Key:        "test-cloud",
		ObjectType: corepermission.Cloud,
	}
	userAccesses, err := st.ReadAllUserAccessForTarget(context.Background(), targetCloud)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(userAccesses, gc.HasLen, 2)
	accessZero := userAccesses[0]
	c.Check(accessZero.Access, gc.Equals, corepermission.AddModelAccess)
	c.Check(accessZero.Object, gc.Equals, names.NewCloudTag("test-cloud"))
	accessOne := userAccesses[1]
	c.Check(accessOne.Access, gc.Equals, corepermission.AddModelAccess)
	c.Check(accessOne.Object, gc.Equals, names.NewCloudTag("test-cloud"))

	c.Check(accessZero.UserID, gc.Not(gc.Equals), accessOne.UserID)
}

func (s *permissionStateSuite) TestReadAllAccessForUserAndObjectTypeCloud(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)

	users, err := st.ReadAllAccessForUserAndObjectType(context.Background(), "bob", corepermission.Cloud)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(users, gc.HasLen, 2)

	var foundTestCloud, foundAnotherCloud bool
	for _, userAccess := range users {
		c.Check(userAccess.UserTag.Id(), gc.Equals, "bob")
		c.Check(userAccess.UserName, gc.Equals, "bob")
		c.Check(userAccess.CreatedBy.Id(), gc.Equals, "admin")
		c.Check(userAccess.UserID, gc.Equals, "123")
		c.Check(userAccess.Access, gc.Equals, corepermission.AddModelAccess)
		if userAccess.Object.Id() == "test-cloud" {
			foundTestCloud = true
		}
		if userAccess.Object.Id() == "another-cloud" {
			foundAnotherCloud = true
		}
	}
	c.Check(foundTestCloud && foundAnotherCloud, jc.IsTrue, gc.Commentf("%+v", users))
}

func (s *permissionStateSuite) TestReadAllAccessForUserAndObjectTypeModel(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)

	users, err := st.ReadAllAccessForUserAndObjectType(context.Background(), "bob", corepermission.Model)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(users, gc.HasLen, 2)

	var admin, write bool
	for _, userAccess := range users {
		c.Check(userAccess.UserTag.Id(), gc.Equals, "bob")
		c.Check(userAccess.UserName, gc.Equals, "bob")
		c.Check(userAccess.CreatedBy.Id(), gc.Equals, "admin")
		c.Check(userAccess.UserID, gc.Equals, "123")
		if userAccess.Access == corepermission.WriteAccess {
			write = true
			c.Check(userAccess.Object.Id(), gc.Equals, "default-model")
		}
		if userAccess.Access == corepermission.AdminAccess {
			admin = true
			c.Check(userAccess.Object.Id(), gc.Equals, "model-uuid")
		}
	}
	c.Assert(admin && write, jc.IsTrue, gc.Commentf("%+v", users))
}

func (s *permissionStateSuite) TestReadAllAccessForUserAndObjectTypeController(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)

	users, err := st.ReadAllAccessForUserAndObjectType(context.Background(), "admin", corepermission.Controller)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(users, gc.HasLen, 1)
	userAccess := users[0]
	c.Check(userAccess.UserTag.Id(), gc.Equals, "admin", gc.Commentf("%+v", users))
	c.Check(userAccess.UserName, gc.Equals, "admin", gc.Commentf("%+v", users))
	c.Check(userAccess.CreatedBy.Id(), gc.Equals, "admin", gc.Commentf("%+v", users))
	c.Check(userAccess.UserID, gc.Equals, "42", gc.Commentf("%+v", users))
	c.Check(userAccess.Access, gc.Equals, corepermission.SuperuserAccess, gc.Commentf("%+v", users))
}

func (s *permissionStateSuite) TestReadAllAccessForUserAndObjectTypeNotFound(c *gc.C) {
	st := NewState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	_, err := st.ReadAllAccessForUserAndObjectType(context.Background(), "bob", corepermission.Cloud)
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionNotFound)
}

func (s *permissionStateSuite) TestUpsertPermissionGrantNewUser(c *gc.C) {
	st := NewState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "admin",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "default-model",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.AdminAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "admin",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "controller",
				ObjectType: corepermission.Controller,
			},
			Access: corepermission.SuperuserAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	target := corepermission.ID{
		ObjectType: corepermission.Model,
		Key:        "default-model",
	}
	arg := access.UpdatePermissionArgs{
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.WriteAccess,
		},
		AddUser: true,
		ApiUser: "admin",
		Change:  corepermission.Grant,
		Subject: "tom",
	}
	err = st.UpsertPermission(context.Background(), arg)
	c.Assert(err, jc.ErrorIsNil)

	obtainedUserAccess, err := st.ReadUserAccessForTarget(context.Background(), "tom", target)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(obtainedUserAccess.UserTag.Id(), gc.Equals, "tom")
	c.Check(obtainedUserAccess.UserName, gc.Equals, "tom")
	c.Check(obtainedUserAccess.CreatedBy.Id(), gc.Equals, "admin")
	c.Check(obtainedUserAccess.UserID, gc.Not(gc.Equals), "")
	c.Check(obtainedUserAccess.Access, gc.Equals, corepermission.WriteAccess)
	c.Check(obtainedUserAccess.Object.Id(), gc.Equals, "default-model")
}

func (s *permissionStateSuite) TestUpsertPermissionGrantExistingUser(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))
	// Bob starts with Write access on "default-model"
	s.setupForRead(c, st)

	target := corepermission.ID{
		ObjectType: corepermission.Model,
		Key:        "default-model",
	}
	arg := access.UpdatePermissionArgs{
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.AdminAccess,
		},
		AddUser: true,
		ApiUser: "admin",
		Change:  corepermission.Grant,
		Subject: "bob",
	}
	err := st.UpsertPermission(context.Background(), arg)
	c.Assert(err, jc.ErrorIsNil)

	obtainedUserAccess, err := st.ReadUserAccessForTarget(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(obtainedUserAccess.UserTag.Id(), gc.Equals, "bob")
	c.Check(obtainedUserAccess.Access, gc.Equals, corepermission.AdminAccess)
}

func (s *permissionStateSuite) TestUpsertPermissionGrantLessAccess(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))
	// Bob starts with Write access on "default-model"
	s.setupForRead(c, st)

	target := corepermission.ID{
		ObjectType: corepermission.Model,
		Key:        "default-model",
	}
	arg := access.UpdatePermissionArgs{
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.ReadAccess,
		},
		AddUser: true,
		ApiUser: "admin",
		Change:  corepermission.Grant,
		Subject: "bob",
	}
	err := st.UpsertPermission(context.Background(), arg)
	c.Assert(err, gc.ErrorMatches, `user "bob" already has "read" access or greater`)
}

func (s *permissionStateSuite) TestUpsertPermissionNotAuthorized(c *gc.C) {
	st := NewState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	arg := access.UpdatePermissionArgs{
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				ObjectType: corepermission.Model,
				Key:        "model-uuid",
			},
		},
		AddUser: false,
		ApiUser: "admin",
		Change:  corepermission.Grant,
		Subject: "bub",
	}
	err := st.UpsertPermission(context.Background(), arg)
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionNotValid)
}

func (s *permissionStateSuite) TestUpsertPermissionRevokeRemovePerm(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))
	s.setupForRead(c, st)
	// Bob starts with Admin access on "default-model".
	// Revoke of Read yields permission removed on the model.
	target := corepermission.ID{
		ObjectType: corepermission.Model,
		Key:        "default-model",
	}
	arg := access.UpdatePermissionArgs{
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.ReadAccess,
		},
		AddUser: true,
		ApiUser: "admin",
		Change:  corepermission.Revoke,
		Subject: "bob",
	}
	err := st.UpsertPermission(context.Background(), arg)
	c.Assert(err, jc.ErrorIsNil)

	_, err = st.ReadUserAccessForTarget(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionNotFound)
}

func (s *permissionStateSuite) TestUpsertPermissionRevoke(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))
	// Sue starts with Admin access on "test-cloud".
	// Revoke of Admin yields AddModel on clouds.
	s.setupForRead(c, st)

	target := corepermission.ID{
		ObjectType: corepermission.Cloud,
		Key:        "test-cloud",
	}
	arg := access.UpdatePermissionArgs{
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.AdminAccess,
		},
		AddUser: false,
		ApiUser: "admin",
		Change:  corepermission.Revoke,
		Subject: "sue",
	}
	err := st.UpsertPermission(context.Background(), arg)
	c.Assert(err, jc.ErrorIsNil)

	obtainedUserAccess, err := st.ReadUserAccessForTarget(context.Background(), "sue", target)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(obtainedUserAccess.UserTag.Id(), gc.Equals, "sue")
	c.Check(obtainedUserAccess.Access, gc.Equals, corepermission.AddModelAccess)
}

func (s *permissionStateSuite) setupForRead(c *gc.C, st *PermissionState) {
	targetCloud := corepermission.ID{
		Key:        "test-cloud",
		ObjectType: corepermission.Cloud,
	}
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: targetCloud,
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "sue",
		AccessSpec: corepermission.AccessSpec{
			Target: targetCloud,
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "model-uuid",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.AdminAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "another-cloud",
				ObjectType: corepermission.Cloud,
			},
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "default-model",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.WriteAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "admin",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "controller",
				ObjectType: corepermission.Controller,
			},
			Access: corepermission.SuperuserAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	if s.debug {
		s.printUsers(c)
		s.printUserAuthentication(c)
		s.printClouds(c)
		s.printPermissions(c)
		s.printRead(c)
	}
}

func (s *permissionStateSuite) ensureUser(c *gc.C, userUUID, name, createdByUUID string) {
	err := s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO user (uuid, name, display_name, removed, created_by_uuid, created_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`, userUUID, name, name, false, createdByUUID, time.Now())
		return err
	})
	c.Assert(err, jc.ErrorIsNil)
	err = s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO user_authentication (user_uuid, disabled)
			VALUES (?, ?)
		`, userUUID, false)
		return err
	})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *permissionStateSuite) ensureModel(c *gc.C, uuid string) {
	err := s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO model_list (uuid)
			VALUES (?)
		`, uuid)
		return err
	})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *permissionStateSuite) ensureCloud(c *gc.C, uuid, name string) {
	err := s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO cloud (uuid, name, cloud_type_id, endpoint, skip_tls_verify)
			VALUES (?, ?, 7, "test-endpoint", true)
		`, uuid, name)
		return err
	})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *permissionStateSuite) printPermissions(c *gc.C) {
	rows, _ := s.DB().Query(`
SELECT uuid, permission_type_id, grant_to, grant_on 
FROM permission
`)
	defer func() { _ = rows.Close() }()
	var (
		userUuid, grantTo, grantOn string
		permissionTypeId           int
	)

	c.Logf("PERMISSIONS")
	for rows.Next() {
		err := rows.Scan(&userUuid, &permissionTypeId, &grantTo, &grantOn)
		c.Assert(err, jc.ErrorIsNil)
		c.Logf("%q, %d, %q, %q", userUuid, permissionTypeId, grantTo, grantOn)
	}

}

func (s *permissionStateSuite) printUsers(c *gc.C) {
	rows, _ := s.DB().Query(`
SELECT u.uuid, u.name, u.created_by_uuid, u.disabled, u.removed
FROM v_user_auth u
`)
	defer func() { _ = rows.Close() }()
	var (
		rowUUID, name     string
		creatorUUID       user.UUID
		disabled, removed bool
	)
	c.Logf("USERS")
	for rows.Next() {
		err := rows.Scan(&rowUUID, &name, &creatorUUID, &disabled, &removed)
		c.Assert(err, jc.ErrorIsNil)
		c.Logf("LINE %q, %q, %q, %t, %t", rowUUID, name, creatorUUID, disabled, removed)
	}
}

func (s *permissionStateSuite) printUserAuthentication(c *gc.C) {
	rows, _ := s.DB().Query(`
SELECT user_uuid, disabled
FROM user_authentication
`)
	defer func() { _ = rows.Close() }()
	var (
		userUUID string
		disabled bool
	)
	c.Logf("USERS AUTHENTICATION")
	for rows.Next() {
		err := rows.Scan(&userUUID, &disabled)
		c.Assert(err, jc.ErrorIsNil)
		c.Logf("LINE %q, %t", userUUID, disabled)
	}
}

func (s *permissionStateSuite) printRead(c *gc.C) {
	q := `
SELECT  p.uuid, p.grant_on, p.grant_to, p.access_type,
        u.uuid, u.name, creator.name
FROM    v_user_auth u
        JOIN user AS creator ON u.created_by_uuid = creator.uuid
        JOIN v_permission p ON u.uuid = p.grant_to
`
	rows, _ := s.DB().Query(q)
	defer func() { _ = rows.Close() }()
	var (
		permUUID, grantOn, grantTo, accessType string
		userUUID, userName, createName         string
	)
	c.Logf("READ")
	for rows.Next() {
		err := rows.Scan(&permUUID, &grantOn, &grantTo, &accessType, &userUUID, &userName, &createName)
		c.Assert(err, jc.ErrorIsNil)
		c.Logf("LINE: %q, %q, %q, %q, %q, %q, %q,", permUUID, grantOn, grantTo, accessType, userUUID, userName, createName)
	}
}

func (s *permissionStateSuite) printClouds(c *gc.C) {
	rows, _ := s.DB().Query(`
SELECT uuid, name
FROM cloud
`)
	defer func() { _ = rows.Close() }()
	var (
		rowUUID, name string
	)

	c.Logf("CLOUDS")
	for rows.Next() {
		err := rows.Scan(&rowUUID, &name)
		c.Assert(err, jc.ErrorIsNil)
		c.Logf("%q, %q", rowUUID, name)
	}
}