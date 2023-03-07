package data

type Tax struct {
	Username   string `bson:"username,omitempty"`
	IsPstValid bool
	IsQstValid bool
	PstNumber  string
	QstNumber  string
	Enterprise string
	Date       string
}

type TaxBody struct {
	Pst string `json:"pst"`
	Qst string `json:"qst"`
}

type User struct {
	Username        string
	IsJwtValid      bool
	PermissionLevel int32
}
