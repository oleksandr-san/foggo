package foggo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func fetchParameters(r *http.Request) (string, float64, error) {
	id := r.FormValue("id")
	t, err := strconv.ParseFloat(r.FormValue("temperature"), 64)
	if err != nil {
		return "", 0., err
	}

	return id, t, nil
}

func HelloPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("parse form error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, t, err := fetchParameters(r)

	if err != nil {
		log.Printf("fetching parameters error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(true)
	if err != nil {
		log.Printf("marshalling error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("id: %s; temperature: %v, response: %v\n", id, t, data)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func ListGet(w http.ResponseWriter, r *http.Request) {
	var data []Data
	data = append(data, Data{"LUL", 36.6, int32(0)})

	mdata, err := json.Marshal(data)
	if err != nil {
		log.Printf("marshalling error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(mdata)
}
