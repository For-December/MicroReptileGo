package utils

import (
	"encoding/json"
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"os"
)

func LoginBiliBili(driver selenium.WebDriver) {
	// 导航到网页
	if err := driver.Get("https://www.bilibili.com"); err != nil {
		fmt.Println("打开网页失败:", err)
		return
	}

	if err := driver.DeleteAllCookies(); err != nil {
		log.Fatalln(err)
		return
	}
	// 读取cookie并使用
	file, err := os.ReadFile("cookies.json")
	if err != nil {
		log.Fatalln(err)
		return
	}
	var cookies = make([]selenium.Cookie, 0)
	err = json.Unmarshal(file, &cookies)
	if err != nil {
		log.Fatalln(err)
		return
	}
	for _, cookie := range cookies {
		err := driver.AddCookie(&cookie)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
	err = driver.Refresh()
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("已成功登录！")
}
