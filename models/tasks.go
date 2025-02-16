package models

import "time"

type TaskStatus int
const (
    Pending TaskStatus = iota
    InProgress
    Completed
    Stashed
)

type Task struct {
    Id          int           `json:"id"`
    Description string        `json:"description"`
    CreatedAt   time.Time     `json:"created_at"`
    Status      TaskStatus    `json:"status"`
    StartedAt   time.Time     `json:"startedAt"`
    Comment     string        `json:"comment"`
    Category    string        `json:"category"`
}
