package response

type PingResponse struct {
	Database string `json:"database"`
	Redis    string `json:"redis"`
}
