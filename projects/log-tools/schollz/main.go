package main

import (
	log "github.com/schollz/logger"
)

func main() {
	log.SetLevel("debug")
	log.Debug("debug")
	log.Info("info")
	log.Error("error")
}
