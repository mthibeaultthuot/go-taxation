package data

type Tax struct {
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
