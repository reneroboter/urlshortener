package handler

type PostRequest struct {
	Url string `json:"url"`
}

type PostResponse struct {
	code string `json:"code"`
}
