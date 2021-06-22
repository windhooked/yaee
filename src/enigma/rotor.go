package enigma

/*
http://users.telenet.be/d.rijmenants/en/enigmatech.htm
//https://de.wikipedia.org/wiki/Enigma-M4

            A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
I           E K M F L G D Q V Z N T O W Y H X U S P A I B R C J
II          A J D K S I R U X B L H W T M C Q G Z N P Y F V O E
III         B D F H J L C P R T X V Z N Y E I W G A K M U S Q O
IV          E S O V P Z J A Y Q U I R H X L N F T G K D C M W B
V           V Z B R G I T Y U P S D N H L X A W M J Q O F E C K
VI          J P G V O U M F Y Q B E N H Z R D K A S X L I C T W
VII         N Z J H G R C X M Y S W B O U F A I V L P E K Q D T
VIII        F K Q H T L X O C B J S P D Z R A M E W N I U Y G V

            A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
Beta        L E Y J V C N I X W P B Q M D R T A K Z G F U H O S
Gamma       F S O K A N U E R H M B T I Y C W L Q P Z X V G J D

*/
var (
	W_I   = Rotor{Map: "EKMFLGDQVZNTOWYHXUSPAIBRCJ", Notch: 'Y', Window: 'Q'}
	W_II  = Rotor{Map: "AJDKSIRUXBLHWTMCQGZNPYFVOE", Notch: 'M', Window: 'E'}
	W_III = Rotor{Map: "BDFHJLCPRTXVZNYEIWGAKMUSQO", Notch: 'D', Window: 'V'}
	W_IV  = Rotor{Map: "ESOVPZJAYQUIRHXLNFTGKDCMWB", Notch: 'R', Window: 'J'}
	W_V   = Rotor{Map: "VZBRGITYUPSDNHLXAWMJQOFECK", Notch: 'H', Window: 'Z'}
)

type (
	//Walze
	Rotor struct {
		Map    []byte // Wire routes
		Notch  byte   // Notch position
		Window byte
		index  uint8
	}
)

func (h *Rotor) Route(in rune) (out rune, notch bool) {
	notch = false
	if index <= numChars {
		out := h.Map[h.index]
		h.index++
		if out == h.notch {
			notch = true
		}
	} else if h.index < 0 {
		// should never be
		h.index = 0

	} else if index > numChars {
		// start over
		h.index = 0
	}
}
