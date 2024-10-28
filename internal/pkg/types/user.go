package types

type User struct {
	ID        string `bson:"_id" json:"id"`
	FIRSTNAME string `bson:"firstName" json:"firstName"`
	LASTNAME  string `bson:"lastName" json:"lastName"`
}
