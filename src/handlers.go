package foggo

import (
	"encoding/json"
	"time"
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

	data := Data{
		Id: id,
		Temperature: float32(t),
		Timestamp: int32(time.Now().Unix()),
	}

	err = AddData(data)
	if err != nil {
		log.Printf("adding data: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(true)
	if err != nil {
		log.Printf("marshalling error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func ListGet(w http.ResponseWriter, r *http.Request) {
	timestamp := int(time.Now().Add(time.Minute * -5).Unix())
	data, err := GetData(timestamp)

	if err != nil {
		log.Printf("database error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
