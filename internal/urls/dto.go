package urls

type ShortenRequest struct {
	URL string `json:"url"`
}

type RedirectRequest struct{}

type ShortenResponse struct {
	Short string `json:"short"`
}
