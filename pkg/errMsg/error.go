package errMsg

type Error struct {
	Error       string `json:"error"`
	StatusCode  int    `json:"statusCode"`
	Description string `json:"description"`
}
