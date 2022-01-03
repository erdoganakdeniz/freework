package main

import (
	"encoding/json"
	"fmt"
	"freework/models"
	"freework/service"
	"freework/store"
	"freework/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var s *store.Store
var PORT = ":8080"

func main() {
	s = store.New(35*time.Second, 35*time.Second)
	utils.LoadfromFile(s)
	utils.SavetoFile(s)

	go func() {
		//Running it synchronously.
		utils.SaveInterval(s)
	}()

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
	var keyvalue models.KeyValue
	err = json.Unmarshal(d, &keyvalue)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error : ", http.StatusBadRequest)
		return
	}
	era := service.Set(s, keyvalue)
	res, _ := json.Marshal(era)
	w.Write(res)
	utils.SavetoFile(s)
}

// GET all the stored store
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Serving:", r.URL.Path, " from", r.Host, r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Error : ", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	era := service.Get(s)
	res, _ := json.Marshal(era)
	w.Write(res)
	utils.SavetoFile(s)

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
	service.Flush(s)
	utils.SavetoFile(s)

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
	era := service.GetByKey(s, key)
	res, _ := json.Marshal(era)
	w.Write(res)
	utils.SavetoFile(s)

}

//Delete Value by Key
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Serving:", r.URL.Path, " from", r.Host, r.Method)
	if r.Method != http.MethodDelete {
		http.Error(w, "Error : ", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	key := strings.TrimPrefix(r.URL.Path, "/delete/")
	era := service.Delete(s, key)
	res, _ := json.Marshal(era)
	w.Write(res)
	utils.SavetoFile(s)

}
