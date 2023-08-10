package main

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
)

type IntegerBody struct {
    Int1 int
    Int2 int
}

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  // return the sum of tow integers
  r.POST("/add", func(c *gin.Context) {
    var requestBody IntegerBody
    if err := c.BindJSON(&requestBody); err != nil {
        fmt.Println("error")
    }
    var sum int
    sum = requestBody.Int1 + requestBody.Int2
    c.JSON(http.StatusOK, gin.H{
      "sum is": sum,
    })
  })
  
  r.Run(":10000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
