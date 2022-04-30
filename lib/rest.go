package lib

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestOptions struct {
	Url     string
	Method  string
	Body    io.Reader
	Query   map[string]string
	Headers map[string]string
}

func Request(options *RequestOptions, result interface{}) {
	client := &http.Client{}

	req, err := http.NewRequest(options.Method, options.Url, options.Body)
	if err != nil {
		panic(err)
	}

	if options.Query != nil {
		q := req.URL.Query()
		for index, value := range options.Query {
			q.Add(index, value)
		}

		req.URL.RawQuery = q.Encode()
	}

	if options.Headers != nil {
		for prop, value := range options.Headers {
			req.Header.Add(prop, value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(bytes))

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		panic(err)
	}
}
