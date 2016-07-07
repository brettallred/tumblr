package tumblr

type Activity struct {
	Id     string `json:"id"`
	Action string `json:"action"`
	Blog   Blog   `json:"blog"`
}
