package artifactory

import "encoding/json"

type Version struct {
	Version string `json:"version"`
	Revision string `json:"revision"`
	Addons []string `json:"addons"`
	License string `json:"license"`
}

func (c *ArtifactoryClient) GetVersion() (Version, error) {
	var res Version
	d, e := c.Get("/api/system/version", make(map[string]string))
	if e != nil {
		return res, e
	} else {
		err := json.Unmarshal(d, &res)
		if err != nil {
			return res, err
		} else {
			return res, e
		}
	}
}
