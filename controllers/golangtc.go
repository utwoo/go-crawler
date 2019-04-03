package controllers

import (
	"fmt"
	"log"
	"net/http"

	"utwoo.com/go-crawler/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

//ScrapyGolangTcTopics scrapys topics
func ScrapyGolangTcTopics(c *gin.Context) {
	topics := make([]models.GolangTcTopic, 0, 200)
	urlPages := "https://www.golangtc.com/?p=%d"

	for page := 1; page <= 10; page++ {
		response, err := http.Get(fmt.Sprintf(urlPages, page))
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("article").Each(func(i int, contentSelection *goquery.Selection) {
			img, _ := contentSelection.Find("img").Attr("src")
			url, _ := contentSelection.Find(".title").Find("a").Attr("href")
			title := contentSelection.Find(".title").Text()

			topics = append(topics, models.GolangTcTopic{ImgSrc: img, URL: url, Title: title})
		})
	}

	c.HTML(http.StatusOK, "golangtcTopics.html", topics)
}
