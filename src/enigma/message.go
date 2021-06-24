package enigma

import (
	"fmt"
)

/*


X = Period / Dot
Y = Comma
UD = Question Mark
XX = Colon
YY = Dash/Hyphen/Slant
KK**KK = Parenthesis
J******J = Stress Mark

SSD= "SEHR SEHR DRINGEND" or "EXTREMELY URGENT"
Note: BINE" (or "BIENE")  - "MUKE" (or "MUECKE") - and "WESPE" are three codewords that were optionally used as substitutes for the urgency sign "SSD".

UUU = Uboat / Submarine
VVV VON = From

NUL NULL = Zero
EINS = One
ZWO ZWEI = Two
DREI = Three
VIR VIER = Four
FUNF FUENF = Five
SECHS SEQS = Six
SIBEN SIEBEN = Seven
ACHT AQT = Eight
NEUN = Nine
*/
// Decode and apply the various substitutions
func Format(m string) string {
	i := 6
	prefix := m[0:i]
	fmt.Printf("%v\n", prefix)

	m = m[i:]
	n := 5
	for k := 0; k < len(m); k += n {
		if (k + 5) > len(m) {
			n = len(m) - k
		}
		fmt.Printf("%v\n", m[k:k+n])
	}
	return ""

	//	re := regexp.MustCompile(`(\S{3})`)
	//	x := re.FindAllStringSubmatch(m, -1)

}
