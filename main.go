
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func sayHello( c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"message":"Hello"})
}

func main (){
	router :=gin.Default();
	router.GET("/hello", sayHello)
	router.Run("localhost:8080")
}

