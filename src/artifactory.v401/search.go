package artifactory

import (
	"encoding/json"
	"strings"
	"time"
	"strconv"
)

type Gavc struct {
	GroupID    string
	ArtifactID string
	Version    string
	Classifier string
	Repos      []string
}

func (c *ArtifactoryClient) GAVCSearch(coords *Gavc) (files []FileInfo, e error) {
	url := "/api/search/gavc"
	params := make(map[string]string)
	if len(coords.GroupID) > 0 {
		params["g"] = coords.GroupID
	}
	if len(coords.ArtifactID) > 0 {
		params["a"] = coords.ArtifactID
	}
	if len(coords.Version) > 0 {
		params["v"] = coords.Version
	}
	if len(coords.Classifier) > 0 {
		params["c"] = coords.Classifier
	}
	if len(coords.Repos) > 0 {
		params["repos"] = strings.Join(coords.Repos, ",")
	}
	d, err := c.Get(url, params)
	if err != nil {
		return files, err
	} else {
		var dat SearchResults
		err := json.Unmarshal(d, &dat)
		if err != nil {
			return files, err
		} else {
			files = dat.Results
			return files, nil
		}
	}
}

type DateSearchParams struct {
	From time.Time
	To time.Time
	Repos []string
	SearchCreated bool
	SearchLastModified bool
	SearchLastDownloaded bool
}

func (c *ArtifactoryClient) DateSearch(searchParams *DateSearchParams) (files []FileInfo, e error) {
	url := "/api/search/dates"
	params := make(map[string]string)
	if !searchParams.From.IsZero() {
		params["from"] = strconv.FormatInt(searchParams.From.Unix() * 1000, 10)
	}
	if !searchParams.To.IsZero() {
		params["to"] = strconv.FormatInt(searchParams.To.Unix() * 1000, 10)
	}
	if len(searchParams.Repos) > 0 {
		params["repos"] = strings.Join(searchParams.Repos, ",")
	}
	datefields := make([]string, 0, 3)
	if searchParams.SearchCreated {
		datefields = append(datefields, "created")
	}
	if searchParams.SearchLastModified {
		datefields = append(datefields, "lastModified")
	}
	if searchParams.SearchLastDownloaded {
		datefields = append(datefields, "lastDownloaded")
	}
	if len(datefields) > 0 {
		params["dateFields"] = strings.Join(datefields, ",")
	}
	d, err := c.Get(url, params)
	if err != nil {
		return files, err
	} else {
		var dat SearchResults
		err := json.Unmarshal(d, &dat)
		if err != nil {
			return files, err
		} else {
			files = dat.Results
			return files, nil
		}
	}
}
