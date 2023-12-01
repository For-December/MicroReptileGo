package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tebeka/selenium"
	"log"
)

var driver selenium.WebDriver

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	router := gin.Default()
	biliBiliSummary(router)
	webScreen(router)

	if err := router.Run(":9876"); err != nil {
		log.Fatalln(err)
		return
	}
}

func webScreen(router *gin.Engine) {
	// 定义一个路由处理函数
	router.POST("/screen", webScreenHandler)

}

func localSummary(router *gin.Engine) {
	// 定义一个路由处理函数
	router.POST("/ai", localSummaryHandler)

}

func biliBiliSummary(router *gin.Engine) {
	// 定义一个路由处理函数
	router.POST("/ai", biliSummaryHandler)

}
