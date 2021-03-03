package gopwn

import (
	"io"
	"log"
	"net"
	"os/exec"
)

// connect to remote host.
func Remote(address string) *AppProcess{
	c, err := net.Dial("tcp", address)

	if err != nil{
		log.Fatalf("Remote Connection Error %s", err)
	}

	return NewProcess(c,c)
}

func Process(binaryName string) *AppProcess{
	if !FileExists(binaryName){
		log.Fatalf("File %s dooesn't exist.", binaryName)
	}
	cmd := exec.Command(binaryName)

	stderr, err := cmd.StderrPipe()
	if err != nil{
		log.Fatal("Stderr error: ",err)
	}

	stdin, err := cmd.StdinPipe()
	if err != nil{
		log.Fatal("Stdin error: ",err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil{
		log.Fatal("Stdout error: ",err)
	}

	err = cmd.Start()
	if err != nil{
		log.Fatal("Start error: ",err)
	}

	return NewProcess(io.MultiReader(stdout, stderr), stdin)
}