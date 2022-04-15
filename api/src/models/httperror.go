package models

type HttpError struct {
	Timestamp string `json:"timestamp"`
	Error     string `json:"error"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}
