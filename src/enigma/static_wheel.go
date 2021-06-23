package enigma

/*
https://de.wikipedia.org/wiki/Enigma-Walzen#Walzenverdrahtung
ETW = Input Rotor ( Eintrittswalze )

*/
var (
	ETW_Reichsbahn = ETW{W{setting: []byte("QWERTZUIOASDFGHJKPYXCVBNML")}}
	ETW_Enigma1    = ETW{W{setting: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")}}
	ETW_M4         = ETW{W{setting: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")}} //		M4
)

type (
	ETW struct {
		W
	}
)

func init() {
	ETW_M4.Build()
	ETW_Enigma1.Build()
	ETW_Reichsbahn.Build()
}

func (h *W) Encode(in uint8) (out byte) {
	a := h.GetIndex(in)
	return h.GetChar(a)
}
