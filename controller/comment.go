package controller

import (
	"go-pg/model"
	"log"

	"github.com/kataras/iris/v12"
)

func GetComments(ctx iris.Context) {
	var comments []model.Comment
	err := DB.Model(&comments).Relation("User").Relation("Post").Select()
	if err != nil {
		panic(err)
	}
	ctx.JSON(comments)
}

func GetCommentById(ctx iris.Context) {
	id := ctx.Params().Get("userId")

	var comment model.Comment

	err := DB.Model(&comment).Relation("Post").Relation("User").Where("id = ?", id).Select()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(comment)

}

func UpdateComment(ctx iris.Context) {
	var data map[string]interface{}
	ctx.ReadJSON(&data)

	id := ctx.Params().Get("id")

	_, err := DB.Model(&data).TableExpr("blog.comment").Where("id = ?", id).Update()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}

	ctx.JSON("Update thành công")
}

func DeleteComment(ctx iris.Context) {
	id := ctx.Params().Get("id")
	_, err := DB.Model((*model.Comment)(nil)).Where("id = ?", id).Delete()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.JSON("Xóa thành công")
}

func CreateComment(ctx iris.Context) {
	var data map[string]interface{}
	ctx.ReadJSON(&data)

	data["user_id"] = 1
	data["post_id"] = 1

	_, err := DB.Model(&data).TableExpr("blog.comment").Insert()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON("Comment thành công")
}
