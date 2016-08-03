package artifactory

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestGetVersion(t *testing.T) {
	responseFile, err := os.Open("assets/test/version.json")
	if err != nil {
		t.Fatalf("Unable to read test data: %s", err.Error())
	}
	defer responseFile.Close()
	responseBody, _ := ioutil.ReadAll(responseFile)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	conf := &ClientConfig{
		BaseURL:   "http://127.0.0.1:8080/",
		Username:  "username",
		Password:  "password",
		VerifySSL: false,
		Transport: transport,
	}

	client := NewClient(conf)
	version, err := client.GetVersion()
	assert.NoError(t, err, "should not return an error")
	assert.Equal(t, version.Version, "4.0.1")
	assert.Equal(t, version.Revision, "40110")
	assert.Equal(t, version.License, "fb870be54fd3f605fa1e5f036052ca800dd5bac5b")
	assert.Contains(t, version.Addons, "build")
	assert.Contains(t, version.Addons, "debian")
	assert.Contains(t, version.Addons, "docker")
	assert.Contains(t, version.Addons, "filestore")
	assert.Contains(t, version.Addons, "filtered-resources")
	assert.Contains(t, version.Addons, "yum")
}
