package util

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongoDB() *mongo.Database {
	dbUrl := fmt.Sprintf("mongodb://%s:%d", Config.Database.Host, Config.Database.Port)
	credential := options.Credential{Username: Config.Database.User, Password: Config.Database.Pwd, AuthSource: Config.Database.Name}
	client, err := mongo.Connect(context.TODO(), options.Client().SetAuth(credential).ApplyURI(dbUrl))
	if err != nil {
		log.Error().Stack().Err(err).Str("dbUrl", dbUrl).Msg("Erreur de connexion à la base de données MongoDB")
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Error().Stack().Err(err).Str("dbUrl", dbUrl).Msg("Erreur au ping la base de données MongoDB")
		panic(err)
	}
	log.Info().Str("dbUrl", dbUrl).Msg("Connexion MongoDB OK")
	return client.Database(Config.Database.Name)
}
