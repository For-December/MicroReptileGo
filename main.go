package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tebeka/selenium"
	"log"
)

var driver selenium.WebDriver

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	localSummary()
}

func localSummary() {
	router := gin.Default()
	// 定义一个路由处理函数
	router.POST("/ai", localSummaryHandler)

	if err := router.Run(":9876"); err != nil {
		log.Fatalln(err)
		return
	}
}

func biliBiliSummary() {
	router := gin.Default()
	// 定义一个路由处理函数
	router.POST("/ai", biliSummaryHandler)

	if err := router.Run(":9876"); err != nil {
		log.Fatalln(err)
		return
	}
}
