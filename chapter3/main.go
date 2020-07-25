package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// URL is a local server address
const URL = "http://localhost:18888"

// func main() {
// 	values := url.Values{
// 		"query": {"hello world"},
// 	}

// 	resp, _ := http.Get(URL + "?" + values.Encode())
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	log.Println(string(body))
// }

// x-www-form-urlencoded形式
// func main() {
// 	values := url.Values{
// 		"test": {"value"},
// 	}

// 	resp, _ := http.PostForm(URL, values)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	log.Println(string(body))
// 	defer resp.Body.Close()
// }

// func main() {
// 	file, _ := os.Open("main.go")
// 	resp, _ := http.Post(URL, "text/plain", file)
// 	log.Println(resp.StatusCode)
// }

// func main() {
// 	reader := strings.NewReader("テキスト")
// 	resp, _ := http.Post(URL, "text/plain", reader)
// 	log.Println(resp.Status)
// }

// multipart/form-data形式
func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Jack")

	fileWriter, _ := writer.CreateFormFile("sample", "main.go")
	readFile, _ := os.Open("main.go")
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()
	log.Println(writer.FormDataContentType())

	resp, _ := http.Post(URL, writer.FormDataContentType(), &buffer)
	log.Println(resp.Status)
}
