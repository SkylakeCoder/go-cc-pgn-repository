package repository

import (
	"log"
	"net/http"
)

func StartServer(host string, port string) {
	http.HandleFunc("/", handleRequest)
	log.Printf("server listening at: %s:%s\n", host, port)
	log.Fatalln(http.ListenAndServe(host + ":" + port, nil))
}

func handleRequest(resp http.ResponseWriter, req *http.Request) {
	log.Printf("request coming...")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	chessBoard := req.FormValue("chess")
	result := Search(chessBoard)
	resp.Write([]byte(result))
}
