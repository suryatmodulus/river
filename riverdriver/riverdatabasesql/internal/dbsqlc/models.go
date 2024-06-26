// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dbsqlc

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type JobState string

const (
	RiverJobStateAvailable JobState = "available"
	RiverJobStateCancelled JobState = "cancelled"
	RiverJobStateCompleted JobState = "completed"
	RiverJobStateDiscarded JobState = "discarded"
	RiverJobStatePending   JobState = "pending"
	RiverJobStateRetryable JobState = "retryable"
	RiverJobStateRunning   JobState = "running"
	RiverJobStateScheduled JobState = "scheduled"
)

func (e *JobState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = JobState(s)
	case string:
		*e = JobState(s)
	default:
		return fmt.Errorf("unsupported scan type for JobState: %T", src)
	}
	return nil
}

type NullJobState struct {
	JobState JobState
	Valid    bool // Valid is true if JobState is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullJobState) Scan(value interface{}) error {
	if value == nil {
		ns.JobState, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.JobState.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullJobState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.JobState), nil
}

type RiverJob struct {
	ID          int64
	Args        []byte
	Attempt     int16
	AttemptedAt *time.Time
	AttemptedBy []string
	CreatedAt   time.Time
	Errors      []AttemptError
	FinalizedAt *time.Time
	Kind        string
	MaxAttempts int16
	Metadata    json.RawMessage
	Priority    int16
	Queue       string
	State       JobState
	ScheduledAt time.Time
	Tags        []string
}

type RiverLeader struct {
	ElectedAt time.Time
	ExpiresAt time.Time
	LeaderID  string
	Name      string
}

type RiverMigration struct {
	ID        int64
	CreatedAt time.Time
	Version   int64
}
