package webserver

// Default Go Web Server
import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
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
	_, file, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Image: could not get file", http.StatusInternalServerError)
		return
	}
	args, err := ImageSaver(file)
	if err != nil {
		http.Error(w, "Image: could not save image", http.StatusInternalServerError)
		return
	}
	msg, err := CreateHandler("net/http", "image", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(msg))
}

func SleepHandler(w http.ResponseWriter, _ *http.Request) {
	msg, err := CreateHandler("net/http", "sleep", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// omit a text response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(msg))
}

func EmptyHandler(w http.ResponseWriter, _ *http.Request) {
	msg, err := CreateHandler("net/http", "empty", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(msg))
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
	fmt.Fprint(w, string(msg))
}
