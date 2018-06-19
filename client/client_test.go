package client

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"
)

const repo = "debarc/golang-server"
const port = 8000

func ExecuteCommand(t *testing.T, command string, logString string) {
	t.Helper()
	cmd := exec.Command("/bin/sh", "-c", command)
	fmt.Println(logString)
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSlashIndexEndpointForGolangServer(t *testing.T) {
	ExecuteCommand(t, fmt.Sprintf("docker run -d -p %d:%d %v", port, port, repo), "Running container from built image...")
	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("curl localhost:%d/index", port)).Output()
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput, _ := ioutil.ReadFile("../public/index.html")
	if string(output) != string(expectedOutput) {
		t.Errorf("expected %v got %v", string(expectedOutput), string(output))
	}
	ExecuteCommand(t, fmt.Sprintf(`docker kill $(docker ps | grep "%v" | awk '{print $1}')`, repo), "Killing image...")
}

func TestSlashEndpointForGolangServer(t *testing.T) {
	ExecuteCommand(t, fmt.Sprintf("docker run -d -p %d:%d %v", port, port, repo), "Running container from built image")
	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("curl localhost:%d/", port)).Output()
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput := "404 page not found\n"
	if string(output) != expectedOutput {
		t.Errorf("expected %v got %v", expectedOutput, string(output))
	}
	ExecuteCommand(t, fmt.Sprintf(`docker kill $(docker ps | grep "%v" | awk '{print $1}')`, repo), "Killing image...")
}
