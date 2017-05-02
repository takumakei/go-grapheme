package grapheme

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleBreak() {
	a := Break([]rune("👨‍👨‍👧‍👧🇯🇵"))
	fmt.Printf("len = %d\n", len(a))
	for _, e := range a {
		fmt.Printf("%s %04x\n", string(e), e)
	}
	// Output:
	// len = 2
	// 👨‍👨‍👧‍👧 [1f468 200d 1f468 200d 1f467 200d 1f467]
	// 🇯🇵 [1f1ef 1f1f5]
}

func ExampleProperty() {
	fmt.Println(Property('\n'))
	// Output:
	// LF
}

func ExampleProperties() {
	a := Properties([]rune("A👨🇵"))
	for _, e := range a {
		fmt.Printf("%s\n", e)
	}
	// Output:
	// Any
	// E_Base_GAZ
	// Regional_Indicator
}

type testCase struct {
	label string
	input string
	want  []string
}

func TestBreak(t *testing.T) {
	check(t, testCases)

	check(t, []testCase{
		{"empty", "", []string{}},
		{"label", "aa", []string{"a", "a"}},
	})
}

func check(t *testing.T, ts []testCase) {
	for _, e := range ts {
		a := ss(t, Break(toUTF32(t, e.input)))
		if !eq(a, e.want) {
			t.Errorf(`Break(%s) = %s, want %s // %s`, hex(e.input), pp(a), pp(e.want), e.label)
		}
	}
}

func ss(t *testing.T, a [][]rune) []string {
	var r []string
	for _, e := range a {
		r = append(r, toUTF8(t, e))
	}
	return r
}

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}

func hex(s string) string {
	return fmt.Sprintf(`"%s"`, strings.Replace(fmt.Sprintf(" % 02x", s), " ", `\x`, -1))
}

func pp(a []string) string {
	var s []string
	for _, e := range a {
		s = append(s, hex(e))
	}
	return strings.Join(s, ", ")
}

func toUTF32(t *testing.T, s string) []rune {
	r, err := decodeUTF8(s)
	if err != nil {
		t.Log(err)
	}
	return r
}

func toUTF8(t *testing.T, s []rune) string {
	r, err := encodeUTF8(s)
	if err != nil {
		t.Log(err)
	}
	return r
}

