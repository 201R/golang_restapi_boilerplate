package middlewares

import (
	"net/http"
	"strings"

	"github.com/201R/go_api_boilerplate/packages/app"
	"github.com/201R/go_api_boilerplate/packages/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		var (
			success bool = false
			msg     string
			err     error
		)

		app.HTTPRes(c, http.StatusAccepted, http.StatusText(http.StatusAccepted), "Authenticated...")

		token := extractToken(c.Request)
		if token == "" {
			msg = "error missing jwt"
		} else {
			_, err = helpers.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					msg = "error jwt validation"
				default:
					msg = "error expired jwt"
				}
			}
		}

		if !success {
			app.HTTPRes(c, http.StatusNotAcceptable, http.StatusText(http.StatusNotAcceptable), msg)
			c.Abort()
			return
		}

		c.Next()
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
