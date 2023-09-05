// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package changestream

import (
	"context"
	"database/sql"

	"github.com/canonical/sqlair"
	"github.com/juju/clock"
	"github.com/juju/errors"
	"github.com/juju/worker/v3"
	"github.com/juju/worker/v3/catacomb"

	"github.com/juju/juju/core/changestream"
	coredatabase "github.com/juju/juju/core/database"
	"github.com/juju/juju/worker/changestream/eventmultiplexer"
	"github.com/juju/juju/worker/changestream/stream"
)

// WatchableDBWorker is the interface that the worker uses to interact with the
// watchable database.
type WatchableDBWorker interface {
	worker.Worker
	changestream.WatchableDB
}

// WatchableDB is a worker that is responsible for managing the lifecycle
// of both the DBStream and the EventQueue.
type WatchableDB struct {
	catacomb catacomb.Catacomb

	db  coredatabase.TxnRunner
	mux *eventmultiplexer.EventMultiplexer
}

// NewWatchableDB creates a new WatchableDB.
func NewWatchableDB(
	tag string,
	db coredatabase.TxnRunner,
	fileNotifier FileNotifier,
	clock clock.Clock,
	metrics NamespaceMetrics,
	logger Logger,
) (WatchableDBWorker, error) {
	stream := stream.New(tag, db, fileNotifier, clock, metrics, logger)

	mux, err := eventmultiplexer.New(stream, clock, metrics, logger)
	if err != nil {
		return nil, errors.Trace(err)
	}

	w := &WatchableDB{
		db:  db,
		mux: mux,
	}

	if err := catacomb.Invoke(catacomb.Plan{
		Site: &w.catacomb,
		Work: w.loop,
		Init: []worker.Worker{
			stream,
			mux,
		},
	}); err != nil {
		return nil, errors.Trace(err)
	}

	return w, nil
}

// Kill is part of the worker.Worker interface.
func (w *WatchableDB) Kill() {
	w.catacomb.Kill(nil)
}

// Wait is part of the worker.Worker interface.
func (w *WatchableDB) Wait() error {
	return w.catacomb.Wait()
}

// Txn manages the application of a SQLair transaction within which the
// input function is executed. See https://github.com/canonical/sqlair.
// The input context can be used by the caller to cancel this process.
func (w *WatchableDB) Txn(ctx context.Context, fn func(context.Context, *sqlair.TX) error) error {
	return w.db.Txn(ctx, fn)
}

// StdTxn manages the application of a standard library transaction within
// which the input function is executed.
// The input context can be used by the caller to cancel this process.
func (w *WatchableDB) StdTxn(ctx context.Context, fn func(context.Context, *sql.Tx) error) error {
	return w.db.StdTxn(ctx, fn)
}

// Subscribe returns a subscription for the input options.
// The subscription is then used to drive watchers.
func (w *WatchableDB) Subscribe(opts ...changestream.SubscriptionOption) (changestream.Subscription, error) {
	return w.mux.Subscribe(opts...)
}

func (w *WatchableDB) loop() error {
	<-w.catacomb.Dying()
	return w.catacomb.ErrDying()
}
