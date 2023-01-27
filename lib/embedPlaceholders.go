package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/model"
)

func EmbedPlaceholders(site_name string, site_name_url string, title string, description string, author string, author_url string, token string, fileSize string, fileName string, delay time.Duration, path string) model.EmbedStruct {
	var placeholders [7]string

	placeholders[0] = "$user$"
	placeholders[1] = "$uploads$"
	placeholders[2] = "$delay$"
	placeholders[3] = "$size$"
	placeholders[4] = "$filename$"
	placeholders[5] = "$uid$"
	placeholders[6] = "$path$"
	uploadCounter, _ := database.GetUploadCounter(token)
	userClaims := database.VerifyUser(token)

	var values [7]any
	values[0] = userClaims.Username
	values[1] = uploadCounter + 1 // plus one counting this upload
	values[2] = delay
	values[3] = fileSize
	values[4] = fileName
	values[5] = userClaims.Uid
	values[6] = path
	substituteValue := func(str string) string {
		newStr := str
		for i := 0; len(placeholders) > i; i++ {
			placeholder := placeholders[i]
			value := values[i]
			newStr = strings.ReplaceAll(newStr, placeholder, fmt.Sprintf("%v", value))
		}
		return newStr
	}

	new_site_name := substituteValue(site_name)
	new_title := substituteValue(title)
	new_description := substituteValue(description)
	new_author := substituteValue(author)
	return model.EmbedStruct{Title: new_title, Description: new_description, Author: new_author, SiteName: new_site_name, SiteNameUrl: site_name_url, AuthorUrl: author_url}
}
