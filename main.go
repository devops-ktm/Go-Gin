// main.go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })

    r.Run() // default listens on :8080
}
