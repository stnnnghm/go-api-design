package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stnnnghm/go-api-design/cmd/api/handlers/getuser"
	"github.com/stnnnghm/go-api-design/pkg/application"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users", getuser.Do(app))
	return mux
}
