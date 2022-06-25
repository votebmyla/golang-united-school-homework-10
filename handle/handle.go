package handle

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

func HandleData(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	param := fmt.Sprintf("I got message:\n%s", string(data))

	w.Write([]byte(param))
}

func HandleHeader(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		log.Fatal(err)
	}
	result := strconv.Itoa(a + b)
	w.Header().Set("a+b", result)
}
