package domain

type ShortURL struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

func (s ShortURL) isValidCode() bool {
	return IsValidCode(s.Code)
}

func (s ShortURL) isValidURL() bool {
	return IsValidURL(s.URL)
}

func (s ShortURL) Validate() {
	s.isValidURL()
	s.isValidCode()
}
