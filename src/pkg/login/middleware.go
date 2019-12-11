package login

import (
	"context"
	"log"
)

// MiddlewareService describes a service middleware.
type MiddlewareService func(Service) Service

// LoggingMiddleware takes a logger as a dependency and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) MiddlewareService {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) Login(ctx context.Context, user User) (s string, err error) {
	defer func() {
		mw.logger.Print("method", "Login", "userName", user.Name, "userPw", user.Password, "userRole", user.Role, "timestamp", user.Timestamp)
	}()
	return mw.next.Login(ctx, user)
}

// InstrumentingMiddleware returns a service middleware that instruments
// the number of successful logins done over the lifetime of the service.
func InstrumentingMiddleware(ints int) MiddlewareService {
	return func(next Service) Service {
		return instrumentingMiddleware{
			ints: ints,
			next: next,
		}
	}
}

type instrumentingMiddleware struct {
	ints int
	next Service
}

func (mw instrumentingMiddleware) Login(ctx context.Context, user User) (string, error) {
	s, err := mw.next.Login(ctx, user)
	if s == "success" {
		mw.ints = mw.ints + 1
	}
	return s, err
}

// MiddlewareEndpoint describes an endpoint middleware.
type MiddlewareEndpoint func(Endpoint) Endpoint

// InstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
/*
func InstrumentingMiddleware(duration metrics.Histogram) MiddlewareEndpoint {
	return func(next Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With("success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}*/

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
/*
func LoggingMiddleware(logger log.Logger) MiddlewareEndpoint {
	return func(next Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}*/

/*
func CircuitBreakingMiddleware(cb *gobreaker.CircuitBreaker) MiddlewareEndpoint {
	return func(next Endpoint) Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			//return cb.Execute(func() (interface{}, error) {
			return next(ctx, req)
			//})
		}
	}
}*/

/*
func ThrottlingMiddleware(b *ratelimit.Bucket) MiddlewareEndpoint {
	return func(next Endpoint) Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			//time.Sleep(b.Take(1))
			return next(ctx, req)
		}
	}
}*/
