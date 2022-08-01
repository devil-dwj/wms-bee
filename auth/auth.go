package auth

// import (
// 	"context"
// 	"errors"

// 	"github.com/devil-dwj/wms-bee/token"
// 	"github.com/devil-dwj/wms/middleware"
// 	"github.com/devil-dwj/wms/runtime/http"
// 	"github.com/dgrijalva/jwt-go"
// )

// type Option func(*options)

// type options struct {
// 	claims func() jwt.Claims
// }

// func WithClaims(f func() jwt.Claims) Option {
// 	return func(o *options) {
// 		o.claims = f
// 	}
// }

// func Auth(secret string, opts ...Option) middleware.Middleware {
// 	o := &options{
// 		claims: func() jwt.Claims {
// 			return jwt.MapClaims{}
// 		},
// 	}
// 	for _, opt := range opts {
// 		opt(o)
// 	}
// 	return func(h middleware.Handler) middleware.Handler {
// 		return func(ctx context.Context, req interface{}) (interface{}, error) {
// 			c, ok := http.HttpContextFromContext(ctx)
// 			if !ok {
// 				return nil, errors.New("not find http context")
// 			}
// 			t, err := token.VerityExtractTokenFromRequest(c.Request(), o.claims, secret)
// 			if err != nil {
// 				return h(ctx, req)
// 			}
// 			ctx = NewClaimsContext(ctx, t.Claims)
// 			return h(ctx, req)
// 		}
// 	}
// }

// type claimsKey struct{}

// func NewClaimsContext(ctx context.Context, c jwt.Claims) context.Context {
// 	return context.WithValue(ctx, claimsKey{}, c)
// }

// func ClaimsFromContext(ctx context.Context) (c jwt.Claims, ok bool) {
// 	c, ok = ctx.Value(claimsKey{}).(jwt.Claims)
// 	return
// }
