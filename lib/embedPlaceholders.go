package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/model"
)

func EmbedPlaceholders(title string, description string, name string, author string, token string, fileSize int64, fileName string, delay time.Duration) model.EmbedStruct {
	var placeholders [6]string

	placeholders[0] = "$user$"
	placeholders[1] = "$uploads$"
	placeholders[2] = "$delay$"
	placeholders[3] = "$size$"
	placeholders[4] = "$filename$"
	placeholders[5] = "$uid$"

	fmt.Println(placeholders[1])
	uploadCounter, _ := database.GetUploadCounter(token)
	userClaims := database.VerifyUser(token)
	
	var values [6]any
	values[0] = userClaims.Username
	values[1] = uploadCounter+1 // plus one counting this upload
	values[2] = delay
	values[3] = fileSize
	values[4] = fileName 
	values[5] = userClaims.Uid

	substituteValue := func(str string) string {
		newStr := str
		for i := 0; len(placeholders) > i; i++ {
			placeholder := placeholders[i]
			value := values[i]
			newStr = strings.ReplaceAll(newStr, placeholder, fmt.Sprintf("%v", value))
		}
		return newStr
	}

	substituteValue("e")
	

	newTitle := substituteValue(title)
	newDescription := substituteValue(description)
	newName := substituteValue(name)
	newAuthor := substituteValue(author)


	return model.EmbedStruct{Title: newTitle, Description: newDescription, Name: newName, Author: newAuthor}
}