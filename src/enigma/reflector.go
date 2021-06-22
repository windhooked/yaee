package enigma

/*
https://de.wikipedia.org/wiki/Enigma-Walzen#Walzenverdrahtung

ETW	ABCDEFGHIJKLMNOPQRSTUVWXYZ		M4
UKW Bruno	ENKQAUYWJICOPBLMDXZVFTHRGS	1. Februar 1942	M4
UKW Cäsar	RDOBJNTKVEHMLFCWZAXGYIPSUQ	1. Juli 1943	M4
Beta	LEYJVCNIXWPBQMDRTAKZGFUHOS	1. Februar 1942	M4
Gamma	FSOKANUERHMBTIYCWLQPZXVGJD	1. Juli 1943	M4

ETW = Input Rotor
UKW = Reflector

*/
var (
	UKW_A     = UKW{lut: []byte("EJMZALYXVBWFCRQUONTSPIKHGD")}
	UKW_B     = UKW{lut: []byte("YRUHQSLDPXNGOKMIEBFZCWVJAT")} // 	2. November 193"
	UKW_C     = UKW{lut: []byte("FVPJIAOYEDRZXWGCTKUQSBNMHL")} // 	1940"
	UKW_Bruno = UKW{lut: []byte("YRUHQSLDPXNGOKMIEBFZCWVJAT")}
	UKW_Casar = UKW{lut: []byte("RDOBJNTKVEHMLFCWZAXGYIPSUQ")}
	UKW_Beta  = UKW{lut: []byte("LEYJVCNIXWPBQMDRTAKZGFUHOS")}
	UKW_Gamma = UKW{lut: []byte("FSOKANUERHMBTIYCWLQPZXVGJD")}
)

type (
	UKW struct {
		lut []byte
	}
)

func (h *UKW) Encode(in byte) (out byte) {
	return h.lut[in-'A']
}
