package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

var (
	baseURL = os.Getenv("BASE_URL")
	port    = os.Getenv("PORT")
)

func init() {
	if baseURL == "" {
		baseURL = "http://localhost:8888"
	}
	if port == "" {
		port = "8888"
	}
}

func main() {
	http.HandleFunc("/qr/create", func(w http.ResponseWriter, r *http.Request) {
		// get the url from the post form
		embedURL := r.PostFormValue("url")
		// url safe base64 encode the url
		embedURLB64 := base64.URLEncoding.EncodeToString([]byte(embedURL))

		// return the view url as json
		w.Header().Set("Content-Type", "application/json")
		res := map[string]string{
			"url": fmt.Sprintf("%s/qr/view/%s", baseURL, embedURLB64),
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/qr/view/", func(w http.ResponseWriter, r *http.Request) {
		// get the url from the path
		embedURLB64 := r.URL.Path[len("/qr/view/"):]
		// url safe base64 decode the url
		embedURL, err := base64.URLEncoding.DecodeString(embedURLB64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// create the qr code
		qr, err := qrcode.Encode(string(embedURL), qrcode.Low, 256)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// write the qr code to the response
		w.Header().Set("Content-Type", "image/png")
		_, err = w.Write(qr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tmpl/index.html")
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
