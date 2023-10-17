package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"ecatrom/internal/domain/appcontext"
	"ecatrom/internal/domain/ecatrom"
	"ecatrom/internal/infrastructure/structx"

	"github.com/gin-gonic/gin"
)

func MakeEntriesRoute(r *gin.Engine, managerUseCases ecatrom.UseCases) {
	grp := r.Group("/entries")

	grp.POST("/", func(c *gin.Context) {
		createEntry(c, managerUseCases)
	})

	grp.GET("/all", func(c *gin.Context) {
		listEntries(c, managerUseCases)
	})
}

func createEntry(c *gin.Context, managerUseCases ecatrom.UseCases) {
	context := getContext(c)
	var questionEntity structx.Messages
	received, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(received, &questionEntity)

	result, err := managerUseCases.Create(context, questionEntity)
	respond(c, result, err)
}

func listEntries(c *gin.Context, managerUseCases ecatrom.UseCases) {
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
