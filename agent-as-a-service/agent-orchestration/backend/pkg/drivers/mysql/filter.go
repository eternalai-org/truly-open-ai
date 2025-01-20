package mysql

import (
	"fmt"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
	"gorm.io/gorm"
)

type (
	FilterItem struct {
		_         struct{}
		FieldName string
		Operator  FilterOperator
		Value     interface{}

		// these fields are used for this statement: db.Where(a.RawQuery, a.Args...)
		RawQuery string
		Args     []interface{}
	}

	FilterOrderBy struct{}

	PreloadItem struct {
		_      struct{}
		Column string
		Args   []interface{}
	}

	// Example
	// func(db *gorm.DB) *gorm.DB {
	//	  return db.Where("info_type = 0") // schema_markup
	// }
	PreloadArgsFunc func(db *gorm.DB) *gorm.DB

	FilterOperator string

	Filters  []FilterItem
	Preloads []PreloadItem
	Limit    int
	OrderBy  string
)

const (
	EqualOperator           = FilterOperator("=")
	LikeOperator            = FilterOperator("like")
	InOperator              = FilterOperator("in")
	LessThanOperator        = FilterOperator("<")
	GreaterThanOperator     = FilterOperator(">")
	NotInOperator           = FilterOperator("not in")
	LessThanOrEqualOperator = FilterOperator("<=")

	// zero value
	ZeroLimit   = Limit(-1)
	ZeroOrderBy = OrderBy("")

	IDFieldName = "id"
)

func NewLimit(i interface{}) Limit {
	if i == nil {
		return ZeroLimit
	}

	float64Value, ok := i.(float64)
	if !ok {
		return ZeroLimit
	}

	return Limit(int(float64Value))
}

func (a *FilterItem) makeWhereQuery() string {
	switch a.Operator {
	case InOperator, NotInOperator:
		return fmt.Sprintf("%s %s (?)", a.FieldName, a.Operator)
	}

	return fmt.Sprintf("%s %s ?", a.FieldName, a.Operator)
}

func (a FilterItem) setCondition(db *gorm.DB) *gorm.DB {
	if utils.IsStringNotEmpty(a.RawQuery) {
		return db.Where(a.RawQuery, a.Args...)
	}

	return db.Where(a.makeWhereQuery(), a.Value)
}

func (a Filters) setCondition(db *gorm.DB) *gorm.DB {
	newDB := db
	for _, item := range a {
		newDB = item.setCondition(newDB)
	}
	return newDB
}

func (a Preloads) setCondition(db *gorm.DB) *gorm.DB {
	newDB := db
	for _, item := range a {
		newDB = item.setCondition(newDB)
	}

	return newDB
}

func (a PreloadItem) setCondition(db *gorm.DB) *gorm.DB {
	if a.Args == nil {
		return db.Preload(a.Column)
	}

	return db.Preload(a.Column, a.Args...)
}

func (a Limit) SetCondition(db *gorm.DB) *gorm.DB {
	newDB := db.Limit(a.toInt())
	return newDB
}

func (a Limit) toInt() int {
	return int(a)
}

func (a Limit) IsZero() bool {
	return a == ZeroLimit
}

func (a OrderBy) SetCondition(db *gorm.DB) *gorm.DB {
	newDB := db.Order(string(a))
	return newDB
}

func (a OrderBy) IsZero() bool {
	return a == ZeroOrderBy
}

func GenPreloadItemHandlerFunc(handler func(db *gorm.DB) *gorm.DB) []interface{} {
	return []interface{}{handler}
}

func GenPreloadItemArgs(args ...interface{}) []interface{} {
	return args
}
