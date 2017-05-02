package grapheme

// Break slices s into grapheme clusters and returns a slice of them.
func Break(s []rune) [][]rune {
	if len(s) == 0 {
		return nil
	}

	p := properties(s)

	r := make([][]rune, 0, len(s))

	for {
		i := next(p)
		r = append(r, s[:i])
		if len(s) == i {
			break
		}
		s = s[i:]
		p = p[i:]
	}

	return r
}

func Properties(s []rune) []Prop {
	if len(s) == 0 {
		return nil
	}
	return properties(s)
}

func properties(s []rune) []Prop {
	r := make([]Prop, len(s))
	for i, v := range s {
		r[i] = Property(v)
	}
	return r
}

func next(p []Prop) int {
	e := len(p)
	// if e == 0 {
	// 	panic("next(p); len(p) is 0, want > 0")
	// }
	i := 0
	for {
		lhs := p[i]
		i++
		if i == e {
			return i
		}
		rhs := p[i]

		// GB3
		if lhs == CR && rhs == LF {
			continue
		}
		// GB4
		if lhs == Control || lhs == CR || lhs == LF {
			return i
		}
		// GB5
		if rhs == Control || rhs == CR || rhs == LF {
			return i
		}
		// GB6
		if lhs == L && (rhs == L || rhs == V || rhs == LV || rhs == LVT) {
			continue
		}
		// GB7
		if (lhs == LV || lhs == V) && (rhs == V || rhs == T) {
			continue
		}
		// GB8
		if (lhs == LVT || lhs == T) && rhs == T {
			continue
		}
		// GB10 (この実装ではGB9より前に検査する必要がある)
		if lhs == E_Base || lhs == E_Base_GAZ {
			for rhs == Extend {
				i++
				if i == e {
					return i
				}
				rhs = p[i]
			}
			if rhs == E_Modifier {
				continue
			}
			if doNotBreakBefore(rhs) {
				continue
			}
			return i
		}
		// GB9
		if rhs == Extend || rhs == ZWJ {
			continue
		}
		// GB9a
		if rhs == SpacingMark {
			continue
		}
		// GB9b
		if lhs == Prepend {
			continue
		}
		// GB11
		if lhs == ZWJ && (rhs == Glue_After_Zwj || rhs == E_Base_GAZ) {
			continue
		}
		// GB12 and GB13
		if lhs == Regional_Indicator && rhs == Regional_Indicator {
			i++
			if i == e {
				return i
			}
			rhs = p[i]
			if doNotBreakBefore(rhs) {
				continue
			}
			return i
		}
		// GB999
		return i
	}
}

func doNotBreakBefore(p Prop) bool {
	// GB9 and GB9a
	return p == Extend || p == ZWJ || p == SpacingMark
}
