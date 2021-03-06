package numgothai

import "testing"

func TestIntBaht(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "หนึ่งสตางค์"},
		{11, "สิบเอ็ดสตางค์"},
		{100, "หนึ่งบาทถ้วน"},
		{1000, "สิบบาทถ้วน"},
		{101, "หนึ่งบาทหนึ่งสตางค์"},
		{110, "หนึ่งบาทสิบสตางค์"},
		{2020, "ยี่สิบบาทยี่สิบสตางค์"},
		{450022, "สี่พันห้าร้อยบาทยี่สิบสองสตางค์"},
		{4239320, "สี่หมื่นสองพันสามร้อยเก้าสิบสามบาทยี่สิบสตางค์"},
		{9223372036854775807, "เก้าหมื่นสองพันสองร้อยสามสิบสามล้านเจ็ดแสนสองหมื่นสามร้อยหกสิบแปดล้านห้าแสนสี่หมื่นเจ็ดพันเจ็ดร้อยห้าสิบแปดบาทเจ็ดสตางค์"},
		{9223372036854775800, "เก้าหมื่นสองพันสองร้อยสามสิบสามล้านเจ็ดแสนสองหมื่นสามร้อยหกสิบแปดล้านห้าแสนสี่หมื่นเจ็ดพันเจ็ดร้อยห้าสิบแปดบาทถ้วน"},
	}

	for _, testCase := range cases {
		r := IntBaht(testCase.in)
		if r != testCase.want {
			t.Errorf("want %q, got %q\n", testCase.want, r)
		}
	}
}

func BenchmarkIntBaht(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntBaht(8223372236854775811)
	}
}
