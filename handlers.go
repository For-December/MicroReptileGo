package main

import (
	"MicroReptileGo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tebeka/selenium"
	"io"
	"log"
	"time"
)

func webScreenHandler(c *gin.Context) {
	// 启动 Chrome 浏览器
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, _ = selenium.NewRemote(caps, "")
	err := driver.SetImplicitWaitTimeout(30 * time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer func(driver selenium.WebDriver) {
		log.Println("浏览器推出了？")
		err := driver.Quit()
		if err != nil {
			log.Fatalln(err)
		}
	}(driver)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}
	url := string(body)
	println(url)

	// 导航到网页
	if err := driver.Get(url); err != nil {
		fmt.Println("打开网页失败:", err)
		return
	}

	// 最大化浏览器窗口
	if err := driver.MaximizeWindow(""); err != nil {
		log.Fatal("无法最大化窗口：", err)
	}
	time.Sleep(7 * time.Second)
	screenshot, err := driver.Screenshot()
	if err != nil {
		log.Println(err)
		return
	}
	num, err := c.Writer.Write(screenshot)
	if err != nil {
		log.Println(num)
		return
	}

}
func localSummaryHandler(c *gin.Context) {

	// 启动 Chrome 浏览器
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, _ = selenium.NewRemote(caps, "")
	err := driver.SetImplicitWaitTimeout(30 * time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer func(driver selenium.WebDriver) {
		log.Println("浏览器推出了？")
		err := driver.Quit()
		if err != nil {
			log.Fatalln(err)
		}
	}(driver)

	// 导航到网页
	if err := driver.Get("http://localhost:3000/"); err != nil {
		fmt.Println("打开网页失败:", err)
		return
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}
	url := string(body)
	println(url)

	input, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"__next\"]/div[1]/main/div/form/input")
	if err != nil {
		log.Fatalln(err)
		return
	}
	if err := input.SendKeys(url); err != nil {
		log.Fatalln(err)
		return
	}
	time.Sleep(2 * time.Second)
	button, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"__next\"]/div[1]/main/div/form/button")
	if err != nil {
		log.Fatalln(err)
		return
	}

	if err := button.Click(); err != nil {
		log.Fatalln(err)
		return
	}

	time.Sleep(35 * time.Second)

	summary := ""
	markdown, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"__next\"]/div[1]/main/div/div/div[2]/div")
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(markdown)

	ps, err := markdown.FindElements(selenium.ByTagName, "p")
	if err != nil {
		return
	}
	for _, p := range ps {
		text, _ := p.Text()
		summary += text + "\n"
	}
	summary += "\n"
	lis, err := markdown.FindElements(selenium.ByTagName, "li")
	if err != nil {
		return
	}
	for _, li := range lis {
		text, _ := li.Text()
		summary += text + "\n"
	}

	num, err := c.Writer.Write([]byte(summary))
	if err != nil {
		log.Fatalln(num, err)
		return
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"message": summary,
	//})

}
func biliSummaryHandler(c *gin.Context) {

	// 启动 Chrome 浏览器
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, _ = selenium.NewRemote(caps, "")
	err := driver.SetImplicitWaitTimeout(5 * time.Second)
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

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}
	url := string(body)
	println(url)

	if err := driver.Get(url); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(5 * time.Second)

	element, err := driver.FindElement(selenium.ByCSSSelector, ".video-ai-assistant-icon.video-toolbar-item-icon")
	if err != nil {
		log.Fatalln(err)
	}
	if err := element.Click(); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(5 * time.Second)
	summary := ""
	aiSummary, err := driver.FindElement(selenium.ByCSSSelector, ".ai-summary-popup-body-abstracts")
	if err == nil {
		text, _ := aiSummary.Text()
		summary += text + "\n\n"
	} else {
		log.Println(err)
	}

	if outline, err := driver.FindElement(
		selenium.ByCSSSelector,
		".ai-summary-popup-body-outline"); err == nil {
		sections, err := outline.FindElements(selenium.ByCSSSelector, ".section")
		if err != nil {
			log.Println(err)
		} else {
			for _, section := range sections {
				titleElement, err := section.FindElement(selenium.ByCSSSelector, ".section-title")
				if err != nil {
					log.Println(err)
				} else {
					title, _ := titleElement.Text()
					summary += title + "\n"
				}
				contentElements, err := section.FindElements(selenium.ByCSSSelector, ".bullet")
				if err != nil {
					log.Println(err)
				} else {
					for _, contentElement := range contentElements {
						timeTagElem, err := contentElement.FindElement(selenium.ByCSSSelector, ".timestamp-inner")
						if err != nil {
							log.Println(err)
						}
						contentElem, err := contentElement.FindElement(selenium.ByCSSSelector, ".content")
						if err != nil {
							log.Println(err)
						}
						timeTag, _ := timeTagElem.Text()
						content, _ := contentElem.Text()
						summary += timeTag + "  "
						summary += content + "\n"
					}
				}
				summary += "\n"

			}
		}
	}

	if summary == "" {

		_, err := c.Writer.Write([]byte("该视频不能总结捏~"))
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	num, err := c.Writer.Write([]byte(summary))
	if err != nil {
		log.Fatalln(num)
		return
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"message": text,
	//})
}
