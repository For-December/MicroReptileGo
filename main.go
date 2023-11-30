package main

import (
	"MicroReptileGo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tebeka/selenium"
	"log"
	"time"
)

var driver selenium.WebDriver

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	localSummary()
}

func localSummary() {
	// 启动 Chrome 浏览器
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, _ = selenium.NewRemote(caps, "")
	err := driver.SetImplicitWaitTimeout(30 * time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 导航到网页
	if err := driver.Get("https://www.bilibili.com"); err != nil {
		fmt.Println("打开网页失败:", err)
		return
	}

	defer func(driver selenium.WebDriver) {
		log.Println("浏览器推出了？")
		err := driver.Quit()
		if err != nil {
			log.Fatalln(err)
		}
	}(driver)

	router := gin.Default()
	// 定义一个路由处理函数
	router.POST("/ai", localSummaryHandler)

	if err := router.Run(":9876"); err != nil {
		log.Fatalln(err)
		return
	}
}

func biliBiliSummary() {

	// 启动 Chrome 浏览器
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, _ = selenium.NewRemote(caps, "")
	err := driver.SetImplicitWaitTimeout(30 * time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer func(driver selenium.WebDriver) {
		err := driver.Quit()
		if err != nil {
			log.Fatalln(err)
		}
	}(driver)
	utils.LoginBiliBili(driver)

	router := gin.Default()
	// 定义一个路由处理函数
	router.POST("/ai", aiSummaryHandler)

	if err := router.Run(":9876"); err != nil {
		log.Fatalln(err)
		return
	}
}
