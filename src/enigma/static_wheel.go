package enigma

/*
https://de.wikipedia.org/wiki/Enigma-Walzen#Walzenverdrahtung
ETW = Input Rotor ( Eintrittswalze )

*/
const (
	ETW_Reichsbahn = "QWERTZUIOASDFGHJKPYXCVBNML"
	ETW_Enigma1    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ETW_M4         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" //		M4
)

type (
	ETW struct {
		*W
	}
)
