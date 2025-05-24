package model

type Task struct {
	Title    string `json:"title"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
}
