package application

import (
	"github.com/akmittal/go-rest-boilerplate/pkg/config"
	"github.com/akmittal/go-rest-boilerplate/pkg/db"
	"github.com/akmittal/go-rest-boilerplate/pkg/router"
	"github.com/akmittal/go-rest-boilerplate/pkg/server"
	"github.com/akmittal/go-rest-boilerplate/pkg/user"
)

type Application struct {
	DB     *db.DB
	Cfg    *config.Config
	Router *router.Router
	Server *server.Server
}

func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}
	router, err := router.Get()
	server := server.Get(cfg.GetAppHost(), router)

	return &Application{
		DB:     db,
		Cfg:    cfg,
		Router: router,
		Server: server,
	}, nil
}

func (a *Application) Start() error {
	a.RegisterRoutes()
	return a.Server.Start()
}

func (a *Application) RegisterRoutes() {

	a.Router.Get("/", user.UserController())
}
