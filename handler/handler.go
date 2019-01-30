package handler

import (
	"net/http"
	"regexp"

	log "github.com/sirupsen/logrus"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func NewRegexpHandler() *RegexpHandler {
	return &RegexpHandler{
		routes: []*route{},
	}
}

func NewRoute(pattern_txt string, handler http.Handler) *route {
	log.Infof("Creating Regexp(%s)", pattern_txt)
	if pattern, err := regexp.Compile(pattern_txt); err == nil {
		return &route{
			pattern,
			handler,
		}
	}
	return nil
}

func (h *RegexpHandler) Handler(pattern string, handler http.Handler) {
	h.routes = append(h.routes, NewRoute(pattern, handler))
}

func (h *RegexpHandler) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, NewRoute(pattern, http.HandlerFunc(handler)))
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Infof("Matching against '%s'", r.URL.Path)
	for _, route := range h.routes {
		if route != nil && route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}
