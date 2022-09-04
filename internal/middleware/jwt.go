package middleware

import (
	"context"
	"github.com/buddyxiao/navmark-go/internal/model/po"
	"github.com/buddyxiao/navmark-go/internal/model/req"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"time"
)

var (
	_ app.HandlerFunc = (*&jwt.HertzJWTMiddleware{}).MiddlewareFunc()
)

const jwtSecret = "navmark-jwt"

// 模拟数据库中的用户
var defaultUser = po.User{
	ID:       "1000",
	Username: "admin",
	Password: "admin123",
}

func AuthMiddleware() app.HandlerFunc {
	jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(jwtSecret),
		Timeout:    2 * time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals req.LoginReq
			if err := c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			if username == defaultUser.Username && password == defaultUser.Password {
				return &User{
					UserName:  userID,
					LastName:  "Hertz",
					FirstName: "CloudWeGo",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// 授权
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(*po.User); ok && v.Username == defaultUser.Username {
				return true
			}
			return false
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
	})
}
