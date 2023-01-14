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
	endpoint := os.Getenv("BUCKET_ENDPOINT")
	accessKeyID := os.Getenv("BUCKET_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("BUCKET_SECRET_ACCESS_KEY")
	region := os.Getenv("BUCKET_REGION")
	useSSL, _ := strconv.ParseBool(os.Getenv("BUCKET_SSL"))

	var err error
	Bucket, err = minio.New(endpoint, &minio.Options{
		Region: region,
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

}
