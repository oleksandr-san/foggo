package server

import (
	"encoding/json"
	"time"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"io/ioutil"

	"golang.org/x/net/context"

	"../common"
	cl "../client"
)

const addressServer = "http://lab.ytdev.com/"

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

	data := common.Data{
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

func getIPAdresses() ([]string, error) {
	response, err := http.Get(addressServer)
    if err != nil {
        return nil, err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var addresses []string
	if err := json.Unmarshal(contents, &addresses); err != nil {
		return nil, err
	}

	return addresses, nil
}

func getDataFromServer(ip string) ([]common.Data, error) {
	config := cl.NewConfiguration()
	client := cl.NewAPIClient(config) 

	client.ChangeBasePath("http://" + ip + ":3001")

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	data, _, err := client.DefaultApi.ListGet(ctx)
	if err != nil {
		return nil, err
	}

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("done")
	}

	return data, nil
}


func Discover(w http.ResponseWriter, r *http.Request) {
	ips, err := getIPAdresses()
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		data, err := getDataFromServer(ip)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range data {
			newItem := common.Data{
				Id: item.Id,
				Temperature: item.Temperature,
				Timestamp: item.Timestamp,
			}
			
			err = AddData(newItem)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}

	js, err := json.Marshal(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}