package controllers

import (
	"log"
	"net/http"

	"utwoo.com/go-crawler/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

//ScrapyOpenSourceArticles scrapys articles
func ScrapyOpenSourceArticles(c *gin.Context) {
	articles := make([]models.OpenSourceAticles, 0, 8)
	url := "https://opensource.com/tags/go"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".teaser__second").Each(func(i int, contentSelection *goquery.Selection) {
		article := contentSelection.Find("a").First()
		title := article.Text()
		url, _ := article.Attr("href")
		description := contentSelection.Find(".field-item").Text()

		articles = append(articles, models.OpenSourceAticles{Title: title, URL: url, Description: description})
	})

	c.HTML(http.StatusOK, "opensourceArticles.html", articles)
}
