package pdf

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	u "net/url"
	"strings"
)

// Gen generates a pdf
// Deprecated: use Create instead
func Gen(url string, key string, isMarginless bool) (*bytes.Buffer, error) {
	marginless := ""
	if isMarginless {
		marginless = "marginless"
	}
	return Create(url, key, marginless)
}

func Create(urlOrText string, key string, opts ...string) (*bytes.Buffer, error) {
	form := u.Values{}
	form.Add("url", urlOrText)

	for _, opt := range opts {
		if opt == "marginless" {
			form.Add("nomargin", "1")
		}

		if opt == "nomargin" {
			form.Add("nomargin", "1")
		}

		if opt == "landscape" {
			form.Add("orientation", "landscape")
		}
	}

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
		return nil, fmt.Errorf("Request failed status code: %d %s", resp.StatusCode, string(b))
	}
	return bytes.NewBuffer(b), err
}
