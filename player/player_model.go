package player

import "but3/go-score/base"

type Tour string

const (
	ATP Tour = "ATP"
	WTA Tour = "WTA"
)

type Player struct {
	base.Entity `bson:"inline"`
	FirstName   string `bson:"firstName,string" json:firstName`
	LastName    string `bson:"lastName,string" json:lastName`
	Tour        Tour   `bson:"tour,string" json:tour`
	Country     string `bson:"country,string" json:country`
}
