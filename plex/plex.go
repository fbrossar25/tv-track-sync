package plex

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"strings"
)

func InitPlex(router *gin.Engine) {
	router.POST("plex/webhook", webhook)
}

// mostly taken from github.com/jrudio/go-plex-client

type WebhookData struct {
	Event   string `json:"event"`
	User    bool   `json:"user"`
	Owner   bool   `json:"owner"`
	Account struct {
		ID    int    `json:"id"`
		Thumb string `json:"thumb"`
		Title string `json:"title"`
	} `json:"Account"`
	Server struct {
		Title string `json:"title"`
		UUID  string `json:"uuid"`
	} `json:"Server"`
	Player struct {
		Local         bool   `json:"local"`
		PublicAddress string `json:"PublicAddress"`
		Title         string `json:"title"`
		UUID          string `json:"uuid"`
	} `json:"Player"`
	Metadata struct {
		LibrarySectionType   string `json:"librarySectionType"`
		RatingKey            string `json:"ratingKey"`
		Key                  string `json:"key"`
		ParentRatingKey      string `json:"parentRatingKey"`
		GrandparentRatingKey string `json:"grandparentRatingKey"`
		GUID                 []struct {
			Id string `json:"id"`
		} `json:"Guid"`
		LibrarySectionID int    `json:"librarySectionID"`
		MediaType        string `json:"type"`
		Title            string `json:"title"`
		GrandparentKey   string `json:"grandparentKey"`
		ParentKey        string `json:"parentKey"`
		GrandparentTitle string `json:"grandparentTitle"`
		ParentTitle      string `json:"parentTitle"`
		Summary          string `json:"summary"`
		Index            int    `json:"index"`
		ParentIndex      int    `json:"parentIndex"`
		RatingCount      int    `json:"ratingCount"`
		Thumb            string `json:"thumb"`
		Art              string `json:"art"`
		ParentThumb      string `json:"parentThumb"`
		GrandparentThumb string `json:"grandparentThumb"`
		GrandparentArt   string `json:"grandparentArt"`
		AddedAt          int    `json:"addedAt"`
		UpdatedAt        int    `json:"updatedAt"`
	} `json:"Metadata"`
}

func webhook(context *gin.Context) {
	sublog := log.With().Str("method", "plex/webhook").Logger()
	var data WebhookData
	form, err := context.MultipartForm()
	if err != nil {
		b, payloadErr := io.ReadAll(context.Request.Body)
		if payloadErr != nil {
			sublog.Error().Err(err).Str("payloadReadError", payloadErr.Error()).Msg("Lecture payload impossible")
		} else {
			sublog.Error().Err(err).Str("payload", string(b)).Send()
		}
		return
	}
	payload, hasPayload := form.Value["payload"]
	if hasPayload {
		if !strings.Contains(payload[0], "\"media.scrobble\"") {
			// not a scrobble event
			return
		}
		fmt.Printf("payload:\n%s\n", payload[0])
		if err := json.Unmarshal([]byte(payload[0]), &data); err != nil {
			log.Error().Stack().Err(err).Msg(fmt.Sprintf("can not parse json"))
			return
		}
		log.Debug().Interface("guid", data.Metadata.GUID).Msg("parsed")
		handleEvent(&data)
	}
}

func handleEvent(data *WebhookData) {
	if data.Event != "media.scrobble" {
		return
	}
	if data.Metadata.MediaType == "episode" {
		log.Info().Str("user", data.Account.Title).Str("title", fmt.Sprintf("%s S%dE%d", data.Metadata.GrandparentTitle, data.Metadata.ParentIndex, data.Metadata.Index))
	} else if data.Metadata.LibrarySectionType == "movie" {

	} else {
		return
	}
}
