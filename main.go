package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func databases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // for CORS
	w.WriteHeader(http.StatusOK)
	test := []string{}
	test = append(test, "Hello Now Testing Pictures")
	test = append(test, "World")

	// json.NewEncoder(w).Encode(test)
	err := json.NewEncoder(w).Encode(test)
if err != nil {
    // Handle the error, for example, log it or send an error response.
    log.Println("Error encoding JSON:", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
}


}

func main() {

	//  mime.AddExtensionType(".js", "application/javascript")

	http.Handle("/test", http.HandlerFunc(databases))
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
