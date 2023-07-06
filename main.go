package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Define the path to the file that stores the brightness
const brightnessFile = "/tmp/brightness"

func runCmd(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.Output()
	return string(stdout), err
}

// setBrightness sets the brightness to the given value
func setBrightness(value int) error {
	_, err := runCmd("brightnessctl", "set", fmt.Sprintf("%d", value))
	return err
}

// getBrightness gets the current brightness
func getBrightness() (int, error) {
	brightness, err := runCmd("brightnessctl", "g")
	if err != nil {
		return 0, err
	}
	brightness = strings.TrimSpace(brightness)
	brightnessInt, err := strconv.Atoi(brightness)
	if err != nil {
		return 0, err
	}
	return brightnessInt, nil
}

// isVideoPlaying checks if a video is playing on any of the players
func isVideoPlaying() bool {
	players, err := runCmd("playerctl", "-l")
	if err != nil {
		fmt.Printf("Error running playerctl: %v\n", err)
		return false
	}
	// I'm forced to use a regex because playerctl returns a newline character even if there is no output
	re := regexp.MustCompile(`\S`) // (Matches any non-whitespace character)

	for _, player := range strings.Fields(players) {
		// Get the metadata for the current player to check if it's playing a video (If there is an album it is a song, if there is no album it is a video)
		metadata, err := runCmd("playerctl", "-p", player, "metadata", "xesam:album")

		// Basic error handling
		if err != nil {
			fmt.Printf("Error running playerctl: %v\n", err)
			return false
		}

		// If there is no album, it is a video
		if !re.MatchString(metadata) {
			fmt.Printf("Video is playing on %s\n", player)
			return true // If a video is playing, return true
		}

		// If there is an album, it is a song
		if re.MatchString(metadata) {
			fmt.Printf("Audio is playing on %s\n", player)
		}
	}

	return false // If no media is playing, or if audio is playing, return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "dim":
		// Use the isVideoPlaying function to check if a video is playing before dimming
		if isVideoPlaying() {
			fmt.Println("Video is playing, not dimming.")
			return
		}
		currentBrightness, err := getBrightness()
		if err != nil {
			fmt.Printf("Error getting brightness: %v\n", err)
			os.Exit(1)
		}

		file, err := os.Create(brightnessFile)
		if err != nil {
			fmt.Printf("Error creating brightness file: %v\n", err)
			os.Exit(1)
		}
		_, err = io.WriteString(file, fmt.Sprintf("%d", currentBrightness))
		file.Close()

		if err != nil {
			fmt.Printf("Error writing to brightness file: %v\n", err)
			os.Exit(1)
		}

		newBrightness := currentBrightness / 10
		err = setBrightness(newBrightness)
		if err != nil {
			fmt.Printf("Error setting brightness: %v\n", err)
			os.Exit(1)
		}
	case "restore":
		data, err := os.ReadFile(brightnessFile)
		if err != nil {
			fmt.Printf("Error reading brightness file: %v\n", err)
			os.Exit(1)
		}

		restoreBrightness, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Printf("Error parsing brightness data: %v\n", err)
			os.Exit(1)
		}

		err = setBrightness(restoreBrightness)
		if err != nil {
			fmt.Printf("Error restoring brightness: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Invalid argument. Please use 'dim' or 'restore'")

		os.Exit(1)
	}
}
