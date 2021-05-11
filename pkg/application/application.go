package application

import (
	"github.com/stnnnghm/go-api-design/db"
	"github.com/stnnnghm/go-api-design/pkg/config"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  db,
		Cfg: cfg,
	}, nil
}
