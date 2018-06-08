package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/vothanhkiet/noop/libs"
)

var log = logrus.New()

func main() {
	token := os.Getenv("LOGENTRIES_TOKEN")
	if token != "" {
		logentriesHook, err := libs.NewLogentriesHook(token)
		if err == nil {
			logentriesHook.SetFormatter(&logrus.JSONFormatter{})
			log.AddHook(logentriesHook)
		}
	}
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = ioutil.Discard

	log.Info("No-op http server is running")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ping", healthHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	params, _ := url.ParseQuery(r.URL.RawQuery)
	log.WithFields(logrus.Fields{
		"headers":        convert(r.Header),
		"method":         r.Method,
		"host":           r.Host,
		"protocol":       r.Proto,
		"remote":         r.RemoteAddr,
		"content_length": r.ContentLength,
		"params":         convert(params),
	}).Info(r.Method + " " + r.URL.Path)
	w.WriteHeader(418)
	fmt.Fprint(w, "")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprint(w, "")
}

func convert(input map[string][]string) map[string]string {
	ret := make(map[string]string)
	for key, value := range input {
		ret[key] = value[0]
	}

	return ret
}
