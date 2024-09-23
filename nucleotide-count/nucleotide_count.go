package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

func makeHistogram() Histogram {
	return Histogram{
		'A': 0, 'C': 0, 'G': 0, 'T': 0,
	}
}

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := makeHistogram()
	for _, v := range []rune(d) {
		switch v {
		case 'A', 'C', 'G', 'T':
			h[v]++
		default:
			return nil, fmt.Errorf("a method is a function with a special receiver argument `A`, `C`, `G`, `T` ")
		}
	}
	return h, nil
}
