package mysql

type Action int

const (
	WriteOrRead = Action(1)
	ReadOnly = Action(2)
)
