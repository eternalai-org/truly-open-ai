package utils

import (
	"errors"

	mongo2 "go.mongodb.org/mongo-driver/mongo"
)

func IsErrNoDocuments(err error) bool {
	return errors.Is(err, mongo2.ErrNoDocuments)
}
