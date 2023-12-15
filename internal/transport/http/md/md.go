package md

import (
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"openapi/internal/logger"
	"time"
)

func Log(c *fiber.Ctx) error {

	if c.OriginalURL() == "/favicon.ico" {
		return nil
	}

	source := c.Get("source", "")
	trackId := c.Get("track-id", "")

	if trackId == "" {
		seed := uuid.NewV4()
		UUID := uuid.NewV5(seed, "track-id").String()
		trackId = UUID
	}

	startTime := time.Now()
	// 调用处理程序之前的逻辑
	logger.Logger.Infof("HTTP Source: %s, TrackId: %s, Method: %s, started at: %s, Request: %s\n", source, trackId, c.OriginalURL(), startTime.Format(time.RFC3339), string(c.Request().Body()))
	err := c.Next()
	// 调用处理程序之后的逻辑
	logger.Logger.Infof("HTTP Source: %s, TrackId: %s, Method: %s, completed in: %s Response: %v, Error: %v\n", source, trackId, c.OriginalURL(), time.Since(startTime), string(c.Response().Body()), err)
	return err
}
