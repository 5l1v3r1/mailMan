package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// check args
	if len(os.Args) < 2 {
		log.Fatal(errors.New("Usage: ./" + os.Args[0] + " <servicename://host:port> ...\n"))
	}

	// create services map
	services := make(map[string]string)
	for index, arg := range os.Args {
		if index < 1 {
			continue
		}
		argParts := strings.Split(arg, "://")
		services[argParts[0]] = argParts[1]
	}

	// create HTTP client obj
	client := &http.Client{
		Timeout: time.Millisecond * 5,
	}

	// define handler
	http.HandleFunc("/msf", func(w http.ResponseWriter, r *http.Request) {

		// create request object
		request, err := http.NewRequest(r.Method, "https://"+services[r.URL.Path], r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, err, 404)
			return
		}

		// copy request data to request object
		for key, values := range r.Header {
			for _, value := range values {
				request.Header.Add(key, value)
			}
		}

		// send request
		resp, err := client.Do(request)
		if err != nil {
			w.Write([]byte(err.Error() + "\n"))
			http.Error(w, err, 500)
			return
		}

		// copy client result as response
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		// write response
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	// start!
	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
