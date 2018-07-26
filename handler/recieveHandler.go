package handler

import (
	"bitbucket.org/biglog/channel"

	"io/ioutil"
	"net/http"
)

func RevieveHandler(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if r.Method == "POST" {
		go channel.SendLog(result)
	}
	w.WriteHeader(http.StatusOK)
}
