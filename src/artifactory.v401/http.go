package artifactory

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *ArtifactoryClient) Get(path string, options map[string]string) ([]byte, error) {
	return c.makeRequest("GET", path, options, nil)
}

func (c *ArtifactoryClient) Post(path string, data string, options map[string]string) ([]byte, error) {
	body := strings.NewReader(data)
	return c.makeRequest("POST", path, options, body)
}

func (c *ArtifactoryClient) Put(path string, data string, options map[string]string) ([]byte, error) {
	body := strings.NewReader(data)
	return c.makeRequest("PUT", path, options, body)
}

func (c *ArtifactoryClient) Delete(path string) error {
	_, err := c.makeRequest("DELETE", path, make(map[string]string), nil)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *ArtifactoryClient) makeRequest(method string, path string, options map[string]string, body io.Reader) ([]byte, error) {
	qs := url.Values{}
	for q, p := range options {
		qs.Add(q, p)
	}
	base_req_path := strings.TrimRight(c.Config.BaseURL, "/") + path
	u, err := url.Parse(base_req_path)
	if err != nil {
		var data bytes.Buffer
		return data.Bytes(), err
	}
	if len(options) != 0 {
		u.RawQuery = qs.Encode()
	}
	buf := new(bytes.Buffer)
	if body != nil {
		buf.ReadFrom(body)
	}
	req, _ := http.NewRequest(method, u.String(), bytes.NewReader(buf.Bytes()))
	if body != nil {
		h := sha1.New()
		h.Write(buf.Bytes())
		chkSum := h.Sum(nil)
		req.Header.Add("X-Checksum-Sha1", fmt.Sprintf("%x", chkSum))
	}
	req.Header.Add("user-agent", "artifactory-go." + CLIENT_VERSION)
	req.Header.Add("X-Result-Detail", "info, properties")
	req.SetBasicAuth(c.Config.Username, c.Config.Password)
	r, err := c.Client.Do(req)
	if err != nil {
		var data bytes.Buffer
		return data.Bytes(), err
	} else {
		defer r.Body.Close()
		data, err := ioutil.ReadAll(r.Body)
		if r.StatusCode < 200 || r.StatusCode > 299 {
			var ej ErrorsJson
			uerr := json.Unmarshal(data, &ej)
			if uerr != nil {
				return data, &ArtifactoryHTTPError{HttpStatus:r.StatusCode, Body: data}
			} else {
				return data, &ArtifactoryHTTPError{HttpStatus:r.StatusCode, Body: data, Errors: ej}
			}
		} else {
			return data, err
		}
	}
}
