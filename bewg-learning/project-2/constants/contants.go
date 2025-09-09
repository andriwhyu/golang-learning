package constants

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type (
	postKey string
	userKey string
)

const (
	MaxBytesRequest         = 1_048_578 // 1Mb = 1048578
	PostCtx         postKey = "post"
	UserCtx         userKey = "user"
	QueryTimeout            = time.Second * 5
)

var (
	ErrDataNotFoundByID = errors.New("no data found for given id")
	ErrConflict         = errors.New("conflict mismatch version")
)

var Validator *validator.Validate

func init() {
	Validator = validator.New(validator.WithRequiredStructEnabled())
}
