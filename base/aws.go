package base

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

var S3 *s3.Client
var S3BucketName *string
var Cloudfront *cloudfront.Client
var CloudfrontDistributionId *string

func AwsInit() {
	// Awsの初期化
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if GetEnv("APP_ENV", "develop") == "develop" {
			return aws.Endpoint{
				URL: "http://minio:9000",
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	rand.Seed(time.Now().UnixNano())
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(GetEnv("AWS_REGION", "ap-northeast-1")),
		config.WithEndpointResolverWithOptions(resolver),
	)
	if err != nil {
		panic(err)
	}
	if GetEnv("AWS_ASSUME_ROLE", "") != "" { // assume roleが設定されている場合のみ実行
		sc := sts.NewFromConfig(cfg)
		if rs, err := sc.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
			RoleArn:         aws.String(GetEnv("AWS_ASSUME_ROLE", "")),
			RoleSessionName: aws.String("role-identifier-" + strconv.Itoa(10000+rand.Intn(25000))),
		}); err == nil {
			cfg, err = config.LoadDefaultConfig(
				context.TODO(),
				config.WithRegion(GetEnv("AWS_REGION", "ap-northeast-1")),
				config.WithCredentialsProvider(
					credentials.NewStaticCredentialsProvider(
						*rs.Credentials.AccessKeyId,
						*rs.Credentials.SecretAccessKey,
						*rs.Credentials.SessionToken,
					),
				),
			)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	// s3のバケット名の取得
	S3BucketName = aws.String(GetEnv("S3_BUCKET_NAME", "test-bucket"))
	if S3BucketName != nil && *S3BucketName != "" {
		// S3クライアントの初期化
		S3 = s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		})
	}

	// cloudfrontのDistributionIdの取得
	CloudfrontDistributionId = aws.String(GetEnv("CLOUDFRONT_DISTRIBUTION_ID", ""))
	if CloudfrontDistributionId != nil && *CloudfrontDistributionId != "" {
		// cloudfrontクライアントの初期化
		Cloudfront = cloudfront.NewFromConfig(cfg)
	}
}
