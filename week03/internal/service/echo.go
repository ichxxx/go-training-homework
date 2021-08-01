package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	// 模拟耗时
	time.Sleep(5 * time.Second)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("service: read http body failed,", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		log.Println("service: write http body failed,", err.Error())
	}
}