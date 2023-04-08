package tautulli

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"tv-track-sync/util"
)

func InitTautulli(router *gin.Engine) {
	router.POST("tautulli/webhook", webhook)
}

// Event mostly taken from github.com/jrudio/go-plex-client
type Event struct {
	// Action only "watched" action is consumed
	Action string `json:"action" bson:"action"`
	// MediaType only "movie" and "episode" are consumed
	MediaType    string `json:"media_type" bson:"media_type"`
	User         string `json:"user" bson:"user"`
	SeasonNum    string `json:"season_num" bson:"season_num"`
	EpisodeNum   string `json:"episode_num" bson:"episode_num"`
	Title        string `json:"title" bson:"title"`
	PlexId       string `json:"plex_id" bson:"plex_id"`
	ImdbId       string `json:"imdb_id" bson:"imdb_id"`
	ThetvdbId    string `json:"thetvdb_id" bson:"thetvdb_id"`
	ThemoviedbId string `json:"themoviedb_id" bson:"themoviedb_id"`
	TraktUrl     string `json:"trakt_url" bson:"trakt_url"`
}

func webhook(context *gin.Context) {
	sublog := log.With().Str("method", "tautulli/webhook").Logger()
	var data Event
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

func handleEvent(data *Event) {
	if data.Action != "watched" || (data.MediaType != "movie" && data.MediaType != "episode") {
		return
	}
	saveData(data)
}

var EventsCollection = "events"

func saveData(data *Event) {
	db := util.ConnectToMongoDB()
	_, err := db.Collection(EventsCollection).InsertOne(context.TODO(), data)
	if err != nil {
		log.Error().
			Stack().
			Err(err).
			Str("collection", EventsCollection).
			Interface("event", data).
			Msg("Erreur à l'insertion de l'évènement Tautulli")
		return
	}
}
