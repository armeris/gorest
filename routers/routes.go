package routers

import (
	"armeris/gorest/controllers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetToken",
		"GET",
		"/get-token",
		controllers.GetTokenHandler,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		controllers.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoID}",
		controllers.TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		controllers.TodoCreate,
	},
}
