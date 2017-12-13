package dto

type ReponseDTO struct {
	Code uint16	`json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

const(
	RIGHT_CODE=10000
	ERROR_CODE=10001
)