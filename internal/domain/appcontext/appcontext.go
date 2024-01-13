package appcontext

import (
	"context"
	"strconv"
	"time"

	"polarisai/internal/infrastructure/logger/logwrapper"

	"github.com/gin-gonic/gin"
)

// ContextKey is the key for the app context
type ContextKey string

const (
	// AppContextKey is the key for the app context
	AppContextKey               ContextKey = "appContext"
	ginContextKey               ContextKey = "ginContext"
	defaultBackgroundContextKey ContextKey = "ctx"
)

// Context is wrapper for gin context 'n go context. Provide logger with trace_id.
type Context interface {
	Done()
	SetLogger(logger logwrapper.LoggerWrapper)
	Logger() logwrapper.LoggerWrapper
	Context() context.Context
	TraceID() string
	SpanID() string
	WithValue(key, val interface{})
	Value(key interface{}) interface{}
	TTL() *int64
}

// New returns a new app context
func New(ctx context.Context, ginContext *gin.Context) Context {
	return &appContext{
		defaultBackgroundContext: ctx,
		ginContext:               ginContext,
	}
}

// NewBackground returns a new background void context
func NewBackground(g *gin.Context) Context {
	ctx := context.Background()

	return &appContext{
		defaultBackgroundContext: ctx,
		ginContext:               g,
	}
}

// GetAppContext returns the app context
func GetAppContext(c *gin.Context) Context {
	return c.MustGet(string(AppContextKey)).(Context)
}

type appContext struct {
	logger                   logwrapper.LoggerWrapper
	defaultBackgroundContext context.Context
	ginContext               *gin.Context
}

func (appContext *appContext) SetLogger(logger logwrapper.LoggerWrapper) {
	appContext.logger = logger
}

func (appContext *appContext) Logger() logwrapper.LoggerWrapper {
	return appContext.logger
}

func (appContext *appContext) Context() context.Context {
	return appContext.defaultBackgroundContext
}

func (appContext *appContext) Done() {
	appContext.ginContext = nil
	appContext.defaultBackgroundContext = nil
	appContext.logger = nil
}

func (appContext *appContext) TraceID() string {
	return "nil"
}

func (appContext *appContext) SpanID() string {
	return "span.ID"
}

func (appContext *appContext) WithValue(key, val interface{}) {
	appContext.defaultBackgroundContext = context.WithValue(appContext.defaultBackgroundContext, key, val)
}

func (appContext *appContext) Value(key interface{}) interface{} {
	return appContext.defaultBackgroundContext.Value(key)
}

func (appContext *appContext) TTL() *int64 {
	ttl := appContext.Value("ttl")

	if ttl == nil {
		return nil
	}

	ttlInt, _ := strconv.Atoi(ttl.(string))
	timeUnix := time.Now().Add(time.Duration(ttlInt) * time.Minute).Unix()

	return &timeUnix
}
