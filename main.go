package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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

func setBrightness(value int) error {
	_, err := runCmd("brightnessctl", "set", fmt.Sprintf("%d", value))
	return err
}

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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "dim":
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
