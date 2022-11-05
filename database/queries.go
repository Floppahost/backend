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

func VerifyUser(jwt string) model.UserValidation {
	db := DB
	result := map[string]any{}
	db.Model(&model.Users{}).Select("blacklist, admin").Where("token = ?", jwt).Find(&result)
	if result == nil {
		return model.UserValidation{Admin: false, ValidUser: false, Blacklisted: false}
	}
	admin, _ := strconv.ParseBool(fmt.Sprintf("%v", result["admin"]))
	blacklisted := result["blacklist"] != nil
	return model.UserValidation{Admin: admin, Blacklisted: blacklisted, ValidUser: true}
}

func Login(username string, password string) (string, error) {
	db := DB

	// result map
	result := map[string]any{}

	// we make a query, getting the password and id
	db.Model(&model.Users{}).Select("password, id, token, blacklist").Where("username = ?", username).Find(&result)

	// Em SQL: SELECT senha, id FROM users WHERE usuario=usuario
	// we pass the hashedPass to string
	hashedPassword := fmt.Sprintf("%v", result["password"])

	// we verify if the password is correct
	err := argonpass.Verify(password, hashedPassword)

	fmt.Println(hashedPassword)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	fmt.Println(result)
	if result["blacklist"] != nil {
		return "", errors.New("the user is currently blacklisted")
	}
	return fmt.Sprintf("%s", result["token"]), nil
}

func Register(username string, password string, email string, inviteCode string) error {
	db := DB

	result := map[string]any{}

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

	token, err := jwt.GenerateUserToken("system", username)

	if err != nil {
		return err
	}

	user := model.Users{Username: username, Email: email, Password: hashedPassword, ApiKey: uuid.NewString(), Token: token}

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

	userClaims := VerifyUser(jwt)

	if !userClaims.ValidUser || !userClaims.Admin || userClaims.Blacklisted {
		return errors.New("you don't have permission to perform this action")
	}

	var result any
	db.Raw("INSERT INTO invites (user_id) (SELECT id FROM users)").Scan(&result)

	if result != nil {
		return errors.New("something unexpected happened")
	}
	return nil
}

func GenerateInvite(jwt string, username string) error {
	db := DB

	userClaims := VerifyUser(jwt)
	if !userClaims.ValidUser || !userClaims.Admin || userClaims.Blacklisted {
		return errors.New("you don't have permission to perform this action")
	}

	VerifyUser(jwt)
	result := map[string]any{}
	db.Model(&model.Users{}).Select("id").Where("username = ?", username).Find(&result)
	if len(result) == 0 {
		return errors.New("the requested user doesn't exist")
	}

	id, err := strconv.ParseInt(fmt.Sprintf("%v", result["id"]), 10, 64)
	if err != nil {
		return errors.New("something weird happened")
	}

	invite := model.Invites{UserID: int(id)}
	create := db.Create(&invite)

	if create.Error != nil {
		return errors.New("something weird happened")
	}

	return nil
}

func BlacklistUser(jwt string, username string, reason string) error {
	db := DB

	userClaims := VerifyUser(jwt)

	if !userClaims.ValidUser || !userClaims.Admin || userClaims.Blacklisted {
		return errors.New("you don't have permission to perform this action")
	}

	query := db.Model(&model.Users{}).Where("username = ?", username).Updates(model.Users{Blacklist: reason})
	
	if query.Error != nil {
		return errors.New("something weird happened")
	}
	
	return nil
}

func UnblacklistUser(jwt string, username string, reason string) error {
	db := DB

	userClaims := VerifyUser(jwt)

	if !userClaims.ValidUser || !userClaims.Admin || userClaims.Blacklisted {
		return errors.New("you don't have permission to perform this action")
	}

	query := db.Model(&model.Users{}).Where("username = ?", username).Updates(model.Users{Blacklist: ""})
	
	if query.Error != nil {
		return errors.New("something weird happened")
	}
	
	return nil
}