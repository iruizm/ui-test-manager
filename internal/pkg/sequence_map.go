package pkg

import (
	"sync"
)

type MyData struct {
	MapData map[string]interface{} `json:"map_data"`
	mu      sync.Mutex
}
