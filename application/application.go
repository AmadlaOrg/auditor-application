package application

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
)

type IApplication interface {
	FindApplication(name string, isContainer bool) (Application, error)
}

type SApplication struct{}

// FindApplication
func (s *SApplication) FindApplication(name string, isContainer bool) (Application, error) {
	if runtime.GOOS == "linux" {
		_, err := s.linuxLocal(name)
		if err != nil {
			return Application{}, err
		}
	} else if runtime.GOOS == "windows" {
		_, err := s.windowsLocal(name)
		if err != nil {
			return Application{}, err
		}
	}
	return Application{}, nil
}

// linuxLocal
func (s *SApplication) linuxLocal(name string) (string, error) {
	if _, err := exec.LookPath("which"); err != nil {
		return "", errors.New("the command 'which' is not available on this system")
	}

	isValidName := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString
	if !isValidName(name) {
		return "", fmt.Errorf("invalid command name: %s", name)
	}

	cmd := exec.Command("which", name)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not run command 'which %s': %v", name, err)
	}

	return string(out), nil
}

// windowsLocal checks if a given command exists in the system's PATH on Windows.
func (s *SApplication) windowsLocal(name string) (string, error) {
	// Check if 'where' command is available
	if _, err := exec.LookPath("where"); err != nil {
		return "", errors.New("the command 'where' is not available on this system")
	}

	// Validate the command name
	isValidName := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString
	if !isValidName(name) {
		return "", fmt.Errorf("invalid command name: %s", name)
	}

	// Execute the 'where' command
	cmd := exec.Command("where", name)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not run command 'where %s': %v", name, err)
	}

	return string(out), nil
}

// container checks if a container image exists locally and remotely using the specified container engine (Docker or Podman).
func (s *SApplication) container(name string, containerEngine string) (ContainerApplication, error) {
	var result ContainerApplication

	// Verify the container engine (Docker or Podman)
	if containerEngine != "docker" && containerEngine != "podman" {
		return result, fmt.Errorf("unsupported container engine: %s", containerEngine)
	}

	// Check if the container engine is available on the system
	if _, err := exec.LookPath(containerEngine); err != nil {
		return result, fmt.Errorf("the container engine '%s' is not available on this system", containerEngine)
	}

	// Check if the image exists remotely by pulling it with a dry-run or checking if it can be pulled
	existsCmd := exec.Command(containerEngine, "pull", "--quiet", name)
	if err := existsCmd.Run(); err == nil {

		// TODO: Check if the HTTP Status error is 200
		result.Exists = true
		result.Source = fmt.Sprintf("https://hub.docker.com/r/%s", name) // Placeholder URL, adjust if necessary
	} else {
		result.Exists = false
		result.Source = ""
	}

	// Check if the image is available locally
	haveCmd := exec.Command(containerEngine, "image", "inspect", name)
	if err := haveCmd.Run(); err == nil {
		result.Have = true
	} else {
		result.Have = false
	}

	return result, nil
}

// containerImages
func containerImages(containerEngineName string) (string, error) {
	isValidContainerEngineName := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString
	if !isValidContainerEngineName(containerEngineName) {
		return "", fmt.Errorf("invalid container name: %s", containerEngineName)
	}

	cmd := exec.Command(containerEngineName, "images", "--format", "json")

	// Redirect stderr to /dev/null to suppress error messages
	cmd.Stderr = nil

	// Capture stdout output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute podman images: %v", err)
	}

	// Return the output as a string
	return out.String(), nil
}
