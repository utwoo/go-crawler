package models

//StudyGolangTopic is crawlers topics in studygolang web
type StudyGolangTopic struct {
	URL    string
	ImgSrc string
	Title  string
}

//StudyGolangAticles is crawlers articles in studygolang web
type StudyGolangAticles struct {
	URL         string
	Title       string
	Description string
}
