package scheduler

import (
	"context"

	"github.com/leg100/otf/internal"
)

// interfaces purely for faking purposes
type queueFactory interface {
	newQueue(opts queueOptions) eventHandler
}

type eventHandler interface {
	handleEvent(context.Context, internal.Event) error
}