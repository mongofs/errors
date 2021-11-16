package errors

// 通用: 基本错误
// Code must start with 1xxxxx
const (
	// ErrSuccess - 200: OK.
	ErrSuccess int = iota + 100001

	// ErrUnknown - 500: Internal server error.
	ErrUnknown

	// ErrBind - 400: Error occurred while binding the request body to the struct.
	ErrBind

	// ErrValidation - 400: Validation failed.
	ErrValidation

	// ErrTokenInvalid - 401: Token invalid.
	ErrTokenInvalid

	// ParamsErr
	ParamsErr
)

// 通用：数据库类错误
const (
	// ErrDatabase - 500: Database error.
	ErrDatabase int = iota + 100101
)

// 通用：认证授权类错误
const (
	// ErrEncrypt - 401: Error occurred while encrypting the user password.
	ErrEncrypt int = iota + 100201

	// ErrSignatureInvalid - 401: Signature is invalid.
	ErrSignatureInvalid

	// ErrExpired - 401: Token expired.
	ErrExpired

	// ErrInvalidAuthHeader - 401: Invalid authorization header.
	ErrInvalidAuthHeader

	// ErrMissingHeader - 401: The `Authorization` header was empty.
	ErrMissingHeader

	// ErrorExpired - 401: Token expired.
	ErrorExpired

	// ErrPasswordIncorrect - 401: Password was incorrect.
	ErrPasswordIncorrect

	// PermissionDenied - 403: Permission denied.
	ErrPermissionDenied
)

// 通用：编解码类错误
const (
	// ErrEncodingFailed - 500: Encoding failed due to an error with the data.
	ErrEncodingFailed int = iota + 100301

	// ErrDecodingFailed - 500: Decoding failed due to an error with the data.
	ErrDecodingFailed

	// ErrInvalidJSON - 500: Data is not valid JSON.
	ErrInvalidJSON

	// ErrEncodingJSON - 500: JSON data could not be encoded.
	ErrEncodingJSON

	// ErrDecodingJSON - 500: JSON data could not be decoded.
	ErrDecodingJSON

	// ErrInvalidYaml - 500: Data is not valid Yaml.
	ErrInvalidYaml

	// ErrEncodingYaml - 500: Yaml data could not be encoded.
	ErrEncodingYaml

	// ErrDecodingYaml - 500: Yaml data could not be decoded.
	ErrDecodingYaml

	// ErrEncodingProto
	ErrEncodingProto

	// ErrDecodingProto
	ErrDecodingProto
)

func init() {
	//base 类
	Register(ErrSuccess, 200, "errors_map:  OK.")
	Register(ErrUnknown, 500, "errors_map:  Internal server error.")
	Register(ErrBind, 400, "errors_map:  Error occurred while binding the request body to the struct.")
	Register(ErrValidation, 400, "errors_map:  Validation failed.")
	Register(ErrTokenInvalid, 500, "errors_map:  Internal server error.")
	Register(ParamsErr, 401, "params error")
	// database
	Register(ErrDatabase, 500, "errors_map:  database server error.")
	// 通用：认证授权类错误
	Register(ErrEncrypt, 401, "errors_map:  Error occurred while encrypting the user password.")
	Register(ErrSignatureInvalid, 401, "errors_map:  Signature is invalid.")
	Register(ErrExpired, 401, "errors_map:  Token expired.")
	Register(ErrInvalidAuthHeader, 401, "errors_map:  Invalid authorization header.")
	Register(ErrMissingHeader, 401, "errors_map:  The `Authorization` header was empty.")
	Register(ErrorExpired, 401, "errors_map:  Token expired.")
	Register(ErrPasswordIncorrect, 401, "errors_map:  Password was incorrect.")
	Register(ErrPermissionDenied, 403, "errors_map:  Permission denied.")
	// 编解码错误
	Register(ErrEncodingFailed, 500, "errors_map:  Encoding failed due to an error with the data.")
	Register(ErrDecodingFailed, 500, "errors_map:  Decoding failed due to an error with the data.")
	Register(ErrInvalidJSON, 500, "errors_map:  Data is not valid JSON.")
	Register(ErrEncodingJSON, 500, "errors_map:  JSON data could not be encoded.")
	Register(ErrDecodingJSON, 500, "errors_map:  JSON data could not be decoded.")
	Register(ErrInvalidYaml, 500, "errors_map:  Data is not valid Yaml.")
	Register(ErrEncodingYaml, 500, "errors_map:  Yaml data could not be encoded.")
	Register(ErrDecodingYaml, 500, "errors_map:  Yaml data could not be decoded.")

}
