package numgothai

// IntBaht create new string in Thai currency text. Value must be an integer represent Thai Baht with Satang.
// Digits before 2 last position of value is Baht and the 2 last position of value is Satang.
func IntBaht(v int) string {

	s := ""
	b := v / 100 //Baht part
	if b != 0 {
		s = IntText(int(b)) + "บาท"
	}
	st := v % 100 //Satang part

	switch {
	case b == 0 && st == 0: //0 Baht
		s = numToText[0] + "บาท"
	case b != 0 && st == 0: //No Satang
		s = s + "ถ้วน"
	default:
		s = s + IntText(int(st)) + "สตางค์"
	}

	return s
}
