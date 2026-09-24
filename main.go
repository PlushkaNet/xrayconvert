package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	share "github.com/xtls/libxray/share"
)

func ParseUri(links string) (json.RawMessage, error) {
	return share.ConvertShareLinksToXrayJson(links, "")
}

func HandleUriRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uri := r.FormValue("uri")
	if uri == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	conf, err := ParseUri(uri)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(conf.String()))
}

func main() {
	addr := flag.String("address", "127.0.0.1:10100", "address:port to listen on")
	flag.Parse()
	log.Printf("Starting server on %s\n", *addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleUriRequest)
	log.Fatalf("Fatal: %s\n", http.ListenAndServe(*addr, mux).Error())
}
