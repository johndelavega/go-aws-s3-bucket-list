//
// original source:
// https://github.com/aws/aws-sdk-go/blob/master/example/service/s3/listObjects/listObjects.go
// github.com/aws/aws-sdk-go/example/service/s3/listObjects/listObjects.go
// Release v1.17.9 - f2bb620afb315242706942e6ca4c7d26ee5ed627
//
// go-aws-s3-bucket-list
//
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const _appVersion = "v0.0.3"

//
// go run . mybucket
//

/*
// bucket policy to allow AIM user to list bucket content
{
	"Sid": "*",
	"Effect": "Allow",
	"Principal": {
		"AWS": "arn:aws:iam::594...221:user/john"
	},
	"Action": "s3:ListBucket",
	"Resource": "arn:aws:s3:::mybucket"
},
*/

// Lists all objects in a bucket using pagination
//
// Usage:
// listObjects <bucket>
func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must specify a bucket")
		return
	}

	// // Either hard-code the region or use AWS_REGION.
	// // Region: aws.String("us-east-2"),
	// // Region: aws.String("us-west-2"), // Oregon

	// // credentials.NewEnvCredentials assumes two environment variables are
	// // present:
	// // 1. AWS_ACCESS_KEY_ID, and
	// // 2. AWS_SECRET_ACCESS_KEY.
	// // or custom code that uses
	// // AWS_SECURITY_CREDENTIALS_FILE = path to accessKeys.csv created by AWS console credential settings

	// c := &aws.Config{
	// 	Credentials: credentials.NewEnvCredentials(),
	// }
	// sess := session.Must(session.NewSession(c))

	c := &aws.Config{
		Region: aws.String("us-west-2"), // Oregon
	}

	// this will use:
	// ~/.aws/credentials (Linux)
	// %UserProfile%\.aws\credentials (Windows)
	sess := session.Must(session.NewSession(c))

	svc := s3.New(sess)

	i := 0
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &os.Args[1],
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", i)
		i++

		for _, obj := range p.Contents {
			fmt.Println("Object:", *obj.Key)
		}
		return true
	})
	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
}
