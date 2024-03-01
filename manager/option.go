package manager

import (
	"time"

	"github.com/eapache/go-resiliency/retrier"
)

type Option func(m *Manager)

func WithCancelOnError() Option {
	return func(m *Manager) {
		*m.pool = *m.pool.WithCancelOnError()
	}
}

func WithFirstError() Option {
	return func(m *Manager) {
		*m.pool = *m.pool.WithFirstError()
	}
}

func WithMaxTasks(n int) Option {
	return func(m *Manager) {
		*m.pool = *m.pool.WithMaxGoroutines(n)
	}
}

func WithConstantRetry(attempts int, backoffDuration time.Duration) Option {
	return func(m *Manager) {
		backoff := retrier.ConstantBackoff(attempts, backoffDuration)
		m.retrier = retrier.New(backoff, retrier.DefaultClassifier{})
	}
}

func WithExpotentialRetry(attempts int, minBackoffDuration time.Duration, maxBackoffDuration time.Duration) Option {
	return func(m *Manager) {
		backoff := retrier.LimitedExponentialBackoff(attempts, minBackoffDuration, maxBackoffDuration)
		m.retrier = retrier.New(backoff, retrier.DefaultClassifier{})
	}
}
