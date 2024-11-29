package repository

import (
	"go-crud/config"
	"go-crud/models"
	"log"
)

func CreateUser(user models.User) (models.User, error) {
	sql := `INSERT INTO users(name,email) values ($1,$2) RETURNING id`
	var id int
	db := config.ConfigDatabase()
	err := db.QueryRow(sql, user.Name, user.Email).Scan(&id)
	if err != nil {
		log.Print("Error creating user: %v", err)
		return models.User{}, err
	}
	user.ID = id
	db.Close()
	return user, err
}

func GetUserByID(id int) (models.User, error) {
	var user models.User
	sql := `SELECT id, name, email FROM users WHERE id = $1`
	db := config.ConfigDatabase()
	err := db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Print("Error fetching user: %v", err)
		return user, err
	}
	return user, nil
}

func UpdateUser(id int, user models.User) (models.User, error) {
	sql := `UPDATE users SET name=$1,email=$2 where id=$3`
	db := config.ConfigDatabase()
	_, err := db.Exec(sql, user.Name, user.Email, id)
	if err != nil {
		log.Print("Error updating users: %v", err)
		return models.User{}, err
	}
	user.ID = id
	return user, nil
}

func DeleteUser(id int) error {
	sql := `DELETE FROM users WHERE id=$1`
	db := config.ConfigDatabase()
	_, err := db.Exec(sql, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
