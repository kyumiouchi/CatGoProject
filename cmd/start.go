package cmd

import (
	"cat-project/config"
	"cat-project/internal/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "world")
}

func nameHello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func Start() {
	ycg := config.NewYamlConfigurationGetter()
	cfg, err := ycg.GetConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Failed get config") // panic
		return
	}
	logger := logging.New(logging.Config{LogLevel: cfg.Logging.Level})
	logger.Infof("Message hello world port: %s", cfg.Http.Port)

	mux := httprouter.New()
	mux.GET("/hello/:name", nameHello)
	mux.GET("/hello", hello)
	mux.GET("/world", world)

	server := http.Server{
		Addr: "127.0.0.1:" + cfg.Http.Port,
		Handler: mux,
	}
	server.ListenAndServe()
}