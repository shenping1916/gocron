package gocron

import "time"

type config struct {
	local  *time.Location

	log  logger
}
