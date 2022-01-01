package main

import (
	"encoding/json"
	"fmt"
	"freework/cache"
	"freework/model"
	"freework/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var c *cache.Cache

var PORT = ":8080"

func main() {
	c = cache.New(2*time.Minute, 2*time.Minute)
	fileName := "TIMESTAMP-data.gob"
	_, erra := os.Open("tmp/" + fileName)
	if erra != nil {

	}
	c.LoadFile("tmp/" + fileName)

	mux := http.NewServeMux()
	mux.Handle("/get", http.HandlerFunc(Get))
	mux.Handle("/set", http.HandlerFunc(Set))
	mux.Handle("/flush", http.HandlerFunc(Flush))
	mux.Handle("/delete/", http.HandlerFunc(Delete))
	mux.Handle("/get/", http.HandlerFunc(GetByKey))

	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	fmt.Println("Ready to serve at", PORT)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// SET key and value in store
func Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Serving :", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Error :", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error : ", http.StatusInternalServerError)
		return
	}
	var keyvalue model.KeyValue
	err = json.Unmarshal(d, &keyvalue)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error : ", http.StatusBadRequest)
		return
	}
	era := service.Set(c, keyvalue)
	res, _ := json.Marshal(era)
	w.Write(res)
	c.SaveFile("tmp/TIMESTAMP-data.gob")
}

// GET all the stored cache
func Get(w http.ResponseWriter, r *http.Request) {
	c.LoadFile("tmp/TIMESTAMP-data.gob")
	w.Header().Set("Content-Type", "application/json")
	log.Println("Serving:", r.URL.Path, " from", r.Host, r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Error : ", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	era := service.Get(c)
	res, _ := json.Marshal(era)
	w.Write(res)
	c.SaveFile("tmp/TIMESTAMP-data.gob")
}

// FLUSH all the data
func Flush(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Serving:", r.URL.Path, " from", r.Host, r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Error : ", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	service.Flush(c)
	c.SaveFile("tmp/TIMESTAMP-data.gob")

}

// GET Value by Key
func GetByKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Serving:", r.URL.Path, " from", r.Host, r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Error : ", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	key := strings.TrimPrefix(r.URL.Path, "/get/")
	era := service.GetByKey(c, key)
	res, _ := json.Marshal(era)
	w.Write(res)
	c.SaveFile("tmp/TIMESTAMP-data.gob")

}
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Serving:", r.URL.Path, " from", r.Host, r.Method)
	if r.Method != http.MethodDelete {
		http.Error(w, "Error : ", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	key := strings.TrimPrefix(r.URL.Path, "/delete/")
	era := service.Delete(c, key)
	res, _ := json.Marshal(era)
	w.Write(res)
	c.SaveFile("tmp/TIMESTAMP-data.gob")

}
func SaveFile() {
	_, err := os.Create("tmp/TIMESTAMP-data.gob")
	if err != nil {
		fmt.Println(err)
	}
}
