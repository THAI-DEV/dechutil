package dechutil

// * input : []string{"a",""b},"'","'",","  output : 'a','b'
func SliceStringToString(data []string, wrapStr1 string, wrapStr2 string, saperateStr string) string {
	result := ""
	for i, v := range data {
		result = result + wrapStr1 + v + wrapStr2
		if i < len(data)-1 {
			result = result + saperateStr
		}
	}
	// fmt.Println(result)
	return result
}
