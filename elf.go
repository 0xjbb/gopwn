package gopwn

import "log"

type elf struct{
	symbols[] []byte
}

func ELF(binaryName string) *elf{

	if !FileExists(binaryName){
		log.Fatalf("File %s dooesn't exist.", binaryName)
	}


	elf := &elf{

	}
	elf.parse()

	return elf
}

func (e *elf) parse(){

}