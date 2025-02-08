/**
    @author: dongjs
    @date: 2025/2/7
    @description:
**/

package v1

import (
	"io"
	"io/ioutil"
	"net/http"
)

// 发起请求
func (c *Client) doRequest(url string, buf io.Reader) ([]byte, error) {
	req, _ := http.NewRequest("POST", c.GetRequestUrl()+url, buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("sys_code", c.sysCode)
	req.Header.Set("x-access-key", c.accessKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}
