package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

//StartApp is func
func StartApp() {
	mapUrls()
	router.Run(":8080")

}
