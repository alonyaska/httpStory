package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

type authRoundTripper struct {
	token string
	next  http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s \n ", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}
func (a authRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", "Bearer "+a.token)
	fmt.Println(string(a.token))
	return a.next.RoundTrip(r)
}
func generateToken() string {
	b := make([]byte, 8) // даст токен длиной 16 символов
	rand.Read(b)
	return hex.EncodeToString(b)
}

func main() {
	jar, _ := cookiejar.New(nil) 

	token := generateToken()

	client := &http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.Status)
			fmt.Println("redirect")
			return nil
		},
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next: authRoundTripper{
				token: token,
				next:  http.DefaultTransport,

			
			},
		},
	}

	resp, err := client.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("response status -", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
