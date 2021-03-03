package gopwn

import (
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

func Local(binaryName string){
	if !FileExists(binaryName){
		log.Fatalf("File %s dooesn't exist.", binaryName)
	}
	cmd := exec.Command(binaryName)






}