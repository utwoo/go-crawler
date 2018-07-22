package models

//StudyGolangTopic crawlers topics in studygolang web
type StudyGolangTopic struct {
	URL    string
	ImgSrc string
	Title  string
}

//StudyGolangAticles crawlers articles in studygolang web
type StudyGolangAticles struct {
	URL         string
	Title       string
	Description string
}
