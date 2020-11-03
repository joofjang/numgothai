package numgothai

var numToText = map[uint8]string{
	0: "ศูนย์",
	1: "หนึ่ง",
	2: "สอง",
	3: "สาม",
	4: "สี่",
	5: "ห้า",
	6: "หก",
	7: "เจ็ด",
	8: "แปด",
	9: "เก้า",
}

// IntText create new string with Thai text from integer.
func IntText(v int) string {

	if v == 0 {
		return numToText[0]
	}

	s := ""

	isOnotherOfOne := false // If 1 is spelled "ed".

OuterLoop:
	for {
		switch {

		case v >= 1000000:
			s = s + IntText(v/1000000) + "ล้าน"
			v = v % 1000000

		case v >= 100000:
			s = s + numToText[uint8(v/100000)] + "แสน"
			v = v % 100000

		case v >= 10000:
			s = s + numToText[uint8(v/10000)] + "หมื่น"
			v = v % 10000

		case v >= 1000:
			s = s + numToText[uint8(v/1000)] + "พัน"
			v = v % 1000

		case v >= 100:
			s = s + numToText[uint8(v/100)] + "ร้อย"
			v = v % 100

		case v >= 10:
			isOnotherOfOne = true
			switch {
			case v/10 == 1:
				s = s + "สิบ"
			case v/10 == 2:
				s = s + "ยี่สิบ"
			default:
				s = s + numToText[uint8(v/10)] + "สิบ"
			}
			v = v % 10

		default:
			switch {
			case v != 0 && !(v == 1 && isOnotherOfOne):
				s = s + numToText[uint8(v)]
			case v == 1 && isOnotherOfOne:
				s = s + "เอ็ด"
			}
			break OuterLoop
		}
	}

	return s
}
