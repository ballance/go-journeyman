package func main() {
	s := "世界 means world"

	var buf [utf8]
	
	for i, r, := range s {
		rl := utf8.RuneLen(r)

		si := i + rl

		copy(buf[:], s[i:si])

	}

}