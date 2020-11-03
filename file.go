package snap

import "time"

const (
	TYPE_DIR  = 0
	TYPE_FILE = 1
)

type File struct {
	Uuid  string `json:"uuid"`
	PUuid string `json:"puuid"`
	Name  string `json:"name"`
	Md5   string `json:"md5"`
	Size  int64  `json:"size"`
	Times int64  `json:"times"`

	UpdateTime time.Time `json:"updateTime"`
	CreateTime time.Time `json:"createTime"`

	Type int8 `json:"type"`
}
