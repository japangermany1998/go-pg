package controller

import (
	"go-pg/model"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

var DB *pg.DB

func GetUsers(ctx iris.Context) {
	var user []model.User
	err := DB.Model(&user).Relation("Posts").Select()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.JSON(user)
}

func GetUserById(ctx iris.Context) {
	id := ctx.Params().Get("userId")

	var user model.User

	err := DB.Model(&user).Relation("Posts").Where("id = ?", id).Select()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(user)
}

func Register(ctx iris.Context) {
	var data map[string]string
	ctx.ReadJSON(&data)

	if data["password"] != data["passwordconfirm"] {
		ctx.StatusCode(400)
		ctx.JSON(map[string]string{
			"message": "password doesn't match",
		})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := model.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  string(password),
	}
	_, err := DB.Model(&user).Insert()
	if err != nil {
		panic(err)
	}
	ctx.JSON(user)
}

func UpdateUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var data map[string]interface{}
	ctx.ReadJSON(&data)
	_, err := DB.Model(&data).TableExpr("auth.users").Where("id = ?", id).Update()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Cập nhật thành công")
}
