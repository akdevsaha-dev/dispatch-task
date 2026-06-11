package handlers

import (
	"fmt"
	"net/http"
)

func JobHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This route is working")
}
