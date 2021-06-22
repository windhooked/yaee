package enigma

import "strings"

/*
https://de.wikipedia.org/wiki/Enigma-Walzen#Walzenverdrahtung
ETW = Input Rotor ( Eintrittswalze )

*/
var (
	ETW_Reichsbahn = ETW{lut: []byte("QWERTZUIOASDFGHJKPYXCVBNML")}
	ETW_Enigma1    = ETW{lut: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")}
	ETW_M4         = ETW{lut: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")} //		M4
)

type (
	ETW struct{ lut []byte }
)

func (h *ETW) Encode(in byte) (out byte) {
	return h.lut[in-'A']
}

func (h *ETW) Decode(in byte) (out byte) {
	k := strings.Index(string(h.lut), string(in))
	return h.lut[k]
}
func (h *ETW) Lut() []byte {
	return h.lut
}
