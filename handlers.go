package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tebeka/selenium"
	"io"
	"log"
	"net/http"
	"time"
)

func localSummaryHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}
	url := string(body)
	println(url)

}
func aiSummaryHandler(c *gin.Context) {
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
	aiSummary, err := driver.FindElement(selenium.ByCSSSelector, ".ai-summary-popup-body-abstracts")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "抱歉，该视频暂时不能ai总结捏~",
		})
	}

	text, err := aiSummary.Text()
	log.Println(err)

	c.JSON(http.StatusOK, gin.H{
		"message": text,
	})
}
