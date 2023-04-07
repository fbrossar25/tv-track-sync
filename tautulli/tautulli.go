package tautulli

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
)

func InitTautulli(router *gin.Engine) {
	router.POST("tautulli/webhook", webhook)
}

// mostly taken from github.com/jrudio/go-plex-client

type Data struct {
	// Action only "watched" action is consumed
	Action string `json:"action"`
	// MediaType only "movie" and "episode" are consumed
	MediaType    string `json:"media_type"`
	SeasonNum    string `json:"season_num"`
	EpisodeNum   string `json:"episode_num"`
	Title        string `json:"title"`
	PlexId       string `json:"plex_id"`
	ImdbId       string `json:"imdb_id"`
	ThetvdbId    string `json:"thetvdb_id"`
	ThemoviedbId string `json:"themoviedb_id"`
	TraktUrl     string `json:"trakt_url"`
}

func webhook(context *gin.Context) {
	sublog := log.With().Str("method", "tautulli/webhook").Logger()
	var data Data
	if err := context.BindJSON(&data); err != nil {
		b, payloadErr := io.ReadAll(context.Request.Body)
		if payloadErr != nil {
			sublog.Error().Err(err).Str("payloadReadError", payloadErr.Error()).Msg("Lecture payload impossible")
		} else {
			sublog.Error().Err(err).Str("payload", string(b)).Send()
		}
		return
	}
	handleEvent(&data)
}

func handleEvent(data *Data) {
	if data.Action != "watched" || (data.MediaType != "movie" && data.MediaType != "episode") {
		return
	}
	log.Debug().Interface("data", data).Msg("watched")
}
