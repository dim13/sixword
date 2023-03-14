package sixword

import "testing"

func TestSixWord(t *testing.T) {
	testCases := []struct {
		n uint64
		s string
	}{
		{0x9e876134d90499dd, "INCH SEA ANNE LONG AHEM TOUR"},
		{0x7965e05436f5029f, "EASE OIL FUM CURE AWRY AVIS"},
		{0x50fe1962c4965880, "BAIL TUFT BITS GANG CHEF THY"},
		{0x87066dd9644bf206, "FULL PEW DOWN ONCE MORT ARC"},
		{0x7cd34c1040add14b, "FACT HOOF AT FIST SITE KENT"},
		{0x5aa37a81f212146c, "BODE HOP JAKE STOW JUT RAP"},
		{0xf205753943de4cf9, "ULAN NEW ARMY FUSE SUIT EYED"},
		{0xddcdac956f234937, "SKIM CULT LOB SLAM POE HOWL"},
		{0xb203e28fa525be47, "LONG IVY JULY AJAR BOND LEE"},
		{0xbb9e6ae1979d8ff4, "MILT VARY MAST OK SEES WENT"},
		{0x63d936639734385b, "CART OTTO HIVE ODE VAT NUT"},
		{0x87fec7768b73ccf9, "GAFF WAIT SKID GIG SKY EYED"},
		{0xad85f658ebe383c9, "LEST OR HEEL SCOT ROB SUIT"},
		{0xd07ce229b5cf119b, "RITE TAKE GELD COST TUNE RECK"},
		{0x27bc71035aaf3dc6, "MAY STAR TIN LYON VEDA STAN"},
		{0xd51f3e99bf8e6f0b, "RUST WELT KICK FELL TAIL FRAU"},
		{0x82aeb52d943774e4, "FLIT DOSE ALSO MEW DRUM DEFY"},
		{0x4f296a74fe1567ec, "AURA ALOE HURL WING BERG WAIT"},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			s := Encode(tc.n)
			if s != tc.s {
				t.Errorf("got %v, want %v", s, tc.s)
			}
			n, ok := Decode(tc.s)
			if !ok || n != tc.n {
				t.Errorf("got %v, want %v", n, tc.n)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(0x9e876134d90499dd)
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("INCH SEA ANNE LONG AHEM TOUR")
	}
}

func FuzzSixWord(f *testing.F) {
	f.Fuzz(func(t *testing.T, n uint64) {
		s := Encode(n)
		d, ok := Decode(s)
		if !ok {
			t.Errorf("%v not ok", n)
		}
		if n != d {
			t.Errorf("want %v, got %v", n, d)
		}
	})
}
