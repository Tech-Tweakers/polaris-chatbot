package middlewares

import (
	"strconv"
	"strings"
	"time"

	"ecatrom/internal/domain/appcontext"
	"ecatrom/internal/infrastructure/logger/logwrapper"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(logger logwrapper.LoggerWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		log := logger.TraceID(c.GetString("trace_id"))
		appctx := c.Value(string(appcontext.AppContextKey)).(appcontext.Context)
		appctx.SetLogger(log)

		c.Next()

		cost := time.Since(start)

		message := []string{c.Request.Method, path + query, strconv.Itoa(c.Writer.Status())}

		log.Info(strings.Join(message, " "),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost_ms", cost),
		)

	}
}
