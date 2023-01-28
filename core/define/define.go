package define

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

// JwtKey secret key of JWT token
var JwtKey = []byte("cloud-disk-key")

// CodeLength mail code length
var CodeLength = 6

// CodeExpireTime expire time of email code
var CodeExpireTime = time.Second * 300

// BucketName bucket name of MinIO
var BucketName = "bucket-cloud-disk"

var (
	MinioKey      string
	EmailSender   string
	EmailPassword string
	RedisPassword string
	MinioId       string
)

func init() {
	initEnv()
	MinioId = os.Getenv("MINIO_ID")
	MinioKey = os.Getenv("MINIO_KEY")

	EmailSender = os.Getenv("EMAIL_SENDER")
	EmailPassword = os.Getenv("EMAIL_PASSWORD")

	RedisPassword = os.Getenv("REDIS_PASSWORD")
}
