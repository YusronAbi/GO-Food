package helper

import "fmt"

// Struct untuk response umum
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// Response sukses dengan data
func SuccessResponse(message string, data any) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// Response gagal tanpa data
func FailedResponse(message string) Response {
	return Response{
		Success: false,
		Message: message,
	}
}

// Response untuk autentikasi
func AuthResponse(message, token string) Response {
	return Response{
		Success: true,
		Message: message,
		Token:   token,
	}
}

// Struct untuk pesan sistem
type ResponseMessage struct{}

// Metode untuk pesan sukses
func (*ResponseMessage) CreateSuccess(name string) string {
	return fmt.Sprintf("Create data %s successfully", name)
}
func (*ResponseMessage) GetSuccess(name string) string {
	return fmt.Sprintf("Fetch data %s successfully", name)
}
func (*ResponseMessage) UpdateSuccess(name string) string {
	return fmt.Sprintf("Update data %s successfully", name)
}
func (*ResponseMessage) DeleteSuccess(name string) string {
	return fmt.Sprintf("Delete %s successfully", name)
}
func (*ResponseMessage) DeleteAllSuccess(name string) string {
	return fmt.Sprintf("Delete all %s successfully", name)
}
func (*ResponseMessage) LoginSuccess() string    { return "Login successfully" }
func (*ResponseMessage) RegisterSuccess() string { return "Register successfully" }

// Metode untuk pesan gagal
func (*ResponseMessage) CreateFailed(name string) string {
	return fmt.Sprintf("Create data %s failed", name)
}
func (*ResponseMessage) GetFailed(name string) string {
	return fmt.Sprintf("Fetch data %s failed", name)
}
func (*ResponseMessage) UpdateFailed(name string) string {
	return fmt.Sprintf("Update data %s failed", name)
}
func (*ResponseMessage) DeleteFailed(name string) string {
	return fmt.Sprintf("Delete %s failed", name)
}
func (*ResponseMessage) DeleteAllFailed(name string) string {
	return fmt.Sprintf("Delete all %s failed", name)
}
func (*ResponseMessage) RequestFailed(name string) string {
	return fmt.Sprintf("Invalid %s request payload", name)
}
func (*ResponseMessage) NotFound(name string) string { return fmt.Sprintf("Fetch %s not found", name) }
func (*ResponseMessage) IDInvalid(id string) string  { return fmt.Sprintf("Invalid %s ID", id) }

// Metode untuk pesan error autentikasi
func (*ResponseMessage) LoginFailed() Response {
	return FailedResponse("Login failed, please check email and password!")
}
func (*ResponseMessage) LoginFailedEntity() Response {
	return FailedResponse("Email and password are not valid!")
}
func (*ResponseMessage) RegisterFailed() Response {
	return FailedResponse("Register failed, email already exists!")
}
func (*ResponseMessage) RegisterFailedEntity() Response {
	return FailedResponse("Email and password are not valid!")
}
