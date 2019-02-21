package strand

import (
	"fmt"
)

var transcription = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA translates DNA sequence to RNA sequence
func ToRNA(dna string) (rna string) {
	for _, d := range dna {
		r, ok := transcription[d]
		if !ok {
			panic(fmt.Errorf("Wrong sequence value %v", r))
		}
		rna = rna + string(r)
	}
	return
}
