package main

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
	//"os"
	"io/ioutil"
	"encoding/binary"
)

type IntegerBody struct {
    Int1 int
    Int2 int
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// return the sum of tow integers
  r.POST("/add", func(c *gin.Context) {
    var requestBody IntegerBody
    if err := c.BindJSON(&requestBody); err != nil {
        fmt.Println("error")
    }
    var sum int
    sum = requestBody.Int1 + requestBody.Int2

    //file, err := os.OpenFile("/root/project/job/watiinterview/sum.txt", os.O_WRONLY|os.O_CREATE, 0666)
	data := int64(sum)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
    if err := ioutil.WriteFile("/root/project/job/watiinterview/sum.txt", bytebuf.Bytes(), 0666); err != nil{
		fmt.Println("Writefile Error =", err)
		return
	}

    c.JSON(http.StatusOK, gin.H{
      "sum is": sum,
    })
  })
    return r
}

func main() {
	r := setupRouter()
	r.Run(":10000")
}

func TestAddRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	jsonBody := []byte(`{"Int1": 20, "Int2": 15}`)
 	bodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "/add", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
