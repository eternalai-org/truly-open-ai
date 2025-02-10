package sort

import "strings"

type Type string

const (
	ASC  Type = "asc"
	DESC Type = "desc"
)

func (s Type) Ordinal() string {
	return string(s)
}

type Sort struct {
	Field string `json:"field"`
	Type  Type   `json:"type"`
}

func (s Sort) GetSortType() int64 {
	if strings.EqualFold(s.Type.Ordinal(), ASC.Ordinal()) {
		return 1
	}
	return -1
}
