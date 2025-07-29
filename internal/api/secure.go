package api

import (
	"fmt"
	"net/http"
)

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You have accessed the following endpoints!")
}
