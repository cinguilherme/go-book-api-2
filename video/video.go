package video

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	Name     string          `json:"name"`
	Source   string          `json:"source"`
	Length   float32         `json:"length"`
	Free     bool            `json:"free"`
	Metadata json.RawMessage `json:"metadata"`
}
