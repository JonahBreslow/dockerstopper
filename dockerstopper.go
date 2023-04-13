package main

import (
    "fmt"
    "os"
    "os/exec"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    validateInputs()
    input := os.Args[1]
    indices := parseInputIndices(input, []int{})
    containerIDs := getContainersToStop(indices, []string{})
    stopContainers(containerIDs)
}

func validateInputs() {
	if len(os.Args) < 2 {
		fmt.Println("No input provided")
		os.Exit(1)
	}
}

func stopContainers(containerIDs []string) {
	for _, containerID := range containerIDs {
		cmd := exec.Command("docker", "stop", containerID)
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error stopping container %s: %v\n", containerID, err)
			continue
		}
		cmd = exec.Command("docker", "rm", "-f", containerID)
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error removing container %s: %v\n", containerID, err)
			continue
		}
		fmt.Printf("Stopped container: %s\n", containerID)
	}
}

func getContainersToStop(indices []int, containerIDs []string) []string {
	out, err := exec.Command("docker", "ps", "-q").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i, s := range strings.Split(strings.Trim(string(out), "\n"), "\n") {
		if contains(indices, i+1) {
			containerIDs = append(containerIDs, strings.TrimSpace(s))
		}
	}
	return containerIDs
}

func parseInputIndices(input string, indices []int) []int {
	if matched, _ := regexp.MatchString("^[0-9]+$", input); matched {
		return append(indices, atoi(input))
	}
	if matched, _ := regexp.MatchString("^([0-9]+(,[0-9]+)*)$", input); matched {
		parts := strings.Split(input, ",")
		for _, s := range parts {
			indices = append(indices, atoi(s))
		}
		return indices
	}
	fmt.Printf("Invalid input: %s\n", input)
	os.Exit(1)
	return nil  // unreachable code
}

func atoi(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func contains(a []int, x int) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}
