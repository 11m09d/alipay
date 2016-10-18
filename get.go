package alipay

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

// doRequest get the order in xml format with a sign
func DoHttpGet(targetUrl string) ([]byte, error) {
	//req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(targetUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respData, nil
}
