package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type SkipperFunc func(*gin.Context) bool

// escape handler
func SkipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}

// Allow some prefix escaping an URL
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowHandlerSkipper skips handlers by name
func AllowHandlerSkipper(handlers ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		fmt.Println("AllowHandlerSkipper", c.HandlerName())
		for _, h := range handlers {
			if strings.HasSuffix(c.HandlerName(), h) {
				return true
			}
		}
		return false
	}
}
