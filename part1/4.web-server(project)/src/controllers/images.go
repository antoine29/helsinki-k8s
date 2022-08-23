package controllers

import (
	fileHelpers "antoine29/go/web-server/src"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var picSumUrl = "https://picsum.photos/id/%s/400"

func GetImage(c *gin.Context) {
	name := c.Param("name")
	if fileHelpers.LocalFileExists(name) {
		fmt.Printf("Getting '%s' from local cache\n", name)
		c.File(fmt.Sprintf("/tmp/%s", name))
		return
	}

	downloadPath := fmt.Sprintf(picSumUrl, name)
	downloadError := fileHelpers.DownloadFile(downloadPath, "/tmp/"+name)
	if downloadError == nil {
		fmt.Printf("Downloading '%s' from picsum", name)
		if fileHelpers.LocalFileExists(name) {
			c.File(fmt.Sprintf("/tmp/%s", name))
			return
		}

		fmt.Println("file not found after downloading it")
		c.Status(http.StatusNotFound)
		return
	}

	fmt.Printf("error downloading '%s'\n", downloadPath)
	c.Status(http.StatusNotFound)
}
