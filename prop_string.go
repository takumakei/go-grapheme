// Code generated by "stringer -type=Prop property.go"; DO NOT EDIT

package grapheme

import "fmt"

const _Prop_name = "OutOfRangeAnyPrependCRLFControlExtendRegional_IndicatorSpacingMarkLVTLVLVTE_BaseE_ModifierZWJGlue_After_ZwjE_Base_GAZ"

var _Prop_index = [...]uint8{0, 10, 13, 20, 22, 24, 31, 37, 55, 66, 67, 68, 69, 71, 74, 80, 90, 93, 107, 117}

func (i Prop) String() string {
	if i < 0 || i >= Prop(len(_Prop_index)-1) {
		return fmt.Sprintf("Prop(%d)", i)
	}
	return _Prop_name[_Prop_index[i]:_Prop_index[i+1]]
}