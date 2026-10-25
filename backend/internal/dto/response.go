package dto

type SuccessResponse struct {
	Data any `json:"data"`
}

type FailResponse struct {
	Error FailDetails `json:"error"`
}

type FailDetails struct {
	Code    string `json:"code"`
	Details *any
}
