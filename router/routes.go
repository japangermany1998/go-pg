package router

import (
	"go-pg/controller"

	"github.com/kataras/iris/v12"
)

func AllRoutes(app *iris.Application){
	app.Get("/example",func(c iris.Context) {})
	app.Get("/api/user",controller.GetUsers)

	app.Get("/api/user/{userId}",controller.GetUserById)

	app.Post("/api/register",controller.Register)

	app.Put("/api/user/{id}",controller.UpdateUser)

	app.Get("/api/posts",controller.GetPosts)

	app.Get("/api/posts/{postId}",controller.GetPostById)

	app.Post("/api/post/create",controller.CreatePost)

	app.Put("/api/post/{id}",controller.UpdatePost)

	app.Delete("/api/post/{id}",controller.DeletePost)

	app.Get("/api/comment",controller.GetComments)

	app.Get("/api/comment/{commentId}",controller.GetCommentById)

	app.Post("/api/comment/create",controller.CreateComment)

	app.Put("/api/comment/{id}",controller.UpdateComment)

	app.Delete("/api/comment/{id}",controller.DeleteComment)
}