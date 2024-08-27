package base

import "go.mongodb.org/mongo-driver/bson/primitive"

// type de base pour les entités
type Entity struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}

// accesseur à la propriété Id
func (entity Entity) GetId() primitive.ObjectID {
	return entity.Id
}

// contrainte pour la généricité
type EntityConstraint interface {
	GetId() primitive.ObjectID
}
