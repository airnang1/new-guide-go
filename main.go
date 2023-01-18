package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// "github.com/gorilla/mux"
	"example.com/packages/greetings"
)

func newGenericFunc[age int64 | float32](myAge age) {
	val := int(myAge) + 1
	fmt.Println(val)
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": greetings.GetGreeting(),
		})
	})
	r.Run(":9888") // listen and serve on 0.0.0.0:8080
	fmt.Println("Go Generic Tutorial", greetings.Hello("abc"))

	var testAge int64 = 28
	var testAge2 float32 = 28.4

	newGenericFunc(testAge)
	newGenericFunc(testAge2)
	// r := mux.NewRouter()
	fmt.Println(greetings.GetGreeting())

	// fmt.Println(quote.Go())
	// http.ListenAndServe(":9000", r)

}
