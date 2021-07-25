package controller

import (
	"go-pg/model"
	"log"
	"strconv"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

var DB *pg.DB

func GetUsers(ctx iris.Context) {
	var user []model.User

	err := DB.Model(&user).Relation("Posts").Relation("Profile").Relation("Roles").Select()
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

func CreateRole(ctx iris.Context){
	var data map[string]interface{}
	ctx.ReadJSON(&data)

	_, err := DB.Model(&data).TableExpr("auth.role").Insert()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}

	ctx.JSON("Tạo role thành công")
}

func SetUserRole(ctx iris.Context){
	userId,_ := strconv.Atoi(ctx.Params().Get("userId"))
	var user_role model.UserRole
	ctx.ReadJSON(&user_role)

	user_role.UserId = userId

	_, err := DB.Model(&user_role).Insert()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON("Set quyền thành công")
}

func CreateProfile(ctx iris.Context){
	userId := ctx.Params().Get("userId")
	var data map[string]interface{}
	ctx.ReadJSON(&data)

	data["user_id"]=userId

	_, err := DB.Model(&data).TableExpr("auth.profile").Insert()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON("Tạo profile thành công")
}

func UpdateProfile(ctx iris.Context){

}
