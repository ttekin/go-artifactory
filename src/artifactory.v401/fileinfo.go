package artifactory

import (
	"encoding/json"
)

func (c *ArtifactoryClient) GetFileInfo(repoKey string, filePath string) (fileInfo *FileInfo, e error) {
	url := "/api/storage/" + repoKey + "/" + filePath
	d, e := c.Get(url, make(map[string]string))
	if e != nil {
		return nil, e
	} else {
		var res FileInfo
		err := json.Unmarshal(d, &res)
		if err != nil {
			return &res, err
		} else {
			return &res, e
		}
	}
}
