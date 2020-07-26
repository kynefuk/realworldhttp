package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
)

// URL is a local server address
const URL = "https://localhost:18443"

// func main() {
// 	cert, _ := ioutil.ReadFile("../ca/ca.crt")
// 	certPool := x509.NewCertPool()
// 	certPool.AppendCertsFromPEM(cert)
// 	tlsConfig := &tls.Config{
// 		RootCAs: certPool,
// 	}
// 	tlsConfig.BuildNameToCertificate()

// 	client := &http.Client{
// 		Transport: &http.Transport{
// 			TLSClientConfig: tlsConfig,
// 		},
// 	}

// 	resp, err := client.Get(URL)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer resp.Body.Close()

// 	dump, _ := httputil.DumpResponse(resp, true)
// 	log.Println(string(dump))
// }

func main() {
	cert, _ := tls.LoadX509KeyPair("../client/client.crt", "../client/client.key")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	resp, err := client.Get(URL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	dump, _ := httputil.DumpResponse(resp, true)
	log.Println(string(dump))
}
