package main
import (
  "net/http"
  "strings"
  "io/ioutil"
  "encoding/base64"
  "crypto/hmac"
  "crypto/sha256"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}

func webhook(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
	    // ...
	}
	decoded, err := base64.StdEncoding.DecodeString(req.Header.Get("X-Line-Signature"))
	if err != nil {
	    // ...
	}
	hash := hmac.New(sha256.New, []byte("<channel secret>"))
	hash.Write(body)
	// Compare decoded signature and `hash.Sum(nil)` by using `hmac.Equal`
}

func main() {
  http.HandleFunc("/", sayHello)
  http.HandleFunc("/webhook", webhook)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}

