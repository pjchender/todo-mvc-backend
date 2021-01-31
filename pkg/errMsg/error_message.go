package errMsg

var (
	// Database
	RecordNotFound      = "record not found"
	RecordAlreadyExists = "record already exists"
	GetRecordFail       = "fail to get the record"
	GetRecordListFail   = "fail to get records"
	CreateRecordFail    = "fail to create record"
	UpdateRecordFail    = "fail to update record"
	DeleteRecordFail    = "fail to delete record"

	// Server
	ServerError         = "something went wrong! Please try again later"
	InternalServerError = "internal server error"

	// Auth
	GenerateJWTFailed     = "failed to generate JWT token"
	LoginValidationFailed = "incorrect username or password"
	InvalidToken          = "token is invalid"
	ExpiredToken          = "token is expired"
	TokenNotAllowed       = "token is not allowed to access this api"

	// Other
	InvalidJsonSyntax = "invalid JSON syntax"
)
