package middlewares

import (
	"AMC_gateway/controllers/functions"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

// Middleware to check authorization token
func AuthMiddleware(ctx *context.Context) {
	// Get the authorization token from the request header
	logs.Info("About to check token")

	// Bypass authentication for OPTIONS requests
	if ctx.Input.Method() == "OPTIONS" {
		logs.Info("AuthMiddleware: OPTIONS request detected, bypassing authentication")
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Header("Origin"))
		ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Output.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Authorization")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		ctx.Output.SetStatus(204) // No Content
		// ctx.StopRun()             // Stop further processing
		return
	}

	// Add CORS headers
	ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Header("Origin"))
	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Authorization")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	authorization := ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	if token[0] == "Bearer" {
		verifyToken := functions.VerifyTokenNew(token[1])
		if verifyToken.StatusCode == 200 {
			logs.Info("User details are ", verifyToken.User)
			ctx.Input.SetData("user", verifyToken.User)

			return
		} else {
			ctx.Output.SetStatus(401)
			ctx.Output.JSON(map[string]string{"error": "You are not authorized to access this resource"}, false, false)
			return
		}

	} else {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string{"error": "Invalid authorization token"}, false, false)
		return
	}
	// If the token is valid, proceed to the next handler
}
