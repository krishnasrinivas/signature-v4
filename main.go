package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/minio/s3v2tov4-proxy/s3auth"
)

func main() {
	accessKey := flag.String("access", "", "S3 access key")
	secretKey := flag.String("secret", "", "S3 secret key")
	region := flag.String("region", "us-east-1", "S3 bucket region")
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("incorrect usage")
	}
	auth := s3auth.CredentialsV4{
		AccessKey: *accessKey,
		SecretKey: *secretKey,
		Region:    *region,
	}
	req := struct {
		Method string      `json:"method"`
		Path   string      `json:"path"`
		Query  string      `json:"query"`
		Header http.Header `json:"header"`
	}{}
	err := json.Unmarshal([]byte(flag.Args()[0]), &req)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Del("authorization")
	fmt.Println(auth.Sign(req.Method, req.Path, req.Query, req.Header))
}
