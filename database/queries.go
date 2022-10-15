package database

import (
	"errors"
	"fmt"
	"strconv"

	argonpass "github.com/dwin/goArgonPass"
	"github.com/floppahost/backend/model"
)

func Login(database any, usuario string, senha string) (uint, error) {
	db := DB

	// result map
	result := map[string]any{}

	// we make a query, getting the password and id
	db.Model(database).Select("senha, id").Where("usuario = ?", usuario).Find(&result)
	// Em SQL: SELECT senha, id FROM users WHERE usuario=usuario

	// we pass the hashedPass to strign
	hashedPassword := fmt.Sprintf("%v", result["senha"])

	// we pass the id to int
	id, _ := strconv.ParseUint(fmt.Sprintf("%v", result["id"]), 10, 64)

	// we verify if the password is correct
	err := argonpass.Verify(senha, hashedPassword)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

func GetProfile(user string) (string, error) {
	db := DB

	// result map
	result := map[string]any{}

	// we make a query, getting the password and id
	db.Model(model.Users{}).Where("user = ?", string(user)).First(&result)
	fmt.Println(result)
	id := result["id"]

	if (id == nil) {
		return "", errors.New("this user doesn't exist")
	}

	return fmt.Sprintf("%v",id), nil
}

