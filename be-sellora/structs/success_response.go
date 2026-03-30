package structs

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Meta    any    `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Succcess bool              `json:"success"`
	Message  string            `json:"message"`
	Errors   map[string]string `json:"errors"`
}
