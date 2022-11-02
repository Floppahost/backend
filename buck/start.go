package buck

import (
	"log"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Bucket *minio.Client

func Start() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINION_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINION_SECRET_ACCESS_KEY")
	useSSL, _ := strconv.ParseBool(os.Getenv("MINIO_SSL"))

	var err error
	Bucket, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

}
