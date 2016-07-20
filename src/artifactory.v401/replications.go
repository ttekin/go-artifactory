package artifactory

import (
	"encoding/json"
)

type ReplicationConfiguration struct {
	URL string `json:"url"`
	SocketTimeoutMillis int `json:"socketTimeoutMillis"`
	Username string `json:"username"`
	Password string `json:"password"`
	EnableEventReplication bool `json:"enableEventReplication"`
	Enabled bool `json:"enabled"`
	CronExp string `json:"cronExp"`
	SyncDeletes bool `json:"syncDeletes"`
	SyncProperties bool `json:"syncProperties"`
	RepoKey string `json:"repoKey"`
}

func (c *ArtifactoryClient) GetReplicationConfiguration(repoKey string) ([]ReplicationConfiguration, error) {
	res := make([]ReplicationConfiguration, 0)
	d, e := c.Get("/api/replications/" + repoKey, make(map[string]string))
	if e != nil {
		return nil, e
	} else {
		err := json.Unmarshal(d, &res)
		if err != nil {
			return nil, err
		} else {
			return res, e
		}
	}
}
