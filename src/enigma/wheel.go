package enigma

// the wheel is a generic pattern used by any coding component in the enigma
// the coding components: static input wheel, rotors, plugboard and reflector, all
// have a input to output code associated, translating 26 input to 26 output
const CharacterSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const CharacterSetCount = len(CharacterSet)

type (
	W struct {
		setting []byte

		lutIndexToChar []byte
		lutCharToIndex []byte

		lutIn  [CharacterSetCount]byte // right to left forward
		lutOut [CharacterSetCount]byte // left to right reverse
	}
)

func NewWheel(s string) *W {
	w := &W{setting: []byte(s)}
	w.Build()
	return w
}

func (h *W) Build() {
	h.lutIndexToChar = make([]uint8, 255) //numChars+1)
	h.lutCharToIndex = make([]uint8, 255) //numChars+1)

	for k, v := range h.setting {

		h.lutIndexToChar[k] = v        // forward mapping
		h.lutCharToIndex[v] = uint8(k) // reverse mapping

		index := v - 'A'           //index of letter
		h.lutIn[k] = index         //index of letter
		h.lutOut[index] = uint8(k) //index of letter

	}
}

func (h *W) Encode(in uint8) (out uint8) {
	return h.lutIn[in]
}

func (h *W) Decode(in uint8) (out uint8) {
	return h.lutOut[in]
}

// GetIndex returns the index for given character
// eg. A == 1, Z == 26
func (h *W) GetIndex(in byte) (out uint8) {
	return h.lutCharToIndex[in]
}

// GetChar returns the character for a given index
// eg. 1 == A, 26 == Z
func (h *W) GetChar(in uint8) (out byte) {
	return h.lutIndexToChar[in]
}

//func (h *W) DecodeOffset(in byte, offset uint8) (out byte) {
//	return h.lutIn[h.lutOut[in]]
//}
func (h *W) Lut() []byte {
	return h.setting
}
