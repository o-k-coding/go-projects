package complementary_dna

func CreateComplements() map[rune] rune {
	complements := make(map[rune]rune)
	complements['A'] = 'T'
	complements['T'] = 'A'
	complements['G'] = 'C'
	complements['C'] = 'G'
	return complements
}

func DNAStrand(dna string) string {
	complements := CreateComplements()
	complementaryDna := make([]rune, len(dna))
  // your code here
	for i, base := range dna {
		complementaryDna[i] = complements[base]
	}
	return string(complementaryDna)
}
