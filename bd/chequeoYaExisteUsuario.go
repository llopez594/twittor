package bd

import (
	"context"
	"github.com/llopez594/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/*ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya esta en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email":email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
