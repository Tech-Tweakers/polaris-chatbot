package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"polarisai/internal/domain/appcontext"
	"polarisai/internal/domain/ecatrom"
	"polarisai/internal/infrastructure/structx"

	"github.com/gin-gonic/gin"
)

func MakeEntriesRoute(r *gin.Engine, managerUseCases ecatrom.UseCases) {
	grpChat := r.Group("/chat")

	grpChat.POST("/send", func(c *gin.Context) {
		createChatEntry(c, managerUseCases)
	})

	grpChat.GET("/history", func(c *gin.Context) {
		listChatEntries(c, managerUseCases)
	})

	grpCode := r.Group("/code")

	grpCode.POST("/send", func(c *gin.Context) {
		createCodeEntry(c, managerUseCases)
	})

	grpCode.GET("/history", func(c *gin.Context) {
		listCodeEntries(c, managerUseCases)
	})
}

func createChatEntry(c *gin.Context, managerUseCases ecatrom.UseCases) {
	context := getContext(c)
	var questionEntity structx.Messages
	received, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(received, &questionEntity)

	kind := "chat"
	result, err := managerUseCases.Create(context, questionEntity, kind)
	respond(c, result, err)
}

func listChatEntries(c *gin.Context, managerUseCases ecatrom.UseCases) {
	context := getContext(c)
	result, err := managerUseCases.ListAll(context)
	respond(c, result, err)
}

func createCodeEntry(c *gin.Context, managerUseCases ecatrom.UseCases) {
	context := getContext(c)
	var questionEntity structx.Messages
	received, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(received, &questionEntity)

	kind := "code"
	result, err := managerUseCases.Create(context, questionEntity, kind)
	respond(c, result, err)
}

func listCodeEntries(c *gin.Context, managerUseCases ecatrom.UseCases) {
	context := getContext(c)
	result, err := managerUseCases.ListAll(context)
	respond(c, result, err)
}

func getContext(c *gin.Context) appcontext.Context {
	return c.Value(string(appcontext.AppContextKey)).(appcontext.Context)
}

func respond(c *gin.Context, result interface{}, err error) {
	if err != nil {
		if re, ok := err.(*ecatrom.DomainError); ok {
			fmt.Printf("re: %v\n", re)
			c.JSON(re.StatusCode, gin.H{"error": re.Err.Error(), "retryable": re.Retryable, "message": re.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "retryable": true})
		}
		return
	}
	c.JSON(http.StatusOK, result)
}
