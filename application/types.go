package application

type Application struct {
	Name      string
	Path      string
	Container ContainerApplication
}

type ContainerApplication struct {
	Exists bool
	Have   bool
	Source string
}

type ContainerImagesList struct {
	Id          string   `json:"Id"`          // E.g.: 778af9bd63976fc86be91ab164abe68d1f3ab73300dc38735b9210c74e6b1cb4
	ParentId    string   `json:"ParentId"`    // E.g.: ""
	RepoTags    string   `json:"RepoTags"`    // E.g.: null
	RepoDigests []string `json:"RepoDigests"` // E.g.: ["docker.io/rediscommander/redis-commander@sha256:19cd0c49f418779fa2822a0496c5e6516d0c792effc39ed20089e6268477e40a"]
	Size        int64    `json:"Size"`        // E.g.: 84653752
	SharedSize  int64    `json:"SharedSize"`  // E.g.: 0
	VirtualSize int64    `json:"VirtualSize"` // E.g.: 84653752
	Labels      []any    `json:"Labels"`      // E.g.: map (all strings)
	Containers  int64    `json:"Containers"`  // E.g.: 1
	Digest      string   `json:"Digest"`      // E.g.: "sha256:b2177a8bfe85f89ff403c9f51b8a00a6efd1be8e475bc2637390c36977df994d"
	History     []string `json:"History"`     // E.g.: ["docker.io/hashicorp/vault:latest"]
	Names       []string `json:"Names"`       // E.g.: ["docker.io/hashicorp/vault:latest"]
	Created     int64    `json:"Created"`     // E.g.: 1687190227
	CreatedAt   string   `json:"CreatedAt"`   // E.g.: "2023-06-19T15:57:07Z"
}
