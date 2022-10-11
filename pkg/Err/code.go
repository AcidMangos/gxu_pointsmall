package Err

const (
	SUCCESS                 = 200
	UPDATE_PASSWORD_SUCCESS = 201
	NOT_EXIST_IDENTIFIER    = 202
	ERROR                   = 500
	INVALID_PARAMS          = 400

	ERROR_EXIST_NICK           = 10001
	ERROR_EXIST_USER           = 10002
	ERROR_NOT_EXIST_USER       = 10003
	ERROR_NOT_COMPARE          = 10004
	ERROR_NOT_COMPARE_PASSWORD = 10005
	ERROR_FAIL_ENCRYPTION      = 10006
	ERROR_NOT_EXIST_PRODUCT    = 10007
	ERROR_NOT_EXIST_ADDRESS    = 10008
	ERROR_EXIST_FAVORITE       = 10009

	ERROR_AUTH_CHECK_TOKEN_FAIL       = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT    = 20002
	ERROR_AUTH_TOKEN                  = 20003
	ERROR_AUTH                        = 20004
	ERROR_AUTH_INSUFFICIENT_AUTHORITY = 20005
	ERROR_READ_FILE                   = 20006
	ERROR_SEND_EMAIL                  = 20007
	ERROR_CALL_API                    = 20008
	ERROR_UNMARSHAL_JSON              = 20009

	ERROR_DATABASE = 30001

	ERROR_OSS = 40001

	Error_open_file = 50001
)
