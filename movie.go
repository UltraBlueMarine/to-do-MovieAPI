package server_side

type Movie struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
	Done     bool   `json:"done"`
}
