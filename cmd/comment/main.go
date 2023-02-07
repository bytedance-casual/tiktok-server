package main

import (
	"log"
	comment "tiktok-server/kitex_gen/comment/commentservice"
)

func main() {
	svr := comment.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
