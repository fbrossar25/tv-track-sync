package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"tv-track-sync/tautulli"
	"tv-track-sync/util"
)

func main() {
	log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Démarrage du webservice")
	util.LoadConfig()

	router := gin.Default()
	router.GET("/check", check)
	//plex.InitPlex(router)
	tautulli.InitTautulli(router)
	routerErr := router.Run(":8090")
	if routerErr != nil {
		log.Error().Stack().Err(routerErr).Msg("Erreur au démarrage du webservice")
	}
}

func check(context *gin.Context) {
	context.String(http.StatusOK, "OK")
}
