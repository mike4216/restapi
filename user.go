package TODOList_REST

type User struct {
	Id       int    `json:"_"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
