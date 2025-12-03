package urls

import (
	"net/http"

	"github.com/calqs/gopkg/router/handler"
	"github.com/calqs/gopkg/router/jsonresponse"
	"github.com/calqs/gopkg/router/response"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) ShortenURL(req *handler.Request[ShortenRequest]) response.Response {
	url, err := h.Service.ShortenURL(req.Request.Context(), req.Params.URL)
	if err != nil {
		return response.InternalServerError("could not shorten URL", err)
	}
	short := ShortenResponse{
		Short: url,
	}
	return jsonresponse.StatusCreated(short)
}

func (h *Handler) Redirect(req *handler.Request[RedirectRequest]) response.Response {
	code := req.Request.PathValue("code")
	longURL, err := h.Service.GetOriginalURL(req.Request.Context(), code)
	if err != nil {
		return response.InternalServerError("could not get original URL", err)
	}
	return &RedirectResponse{
		Location: longURL,
		Request:  req.Request,
	}
}

type Response interface {
	Send(http.ResponseWriter)
	SetHeader(string, string)
}

type RedirectResponse struct {
	Location string
	Request  *http.Request
}

func (rr *RedirectResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Location", rr.Location)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	http.Redirect(w, rr.Request, rr.Location, http.StatusFound)
}
func (rr *RedirectResponse) SetHeader(key, value string) {}
