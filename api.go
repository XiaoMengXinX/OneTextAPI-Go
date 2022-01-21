package onetext

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler http handler
func (o *onetext) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	bytes, _ := json.MarshalIndent(o.Random(), "", "    ")
	_, _ = fmt.Fprint(w, string(bytes))
}

// StartServer start a local server
func (o *onetext) StartServer(port int) error {
	http.HandleFunc("/", o.Handler)
	fmt.Printf("Starting an onetext server on 127.0.0.1:%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
