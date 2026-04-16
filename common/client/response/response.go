package response

typpe ApiResponse struct {
	suuccess bool `json:"success"`
	message string `json:"message"`
	data interface{} `json:"data"`
}