package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"text/template"

	qrcode "github.com/skip2/go-qrcode"
)

var (
	baseURL = os.Getenv("BASE_URL")
	port    = os.Getenv("PORT")

	showTmpl *template.Template
)

func init() {
	if baseURL == "" {
		baseURL = "http://localhost:8888"
	}
	if port == "" {
		port = "8888"
	}

	tmpl, err := template.ParseFiles("tmpl/show.html")
	if err != nil {
		panic(err)
	}

	showTmpl = tmpl
}

func main() {
	http.HandleFunc("/qr/create", func(w http.ResponseWriter, r *http.Request) {
		// get the url from the post form
		embedURL := r.PostFormValue("url")
		// url safe base64 encode the url
		embedURLB64 := base64.URLEncoding.EncodeToString([]byte(embedURL))

		// redirect to the show page
		http.Redirect(w, r, fmt.Sprintf("/qr/show/%s", embedURLB64), http.StatusFound)
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

	http.HandleFunc("/qr/show/", func(w http.ResponseWriter, r *http.Request) {
		// get the url from the path
		embedURLB64 := r.URL.Path[len("/qr/show/"):]
		// url safe base64 decode the url
		embedURL, err := base64.URLEncoding.DecodeString(embedURLB64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// write the html response
		w.Header().Set("Content-Type", "text/html")
		qrImageURL := fmt.Sprintf("%s/qr/view/%s", baseURL, embedURLB64)
		// read and serve the template file
		data := struct {
			URL        string
			QRImageURL string
		}{
			URL:        string(embedURL),
			QRImageURL: qrImageURL,
		}

		err = showTmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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
