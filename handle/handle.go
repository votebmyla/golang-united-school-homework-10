package handle

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleName(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)["PARAM"]
	s := fmt.Sprintf("Hello, %s!", p)
	fmt.Fprint(w, s)
}

func HandleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}
