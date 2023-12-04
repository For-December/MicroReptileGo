package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tebeka/selenium"
	"log"
	"os"
	"os/exec"
)

var driver selenium.WebDriver

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	// 启动 selenium server
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前目录：", err)
		return
	}
	batFilePath := currentDir + "\\resources\\run.bat"
	cmd := exec.Command("cmd.exe", "/C", batFilePath)
	// 也可以直接相对路径".\\resources\\run.bat"
	err = cmd.Start()
	if err != nil {
		log.Fatalln(err)
		return
	}
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
