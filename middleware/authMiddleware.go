package middleware

import (
	"strings"
	"tomatoPaper/pkg/jwt"
	"tomatoPaper/web"
)

func AuthMiddleware() web.HandleFunc {
	return func(c *web.Context) {
		auth := c.Req.Header.Get("Authorization")
		if auth == "" {
			c.RespJSON(401, "Unauthorized")
			c.Abort()
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.RespJSON(401, "Unauthorized")
			c.Abort()
			return
		}
		token, err := jwt.ValidateToken(parts[1])
		if err != nil {
			c.RespJSON(401, "Unauthorized")
			c.Abort()
			return
		}
		c.Set(jwt.ContextKeyUserObj, token)
		c.Next()
	}
}
