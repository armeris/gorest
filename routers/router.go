package routers

import (
	"armeris/gorest/controllers"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = handlers.LoggingHandler(os.Stdout, handler)

		if(route.Name != "GetToken") {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(controllers.JwtMiddleware.Handler(handler))
		}else{
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)

		}
	}

	router.Handle("/",handlers.LoggingHandler(os.Stdout, http.FileServer(http.Dir("./views/"))))
	router.PathPrefix("/static/").Handler(handlers.LoggingHandler(os.Stdout, http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))))

	return router
}

