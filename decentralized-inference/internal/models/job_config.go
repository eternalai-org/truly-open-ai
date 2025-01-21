package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type JobConfig struct {
	mgm.DefaultModel `bson:",inline"`
	JobName          string    `json:"job_name" bson:"job_name"`
	Enable           bool      `json:"enable" bson:"enable"`
	Interval         int64     `json:"interval" bson:"interval"`
	LastRun          time.Time `json:"last_run" bson:"last_run"`
	IntervalUnit     string    `json:"interval_unit" bson:"interval_unit"`
}

func (JobConfig) CollectionName() string {
	return "job_config"
}
