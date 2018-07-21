package controllers

import (
	"fmt"
	"log"
	"net/http"

	"utwoo.com/go-crawler/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

//ScrapyStudyGolangTopics scrapys topics
func ScrapyStudyGolangTopics(c *gin.Context) {
	topics := make([]models.StudyGolangTopic, 0, 200)
	urlPages := "http://studygolang.com/topics?p=%d"

	for page := 1; page <= 4; page++ {
		response, err := http.Get(fmt.Sprintf(urlPages, page))
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
			img, _ := contentSelection.Find(".avatar img").Attr("src")
			url, _ := contentSelection.Find(".title a").Attr("href")
			title := contentSelection.Find(".title a").Text()

			topics = append(topics, models.StudyGolangTopic{ImgSrc: img, URL: url, Title: title})
		})
	}

	c.HTML(http.StatusOK, "studygolangTopics.html", topics)
}

//ScrapyStudyGolangArticles scrapys articles
func ScrapyStudyGolangArticles(c *gin.Context) {
	articles := make([]models.StudyGolangAticles, 0, 100)
	urlPages := "http://studygolang.com/articles?p=%d"

	for _, v := range []int{1, 2, 3, 4, 5} {
		response, err := http.Get(fmt.Sprintf(urlPages, v))
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("article").Each(func(i int, contentSelection *goquery.Selection) {
			article := contentSelection.Find(".row a").First()
			title, _ := article.Attr("title")
			url, _ := article.Attr("href")
			description := contentSelection.Find(".row .text").Text()

			articles = append(articles, models.StudyGolangAticles{Title: title, URL: url, Description: description})
		})
	}

	c.HTML(http.StatusOK, "studygolangArticles.html", articles)
}
