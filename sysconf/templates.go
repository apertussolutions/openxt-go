package sysconf

import (
	"encoding/json"

	"github.com/openxt/openxt-go/db"
)

type FileEntry struct {
	Dest string `json:"dest"`
	Mode string `json:"mode"`
	Bind string `json:"bind"`
	ReadOnly string `json:"ro"`
	ReadWrite string `json:"rw"`
}

type Config struct {
	Vars  map[string]string    `json:"vars"`
	Files map[string]FileEntry `json:"files"`
}


func NewConfig(config []byte) (*Config, error) {
	tc := Config{}

	if err := json.Unmarshal(config, &tc); err != nil {
		return nil, err
	}

	return &tc, nil
}

