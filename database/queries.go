package database

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	argonpass "github.com/dwin/goArgonPass"
	"github.com/floppahost/backend/jwt"
	"github.com/floppahost/backend/model"
	"github.com/google/uuid"
)

func Login(database any, usuario string, senha string) (string, error) {
	db := DB

	// result map
	result := map[string]any{}

	// we make a query, getting the password and id
	db.Model(database).Select("password, id, token").Where("username = ?", usuario).Find(&result)
	// Em SQL: SELECT senha, id FROM users WHERE usuario=usuario

	// we pass the hashedPass to string
	hashedPassword := fmt.Sprintf("%v", result["senha"])

	// we verify if the password is correct
	err := argonpass.Verify(senha, hashedPassword)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return fmt.Sprintf("%s", result["token"]), nil
}

func Register(username string, password string, email string, inviteCode string) error {
	db := DB

	result := map[string]interface{}{}

	invite, _ := strconv.ParseBool(os.Getenv("INVITE_ONLY"))

	if invite {
		db.Model(&model.Invites{}).Where("code = ? AND used_by IS NULL AND used_by_id IS NULL", inviteCode).Find(&result)
		// verify if the invite is valid
		if len(result) == 0 {
			return errors.New("invalid invite")
		}
	}

	// hash the password and verify the status
	hashedPassword, err := argonpass.Hash(password, nil)
	if err != nil {
		return err
	}

	// declare our user model and do the create operation

	user := model.Users{}
	token, err := jwt.GenerateUserToken("system", username)

	if err != nil {
		return err
	}
	user = model.Users{Username: username, Email: email, Password: hashedPassword, ApiKey: uuid.NewString(), Token: token}

	create := db.Create(&user)

	// verify if we created the user
	if create.Error != nil {
		return create.Error
	}

	// update the invite row
	if invite {
		db.Model(&model.Invites{}).Where("code = ?", inviteCode).Updates(model.Invites{UsedBy: username, UsedByID: int(user.ID)})
	}

	return nil
}

func GetProfile(user string) (string, error) {
	db := DB

	// result map
	result := map[string]any{}

	// we make a query, getting the password and id
	db.Model(model.Users{}).Where("user = ?", string(user)).First(&result)
	id := result["id"]

	if id == nil {
		return "", errors.New("this user doesn't exist")
	}

	return fmt.Sprintf("%v", id), nil
}

func InviteWave(jwt string) error {
	db := DB

	// result := map[string]any{}
	// db.Model(model.Users{}).Where("token = ? AND admin = ?", jwt, true).Find(&result)

	// if result == nil {
	// 	return errors.New("you don't have permission to perform that")
	// }

	var result any
	db.Raw("INSERT INTO invites (user_id) (SELECT id FROM users)").Scan(&result)

	if result != nil {
		return errors.New("something unexpected happened")
	}
	return nil
}
