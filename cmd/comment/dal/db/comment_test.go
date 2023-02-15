package db

import (
	"testing"
	"tiktok-server/internal/conf"
)

func TestCreateComment(t *testing.T) {
	conf.Init()
	Init()
	//result := DB.Create(&Comment{
	//	VideoId: 7,
	//	UserId:  2,
	//	Content: &Content{
	//		Content: "测试评论内容",
	//	},
	//})
	//if result.Error != nil {
	//	fmt.Println(result.Error)
	//}

	//ctx := context.Background()
	//comments, err := QueryComment(7, ctx)
	//if err != nil {
	//	panic(err)
	//}
	//for _, c := range comments {
	//	fmt.Printf("%v", c)
	//}

	//err := DeleteComment(7, 2, ctx)
	//if err != nil {
	//	panic(err)
	//}
}
