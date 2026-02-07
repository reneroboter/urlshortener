package handler

type PostRequest struct {
	Url string `json:"url"`
}

type PostResponse struct {
	Code string `json:"code"`
}
