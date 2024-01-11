// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package testing

import (
	"time"

	"github.com/juju/worker/v4"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/changestream"
	coredatabase "github.com/juju/juju/core/database"
	"github.com/juju/juju/domain/schema/testing"
	jujutesting "github.com/juju/juju/testing"
)

// ControllerSuite is used to provide a sql.DB reference to tests.
// It is pre-populated with the controller schema.
type ControllerSuite struct {
	testing.ControllerSuite

	watchableDB *TestWatchableDB
}

// SetUpTest is responsible for setting up a testing database suite initialised
// with the controller schema.
func (s *ControllerSuite) SetUpTest(c *gc.C) {
	s.ControllerSuite.SetUpTest(c)

	s.watchableDB = NewTestWatchableDB(c, coredatabase.ControllerNS, s.TxnRunner())
}

func (s *ControllerSuite) TearDownTest(c *gc.C) {
	if s.watchableDB != nil {
		// We could use workertest.DirtyKill here, but some workers are already
		// dead when we get here and it causes unwanted logs. This just ensures
		// that we don't have any addition workers running.
		killAndWait(c, s.watchableDB)
	}
	s.ControllerSuite.TearDownTest(c)
}

// GetWatchableDB allows the ControllerSuite to be a WatchableDBGetter
func (s *ControllerSuite) GetWatchableDB(namespace string) (changestream.WatchableDB, error) {
	return s.watchableDB, nil
}

func killAndWait(c *gc.C, w worker.Worker) {
	wait := make(chan error, 1)
	go func() {
		wait <- w.Wait()
	}()
	select {
	case <-wait:
	case <-time.After(jujutesting.LongWait):
	}
}
