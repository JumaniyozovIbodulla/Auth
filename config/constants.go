package config

const (
	ERR_INFORMATION     = "The server has received the request and is continuing the process"
	SUCCESS             = "The request was successful"
	ERR_REDIRECTION     = "You have been redirected and the completion of the request requires further action"
	ERR_BADREQUEST      = "Bad request"
	ERR_INTERNAL_SERVER = "While the request appears to be valid, the server could not complete the request"
	SmtpServer          = "smtp.gmail.com"
	SmtpPort            = "587"
	SmtpUsername        = "sunuz@gmail.com"
	SmtpPassword        = "zwrp mhjs vvuu pzka"
	MaxFileSize         = 5 * 1024 * 1024 // 5 MB
)

var SignedKey = []byte(`fbbzNqAIdOtJRueSXJtdJFjqmJWZVd`) // length 30
