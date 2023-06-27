package webserver

// Default Go Web Server
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/valerius21/scap/pkg/fns"
)

const sleeperDuration = 1 * time.Second

func NetHttp() {
	// Define the routes and their corresponding handlers
	http.HandleFunc("/image", ImageHandler)
	http.HandleFunc("/sleep", SleepHandler)
	http.HandleFunc("/empty", EmptyHandler)
	http.HandleFunc("/math", MathHandler)

	// Start the server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "imageHandler: not implemented")
}

func SleepHandler(w http.ResponseWriter, r *http.Request) {
	// Sleep for 1 seconds before responding
	fns.SleeperFn(sleeperDuration)

	// omit a text response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "sleepHandler: slept for %v", sleeperDuration)
}

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	// Return an empty response with status code 204
	fns.EmptyFn() // TODO: is this necessary?
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "emptyHandler: empty response")
}

func MathHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameter "number" and calculate its square root
	numberStr := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	result := fns.MathFn(number)
	fmt.Fprintf(w, "Square root of %v: %v", number, result)
}
