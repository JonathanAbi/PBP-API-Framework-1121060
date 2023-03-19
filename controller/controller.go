package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	db := connect()
	defer db.Close()
	fmt.Println("Masuk Sini 1")

	queryy := "SELECT * FROM users"
	name := context.Request.URL.Query()["name"]
	age := context.Request.URL.Query()["age"]
	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name='" + name[0] + "'"
	}
	if age != nil {
		if name != nil {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " age='" + age[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		var response ErrorResponse
		response.Status = http.StatusInternalServerError
		response.Message = http.StatusText(http.StatusInternalServerError)
		response.Data = err.Error()

		context.IndentedJSON(response.Status, response)
		return
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password, &user.UserType); err != nil {
			log.Println(err)
			context.IndentedJSON(400, gin.H{"message": err.Error()})
			return
		} else {
			users = append(users, user)
		}
	}
	if len(users) == 0 {
		var response ResponseData
		response.Status = 200
		response.Message = "Data tidak ditemukan"
		response.Data = users
		context.IndentedJSON(response.Status, response)
	} else if len(users) < 5 {
		var response ResponseData
		response.Status = 200
		response.Message = "Success"
		response.Data = users
		context.IndentedJSON(response.Status, response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Error Array Size Not Correct"
		context.IndentedJSON(response.Status, response)
	}
	fmt.Println("Masuk Sini 2")
}

func InsertUser(context *gin.Context) {
	db := connect()
	defer db.Close()
	//Read from Request Body
	err := context.Request.ParseForm()
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	name := context.Request.Form.Get("name")
	age, _ := strconv.Atoi(context.Request.Form.Get("age"))
	address := context.Request.Form.Get("address")
	email := context.Request.Form.Get("email")
	password := context.Request.Form.Get("password")
	userType, _ := strconv.Atoi(context.Request.Form.Get("userType"))

	_, errQuery := db.Exec("INSERT INTO users(name, age, address, email, password, userType) values (?,?,?,?,?,?)",
		name,
		age,
		address,
		email,
		password,
		userType,
	)
	var user User
	user.Name = name
	user.Age = age
	user.Address = address
	user.Email = email
	user.Password = password
	user.UserType = userType

	var response ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = user
	} else {
		response.Status = 400
		response.Message = "Insert Failed!"
	}
	context.IndentedJSON(response.Status, response)
}

func UpdateUser(context *gin.Context){
	db := connect()
	defer db.Close()

	id := context.Param("id")

	err := context.Request.ParseForm()
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	name := context.Request.Form.Get("name")
	age, _ := strconv.Atoi(context.Request.Form.Get("age"))
	address := context.Request.Form.Get("address")
	email := context.Request.Form.Get("email")
	password := context.Request.Form.Get("password")
	userType, _ := strconv.Atoi(context.Request.Form.Get("userType"))

	_, errQuery := db.Exec("UPDATE users SET name = ?,age = ?, address = ? , email = ?, password = ?, userType = ? WHERE id = ?",
		name,
		age,
		address,
		email,
		password,
		userType,
		id,
	)
	var user User
	user.Name = name
	user.Age = age
	user.Address = address
	user.Email = email
	user.Password = password
	user.UserType = userType

	var response ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = user
	} else {
		response.Status = 400
		response.Message = "Update Failed!"
	}
	context.IndentedJSON(response.Status, response)
}

func DeleteUser(context *gin.Context){
	db := connect()
	defer db.Close()

	err := context.Request.ParseForm()
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	id := context.Param("id")

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?", id)

	var response Response
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Delete Failed"
	}
	context.IndentedJSON(response.Status, response)
}
