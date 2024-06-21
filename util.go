package jsondt

func isEmpty(b []byte) bool {
	return (len(b) == 2 && b[0] == '"' && b[1] == '"') ||
		(len(b) == 4 && b[0] == 'n' && b[1] == 'u' && b[2] == 'l' && b[3] == 'l')
}
