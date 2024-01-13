package middlewares

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"polarisai/internal/domain/appcontext"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.Value(string(appcontext.AppContextKey)).(appcontext.Context).Logger()
		defer func() {
			err := recover()
			if err != nil {
				var brokenPipe bool

				ne, ok := err.(*net.OpError)

				if ok {
					se, ok := ne.Err.(*os.SyscallError)

					if ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)

				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)

					c.Error(err.(error))
					c.Abort()

					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
