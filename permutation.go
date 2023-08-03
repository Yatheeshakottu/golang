package main

func swap(s string, i, j int) string {
	var result []byte
	for k := 0; k < len(s); k++ {
		if k == i {
			result = append(result, s[j])
		} else if k == j {
			result = append(result, s[i])
		} else {
			result = append(result, s[k])
		}
	}
	return string(result)
}
func permutations(str string, i, n int) {
	if i == n-1 {
		println(str)
		return
	}
	for j := i; j < n; j++ {
		str = swap(str, i, j)
		permutations(str, i+1, n)
		str = swap(str, i, j)
	}
}
func main() {
	str := "ABCD"
	permutations(str, 0, len(str))
}
