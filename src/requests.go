package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestResult struct {
	StatusCode   int
	Body         []byte
	ContentType  string
	ResponseTime int64
}

func RunRequest(url string, headers map[string]string) (*RequestResult, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	start := time.Now()
	rres, err := client.Do(req)
	responseTime := time.Since(start)

	if err != nil {
		return nil, err
	}

	if rres.Body != nil {
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(rres.Body)
	}

	res := RequestResult{
		StatusCode:   rres.StatusCode,
		ResponseTime: responseTime.Milliseconds(),
	}

	body, err := ioutil.ReadAll(rres.Body)
	if err != nil {
		return nil, err
	}
	res.Body = body
	res.ContentType = rres.Header.Get("Content-Type")

	return &res, nil
}
