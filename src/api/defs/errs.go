package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HTTPSc int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HTTPSc: 400,
		Error: Err{
			Error:     "request body is not correct",
			ErrorCode: "001",
		},
	}

	ErrorAuthUser = ErrorResponse{
		HTTPSc: 401,
		Error: Err{
			Error:     "user authentictation failed",
			ErrorCode: "002",
		},
	}
)
