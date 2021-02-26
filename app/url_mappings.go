package app

import "github.com/projects/REST-GO-GIN/Controller/users"

func mapUrls() {

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)

}
