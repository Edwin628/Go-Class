package sfu

import (
	"log"
	"os/exec"
)

func StartRecording(streamURL, outputFile string) {
	cmd := exec.Command("ffmpeg", "-i", streamURL, "-c", "copy", outputFile)
	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start recording: %v", err)
	}

	log.Println("Recording started")
}
