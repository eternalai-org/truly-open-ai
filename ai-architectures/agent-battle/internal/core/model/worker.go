package model

import "time"

type WorkerStatus int64

const (
	WorkerStatusDisable WorkerStatus = iota + 1
	WorkerStatusEnable
)

type Worker struct {
	Model    `bson:",inline" json:",inline"`
	Name     string       `json:"name" bson:"name"`
	Code     string       `json:"code" bson:"code"`
	Status   WorkerStatus `json:"status" bson:"status"`
	Crontab  string       `json:"crontab" bson:"crontab"`
	State    *State       `json:"state" bson:"state"`
	Interval int          `json:"interval" bson:"interval"`
}

type State struct {
	LastRunDatetime time.Time `json:"last_run_datetime" bson:"last_run_datetime"`
	IsRunning       bool      `json:"is_running" bson:"is_running"`
}

func (m Worker) CollectionName() string {
	return "worker"
}
