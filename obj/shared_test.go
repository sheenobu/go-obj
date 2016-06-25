package obj

// shared test code

type stringList []string

func (sl stringList) ToByteList() (arr [][]byte) {
	for _, s := range sl {
		arr = append(arr, []byte(s))
	}
	return
}

func compareVertices(v1 Vertex, v2 Vertex) bool {
	return v1.Index == v2.Index && v1.X == v2.X && v1.Y == v2.Y && v1.Z == v2.Z
}

func compareNormals(n1 Normal, n2 Normal) bool {
	return n1.Index == n2.Index && n1.X == n2.X && n1.Y == n2.Y && n1.Z == n2.Z
}

func compareErrors(err error, expected string) bool {
	if expected == "" && err != nil {
		return false
	} else if err != nil && expected != err.Error() {
		return false
	}
	return true
}
