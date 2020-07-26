package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Println("start http listening: 18443")
// 	err := http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil)
// 	log.Println(err)
// }

// func main() {
// 	server := &http.Server{
// 		TLSConfig: &tls.Config{
// 			ClientAuth: tls.RequireAndVerifyClientCert,
// 			MinVersion: tls.VersionTLS12,
// 		},
// 		Addr: ":18443",
// 	}

// 	http.HandleFunc("/", handler)
// 	log.Println("start https listening :18443")
// 	err := server.ListenAndServeTLS("server.crt", "server.key")
// 	log.Println(err)
// }

func handlerChunkedResponse(w http.ResponseWriter, r *http.Request) {
	flusher, _ := w.(http.Flusher)

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "Chunk %d\n", i)
		flusher.Flush()
		time.Sleep(500 * time.Millisecond)
	}
	flusher.Flush()
}

func main() {
	http.HandleFunc("/chunked", handlerChunkedResponse)
	log.Println("start http listening: 18443")
	err := http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil)
	log.Println(err)
}
