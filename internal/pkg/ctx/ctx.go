package ctx

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// CTX extends Google's context to support logging methods
type CTX struct {
	context.Context
	logrus.FieldLogger
}

// Background returns a non-nil, empty Context. It is never canceled, has no values, and
// has no deadline. It is typically used by the main function, initialization, and tests,
// and as the top-level Context for incoming requests
func Background() CTX {
	return CTX{
		Context:     context.Background(),
		FieldLogger: logrus.StandardLogger(),
	}
}

// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(parent CTX, key string, val interface{}) CTX {
	return CTX{
		Context:     context.WithValue(parent, key, val),
		FieldLogger: parent.FieldLogger.WithField(key, val),
	}
}

// WithCancel returns a copy of parent with added cancel function
func WithCancel(parent CTX) (CTX, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	return CTX{
		Context:     ctx,
		FieldLogger: parent.FieldLogger,
	}, cancel
}

// WithTimeout returns a copy of parent with timeout condition and cancel function
func WithTimeout(parent CTX, d time.Duration) (CTX, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent, d)
	return CTX{
		Context:     ctx,
		FieldLogger: parent.FieldLogger,
	}, cancel
}
