package format

func Hex(input string) []byte {
	buf := make([]byte, len(input)*3)
	for i := 0; i < len(input); i += 2 {
		pair := input[i : i+2]
		res := "0x" + pair + " "
		buf = append(buf, []byte(res)...)
	}
	return buf
}
