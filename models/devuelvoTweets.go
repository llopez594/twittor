package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*DevueltoTweets es la estructura con la que devolveremos los Tweets*/
type DevueltoTweets struct {
	ID 		primitive.ObjectID 	`bson:"_id" json:"_id,omitempty"`
	UserID 	string 				`bson:"userid" json:"userId,omitempty"`
	Mensaje string 				`bson:"mensaje" json:"mensaje,omitempty"`
	Fecha 	time.Time 			`bson:"fecha" json:"fecha,omitempty"`
}
