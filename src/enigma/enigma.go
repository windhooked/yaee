package enigma

type (
	Enigma struct {
		Rotors []Rotor
		Reflector
	}
)

func Assemble() {
	e := &Enigma{
		Rotors: []Rotor{RI, RII, RIII},
	}
}
