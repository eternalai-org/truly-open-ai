package cast

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectsToStrings(ids []primitive.ObjectID) (result []string) {
	for _, v := range ids {
		result = append(result, v.String())
	}
	return result
}

func StringsToObjects(ids []string) (result []primitive.ObjectID, err error) {
	for _, v := range ids {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return nil, errors.WithMessage(err, "StringsToObject parse id error")
		}
		result = append(result, id)
	}
	return result, nil
}

func ObjectsToHex(ids []primitive.ObjectID) (result []string) {
	for _, v := range ids {
		result = append(result, v.Hex())
	}
	return result
}
