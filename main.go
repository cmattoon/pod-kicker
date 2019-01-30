package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cmattoon/pod-kicker/config"
	customhandler "github.com/cmattoon/pod-kicker/handler"
	//"github.com/kubernetes/client-go/kubernetes/typed/core/v1"
	log "github.com/sirupsen/logrus"
)

func IsAuthorized(req *http.Request, cfg *config.Config) bool {
	auth := false
	for name, values := range req.Header {
		if strings.ToLower(name) == "x-pk-token" {
			for _, v := range values {
				if cfg.Token == v {
					auth = true
				}
			}
		}
	}
	return cfg.Token != "" && auth
}

func main() {
	cfg := config.NewConfig(os.Args[1:])
	handler := customhandler.NewRegexpHandler()

	handler.HandleFunc("/restart/.+", func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthorized(r, cfg) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("lol nope"))
			log.Errorf("Authorization Failed: Your Token is Wrong and Bad")
			return
		}

		app_name := strings.Replace(r.URL.Path, "/restart/", "", -1)
		log.Infof("Restarting %s", app_name)

		// @todo: 'sudo service %s restart'
		// @todo: 'sudo systemctl %s restart'
		// @todo: 'kubectl delete po $(kubectl get po -l app={SelectorFromQueryString})'

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("restarted %s", app_name)))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Infof("curl -H \"x-pk-token: %s\" http://localhost%s/restart/mysvc", cfg.Token, cfg.GetListenPort())

	http.ListenAndServe(cfg.GetListenPort(), handler)
}
