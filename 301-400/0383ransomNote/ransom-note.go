package ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	b := make([]int, 256)
	for i := 0; i < len(magazine); i++ {
		b[int(magazine[i])]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if b[int(ransomNote[i])] != 0 {
			b[int(ransomNote[i])]--
		} else {
			return false
		}
	}
	return true
}
