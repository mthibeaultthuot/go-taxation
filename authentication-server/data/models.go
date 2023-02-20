package data

type User struct {
	Username   string `bson:"username,omitempty"`
	Password   string `bson:"password,omitempty"`
	Permission int    `bson:"permission,omitempty"`
}
