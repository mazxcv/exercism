package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("unknow codon")
	ErrInvalidBase = errors.New("invalid")
)

// func FromRNA accept rna and return list of codon
func FromRNA(rna string) ([]string, error) {
	codons := []string{}
	for i := 0; i <= len(rna)-3; i = i + 3 {
		rna_string := rna[i : i+3]
		codon, err := FromCodon(rna_string)
		switch err {
		case ErrStop:
			return codons, nil
		case ErrInvalidBase:
			return codons, ErrInvalidBase
		default:
			codons = append(codons, codon)
		}
	}

	return codons, nil
}

// func FromCodon accept codon and return protein
func FromCodon(codon string) (protein string, err error) {
	codonToProtein := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	protein, ok := codonToProtein[codon]
	if protein == "STOP" {
		err = ErrStop
	}
	if !ok {
		err = ErrInvalidBase
	}
	return
}
