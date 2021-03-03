package gopwn

import (
	"bufio"
	"io"
)

type AppProcess struct{
	io.Reader
	io.Writer
	BufReader *bufio.Reader
}

func NewProcess(r io.Reader, w io.Writer) *AppProcess{
	return &AppProcess{
		 Reader: r,
		 Writer: w,
		 BufReader: bufio.NewReader(r),
	}
}


func (ap *AppProcess) Recvline(){}
func (ap *AppProcess) Sendline(data []byte){}

func (ap *AppProcess) write(data []byte){}
func (ap *AppProcess) interactive(){}