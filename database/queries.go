package database

import (
	"errors"
	"fmt"
	"strconv"

	argonpass "github.com/dwin/goArgonPass"
	"github.com/floppahost/backend/lib"
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

func Register(username string, password string, email string, inviteCode string) error {
	db := DB

	result := map[string]interface{}{}
	db.Model(&model.Invites{}).Where("code = ?", inviteCode).Find(&result)

	// verify if the invite is valid
	if result == nil {
		return errors.New("invalid invite")
	}

	// hash the password and verify the status
	hashedPassword, err := argonpass.Hash(password, nil)
	if err != nil {
		return err
	}

	// declare our user model and do the create operation
	user := model.Users{User: username, Email: email, Password: hashedPassword, Token: lib.Token()}
	create := db.Create(&user)

	// verify if we created the user
	if create.Error != nil {
		return create.Error
	}

	// update the invite row
	db.Model(model.Invites{}).Where("code = ?", inviteCode).Updates(model.Invites{UsedBy: username, UsedByID: int(user.ID)})

	return nil
}

func GetProfile(user string) (string, error) {
	db := DB

	// result map
	result := map[string]any{}

	// we make a query, getting the password and id
	db.Model(model.Users{}).Where("user = ?", string(user)).First(&result)
	fmt.Println(result)
	id := result["id"]

	if id == nil {
		return "", errors.New("this user doesn't exist")
	}

	return fmt.Sprintf("%v", id), nil
}
