package controller

import (
	"go-pg/model"
	"log"
	"time"

	"github.com/kataras/iris/v12"
)

func GetPosts(ctx iris.Context) {
	var posts []model.Post
	err := DB.Model(&posts).Relation("User").Select()
	if err != nil {
		panic(err)
	}
	ctx.JSON(posts)
}

func GetPostById(ctx iris.Context) {
	id := ctx.Params().Get("userId")

	var post model.Post

	err := DB.Model(&post).Relation("User").Where("id = ?", id).Select()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(post)

}

func CreatePost(ctx iris.Context) {
	var data map[string]interface{}
	ctx.ReadJSON(&data)

	data["user_id"] = 1

	_, err := DB.Model(&data).TableExpr("blog.post").Insert()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}

	ctx.JSON("Tạo bài viết thành công")
}

func UpdatePost(ctx iris.Context) {
	var data map[string]interface{}
	ctx.ReadJSON(&data)

	id := ctx.Params().Get("id")

	data["updated_at"] = time.Now()
	_, err := DB.Model(&data).TableExpr("blog.post").Where("id = ?", id).Update()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}

	ctx.JSON("Cập nhật thành công")
}

func DeletePost(ctx iris.Context) {
	id := ctx.Params().Get("id")
	_, err := DB.Model((*model.Post)(nil)).Where("id = ?", id).Delete()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.JSON("Xóa thành công")
}
