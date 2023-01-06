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

func VerifyUser(token string) model.UserValidation {
	db := DB
	result := map[string]any{}
	db.Model(&model.Users{}).Select("blacklist, admin, username, id").Where("token = ? OR api_key = ?", token, token).Find(&result)
	if len(result) == 0 {
		return model.UserValidation{Admin: false, ValidUser: false, Blacklisted: false, Username: "", Uid: -1}
	}
	admin, _ := strconv.ParseBool(fmt.Sprintf("%v", result["admin"]))
	blacklisted := result["blacklist"] != nil
	username := fmt.Sprintf("%s", result["username"])
	uid, _ := strconv.ParseInt(fmt.Sprintf("%v", result["id"]), 10, 64)
	return model.UserValidation{Admin: admin, Blacklisted: blacklisted, ValidUser: true, Username: username, Uid: int(uid)}
}

func Login(username string, password string) (string, error) {
	db := DB

	// result map
	result := map[string]any{}

	// we make a query, getting the password and id
	db.Model(&model.Users{}).Select("password, id, token, blacklist").Where("username = ? OR email = ?", username, username).Find(&result)

	// Em SQL: SELECT senha, id FROM users WHERE usuario=usuario
	// we pass the hashedPass to string
	hashedPassword := fmt.Sprintf("%v", result["password"])

	// we verify if the password is correct
	err := argonpass.Verify(password, hashedPassword)

	if err != nil {
		return "", errors.New("invalid credentials")
	}
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
		return errors.New("an error occurred. Please contact an admin")
	}

	// update the invite row
	if invite {
		db.Model(&model.Invites{}).Where("code = ?", inviteCode).Updates(model.Invites{UsedByID: int(user.ID)})
	}

	embed := model.Embeds{UserID: int(user.ID)}
	db.Create(&embed)
	return nil
}

func GetProfile(user string) (string, error) {
	db := DB

	// result map
	result := map[string]any{}

	// we make a query, getting the password and id
	db.Model(model.Users{}).Where("username = ?", user).Scan(&result)
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
		return errors.New("the requested user doesn't exist")
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
		return errors.New("the requested user doesn't exist")
	}
	
	return nil
}

func Upload(author string, name string, description string, title string, enabled bool, userid int, object string, color string, uploadId string, fileName string, file_url string, upload_url string, token string) error {
	db := DB
	
	userClaims := VerifyUser(token)

	if !userClaims.ValidUser {
		return errors.New("unauthorized")
	}
	upload := model.Uploads{EmbedEnabled: enabled, Author: author, Name: name, Description: description, UserID: userid, Object: object, Color: color, UploadID: uploadId, FileName: fileName, UploadUrl: upload_url, FileUrl: file_url}
	query := db.Create(&upload)
	
	if query.Error != nil {
		return errors.New("something wrong happened")
	}
	return nil
}

func GetUpload(uploadId string) (map[string]any, error) {
	db := DB
	result := map[string]any{}
	db.Model(&model.Uploads{}).Select("file_url", "created_at", "color", "upload_id", "object", "user_id", "title", "description", "author", "embed_enabled", "file_name").Where("upload_id = ?", uploadId).Find(&result)
	if len(result) <= 0 {
		return nil, errors.New("the upload doesn't exist")
	}
	return result, nil
}

func GetEmbed(token string) (map[string]any, error) {
	db := DB
	userClaims := VerifyUser(token)

	if !userClaims.ValidUser {
		return nil, errors.New("invalid token")
	}

	result := map[string]any{}
	db.Model(&model.Embeds{}).Where("user_id = ?", userClaims.Uid).Find(&result)

	if len(result) == 0 {
		return nil, errors.New("invalid authorization token")
	}

	return result, nil
}

func UpdateEmbed(token string, author string, description string, title string, name string) error {
	db := DB
	userClaims := VerifyUser(token)

	if !userClaims.ValidUser {
		return errors.New("unauthorized")
	}

	uid := userClaims.Uid
	
	db.Model(&model.Embeds{}).Where("user_id = ?", uid).Updates(model.Embeds{Title: title, Author: author, Description: description, Name: name})
	return nil
}

func UpdateDomain(token string, domain string) error {
	db := DB
	userClaims := VerifyUser(token)

	if !userClaims.ValidUser {
		return errors.New("unauthorized")
	}

	uid := userClaims.Uid
	
	domains := map[string]any{}
	db.Model(&model.Domains{}).Where("domain = ?", domain).Find(&domains)

	validDomain := len(domains) > 0

	if !validDomain {
		return errors.New("invalid domain")
	}
	
	query := db.Model(&model.Embeds{}).Where("user_id = ?", uid).Update("domain", domain)

	if query.Error != nil {
		return errors.New("something unexpected happened; please contact an admin")
	}

	return nil
}

func GetDomains(token string) ([]map[string]any, error) {
	db := DB
	userClaims := VerifyUser(token)

	if !userClaims.ValidUser {
		return nil, errors.New("unauthorized")
	}

	result := []map[string]any{} 
	db.Raw("SELECT domain, wildcard, username FROM domains INNER JOIN users ON domains.by_uid = users.id").Find(&result)

	fmt.Println(result)
	return result, nil
}

