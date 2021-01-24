package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// ハンドラを設定
	http.HandleFunc("/reciever", reciever)

	// ポート番号を設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// HTTPサーバの起動
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
		return
	}
}

// 他所からメッセージを受け取ってラインに流す
func reciever(w http.ResponseWriter, req *http.Request) {
	r := struct {
		Message string `json:"message"`
		From    string `json:"from"`
	}{}

	// POST以外は弾く
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Only Post is allowed")
	}

	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, err)
		return
	}

	if err := sendMessage(r.Message, r.From); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success")
}
