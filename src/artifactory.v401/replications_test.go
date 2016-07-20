package artifactory

import (
	"os"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestGetReplicationConfiguration(t *testing.T) {
	responseFile, err := os.Open("assets/test/replication_configuration.json")
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
	config, err := client.GetReplicationConfiguration("docker-dev")
	assert.NoError(t, err)
	assert.Equal(t, len(config), 2)
	assert.Contains(t, []string{"https://repo001.artifactory.io/artifactory/docker-local", "https://repo002.artifactory.io/artifactory/docker-local"}, config[0].URL)
	assert.Contains(t, []string{"https://repo001.artifactory.io/artifactory/docker-local", "https://repo002.artifactory.io/artifactory/docker-local"}, config[1].URL)
}
