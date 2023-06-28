package webserver

// Default Go Web Server
import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

func NetHttp(receiverPort string) {
	// Define the routes and their corresponding handlers
	http.HandleFunc("/image", ImageHandler)
	http.HandleFunc("/sleep", SleepHandler)
	http.HandleFunc("/empty", EmptyHandler)
	http.HandleFunc("/math", MathHandler)

	// Start the server
	log.Info().Msg("Server listening on port " + receiverPort)
	log.Error().Err(http.ListenAndServe(":"+receiverPort, nil))
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "imageHandler: not implemented")
}

func SleepHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := CreateHandler("net/http", "sleep", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// omit a text response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(msg))
}

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := CreateHandler("net/http", "empty", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(msg))
}

func MathHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameter "number" and calculate its square root
	numberStr := r.URL.Query().Get("number")
	msg, err := CreateHandler("net/http", "math", numberStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(msg))
}
