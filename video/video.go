package video

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

func AllVideos(c *fiber.Ctx) {
	c.Send("all videos endpoint")
}

func SingleVideos(c *fiber.Ctx) {
	c.Send("single videos endpoint")
}

func NewVideos(c *fiber.Ctx) {
	c.Send("create videos endpoint")
}
