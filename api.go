package onetext

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler http handler
func (o *OneText) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	bytes, _ := json.MarshalIndent(o.Random(), "", "    ")
	_, _ = fmt.Fprint(w, string(bytes))
}

// StartServer start a local server
func (o *OneText) StartServer(port int) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", o.Handler)
	fmt.Printf("Starting an OneText server on 127.0.0.1:%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
