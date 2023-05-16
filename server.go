package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tyler-sommer/stick"
	"io"
	"os"
)

func login(c *gin.Context) {

	fsRoot, err := os.Getwd() // Templates are loaded relative to this directory.
	if err == nil {
		fmt.Println("error", err)
	}
	env := stick.New(stick.NewFilesystemLoader(fsRoot))
	p := map[string]stick.Value{"name": "World"}
	env.Execute("templates/login.html.twig", c.Writer, p)
}

func index(c *gin.Context) {
	data, _ := io.ReadAll(c.Request.Body)

	fmt.Println("test")
	fsRoot, _ := os.Getwd() // Templates are loaded relative to this directory.
	env := stick.New(stick.NewFilesystemLoader(fsRoot))
	env.Execute("templates/index.html.twig", c.Writer, nil)

	fmt.Println(string(data))

}
func main() {
	server := gin.Default()
	server.GET("/", login)
	server.POST("/submmited", index)
	server.Run(":8080")
}
