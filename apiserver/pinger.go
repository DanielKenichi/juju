// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package apiserver

import (
	"time"

	"github.com/juju/clock"
	"github.com/juju/errors"
	"github.com/juju/worker/v3"
	"gopkg.in/tomb.v2"

	"github.com/juju/juju/apiserver/facade"
)

// NewPinger returns an object that can be pinged by calling its Ping method.
// If this method is not called frequently enough, the connection will be
// dropped.
func NewPinger(ctx facade.Context) (Pinger, error) {
	pingTimeout, ok := ctx.Resources().Get("pingTimeout").(*pingTimeout)
	if !ok {
		return nullPinger{}, nil
	}
	return pingTimeout, nil
}

// pinger describes a resource that can be pinged and stopped.
type Pinger interface {
	worker.Worker
	Ping()
}

// pingTimeout listens for pings and will call the
// passed action in case of a timeout. This way broken
// or inactive connections can be closed.
type pingTimeout struct {
	tomb tomb.Tomb
	worker.Worker
	action  func()
	clock   clock.Clock
	timeout time.Duration
	reset   chan struct{}
}

// newPingTimeout returns a new pingTimeout instance
// that invokes the given action asynchronously if there
// is more than the given timeout interval between calls
// to its Ping method.
func newPingTimeout(action func(), clock clock.Clock, timeout time.Duration) Pinger {
	pt := &pingTimeout{
		action:  action,
		clock:   clock,
		timeout: timeout,
		reset:   make(chan struct{}),
	}
	pt.tomb.Go(pt.loop)
	return pt
}

// Ping is used by the client heartbeat monitor and resets
// the killer.
func (pt *pingTimeout) Ping() {
	select {
	case <-pt.tomb.Dying():
	case pt.reset <- struct{}{}:
	}
}

func (pt *pingTimeout) Kill() {
	pt.tomb.Kill(nil)
}

func (pt *pingTimeout) Wait() error {
	return pt.tomb.Wait()
}

// loop waits for a reset signal, otherwise it performs
// the initially passed action.
func (pt *pingTimeout) loop() error {
	for {
		select {
		case <-pt.tomb.Dying():
			return tomb.ErrDying
		case <-pt.reset:
		case <-pt.clock.After(pt.timeout):
			go pt.action()
			return errors.New("ping timeout")
		}
	}
}

// nullPinger implements the pinger interface but just does nothing
type nullPinger struct{}

func (nullPinger) Ping() {}

func (nullPinger) Kill() {}

func (nullPinger) Wait() error {
	return nil
}
