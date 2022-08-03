package vade

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
)

type Document struct {
    Id string `json:"Id"`
    Name string `json:"Name"`
    Description string `json:"Description"`
}

var Documents []Document

func AllDocuments(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: AllDocuments")
    json.NewEncoder(w).Encode(Documents)
}

func ShowDocument(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Fprintf(w, "Key: " + key)
    for _, document := range Documents {
        if document.Id == key {
            json.NewEncoder(w).Encode(document)
        }
    }
}

func CreateDocument(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var document Document
    json.Unmarshal(reqBody, &document)
    Documents = append(Documents, document)
    json.NewEncoder(w).Encode(document)
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    for index, document := range Documents {
        if document.Id == id {
            Documents = append(Documents[:index], Documents[index+1:]...)
        }
    }
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/documents", AllDocuments).Methods("GET")
    myRouter.HandleFunc("/documents", CreateDocument).Methods("POST")
    myRouter.HandleFunc("/document/{id}", DeleteDocument).Methods("DELETE")
    myRouter.HandleFunc("/document/{id}", ShowDocument)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Documents = []Document{
        Document{Id: "1", Name: "eak", Description: "yo Description"},
        Document{Id: "2", Name: "oak", Description: "oy Description"},
    }
    handleRequests()
}
