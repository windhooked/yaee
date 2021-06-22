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
	UKW_A     = UKW{W{setting: []byte("EJMZALYXVBWFCRQUONTSPIKHGD")}}
	UKW_B     = UKW{W{setting: []byte("YRUHQSLDPXNGOKMIEBFZCWVJAT")}} // 	2. November 193"
	UKW_C     = UKW{W{setting: []byte("FVPJIAOYEDRZXWGCTKUQSBNMHL")}} // 	1940"
	UKW_Bruno = UKW{W{setting: []byte("YRUHQSLDPXNGOKMIEBFZCWVJAT")}}
	UKW_Casar = UKW{W{setting: []byte("RDOBJNTKVEHMLFCWZAXGYIPSUQ")}}
	UKW_Beta  = UKW{W{setting: []byte("LEYJVCNIXWPBQMDRTAKZGFUHOS")}}
	UKW_Gamma = UKW{W{setting: []byte("FSOKANUERHMBTIYCWLQPZXVGJD")}}
)

type (
	UKW struct {
		W
	}
)

func init() {
	UKW_A.Build()
	UKW_B.Build()
	UKW_C.Build()
}

func (h *UKW) Encode(in byte) (out byte) {
	return h.W.Encode(in)
}
