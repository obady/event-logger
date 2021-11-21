package main

import (
	"io/ioutil"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        w.Header().Set("Content-Type", "application/json")
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        app.errorLog.Printf("cant read request body")
    }
    app.infoLog.Printf(string(body))
}

func (app *application) health(w http.ResponseWriter, r *http.Request) {

    w.Write([]byte("if you can see this then everything is ok"))
}