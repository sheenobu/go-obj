package obj

func splitByToken(bl []byte, token byte) (arr [][]byte) {
	var idx int
	var b byte
	l := 0
	for idx, b = range bl {
		if b == token {
			arr = append(arr, bl[l:idx])
			l = l + len(bl[l:idx]) + 1
		}
	}

	arr = append(arr, bl[l:idx+1])
	return
}
