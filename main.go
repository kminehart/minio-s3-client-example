package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	// ctx := context.Background()

	// Configure to use MinIO Server
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("root", "rootroot", ""),
		Endpoint:         aws.String("http://minio:9000"),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		panic(err)
	}

	s3Client := s3.New(newSession)

	cparams := &s3.CreateBucketInput{
		Bucket: aws.String("test-bucket"), // Required
	}

	// Create a new bucket using the CreateBucket call.
	if _, err := s3Client.CreateBucket(cparams); err != nil {
		log.Println(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		key := uuid.New().String()
		log.Println("Storing new object", key)
		_, err := s3Client.PutObject(&s3.PutObjectInput{
			Body:   strings.NewReader("hello"),
			Bucket: aws.String("test-bucket"),
			Key:    aws.String(key),
		})

		if err != nil {
			log.Println(err)
		}
	}).Methods(http.MethodPost)

	if err := http.ListenAndServe("0.0.0.0:3000", r); err != nil {
		log.Println("Server exiting", err)
	}
}
