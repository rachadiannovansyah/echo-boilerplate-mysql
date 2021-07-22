package models

import (
	"learn-go-echo/database"
	"net/http"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Telepon string `json:"telepon"`
}

func GetAllUser() (Response, error) {
	var obj User
	var arrobj []User
	var res Response

	conn := database.CreateConnection()
	sqlStatement := "SELECT * FROM users"

	rows, err := conn.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Telepon)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
