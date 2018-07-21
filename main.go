package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"utwoo.com/go-crawler/controllers"
)

func main() {
	router := gin.Default()
	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLGlob("views/*.html")

	crawlersGroup := router.Group("/crawlers")
	{
		crawlersGroup.GET("/studygolang/topics", controllers.ScrapyStudyGolangTopics)
		crawlersGroup.GET("/studygolang/articles", controllers.ScrapyStudyGolangArticles)
	}

	router.Run(":8080")
}
