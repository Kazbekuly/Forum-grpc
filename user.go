package Forum

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"Name"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
