package enigma

import "fmt"

/*
https://de.wikipedia.org/wiki/Enigma-M4

----------------------------------------------------------------------
| Mo-  |                                                    | Grund- |
| nats-|      S t e c k e r v e r b i n d u n g e n         | stel-  |
| tag  |                                                    | lung   |
----------------------------------------------------------------------
|  30. |18/26 17/4 21/6 3/16 19/14 22/7 8/1 12/25 5/9 10/15 |H F K D |
|  29. |20/13 2/3 10/4 21/24 12/1 6/5 16/18 15/8 7/11 23/26 |O M S R |
|  28. |9/14 4/5 18/24 3/16 20/26 23/21 12/19 13/2 22/6 1/8 |E Y D X |
|  27. |16/2 25/21 6/20 9/17 22/1 15/4 18/26 8/23 3/14 5/19 |T C X K |
|  26. |20/13 26/11 3/4 7/24 14/9 16/10 8/17 12/5 2/6 15/23 |Y S R B |
*/
var (

	//|  30. |18/26 17/4 21/6 3/16 19/14 22/7 8/1 12/25 5/9 10/15 |H F K D |
	PB_30 = [][]byte{{18, 26}, {17, 4}, {21, 6}, {3, 16}, {19, 14}, {22, 7}, {8, 1}, {12, 25}, {5, 9}, {10, 15}}
)

type (
	//Steckerbrett
	SB struct {
		Setting [][]byte
		lut     [len(CharacterSet)]uint8 // lookup table
	}
)

func NewPlugboard(pb [][]byte) *SB {
	h := new(SB)
	h.Setting = pb
	h.WireUp()
	return h
}

func (h *SB) WireUp() {
	for _, v := range h.Setting {
		if len(v) == 2 {
			if v[0] > 26 { // assume characters, translate to index
				v[0] = v[0] - 'A'
				v[1] = v[1] - 'A'
			}
		} else {
			fmt.Printf("Error decoding plugboard setting %v", v)
		}
		// on 26 chars the lower 13 routes to upper 13 and vice versa
		h.lut[v[0]] = v[1]
		h.lut[v[1]] = v[0]
	}
}

// Encode on the plugboard reroutes if plug wires are inserted, else input bridges to output
func (h *SB) Encode(n uint8) (out uint8) {
	out = h.lut[n]
	return
}
func (h *SB) Decode(n uint8) (out uint8) {
	out = h.lut[n]
	return
}
