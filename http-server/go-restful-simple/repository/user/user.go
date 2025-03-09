package user

import (
	"example.com/go-restful-simple/model"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"strconv"
)

var users = []model.User{
	{ID: 1, Name: "Alice", Age: 25},
	{ID: 2, Name: "Bob", Age: 30},
}

func GetAllUsers(req *restful.Request, resp *restful.Response) {
	err := resp.WriteAsJson(users)
	if err != nil {
		return
	}
}

func GetUser(req *restful.Request, resp *restful.Response) {
	idStr := req.PathParameter("user-id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		err := resp.WriteErrorString(http.StatusBadRequest, fmt.Sprintf("invalid user id, msg: %s", err))
		if err != nil {
			return
		}
		return
	}

	for _, user := range users {
		if user.ID == id {
			err := resp.WriteAsJson(user)
			if err != nil {
				return
			}
			return
		}
	}

	err = resp.WriteErrorString(http.StatusNotFound, "user not found")
	if err != nil {
		return
	}
}
