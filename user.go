package TODOList_REST

type User struct {
	Id       int    `json:"_"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
