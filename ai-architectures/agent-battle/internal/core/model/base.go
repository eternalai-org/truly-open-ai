package model

import (
	"strings"
	"time"

	"agent-battle/pkg/constants/sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Id struct {
	Id string `json:"id" bson:"id" query:"id"`
}
type Ids struct {
	Ids []string `json:"ids" bson:"ids" query:"ids"`
}

type Codes struct {
	Codes []string `json:"codes" bson:"codes" query:"codes"`
}

type Code struct {
	Code string `json:"code" bson:"code" query:"code"`
}

type IdResponse struct {
	Id string `json:"id"`
}

type UrlResponse struct {
	Url string `json:"url"`
}

type IdsResponse struct {
	Ids []string `json:"ids"`
}

type Sort struct {
	Field string    `json:"field"`
	Type  sort.Type `json:"type"`
}

func GetSort(sorts []string) []*Sort {
	if len(sorts) > 0 {
		sortMap := make(map[string]string, 0)
		for _, sort := range sorts {
			arr := strings.Split(sort, ",")
			if len(arr) <= 1 {
				continue
			}
			sortMap[arr[0]] = arr[1]
		}
		resp := make([]*Sort, 0)
		for key, val := range sortMap {
			resp = append(resp, &Sort{
				Field: key,
				Type:  sort.Type(strings.ToLower(val)),
			})
		}
		return resp
	}
	return nil
}

type Model struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`

	DateCreated    time.Time `json:"date_created" bson:"date_created"`
	DateModified   time.Time `json:"date_modified" bson:"date_modified"`
	ModifiedUserId string    `json:"modified_user_id" bson:"modified_user_id"`
	CreatedUserId  string    `json:"created_user_id" bson:"created_user_id"`
	DateDeleted    time.Time `json:"date_deleted" bson:"date_deleted"`
}

func (a *Model) Init() {
	a.Id = primitive.NewObjectID()
	a.DateCreated = time.Now()
	a.DateModified = a.DateCreated
}

type Pagination struct {
	LastId *string  `json:"last_id" query:"last_id"`
	Limit  *int64   `json:"limit" query:"limit"`
	Page   *int64   `json:"page" query:"page"`
	Sort   []string `json:"sort,omitempty" query:"sort"`
	Sorts  []*Sort  `json:"-"`
}

func (p *Pagination) BuildPipeline() (bson.M, bson.M) {
	var (
		limit = int64(30)
		page  = int64(1)
	)

	if p.Limit != nil && *p.Limit > 0 {
		limit = *p.Limit
	}

	limitFilter := bson.M{"$limit": limit}

	if p.Page != nil && *p.Page > 0 {
		page = *p.Page
	}

	skipPipeline := bson.M{"$skip": limit * (page - 1)}
	return skipPipeline, limitFilter
}
