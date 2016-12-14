package elasticsearch

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	url string
}

func (es *Client) PostCall(url string, body []byte) ([]byte, error) {
	rdr := bytes.NewReader(body)
	req, _ := http.NewRequest("POST", es.url+url, rdr)
	//cn.Println("Posting to ", es.url+url)
	req.Header.Set("content-type", "application/json")
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	bts, _ := ioutil.ReadAll(res.Body)
	return bts, nil
}

func (es *Client) DeleteCall(url string, body []byte) ([]byte, error) {
	rdr := bytes.NewReader(body)
	req, _ := http.NewRequest("DELETE", es.url+url, rdr)
	fmt.Println("Posting to ", es.url+url)
	req.Header.Set("content-type", "application/json")
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	bts, _ := ioutil.ReadAll(res.Body)
	return bts, nil
}

func (es *Client) GetCall(url string, body []byte) ([]byte, error) {
	rdr := bytes.NewReader(body)
	req, _ := http.NewRequest("GET", es.url+url, rdr)
	req.Header.Set("content-type", "application/json")
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	bts, _ := ioutil.ReadAll(res.Body)
	return bts, nil
}

func (es *Client) PutCall(url string, body []byte) ([]byte, error) {
	rdr := bytes.NewReader(body)
	req, _ := http.NewRequest("PUT", es.url+url, rdr)
	fmt.Println("Posting to ", es.url+url)
	req.Header.Set("content-type", "application/json")
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	bts, _ := ioutil.ReadAll(res.Body)
	return bts, nil
}

func (es *Client) HeadCall(url string, body []byte) ([]byte, error) {
	res, err := http.Head(es.url + url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	bts, _ := ioutil.ReadAll(res.Body)
	return bts, nil
}

func (es *Client) StatusCode(url string) (int, error) {
	res, err := http.Head(es.url + url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	return res.StatusCode, nil
}

func (es *Client) Search(statement *Statement, indexPath string) ([]byte, error) {
	body := statement.Compile()
	fmt.Println(string(body))
	return es.PostCall(indexPath+"/_search", body)
}

func NewClient(url string) *Client {
	return &Client{url: url}
}