func decodeUTF8(s string) ([]rune, error) {
	x := []byte(s)

	// 0x00(0b0000_0000) <= i <= 0x7f(0b0111_1111) ... 1 byte : U+0000..U+007F
	// 0x80(0b1000_0000) <= i <= 0xbf(0b1011_1111) ... (sub bytes)
	// 0xc0(0b1100_0000) <= i <= 0xc1(0b1100_0001) ... <INVALID CODE>
	// 0xc2(0b1100_0010) <= i <= 0xdf(0b1101_1111) ... 2 bytes : U+0080..U+07FF
	// 0xe0(0b1110_0000) <= i <= 0xef(0b1110_1111) ... 3 bytes : U+0800..U+FFFF
	// 0xf0(0b1111_0000) <= i <= 0xf7(0b1111_0111) ... 4 bytes : U+10000..U+1FFFF
	// 0xf8(0b1111_1000) <= i <= 0xff(0b1111_1111) ... <INVALID CODE>

	e := len(x)
	if e == 0 {
		return nil, nil
	}

	r := make([]rune, 0, e)
	i := 0
	for i < e {
		a := rune(x[i])
		i++

		if a <= 0x7f {
			r = append(r, a)
			continue
		}

		if a <= 0xbf {
			// 先頭バイトの位置に２バイト目以降のビットパターン'0b10xxxxxx' が現れた
			return r, fmt.Errorf("unexpected byte (0x80..0xbf) for first byte")
		}

		if a <= 0xc1 {
			// 0xc0(0b1100_0000) <= i <= 0xc1(0b1100_0001) ... <INVALID CODE>
			return r, fmt.Errorf("invalid byte (0xc0..0xc1)")
		}

		if i == e {
			// ２バイト以上必要なのに１バイトしかデータがない
			return r, fmt.Errorf("unexpected end of string")
		}

		if a <= 0xdf {
			b := rune(x[i])
			if (b & 0xc0) != 0x80 {
				return r, fmt.Errorf("unexpected byte for 2nd of 2")
			}
			i++

			c := ((a & 0x1f) << 6) | (b & 0x3f)
			if c < 0x80 {
				return r, fmt.Errorf("too small code point (..0x80) for 2 bytes")
			}

			r = append(r, c)
			continue
		}

		if a <= 0xef {
			b := rune(x[i])
			if (b & 0xc0) != 0x80 {
				return r, fmt.Errorf("unexpected byte for 2nd of 3")
			}
			i++

			if i == e {
				return r, fmt.Errorf("unexpected end of string")
			}
			c := rune(x[i])
			if (c & 0xc0) != 0x80 {
				return r, fmt.Errorf("unexpected byte for 3rd of 3")
			}
			i++

			d := ((a & 0x0f) << 12) | ((b & 0x3f) << 6) | (c & 0x3f)
			if d < 0x800 {
				return r, fmt.Errorf("too small code point (..0x800) for 3 bytes")
			}

			r = append(r, d)
			continue
		}

		if a <= 0xf7 {
			b := rune(x[i])
			if (b & 0xc0) != 0x80 {
				return r, fmt.Errorf("unexpected byte for 2nd of 4")
			}
			i++

			if i == e {
				return r, fmt.Errorf("unexpected end of string")
			}
			c := rune(x[i])
			if (c & 0xc0) != 0x80 {
				return r, fmt.Errorf("unexpected byte for 3rd of 4")
			}
			i++

			if i == e {
				return r, fmt.Errorf("unexpected end of string")
			}
			d := rune(x[i])
			if (d & 0xc0) != 0x80 {
				return r, fmt.Errorf("unexpected byte for 4th of 4")
			}
			i++

			e := ((a & 0x07) << 18) | ((b & 0x3f) << 12) | ((c & 0x3f) << 6) | (d & 0x3f)
			if e < 0x10000 {
				return r, fmt.Errorf("too small code point (..0x10000) for 4 bytes")
			}

			r = append(r, e)
			continue
		}

		// 0xf8(0b1111_1000) <= i <= 0xff(0b1111_1111) ... <INVALID CODE>
		return r, fmt.Errorf("invalid byte (0xf8..0xff)")
	}

	return r, nil
}

func encodeUTF8(runes []rune) (string, error) {
	var b []byte
	for _, i := range runes {
		if i <= 0x007f {
			b = append(b, byte(i))
			continue
		}

		if i <= 0x07ff {
			b = append(b, byte(((i>>6)&0x1f)|0xc0), byte((i&0x3f)|0x80))
			continue
		}

		if i <= 0xffff {
			b = append(b, byte(((i>>12)&0x0f)|0xe0), byte((i>>6)&0x3f|0x80), byte((i&0x3f)|0x80))
			continue
		}

		if i <= 0x10ffff {
			b = append(b, byte(((i>>18)&0x07)|0xf0), byte((i>>12)&0x3f|0x80), byte((i>>6)&0x3f|0x80), byte((i&0x3f)|0x80))
			continue
		}

		return string(b), fmt.Errorf("too large value for a code point of unicode")
	}
	return string(b), nil
}

func TestUnknownProperty(t *testing.T) {
	x := -1
	y := Prop(x).String()
	w := "Prop(-1)"
	if y != w {
		t.Errorf("Prop(%v).String() = %s want %s", x, y, w)
	}
}

func TestProperty2(t *testing.T) {
	for i := 0; i < 0x110000; i++ {
		r := rune(i)
		a := Property(r)
		b := Property2(r)
		if a != b {
			t.Errorf("Property2(0x%04x) = %s want %s", r, b, a)
		}
	}
}

func BenchmarkProperty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r rune
		for r = 0; r < 0x110000; r++ {
			Property(r)
		}
	}
}

func BenchmarkProperty2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r rune
		for r = 0; r < 0x110000; r++ {
			Property2(r)
		}
	}
}
