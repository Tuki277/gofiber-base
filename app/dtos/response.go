package dto

type Response struct {
	Code     int           `json:"code"`
	Message  string        `json:"message"`
	Data     interface{}   `json:"data"`
	Metadata *ListResponse `json:"metadata"`
	Error    error         `json:"error"`
	Success  bool          `json:"success"`
}

type ListResponse struct {
	TotalItems   int `json:"totalItems"`
	ItemsPerPage int `json:"itemPerPage"`
	TotalPages   int `json:"totalPages"`
	CurrentPage  int `json:"currentPage"`
}
