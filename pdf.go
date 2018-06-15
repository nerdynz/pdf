package pdf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	u "net/url"
	"strings"
)

func Gen(url string, key string) ([]byte, error) {
	form := u.Values{}
	form.Add("url", url)

	client := http.Client{}
	r, err := http.NewRequest("POST", "https://pdf.nerdy.co.nz/make", strings.NewReader(form.Encode()))
	// r, err := http.NewRequest("POST", "http://localhost:5000/api/v1/notes/parse", bytes.NewBufferString(str))
	if err != nil {
		return nil, errors.New("Couldn't make the request" + err.Error())
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("key", key)
	resp, err := client.Do(r)
	if err != nil {
		return nil, errors.New("request failed" + err.Error())
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("request failed" + err.Error())
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed status code: %d", resp.StatusCode)
	}
	return b, err
}
