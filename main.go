
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello\nBuilt with Buildx & optimized base images.\n")
	})

	http.ListenAndServe(":8080", nil)
}
