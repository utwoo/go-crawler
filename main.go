package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"utwoo.com/go-crawler/controllers"
)

func main() {
	path := os.Getenv("GOPATH") + "/src/utwoo.com/go-crawler"

	router := gin.Default()
	router.Static("/static", path+"/static")
	router.LoadHTMLGlob(path + "/views/*.html")

	crawlersGroup := router.Group("/crawlers")
	{
		crawlersGroup.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusPermanentRedirect, "/crawlers/studygolang/topics")
		})

		crawlersGroup.GET("/studygolang/topics", controllers.ScrapyStudyGolangTopics)
		crawlersGroup.GET("/studygolang/articles", controllers.ScrapyStudyGolangArticles)
		crawlersGroup.GET("/golangtc/topics", controllers.ScrapyGolangTcTopics)
		crawlersGroup.GET("/opensource/articles", controllers.ScrapyOpenSourceArticles)
	}

	router.Run(":8081")
}
