package main

import (
	article "blinkable/server/kitex_gen/Article/articleservice"
	i "blinkable/server/service/article/init"
	"log"
)

func main() {
	i.InitLogger()
	i.InitNacos()
	i.InitDB()
	i.InitRedis()
	svr := article.NewServer(new(ArticleServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
