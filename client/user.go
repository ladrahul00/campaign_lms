package client

// import (
// 	"context"

// 	"github.com/micro/go-micro/v2"
// 	"github.com/micro/go-micro/v2/server"
// 	"github.com/wolf00/campaign_lms/constants"
// 	user "github.com/wolf00/golang_lms/user/proto/user"
// )

// type userKey struct{}

// // FromContext retrieves the client from the Context
// func UserFromContext(ctx context.Context) (user.UserService, bool) {
// 	c, ok := ctx.Value(userKey{}).(user.UserService)
// 	return c, ok
// }

// // Client returns a wrapper for the HeimdallClient
// func UserWrapper(service micro.Service) server.HandlerWrapper {
// 	client := user.NewUserService(constants.UserService, service.Client())

// 	return func(fn server.HandlerFunc) server.HandlerFunc {
// 		return func(ctx context.Context, req server.Request, rsp interface{}) error {
// 			ctx = context.WithValue(ctx, userKey{}, client)
// 			return fn(ctx, req, rsp)
// 		}
// 	}
// }
