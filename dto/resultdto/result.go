package dto

type SuccessResult struct {
	Status       string      `json:"status"`
	Data         interface{} `json:"data"`
	Notification interface{} `json:"notification"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
