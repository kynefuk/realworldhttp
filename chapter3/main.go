package main

import (
	"log"
	"net/http"
	"net/http/httputil"
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
// func main() {
// 	var buffer bytes.Buffer
// 	writer := multipart.NewWriter(&buffer)
// 	writer.WriteField("name", "Jack")

// 	fileWriter, _ := writer.CreateFormFile("sample", "main.go")
// 	readFile, _ := os.Open("main.go")
// 	defer readFile.Close()
// 	io.Copy(fileWriter, readFile)
// 	writer.Close()
// 	log.Println(writer.FormDataContentType())

// 	resp, _ := http.Post(URL, writer.FormDataContentType(), &buffer)
// 	log.Println(resp.Status)
// }

// func main() {
// 	jar, _ := cookiejar.New(nil)
// 	client := http.Client{
// 		Jar: jar,
// 	}

// 	for i := 0; i < 2; i++ {
// 		resp, _ := client.Get(URL + "/cookie")
// 		dump, _ := httputil.DumpResponse(resp, true)
// 		log.Println(string(dump))
// 	}
// }

// プロキシサーバを経由してリクエストを送る
// func main() {
// 	proxyURL, _ := url.Parse(URL)
// 	client := http.Client{
// 		Transport: &http.Transport{
// 			Proxy: http.ProxyURL(proxyURL),
// 		},
// 	}

// 	resp, _ := client.Get("http://github.com")
// 	dump, _ := httputil.DumpResponse(resp, true)
// 	log.Println(string(dump))
// }

// func main() {
// 	transport := &http.Transport{}
// 	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
// 	client := http.Client{
// 		Transport: transport,
// 	}
// 	resp, _ := client.Get("file://./main.go")
// 	dump, _ := httputil.DumpResponse(resp, true)
// 	log.Println(string(dump))
// }

// GET,HEAD,POST以外はRequestを作成する必要がある
func main() {
	client := http.Client{}
	request, _ := http.NewRequest("DELETE", URL, nil)
	request.Header.Add("X-HOGE", "hoge")
	request.SetBasicAuth("username", "password")
	resp, _ := client.Do(request)
	dump, _ := httputil.DumpResponse(resp, true)
	log.Println(string(dump))
}
