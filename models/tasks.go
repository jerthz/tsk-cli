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

func (s TaskStatus) String() string {
    switch s {
    case Pending:
        return "\033[34mPending\033[0m" // Jaune
    case InProgress:
        return "\033[33mIn Progress\033[0m" // Bleu
    case Completed:
        return "\033[32mCompleted\033[0m" // Vert
    case Stashed:
        return "\033[35mStashed\033[0m" // Violet
    default:
        return "\033[31mUnknown\033[0m" // Rouge
    }
}
