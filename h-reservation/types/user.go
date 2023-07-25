package types

type User struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"` // Mapping struct with MongoDB(bson) and api response type (json)
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"-"`
}
