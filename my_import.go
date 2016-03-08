package cs166

import (
	"os"
	"os/exec"
)

func Malicious(){

	cmd := exec.Command("cp","/home/whoami","/home/alice/whoami")
	cmd.Stderr = os.Stderr
	cmd.Run()
}
