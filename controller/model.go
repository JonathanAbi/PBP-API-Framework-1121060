package controller

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	UserType int	`json:"usertype"`
	Email 	string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"error"`
	Status  int    `json:"status"`
	Data    string `json:"error_message"`
}
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
type ResponseData struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}