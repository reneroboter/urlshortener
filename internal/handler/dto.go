package handler

type PostRequest struct {
	Url string `json:"url"`
}

type PostResponse struct {
	ID string `json:"id"`
}
