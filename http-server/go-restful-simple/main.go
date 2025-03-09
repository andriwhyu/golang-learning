package main

import (
	"example.com/go-restful-simple/model"
	"example.com/go-restful-simple/repository/user"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	restfulApi := new(restful.WebService)

	restfulApi.Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	restfulApi.Route(
		restfulApi.GET("/").
			To(user.GetAllUsers).
			Doc("Get all users").
			Writes([]model.User{}))

	restfulApi.Route(
		restfulApi.GET("/{user-id}").
			To(user.GetUser).
			Doc("Get a user by ID").
			Param(restfulApi.
				PathParameter("user-id", "ID of the user").
				DataType("int")).
			Writes([]model.User{}))

	restful.Add(restfulApi)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
