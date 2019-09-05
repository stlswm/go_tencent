package tool

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// http get 请求
func HttpGet(url string, urlValue *url.Values) (error, []byte) {
	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}
	if urlValue != nil {
		if strings.Index(url, "?") != -1 {
			url += "&" + urlValue.Encode()
		} else {
			url += "?" + urlValue.Encode()
		}
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, []byte{}
	}
	// 处理返回结果
	response, err := client.Do(request)
	if err != nil {
		return err, []byte{}
	}
	defer func() {
		_ = response.Body.Close()
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err, []byte{}
	}
	return nil, body
}

// http post 请求
func HttpPost() {

}
