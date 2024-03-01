package app

import (
	"sola-test-task/internal/config"
	"sola-test-task/internal/controller"
	"sola-test-task/internal/middleware"
	"sola-test-task/internal/provider/sql/postgres"
	postStRepo "sola-test-task/internal/repository/station/postgres"
	"sola-test-task/internal/router"
	transServ "sola-test-task/internal/service/translation"
	valServ "sola-test-task/internal/service/validation"
	"sola-test-task/pkg/log"
)

func Run() error {
	conf, err := config.NewConfig()
	if err != nil {
		return err
	}
	log := log.NewLogger(conf.Server.LogMode, conf.Server.LogPath)

	postProv, err := postgres.NewPostgresProvider(conf)
	if err != nil {
		return err
	}
	if err = postProv.MigrateUp(); err != nil {
		return err
	}

	stRepo := postStRepo.NewStationPostgresRepo(postProv.Conn())

	valServ := valServ.NewValidationService()
	trServ, err := transServ.NewTranslationService()
	if err != nil {
		return err
	}

	reqMid := middleware.NewRequestCtx(log)

	stCont := controller.NewStationController(stRepo)

	rout := router.NewRouter(log, conf, reqMid, valServ, trServ, stCont)

	return rout.SetupHandlers()
}
