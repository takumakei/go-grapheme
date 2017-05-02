// Code generated from "http://www.unicode.org/Public/9.0.0/ucd/auxiliary/GraphemeBreakProperty.txt"; DO NOT EDIT

package grapheme

//go:generate stringer -type=Prop $GOFILE

type Prop int

const (
	OutOfRange         Prop = iota
	Any                     //   0 rules found
	Prepend                 //   7 rules found
	CR                      //   1 rules found
	LF                      //   1 rules found
	Control                 //  26 rules found
	Extend                  // 313 rules found
	Regional_Indicator      //   1 rules found
	SpacingMark             // 140 rules found
	L                       //   2 rules found
	V                       //   2 rules found
	T                       //   2 rules found
	LV                      // 399 rules found
	LVT                     // 399 rules found
	E_Base                  //  28 rules found
	E_Modifier              //   1 rules found
	ZWJ                     //   1 rules found
	Glue_After_Zwj          //   3 rules found
	E_Base_GAZ              //   1 rules found
)

func Property(r rune) Prop {
	if r < 0x0bb6d {
		if r < 0x01cf2 {
			if r < 0x00d01 {
				if r < 0x00a01 {
					if r < 0x007b1 {
						if r < 0x00606 {
							if r < 0x00483 {
								if r < 0x0007f {
									if r < 0x0000d {
										if r < 0x0000a {
											return Control
										} else { // 0x0000a <= r
											if r < 0x0000b {
												return LF
											} else { // 0x0000b <= r
												return Control
											}
										}
									} else { // 0x0000d <= r
										if r < 0x0000e {
											return CR
										} else { // 0x0000e <= r
											if r < 0x00020 {
												return Control
											} else { // 0x00020 <= r
												return Any
											}
										}
									}
								} else { // 0x0007f <= r
									if r < 0x000ae {
										if r < 0x000a0 {
											return Control
										} else { // 0x000a0 <= r
											if r < 0x000ad {
												return Any
											} else { // 0x000ad <= r
												return Control
											}
										}
									} else { // 0x000ae <= r
										if r < 0x00300 {
											return Any
										} else { // 0x00300 <= r
											if r < 0x00370 {
												return Extend
											} else { // 0x00370 <= r
												return Any
											}
										}
									}
								}
							} else { // 0x00483 <= r
								if r < 0x005c1 {
									if r < 0x005be {
										if r < 0x0048a {
											return Extend
										} else { // 0x0048a <= r
											if r < 0x00591 {
												return Any
											} else { // 0x00591 <= r
												return Extend
											}
										}
									} else { // 0x005be <= r
										if r < 0x005bf {
											return Any
										} else { // 0x005bf <= r
											if r < 0x005c0 {
												return Extend
											} else { // 0x005c0 <= r
												return Any
											}
										}
									}
								} else { // 0x005c1 <= r
									if r < 0x005c6 {
										if r < 0x005c3 {
											return Extend
										} else { // 0x005c3 <= r
											if r < 0x005c4 {
												return Any
											} else { // 0x005c4 <= r
												return Extend
											}
										}
									} else { // 0x005c6 <= r
										if r < 0x005c8 {
											if r < 0x005c7 {
												return Any
											} else { // 0x005c7 <= r
												return Extend
											}
										} else { // 0x005c8 <= r
											if r < 0x00600 {
												return Any
											} else { // 0x00600 <= r
												return Prepend
											}
										}
									}
								}
							}
						} else { // 0x00606 <= r
							if r < 0x006df {
								if r < 0x00660 {
									if r < 0x0061c {
										if r < 0x00610 {
											return Any
										} else { // 0x00610 <= r
											if r < 0x0061b {
												return Extend
											} else { // 0x0061b <= r
												return Any
											}
										}
									} else { // 0x0061c <= r
										if r < 0x0061d {
											return Control
										} else { // 0x0061d <= r
											if r < 0x0064b {
												return Any
											} else { // 0x0064b <= r
												return Extend
											}
										}
									}
								} else { // 0x00660 <= r
									if r < 0x006d6 {
										if r < 0x00670 {
											return Any
										} else { // 0x00670 <= r
											if r < 0x00671 {
												return Extend
											} else { // 0x00671 <= r
												return Any
											}
										}
									} else { // 0x006d6 <= r
										if r < 0x006dd {
											return Extend
										} else { // 0x006dd <= r
											if r < 0x006de {
												return Prepend
											} else { // 0x006de <= r
												return Any
											}
										}
									}
								}
							} else { // 0x006df <= r
								if r < 0x0070f {
									if r < 0x006e9 {
										if r < 0x006e5 {
											return Extend
										} else { // 0x006e5 <= r
											if r < 0x006e7 {
												return Any
											} else { // 0x006e7 <= r
												return Extend
											}
										}
									} else { // 0x006e9 <= r
										if r < 0x006ea {
											return Any
										} else { // 0x006ea <= r
											if r < 0x006ee {
												return Extend
											} else { // 0x006ee <= r
												return Any
											}
										}
									}
								} else { // 0x0070f <= r
									if r < 0x00712 {
										if r < 0x00710 {
											return Prepend
										} else { // 0x00710 <= r
											if r < 0x00711 {
												return Any
											} else { // 0x00711 <= r
												return Extend
											}
										}
									} else { // 0x00712 <= r
										if r < 0x0074b {
											if r < 0x00730 {
												return Any
											} else { // 0x00730 <= r
												return Extend
											}
										} else { // 0x0074b <= r
											if r < 0x007a6 {
												return Any
											} else { // 0x007a6 <= r
												return Extend
											}
										}
									}
								}
							}
						}
					} else { // 0x007b1 <= r
						if r < 0x0094d {
							if r < 0x0085c {
								if r < 0x00824 {
									if r < 0x00816 {
										if r < 0x007eb {
											return Any
										} else { // 0x007eb <= r
											if r < 0x007f4 {
												return Extend
											} else { // 0x007f4 <= r
												return Any
											}
										}
									} else { // 0x00816 <= r
										if r < 0x0081a {
											return Extend
										} else { // 0x0081a <= r
											if r < 0x0081b {
												return Any
											} else { // 0x0081b <= r
												return Extend
											}
										}
									}
								} else { // 0x00824 <= r
									if r < 0x00829 {
										if r < 0x00825 {
											return Any
										} else { // 0x00825 <= r
											if r < 0x00828 {
												return Extend
											} else { // 0x00828 <= r
												return Any
											}
										}
									} else { // 0x00829 <= r
										if r < 0x0082e {
											return Extend
										} else { // 0x0082e <= r
											if r < 0x00859 {
												return Any
											} else { // 0x00859 <= r
												return Extend
											}
										}
									}
								}
							} else { // 0x0085c <= r
								if r < 0x0093a {
									if r < 0x008e3 {
										if r < 0x008d4 {
											return Any
										} else { // 0x008d4 <= r
											if r < 0x008e2 {
												return Extend
											} else { // 0x008e2 <= r
												return Prepend
											}
										}
									} else { // 0x008e3 <= r
										if r < 0x00903 {
											return Extend
										} else { // 0x00903 <= r
											if r < 0x00904 {
												return SpacingMark
											} else { // 0x00904 <= r
												return Any
											}
										}
									}
								} else { // 0x0093a <= r
									if r < 0x0093d {
										if r < 0x0093b {
											return Extend
										} else { // 0x0093b <= r
											if r < 0x0093c {
												return SpacingMark
											} else { // 0x0093c <= r
												return Extend
											}
										}
									} else { // 0x0093d <= r
										if r < 0x00941 {
											if r < 0x0093e {
												return Any
											} else { // 0x0093e <= r
												return SpacingMark
											}
										} else { // 0x00941 <= r
											if r < 0x00949 {
												return Extend
											} else { // 0x00949 <= r
												return SpacingMark
											}
										}
									}
								}
							}
						} else { // 0x0094d <= r
							if r < 0x009be {
								if r < 0x00964 {
									if r < 0x00951 {
										if r < 0x0094e {
											return Extend
										} else { // 0x0094e <= r
											if r < 0x00950 {
												return SpacingMark
											} else { // 0x00950 <= r
												return Any
											}
										}
									} else { // 0x00951 <= r
										if r < 0x00958 {
											return Extend
										} else { // 0x00958 <= r
											if r < 0x00962 {
												return Any
											} else { // 0x00962 <= r
												return Extend
											}
										}
									}
								} else { // 0x00964 <= r
									if r < 0x00984 {
										if r < 0x00981 {
											return Any
										} else { // 0x00981 <= r
											if r < 0x00982 {
												return Extend
											} else { // 0x00982 <= r
												return SpacingMark
											}
										}
									} else { // 0x00984 <= r
										if r < 0x009bc {
											return Any
										} else { // 0x009bc <= r
											if r < 0x009bd {
												return Extend
											} else { // 0x009bd <= r
												return Any
											}
										}
									}
								}
							} else { // 0x009be <= r
								if r < 0x009cb {
									if r < 0x009c5 {
										if r < 0x009bf {
											return Extend
										} else { // 0x009bf <= r
											if r < 0x009c1 {
												return SpacingMark
											} else { // 0x009c1 <= r
												return Extend
											}
										}
									} else { // 0x009c5 <= r
										if r < 0x009c7 {
											return Any
										} else { // 0x009c7 <= r
											if r < 0x009c9 {
												return SpacingMark
											} else { // 0x009c9 <= r
												return Any
											}
										}
									}
								} else { // 0x009cb <= r
									if r < 0x009d7 {
										if r < 0x009cd {
											return SpacingMark
										} else { // 0x009cd <= r
											if r < 0x009ce {
												return Extend
											} else { // 0x009ce <= r
												return Any
											}
										}
									} else { // 0x009d7 <= r
										if r < 0x009e2 {
											if r < 0x009d8 {
												return Extend
											} else { // 0x009d8 <= r
												return Any
											}
										} else { // 0x009e2 <= r
											if r < 0x009e4 {
												return Extend
											} else { // 0x009e4 <= r
												return Any
											}
										}
									}
								}
							}
						}
					}
				} else { // 0x00a01 <= r
					if r < 0x00b62 {
						if r < 0x00ac6 {
							if r < 0x00a51 {
								if r < 0x00a41 {
									if r < 0x00a3c {
										if r < 0x00a03 {
											return Extend
										} else { // 0x00a03 <= r
											if r < 0x00a04 {
												return SpacingMark
											} else { // 0x00a04 <= r
												return Any
											}
										}
									} else { // 0x00a3c <= r
										if r < 0x00a3d {
											return Extend
										} else { // 0x00a3d <= r
											if r < 0x00a3e {
												return Any
											} else { // 0x00a3e <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x00a41 <= r
									if r < 0x00a49 {
										if r < 0x00a43 {
											return Extend
										} else { // 0x00a43 <= r
											if r < 0x00a47 {
												return Any
											} else { // 0x00a47 <= r
												return Extend
											}
										}
									} else { // 0x00a49 <= r
										if r < 0x00a4b {
											return Any
										} else { // 0x00a4b <= r
											if r < 0x00a4e {
												return Extend
											} else { // 0x00a4e <= r
												return Any
											}
										}
									}
								}
							} else { // 0x00a51 <= r
								if r < 0x00a81 {
									if r < 0x00a72 {
										if r < 0x00a52 {
											return Extend
										} else { // 0x00a52 <= r
											if r < 0x00a70 {
												return Any
											} else { // 0x00a70 <= r
												return Extend
											}
										}
									} else { // 0x00a72 <= r
										if r < 0x00a75 {
											return Any
										} else { // 0x00a75 <= r
											if r < 0x00a76 {
												return Extend
											} else { // 0x00a76 <= r
												return Any
											}
										}
									}
								} else { // 0x00a81 <= r
									if r < 0x00abc {
										if r < 0x00a83 {
											return Extend
										} else { // 0x00a83 <= r
											if r < 0x00a84 {
												return SpacingMark
											} else { // 0x00a84 <= r
												return Any
											}
										}
									} else { // 0x00abc <= r
										if r < 0x00abe {
											if r < 0x00abd {
												return Extend
											} else { // 0x00abd <= r
												return Any
											}
										} else { // 0x00abe <= r
											if r < 0x00ac1 {
												return SpacingMark
											} else { // 0x00ac1 <= r
												return Extend
											}
										}
									}
								}
							}
						} else { // 0x00ac6 <= r
							if r < 0x00b3c {
								if r < 0x00ace {
									if r < 0x00aca {
										if r < 0x00ac7 {
											return Any
										} else { // 0x00ac7 <= r
											if r < 0x00ac9 {
												return Extend
											} else { // 0x00ac9 <= r
												return SpacingMark
											}
										}
									} else { // 0x00aca <= r
										if r < 0x00acb {
											return Any
										} else { // 0x00acb <= r
											if r < 0x00acd {
												return SpacingMark
											} else { // 0x00acd <= r
												return Extend
											}
										}
									}
								} else { // 0x00ace <= r
									if r < 0x00b01 {
										if r < 0x00ae2 {
											return Any
										} else { // 0x00ae2 <= r
											if r < 0x00ae4 {
												return Extend
											} else { // 0x00ae4 <= r
												return Any
											}
										}
									} else { // 0x00b01 <= r
										if r < 0x00b02 {
											return Extend
										} else { // 0x00b02 <= r
											if r < 0x00b04 {
												return SpacingMark
											} else { // 0x00b04 <= r
												return Any
											}
										}
									}
								}
							} else { // 0x00b3c <= r
								if r < 0x00b47 {
									if r < 0x00b40 {
										if r < 0x00b3d {
											return Extend
										} else { // 0x00b3d <= r
											if r < 0x00b3e {
												return Any
											} else { // 0x00b3e <= r
												return Extend
											}
										}
									} else { // 0x00b40 <= r
										if r < 0x00b41 {
											return SpacingMark
										} else { // 0x00b41 <= r
											if r < 0x00b45 {
												return Extend
											} else { // 0x00b45 <= r
												return Any
											}
										}
									}
								} else { // 0x00b47 <= r
									if r < 0x00b4d {
										if r < 0x00b49 {
											return SpacingMark
										} else { // 0x00b49 <= r
											if r < 0x00b4b {
												return Any
											} else { // 0x00b4b <= r
												return SpacingMark
											}
										}
									} else { // 0x00b4d <= r
										if r < 0x00b56 {
											if r < 0x00b4e {
												return Extend
											} else { // 0x00b4e <= r
												return Any
											}
										} else { // 0x00b56 <= r
											if r < 0x00b58 {
												return Extend
											} else { // 0x00b58 <= r
												return Any
											}
										}
									}
								}
							}
						}
					} else { // 0x00b62 <= r
						if r < 0x00c4e {
							if r < 0x00bcd {
								if r < 0x00bc0 {
									if r < 0x00b83 {
										if r < 0x00b64 {
											return Extend
										} else { // 0x00b64 <= r
											if r < 0x00b82 {
												return Any
											} else { // 0x00b82 <= r
												return Extend
											}
										}
									} else { // 0x00b83 <= r
										if r < 0x00bbe {
											return Any
										} else { // 0x00bbe <= r
											if r < 0x00bbf {
												return Extend
											} else { // 0x00bbf <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x00bc0 <= r
									if r < 0x00bc6 {
										if r < 0x00bc1 {
											return Extend
										} else { // 0x00bc1 <= r
											if r < 0x00bc3 {
												return SpacingMark
											} else { // 0x00bc3 <= r
												return Any
											}
										}
									} else { // 0x00bc6 <= r
										if r < 0x00bc9 {
											return SpacingMark
										} else { // 0x00bc9 <= r
											if r < 0x00bca {
												return Any
											} else { // 0x00bca <= r
												return SpacingMark
											}
										}
									}
								}
							} else { // 0x00bcd <= r
								if r < 0x00c04 {
									if r < 0x00bd8 {
										if r < 0x00bce {
											return Extend
										} else { // 0x00bce <= r
											if r < 0x00bd7 {
												return Any
											} else { // 0x00bd7 <= r
												return Extend
											}
										}
									} else { // 0x00bd8 <= r
										if r < 0x00c00 {
											return Any
										} else { // 0x00c00 <= r
											if r < 0x00c01 {
												return Extend
											} else { // 0x00c01 <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x00c04 <= r
									if r < 0x00c45 {
										if r < 0x00c3e {
											return Any
										} else { // 0x00c3e <= r
											if r < 0x00c41 {
												return Extend
											} else { // 0x00c41 <= r
												return SpacingMark
											}
										}
									} else { // 0x00c45 <= r
										if r < 0x00c49 {
											if r < 0x00c46 {
												return Any
											} else { // 0x00c46 <= r
												return Extend
											}
										} else { // 0x00c49 <= r
											if r < 0x00c4a {
												return Any
											} else { // 0x00c4a <= r
												return Extend
											}
										}
									}
								}
							}
						} else { // 0x00c4e <= r
							if r < 0x00cc2 {
								if r < 0x00c82 {
									if r < 0x00c62 {
										if r < 0x00c55 {
											return Any
										} else { // 0x00c55 <= r
											if r < 0x00c57 {
												return Extend
											} else { // 0x00c57 <= r
												return Any
											}
										}
									} else { // 0x00c62 <= r
										if r < 0x00c64 {
											return Extend
										} else { // 0x00c64 <= r
											if r < 0x00c81 {
												return Any
											} else { // 0x00c81 <= r
												return Extend
											}
										}
									}
								} else { // 0x00c82 <= r
									if r < 0x00cbd {
										if r < 0x00c84 {
											return SpacingMark
										} else { // 0x00c84 <= r
											if r < 0x00cbc {
												return Any
											} else { // 0x00cbc <= r
												return Extend
											}
										}
									} else { // 0x00cbd <= r
										if r < 0x00cbf {
											if r < 0x00cbe {
												return Any
											} else { // 0x00cbe <= r
												return SpacingMark
											}
										} else { // 0x00cbf <= r
											if r < 0x00cc0 {
												return Extend
											} else { // 0x00cc0 <= r
												return SpacingMark
											}
										}
									}
								}
							} else { // 0x00cc2 <= r
								if r < 0x00cca {
									if r < 0x00cc6 {
										if r < 0x00cc3 {
											return Extend
										} else { // 0x00cc3 <= r
											if r < 0x00cc5 {
												return SpacingMark
											} else { // 0x00cc5 <= r
												return Any
											}
										}
									} else { // 0x00cc6 <= r
										if r < 0x00cc7 {
											return Extend
										} else { // 0x00cc7 <= r
											if r < 0x00cc9 {
												return SpacingMark
											} else { // 0x00cc9 <= r
												return Any
											}
										}
									}
								} else { // 0x00cca <= r
									if r < 0x00cd5 {
										if r < 0x00ccc {
											return SpacingMark
										} else { // 0x00ccc <= r
											if r < 0x00cce {
												return Extend
											} else { // 0x00cce <= r
												return Any
											}
										}
									} else { // 0x00cd5 <= r
										if r < 0x00ce2 {
											if r < 0x00cd7 {
												return Extend
											} else { // 0x00cd7 <= r
												return Any
											}
										} else { // 0x00ce2 <= r
											if r < 0x00ce4 {
												return Extend
											} else { // 0x00ce4 <= r
												return Any
											}
										}
									}
								}
							}
						}
					}
				}
			} else { // 0x00d01 <= r
				if r < 0x01715 {
					if r < 0x00f35 {
						if r < 0x00dd6 {
							if r < 0x00d4f {
								if r < 0x00d45 {
									if r < 0x00d3e {
										if r < 0x00d02 {
											return Extend
										} else { // 0x00d02 <= r
											if r < 0x00d04 {
												return SpacingMark
											} else { // 0x00d04 <= r
												return Any
											}
										}
									} else { // 0x00d3e <= r
										if r < 0x00d3f {
											return Extend
										} else { // 0x00d3f <= r
											if r < 0x00d41 {
												return SpacingMark
											} else { // 0x00d41 <= r
												return Extend
											}
										}
									}
								} else { // 0x00d45 <= r
									if r < 0x00d4a {
										if r < 0x00d46 {
											return Any
										} else { // 0x00d46 <= r
											if r < 0x00d49 {
												return SpacingMark
											} else { // 0x00d49 <= r
												return Any
											}
										}
									} else { // 0x00d4a <= r
										if r < 0x00d4d {
											return SpacingMark
										} else { // 0x00d4d <= r
											if r < 0x00d4e {
												return Extend
											} else { // 0x00d4e <= r
												return Prepend
											}
										}
									}
								}
							} else { // 0x00d4f <= r
								if r < 0x00d84 {
									if r < 0x00d62 {
										if r < 0x00d57 {
											return Any
										} else { // 0x00d57 <= r
											if r < 0x00d58 {
												return Extend
											} else { // 0x00d58 <= r
												return Any
											}
										}
									} else { // 0x00d62 <= r
										if r < 0x00d64 {
											return Extend
										} else { // 0x00d64 <= r
											if r < 0x00d82 {
												return Any
											} else { // 0x00d82 <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x00d84 <= r
									if r < 0x00dcf {
										if r < 0x00dca {
											return Any
										} else { // 0x00dca <= r
											if r < 0x00dcb {
												return Extend
											} else { // 0x00dcb <= r
												return Any
											}
										}
									} else { // 0x00dcf <= r
										if r < 0x00dd2 {
											if r < 0x00dd0 {
												return Extend
											} else { // 0x00dd0 <= r
												return SpacingMark
											}
										} else { // 0x00dd2 <= r
											if r < 0x00dd5 {
												return Extend
											} else { // 0x00dd5 <= r
												return Any
											}
										}
									}
								}
							}
						} else { // 0x00dd6 <= r
							if r < 0x00e47 {
								if r < 0x00df4 {
									if r < 0x00ddf {
										if r < 0x00dd7 {
											return Extend
										} else { // 0x00dd7 <= r
											if r < 0x00dd8 {
												return Any
											} else { // 0x00dd8 <= r
												return SpacingMark
											}
										}
									} else { // 0x00ddf <= r
										if r < 0x00de0 {
											return Extend
										} else { // 0x00de0 <= r
											if r < 0x00df2 {
												return Any
											} else { // 0x00df2 <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x00df4 <= r
									if r < 0x00e33 {
										if r < 0x00e31 {
											return Any
										} else { // 0x00e31 <= r
											if r < 0x00e32 {
												return Extend
											} else { // 0x00e32 <= r
												return Any
											}
										}
									} else { // 0x00e33 <= r
										if r < 0x00e34 {
											return SpacingMark
										} else { // 0x00e34 <= r
											if r < 0x00e3b {
												return Extend
											} else { // 0x00e3b <= r
												return Any
											}
										}
									}
								}
							} else { // 0x00e47 <= r
								if r < 0x00eba {
									if r < 0x00eb2 {
										if r < 0x00e4f {
											return Extend
										} else { // 0x00e4f <= r
											if r < 0x00eb1 {
												return Any
											} else { // 0x00eb1 <= r
												return Extend
											}
										}
									} else { // 0x00eb2 <= r
										if r < 0x00eb3 {
											return Any
										} else { // 0x00eb3 <= r
											if r < 0x00eb4 {
												return SpacingMark
											} else { // 0x00eb4 <= r
												return Extend
											}
										}
									}
								} else { // 0x00eba <= r
									if r < 0x00ec8 {
										if r < 0x00ebb {
											return Any
										} else { // 0x00ebb <= r
											if r < 0x00ebd {
												return Extend
											} else { // 0x00ebd <= r
												return Any
											}
										}
									} else { // 0x00ec8 <= r
										if r < 0x00f18 {
											if r < 0x00ece {
												return Extend
											} else { // 0x00ece <= r
												return Any
											}
										} else { // 0x00f18 <= r
											if r < 0x00f1a {
												return Extend
											} else { // 0x00f1a <= r
												return Any
											}
										}
									}
								}
							}
						}
					} else { // 0x00f35 <= r
						if r < 0x0103b {
							if r < 0x00f86 {
								if r < 0x00f3e {
									if r < 0x00f38 {
										if r < 0x00f36 {
											return Extend
										} else { // 0x00f36 <= r
											if r < 0x00f37 {
												return Any
											} else { // 0x00f37 <= r
												return Extend
											}
										}
									} else { // 0x00f38 <= r
										if r < 0x00f39 {
											return Any
										} else { // 0x00f39 <= r
											if r < 0x00f3a {
												return Extend
											} else { // 0x00f3a <= r
												return Any
											}
										}
									}
								} else { // 0x00f3e <= r
									if r < 0x00f7f {
										if r < 0x00f40 {
											return SpacingMark
										} else { // 0x00f40 <= r
											if r < 0x00f71 {
												return Any
											} else { // 0x00f71 <= r
												return Extend
											}
										}
									} else { // 0x00f7f <= r
										if r < 0x00f80 {
											return SpacingMark
										} else { // 0x00f80 <= r
											if r < 0x00f85 {
												return Extend
											} else { // 0x00f85 <= r
												return Any
											}
										}
									}
								}
							} else { // 0x00f86 <= r
								if r < 0x00fc6 {
									if r < 0x00f98 {
										if r < 0x00f88 {
											return Extend
										} else { // 0x00f88 <= r
											if r < 0x00f8d {
												return Any
											} else { // 0x00f8d <= r
												return Extend
											}
										}
									} else { // 0x00f98 <= r
										if r < 0x00f99 {
											return Any
										} else { // 0x00f99 <= r
											if r < 0x00fbd {
												return Extend
											} else { // 0x00fbd <= r
												return Any
											}
										}
									}
								} else { // 0x00fc6 <= r
									if r < 0x01031 {
										if r < 0x00fc7 {
											return Extend
										} else { // 0x00fc7 <= r
											if r < 0x0102d {
												return Any
											} else { // 0x0102d <= r
												return Extend
											}
										}
									} else { // 0x01031 <= r
										if r < 0x01038 {
											if r < 0x01032 {
												return SpacingMark
											} else { // 0x01032 <= r
												return Extend
											}
										} else { // 0x01038 <= r
											if r < 0x01039 {
												return Any
											} else { // 0x01039 <= r
												return Extend
											}
										}
									}
								}
							}
						} else { // 0x0103b <= r
							if r < 0x01085 {
								if r < 0x0105e {
									if r < 0x01056 {
										if r < 0x0103d {
											return SpacingMark
										} else { // 0x0103d <= r
											if r < 0x0103f {
												return Extend
											} else { // 0x0103f <= r
												return Any
											}
										}
									} else { // 0x01056 <= r
										if r < 0x01058 {
											return SpacingMark
										} else { // 0x01058 <= r
											if r < 0x0105a {
												return Extend
											} else { // 0x0105a <= r
												return Any
											}
										}
									}
								} else { // 0x0105e <= r
									if r < 0x01075 {
										if r < 0x01061 {
											return Extend
										} else { // 0x01061 <= r
											if r < 0x01071 {
												return Any
											} else { // 0x01071 <= r
												return Extend
											}
										}
									} else { // 0x01075 <= r
										if r < 0x01083 {
											if r < 0x01082 {
												return Any
											} else { // 0x01082 <= r
												return Extend
											}
										} else { // 0x01083 <= r
											if r < 0x01084 {
												return Any
											} else { // 0x01084 <= r
												return SpacingMark
											}
										}
									}
								}
							} else { // 0x01085 <= r
								if r < 0x01100 {
									if r < 0x0108e {
										if r < 0x01087 {
											return Extend
										} else { // 0x01087 <= r
											if r < 0x0108d {
												return Any
											} else { // 0x0108d <= r
												return Extend
											}
										}
									} else { // 0x0108e <= r
										if r < 0x0109d {
											return Any
										} else { // 0x0109d <= r
											if r < 0x0109e {
												return Extend
											} else { // 0x0109e <= r
												return Any
											}
										}
									}
								} else { // 0x01100 <= r
									if r < 0x01200 {
										if r < 0x01160 {
											return L
										} else { // 0x01160 <= r
											if r < 0x011a8 {
												return V
											} else { // 0x011a8 <= r
												return T
											}
										}
									} else { // 0x01200 <= r
										if r < 0x01360 {
											if r < 0x0135d {
												return Any
											} else { // 0x0135d <= r
												return Extend
											}
										} else { // 0x01360 <= r
											if r < 0x01712 {
												return Any
											} else { // 0x01712 <= r
												return Extend
											}
										}
									}
								}
							}
						}
					}
				} else { // 0x01715 <= r
					if r < 0x01a7d {
						if r < 0x01923 {
							if r < 0x017c7 {
								if r < 0x01774 {
									if r < 0x01752 {
										if r < 0x01732 {
											return Any
										} else { // 0x01732 <= r
											if r < 0x01735 {
												return Extend
											} else { // 0x01735 <= r
												return Any
											}
										}
									} else { // 0x01752 <= r
										if r < 0x01754 {
											return Extend
										} else { // 0x01754 <= r
											if r < 0x01772 {
												return Any
											} else { // 0x01772 <= r
												return Extend
											}
										}
									}
								} else { // 0x01774 <= r
									if r < 0x017b7 {
										if r < 0x017b4 {
											return Any
										} else { // 0x017b4 <= r
											if r < 0x017b6 {
												return Extend
											} else { // 0x017b6 <= r
												return SpacingMark
											}
										}
									} else { // 0x017b7 <= r
										if r < 0x017be {
											return Extend
										} else { // 0x017be <= r
											if r < 0x017c6 {
												return SpacingMark
											} else { // 0x017c6 <= r
												return Extend
											}
										}
									}
								}
							} else { // 0x017c7 <= r
								if r < 0x0180e {
									if r < 0x017dd {
										if r < 0x017c9 {
											return SpacingMark
										} else { // 0x017c9 <= r
											if r < 0x017d4 {
												return Extend
											} else { // 0x017d4 <= r
												return Any
											}
										}
									} else { // 0x017dd <= r
										if r < 0x017de {
											return Extend
										} else { // 0x017de <= r
											if r < 0x0180b {
												return Any
											} else { // 0x0180b <= r
												return Extend
											}
										}
									}
								} else { // 0x0180e <= r
									if r < 0x01887 {
										if r < 0x0180f {
											return Control
										} else { // 0x0180f <= r
											if r < 0x01885 {
												return Any
											} else { // 0x01885 <= r
												return Extend
											}
										}
									} else { // 0x01887 <= r
										if r < 0x018aa {
											if r < 0x018a9 {
												return Any
											} else { // 0x018a9 <= r
												return Extend
											}
										} else { // 0x018aa <= r
											if r < 0x01920 {
												return Any
											} else { // 0x01920 <= r
												return Extend
											}
										}
									}
								}
							}
						} else { // 0x01923 <= r
							if r < 0x01a1c {
								if r < 0x01933 {
									if r < 0x0192c {
										if r < 0x01927 {
											return SpacingMark
										} else { // 0x01927 <= r
											if r < 0x01929 {
												return Extend
											} else { // 0x01929 <= r
												return SpacingMark
											}
										}
									} else { // 0x0192c <= r
										if r < 0x01930 {
											return Any
										} else { // 0x01930 <= r
											if r < 0x01932 {
												return SpacingMark
											} else { // 0x01932 <= r
												return Extend
											}
										}
									}
								} else { // 0x01933 <= r
									if r < 0x01a17 {
										if r < 0x01939 {
											return SpacingMark
										} else { // 0x01939 <= r
											if r < 0x0193c {
												return Extend
											} else { // 0x0193c <= r
												return Any
											}
										}
									} else { // 0x01a17 <= r
										if r < 0x01a19 {
											return Extend
										} else { // 0x01a19 <= r
											if r < 0x01a1b {
												return SpacingMark
											} else { // 0x01a1b <= r
												return Extend
											}
										}
									}
								}
							} else { // 0x01a1c <= r
								if r < 0x01a60 {
									if r < 0x01a57 {
										if r < 0x01a55 {
											return Any
										} else { // 0x01a55 <= r
											if r < 0x01a56 {
												return SpacingMark
											} else { // 0x01a56 <= r
												return Extend
											}
										}
									} else { // 0x01a57 <= r
										if r < 0x01a58 {
											return SpacingMark
										} else { // 0x01a58 <= r
											if r < 0x01a5f {
												return Extend
											} else { // 0x01a5f <= r
												return Any
											}
										}
									}
								} else { // 0x01a60 <= r
									if r < 0x01a63 {
										if r < 0x01a61 {
											return Extend
										} else { // 0x01a61 <= r
											if r < 0x01a62 {
												return Any
											} else { // 0x01a62 <= r
												return Extend
											}
										}
									} else { // 0x01a63 <= r
										if r < 0x01a6d {
											if r < 0x01a65 {
												return Any
											} else { // 0x01a65 <= r
												return Extend
											}
										} else { // 0x01a6d <= r
											if r < 0x01a73 {
												return SpacingMark
											} else { // 0x01a73 <= r
												return Extend
											}
										}
									}
								}
							}
						}
					} else { // 0x01a7d <= r
						if r < 0x01ba8 {
							if r < 0x01b3c {
								if r < 0x01b04 {
									if r < 0x01ab0 {
										if r < 0x01a7f {
											return Any
										} else { // 0x01a7f <= r
											if r < 0x01a80 {
												return Extend
											} else { // 0x01a80 <= r
												return Any
											}
										}
									} else { // 0x01ab0 <= r
										if r < 0x01abf {
											return Extend
										} else { // 0x01abf <= r
											if r < 0x01b00 {
												return Any
											} else { // 0x01b00 <= r
												return Extend
											}
										}
									}
								} else { // 0x01b04 <= r
									if r < 0x01b35 {
										if r < 0x01b05 {
											return SpacingMark
										} else { // 0x01b05 <= r
											if r < 0x01b34 {
												return Any
											} else { // 0x01b34 <= r
												return Extend
											}
										}
									} else { // 0x01b35 <= r
										if r < 0x01b36 {
											return SpacingMark
										} else { // 0x01b36 <= r
											if r < 0x01b3b {
												return Extend
											} else { // 0x01b3b <= r
												return SpacingMark
											}
										}
									}
								}
							} else { // 0x01b3c <= r
								if r < 0x01b74 {
									if r < 0x01b43 {
										if r < 0x01b3d {
											return Extend
										} else { // 0x01b3d <= r
											if r < 0x01b42 {
												return SpacingMark
											} else { // 0x01b42 <= r
												return Extend
											}
										}
									} else { // 0x01b43 <= r
										if r < 0x01b45 {
											return SpacingMark
										} else { // 0x01b45 <= r
											if r < 0x01b6b {
												return Any
											} else { // 0x01b6b <= r
												return Extend
											}
										}
									}
								} else { // 0x01b74 <= r
									if r < 0x01b83 {
										if r < 0x01b80 {
											return Any
										} else { // 0x01b80 <= r
											if r < 0x01b82 {
												return Extend
											} else { // 0x01b82 <= r
												return SpacingMark
											}
										}
									} else { // 0x01b83 <= r
										if r < 0x01ba2 {
											if r < 0x01ba1 {
												return Any
											} else { // 0x01ba1 <= r
												return SpacingMark
											}
										} else { // 0x01ba2 <= r
											if r < 0x01ba6 {
												return Extend
											} else { // 0x01ba6 <= r
												return SpacingMark
											}
										}
									}
								}
							}
						} else { // 0x01ba8 <= r
							if r < 0x01c24 {
								if r < 0x01be8 {
									if r < 0x01bae {
										if r < 0x01baa {
											return Extend
										} else { // 0x01baa <= r
											if r < 0x01bab {
												return SpacingMark
											} else { // 0x01bab <= r
												return Extend
											}
										}
									} else { // 0x01bae <= r
										if r < 0x01be6 {
											return Any
										} else { // 0x01be6 <= r
											if r < 0x01be7 {
												return Extend
											} else { // 0x01be7 <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x01be8 <= r
									if r < 0x01bee {
										if r < 0x01bea {
											return Extend
										} else { // 0x01bea <= r
											if r < 0x01bed {
												return SpacingMark
											} else { // 0x01bed <= r
												return Extend
											}
										}
									} else { // 0x01bee <= r
										if r < 0x01bf2 {
											if r < 0x01bef {
												return SpacingMark
											} else { // 0x01bef <= r
												return Extend
											}
										} else { // 0x01bf2 <= r
											if r < 0x01bf4 {
												return SpacingMark
											} else { // 0x01bf4 <= r
												return Any
											}
										}
									}
								}
							} else { // 0x01c24 <= r
								if r < 0x01cd3 {
									if r < 0x01c36 {
										if r < 0x01c2c {
											return SpacingMark
										} else { // 0x01c2c <= r
											if r < 0x01c34 {
												return Extend
											} else { // 0x01c34 <= r
												return SpacingMark
											}
										}
									} else { // 0x01c36 <= r
										if r < 0x01c38 {
											return Extend
										} else { // 0x01c38 <= r
											if r < 0x01cd0 {
												return Any
											} else { // 0x01cd0 <= r
												return Extend
											}
										}
									}
								} else { // 0x01cd3 <= r
									if r < 0x01ce2 {
										if r < 0x01cd4 {
											return Any
										} else { // 0x01cd4 <= r
											if r < 0x01ce1 {
												return Extend
											} else { // 0x01ce1 <= r
												return SpacingMark
											}
										}
									} else { // 0x01ce2 <= r
										if r < 0x01ced {
											if r < 0x01ce9 {
												return Extend
											} else { // 0x01ce9 <= r
												return Any
											}
										} else { // 0x01ced <= r
											if r < 0x01cee {
												return Extend
											} else { // 0x01cee <= r
												return Any
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else { // 0x01cf2 <= r
			if r < 0x0b061 {
				if r < 0x0aabe {
					if r < 0x0a80b {
						if r < 0x0270e {
							if r < 0x0200e {
								if r < 0x01df6 {
									if r < 0x01cf8 {
										if r < 0x01cf4 {
											return SpacingMark
										} else { // 0x01cf4 <= r
											if r < 0x01cf5 {
												return Extend
											} else { // 0x01cf5 <= r
												return Any
											}
										}
									} else { // 0x01cf8 <= r
										if r < 0x01cfa {
											return Extend
										} else { // 0x01cfa <= r
											if r < 0x01dc0 {
												return Any
											} else { // 0x01dc0 <= r
												return Extend
											}
										}
									}
								} else { // 0x01df6 <= r
									if r < 0x0200b {
										if r < 0x01dfb {
											return Any
										} else { // 0x01dfb <= r
											if r < 0x01e00 {
												return Extend
											} else { // 0x01e00 <= r
												return Any
											}
										}
									} else { // 0x0200b <= r
										if r < 0x0200c {
											return Control
										} else { // 0x0200c <= r
											if r < 0x0200d {
												return Extend
											} else { // 0x0200d <= r
												return ZWJ
											}
										}
									}
								}
							} else { // 0x0200e <= r
								if r < 0x020d0 {
									if r < 0x0202f {
										if r < 0x02010 {
											return Control
										} else { // 0x02010 <= r
											if r < 0x02028 {
												return Any
											} else { // 0x02028 <= r
												return Control
											}
										}
									} else { // 0x0202f <= r
										if r < 0x02060 {
											return Any
										} else { // 0x02060 <= r
											if r < 0x02070 {
												return Control
											} else { // 0x02070 <= r
												return Any
											}
										}
									}
								} else { // 0x020d0 <= r
									if r < 0x0261e {
										if r < 0x020f1 {
											return Extend
										} else { // 0x020f1 <= r
											if r < 0x0261d {
												return Any
											} else { // 0x0261d <= r
												return E_Base
											}
										}
									} else { // 0x0261e <= r
										if r < 0x026fa {
											if r < 0x026f9 {
												return Any
											} else { // 0x026f9 <= r
												return E_Base
											}
										} else { // 0x026fa <= r
											if r < 0x0270a {
												return Any
											} else { // 0x0270a <= r
												return E_Base
											}
										}
									}
								}
							}
						} else { // 0x0270e <= r
							if r < 0x0309b {
								if r < 0x02d80 {
									if r < 0x02cef {
										if r < 0x02764 {
											return Any
										} else { // 0x02764 <= r
											if r < 0x02765 {
												return Glue_After_Zwj
											} else { // 0x02765 <= r
												return Any
											}
										}
									} else { // 0x02cef <= r
										if r < 0x02cf2 {
											return Extend
										} else { // 0x02cf2 <= r
											if r < 0x02d7f {
												return Any
											} else { // 0x02d7f <= r
												return Extend
											}
										}
									}
								} else { // 0x02d80 <= r
									if r < 0x0302a {
										if r < 0x02de0 {
											return Any
										} else { // 0x02de0 <= r
											if r < 0x02e00 {
												return Extend
											} else { // 0x02e00 <= r
												return Any
											}
										}
									} else { // 0x0302a <= r
										if r < 0x03030 {
											return Extend
										} else { // 0x03030 <= r
											if r < 0x03099 {
												return Any
											} else { // 0x03099 <= r
												return Extend
											}
										}
									}
								}
							} else { // 0x0309b <= r
								if r < 0x0a6a0 {
									if r < 0x0a674 {
										if r < 0x0a66f {
											return Any
										} else { // 0x0a66f <= r
											if r < 0x0a673 {
												return Extend
											} else { // 0x0a673 <= r
												return Any
											}
										}
									} else { // 0x0a674 <= r
										if r < 0x0a67e {
											return Extend
										} else { // 0x0a67e <= r
											if r < 0x0a69e {
												return Any
											} else { // 0x0a69e <= r
												return Extend
											}
										}
									}
								} else { // 0x0a6a0 <= r
									if r < 0x0a802 {
										if r < 0x0a6f0 {
											return Any
										} else { // 0x0a6f0 <= r
											if r < 0x0a6f2 {
												return Extend
											} else { // 0x0a6f2 <= r
												return Any
											}
										}
									} else { // 0x0a802 <= r
										if r < 0x0a806 {
											if r < 0x0a803 {
												return Extend
											} else { // 0x0a803 <= r
												return Any
											}
										} else { // 0x0a806 <= r
											if r < 0x0a807 {
												return Extend
											} else { // 0x0a807 <= r
												return Any
											}
										}
									}
								}
							}
						}
					} else { // 0x0a80b <= r
						if r < 0x0a9b6 {
							if r < 0x0a8f2 {
								if r < 0x0a880 {
									if r < 0x0a825 {
										if r < 0x0a80c {
											return Extend
										} else { // 0x0a80c <= r
											if r < 0x0a823 {
												return Any
											} else { // 0x0a823 <= r
												return SpacingMark
											}
										}
									} else { // 0x0a825 <= r
										if r < 0x0a827 {
											return Extend
										} else { // 0x0a827 <= r
											if r < 0x0a828 {
												return SpacingMark
											} else { // 0x0a828 <= r
												return Any
											}
										}
									}
								} else { // 0x0a880 <= r
									if r < 0x0a8c4 {
										if r < 0x0a882 {
											return SpacingMark
										} else { // 0x0a882 <= r
											if r < 0x0a8b4 {
												return Any
											} else { // 0x0a8b4 <= r
												return SpacingMark
											}
										}
									} else { // 0x0a8c4 <= r
										if r < 0x0a8c6 {
											return Extend
										} else { // 0x0a8c6 <= r
											if r < 0x0a8e0 {
												return Any
											} else { // 0x0a8e0 <= r
												return Extend
											}
										}
									}
								}
							} else { // 0x0a8f2 <= r
								if r < 0x0a960 {
									if r < 0x0a947 {
										if r < 0x0a926 {
											return Any
										} else { // 0x0a926 <= r
											if r < 0x0a92e {
												return Extend
											} else { // 0x0a92e <= r
												return Any
											}
										}
									} else { // 0x0a947 <= r
										if r < 0x0a952 {
											return Extend
										} else { // 0x0a952 <= r
											if r < 0x0a954 {
												return SpacingMark
											} else { // 0x0a954 <= r
												return Any
											}
										}
									}
								} else { // 0x0a960 <= r
									if r < 0x0a983 {
										if r < 0x0a97d {
											return L
										} else { // 0x0a97d <= r
											if r < 0x0a980 {
												return Any
											} else { // 0x0a980 <= r
												return Extend
											}
										}
									} else { // 0x0a983 <= r
										if r < 0x0a9b3 {
											if r < 0x0a984 {
												return SpacingMark
											} else { // 0x0a984 <= r
												return Any
											}
										} else { // 0x0a9b3 <= r
											if r < 0x0a9b4 {
												return Extend
											} else { // 0x0a9b4 <= r
												return SpacingMark
											}
										}
									}
								}
							}
						} else { // 0x0a9b6 <= r
							if r < 0x0aa43 {
								if r < 0x0a9e6 {
									if r < 0x0a9bd {
										if r < 0x0a9ba {
											return Extend
										} else { // 0x0a9ba <= r
											if r < 0x0a9bc {
												return SpacingMark
											} else { // 0x0a9bc <= r
												return Extend
											}
										}
									} else { // 0x0a9bd <= r
										if r < 0x0a9c1 {
											return SpacingMark
										} else { // 0x0a9c1 <= r
											if r < 0x0a9e5 {
												return Any
											} else { // 0x0a9e5 <= r
												return Extend
											}
										}
									}
								} else { // 0x0a9e6 <= r
									if r < 0x0aa31 {
										if r < 0x0aa29 {
											return Any
										} else { // 0x0aa29 <= r
											if r < 0x0aa2f {
												return Extend
											} else { // 0x0aa2f <= r
												return SpacingMark
											}
										}
									} else { // 0x0aa31 <= r
										if r < 0x0aa35 {
											if r < 0x0aa33 {
												return Extend
											} else { // 0x0aa33 <= r
												return SpacingMark
											}
										} else { // 0x0aa35 <= r
											if r < 0x0aa37 {
												return Extend
											} else { // 0x0aa37 <= r
												return Any
											}
										}
									}
								}
							} else { // 0x0aa43 <= r
								if r < 0x0aa7d {
									if r < 0x0aa4d {
										if r < 0x0aa44 {
											return Extend
										} else { // 0x0aa44 <= r
											if r < 0x0aa4c {
												return Any
											} else { // 0x0aa4c <= r
												return Extend
											}
										}
									} else { // 0x0aa4d <= r
										if r < 0x0aa4e {
											return SpacingMark
										} else { // 0x0aa4e <= r
											if r < 0x0aa7c {
												return Any
											} else { // 0x0aa7c <= r
												return Extend
											}
										}
									}
								} else { // 0x0aa7d <= r
									if r < 0x0aab2 {
										if r < 0x0aab0 {
											return Any
										} else { // 0x0aab0 <= r
											if r < 0x0aab1 {
												return Extend
											} else { // 0x0aab1 <= r
												return Any
											}
										}
									} else { // 0x0aab2 <= r
										if r < 0x0aab7 {
											if r < 0x0aab5 {
												return Extend
											} else { // 0x0aab5 <= r
												return Any
											}
										} else { // 0x0aab7 <= r
											if r < 0x0aab9 {
												return Extend
											} else { // 0x0aab9 <= r
												return Any
											}
										}
									}
								}
							}
						}
					}
				} else { // 0x0aabe <= r
					if r < 0x0ada4 {
						if r < 0x0ac39 {
							if r < 0x0abe5 {
								if r < 0x0aaee {
									if r < 0x0aac2 {
										if r < 0x0aac0 {
											return Extend
										} else { // 0x0aac0 <= r
											if r < 0x0aac1 {
												return Any
											} else { // 0x0aac1 <= r
												return Extend
											}
										}
									} else { // 0x0aac2 <= r
										if r < 0x0aaeb {
											return Any
										} else { // 0x0aaeb <= r
											if r < 0x0aaec {
												return SpacingMark
											} else { // 0x0aaec <= r
												return Extend
											}
										}
									}
								} else { // 0x0aaee <= r
									if r < 0x0aaf6 {
										if r < 0x0aaf0 {
											return SpacingMark
										} else { // 0x0aaf0 <= r
											if r < 0x0aaf5 {
												return Any
											} else { // 0x0aaf5 <= r
												return SpacingMark
											}
										}
									} else { // 0x0aaf6 <= r
										if r < 0x0aaf7 {
											return Extend
										} else { // 0x0aaf7 <= r
											if r < 0x0abe3 {
												return Any
											} else { // 0x0abe3 <= r
												return SpacingMark
											}
										}
									}
								}
							} else { // 0x0abe5 <= r
								if r < 0x0abed {
									if r < 0x0abe9 {
										if r < 0x0abe6 {
											return Extend
										} else { // 0x0abe6 <= r
											if r < 0x0abe8 {
												return SpacingMark
											} else { // 0x0abe8 <= r
												return Extend
											}
										}
									} else { // 0x0abe9 <= r
										if r < 0x0abeb {
											return SpacingMark
										} else { // 0x0abeb <= r
											if r < 0x0abec {
												return Any
											} else { // 0x0abec <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x0abed <= r
									if r < 0x0ac01 {
										if r < 0x0abee {
											return Extend
										} else { // 0x0abee <= r
											if r < 0x0ac00 {
												return Any
											} else { // 0x0ac00 <= r
												return LV
											}
										}
									} else { // 0x0ac01 <= r
										if r < 0x0ac1d {
											if r < 0x0ac1c {
												return LVT
											} else { // 0x0ac1c <= r
												return LV
											}
										} else { // 0x0ac1d <= r
											if r < 0x0ac38 {
												return LVT
											} else { // 0x0ac38 <= r
												return LV
											}
										}
									}
								}
							}
						} else { // 0x0ac39 <= r
							if r < 0x0ace1 {
								if r < 0x0ac8d {
									if r < 0x0ac70 {
										if r < 0x0ac54 {
											return LVT
										} else { // 0x0ac54 <= r
											if r < 0x0ac55 {
												return LV
											} else { // 0x0ac55 <= r
												return LVT
											}
										}
									} else { // 0x0ac70 <= r
										if r < 0x0ac71 {
											return LV
										} else { // 0x0ac71 <= r
											if r < 0x0ac8c {
												return LVT
											} else { // 0x0ac8c <= r
												return LV
											}
										}
									}
								} else { // 0x0ac8d <= r
									if r < 0x0acc4 {
										if r < 0x0aca8 {
											return LVT
										} else { // 0x0aca8 <= r
											if r < 0x0aca9 {
												return LV
											} else { // 0x0aca9 <= r
												return LVT
											}
										}
									} else { // 0x0acc4 <= r
										if r < 0x0acc5 {
											return LV
										} else { // 0x0acc5 <= r
											if r < 0x0ace0 {
												return LVT
											} else { // 0x0ace0 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0ace1 <= r
								if r < 0x0ad35 {
									if r < 0x0ad18 {
										if r < 0x0acfc {
											return LVT
										} else { // 0x0acfc <= r
											if r < 0x0acfd {
												return LV
											} else { // 0x0acfd <= r
												return LVT
											}
										}
									} else { // 0x0ad18 <= r
										if r < 0x0ad19 {
											return LV
										} else { // 0x0ad19 <= r
											if r < 0x0ad34 {
												return LVT
											} else { // 0x0ad34 <= r
												return LV
											}
										}
									}
								} else { // 0x0ad35 <= r
									if r < 0x0ad6c {
										if r < 0x0ad50 {
											return LVT
										} else { // 0x0ad50 <= r
											if r < 0x0ad51 {
												return LV
											} else { // 0x0ad51 <= r
												return LVT
											}
										}
									} else { // 0x0ad6c <= r
										if r < 0x0ad88 {
											if r < 0x0ad6d {
												return LV
											} else { // 0x0ad6d <= r
												return LVT
											}
										} else { // 0x0ad88 <= r
											if r < 0x0ad89 {
												return LV
											} else { // 0x0ad89 <= r
												return LVT
											}
										}
									}
								}
							}
						}
					} else { // 0x0ada4 <= r
						if r < 0x0aef5 {
							if r < 0x0ae4c {
								if r < 0x0adf8 {
									if r < 0x0adc1 {
										if r < 0x0ada5 {
											return LV
										} else { // 0x0ada5 <= r
											if r < 0x0adc0 {
												return LVT
											} else { // 0x0adc0 <= r
												return LV
											}
										}
									} else { // 0x0adc1 <= r
										if r < 0x0addc {
											return LVT
										} else { // 0x0addc <= r
											if r < 0x0addd {
												return LV
											} else { // 0x0addd <= r
												return LVT
											}
										}
									}
								} else { // 0x0adf8 <= r
									if r < 0x0ae15 {
										if r < 0x0adf9 {
											return LV
										} else { // 0x0adf9 <= r
											if r < 0x0ae14 {
												return LVT
											} else { // 0x0ae14 <= r
												return LV
											}
										}
									} else { // 0x0ae15 <= r
										if r < 0x0ae30 {
											return LVT
										} else { // 0x0ae30 <= r
											if r < 0x0ae31 {
												return LV
											} else { // 0x0ae31 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0ae4c <= r
								if r < 0x0aea0 {
									if r < 0x0ae69 {
										if r < 0x0ae4d {
											return LV
										} else { // 0x0ae4d <= r
											if r < 0x0ae68 {
												return LVT
											} else { // 0x0ae68 <= r
												return LV
											}
										}
									} else { // 0x0ae69 <= r
										if r < 0x0ae84 {
											return LVT
										} else { // 0x0ae84 <= r
											if r < 0x0ae85 {
												return LV
											} else { // 0x0ae85 <= r
												return LVT
											}
										}
									}
								} else { // 0x0aea0 <= r
									if r < 0x0aebd {
										if r < 0x0aea1 {
											return LV
										} else { // 0x0aea1 <= r
											if r < 0x0aebc {
												return LVT
											} else { // 0x0aebc <= r
												return LV
											}
										}
									} else { // 0x0aebd <= r
										if r < 0x0aed9 {
											if r < 0x0aed8 {
												return LVT
											} else { // 0x0aed8 <= r
												return LV
											}
										} else { // 0x0aed9 <= r
											if r < 0x0aef4 {
												return LVT
											} else { // 0x0aef4 <= r
												return LV
											}
										}
									}
								}
							}
						} else { // 0x0aef5 <= r
							if r < 0x0afb8 {
								if r < 0x0af49 {
									if r < 0x0af2c {
										if r < 0x0af10 {
											return LVT
										} else { // 0x0af10 <= r
											if r < 0x0af11 {
												return LV
											} else { // 0x0af11 <= r
												return LVT
											}
										}
									} else { // 0x0af2c <= r
										if r < 0x0af2d {
											return LV
										} else { // 0x0af2d <= r
											if r < 0x0af48 {
												return LVT
											} else { // 0x0af48 <= r
												return LV
											}
										}
									}
								} else { // 0x0af49 <= r
									if r < 0x0af80 {
										if r < 0x0af64 {
											return LVT
										} else { // 0x0af64 <= r
											if r < 0x0af65 {
												return LV
											} else { // 0x0af65 <= r
												return LVT
											}
										}
									} else { // 0x0af80 <= r
										if r < 0x0af9c {
											if r < 0x0af81 {
												return LV
											} else { // 0x0af81 <= r
												return LVT
											}
										} else { // 0x0af9c <= r
											if r < 0x0af9d {
												return LV
											} else { // 0x0af9d <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0afb8 <= r
								if r < 0x0b00c {
									if r < 0x0afd5 {
										if r < 0x0afb9 {
											return LV
										} else { // 0x0afb9 <= r
											if r < 0x0afd4 {
												return LVT
											} else { // 0x0afd4 <= r
												return LV
											}
										}
									} else { // 0x0afd5 <= r
										if r < 0x0aff0 {
											return LVT
										} else { // 0x0aff0 <= r
											if r < 0x0aff1 {
												return LV
											} else { // 0x0aff1 <= r
												return LVT
											}
										}
									}
								} else { // 0x0b00c <= r
									if r < 0x0b029 {
										if r < 0x0b00d {
											return LV
										} else { // 0x0b00d <= r
											if r < 0x0b028 {
												return LVT
											} else { // 0x0b028 <= r
												return LV
											}
										}
									} else { // 0x0b029 <= r
										if r < 0x0b045 {
											if r < 0x0b044 {
												return LVT
											} else { // 0x0b044 <= r
												return LV
											}
										} else { // 0x0b045 <= r
											if r < 0x0b060 {
												return LVT
											} else { // 0x0b060 <= r
												return LV
											}
										}
									}
								}
							}
						}
					}
				}
			} else { // 0x0b061 <= r
				if r < 0x0b5f4 {
					if r < 0x0b31d {
						if r < 0x0b1cc {
							if r < 0x0b109 {
								if r < 0x0b0b5 {
									if r < 0x0b098 {
										if r < 0x0b07c {
											return LVT
										} else { // 0x0b07c <= r
											if r < 0x0b07d {
												return LV
											} else { // 0x0b07d <= r
												return LVT
											}
										}
									} else { // 0x0b098 <= r
										if r < 0x0b099 {
											return LV
										} else { // 0x0b099 <= r
											if r < 0x0b0b4 {
												return LVT
											} else { // 0x0b0b4 <= r
												return LV
											}
										}
									}
								} else { // 0x0b0b5 <= r
									if r < 0x0b0ec {
										if r < 0x0b0d0 {
											return LVT
										} else { // 0x0b0d0 <= r
											if r < 0x0b0d1 {
												return LV
											} else { // 0x0b0d1 <= r
												return LVT
											}
										}
									} else { // 0x0b0ec <= r
										if r < 0x0b0ed {
											return LV
										} else { // 0x0b0ed <= r
											if r < 0x0b108 {
												return LVT
											} else { // 0x0b108 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0b109 <= r
								if r < 0x0b15d {
									if r < 0x0b140 {
										if r < 0x0b124 {
											return LVT
										} else { // 0x0b124 <= r
											if r < 0x0b125 {
												return LV
											} else { // 0x0b125 <= r
												return LVT
											}
										}
									} else { // 0x0b140 <= r
										if r < 0x0b141 {
											return LV
										} else { // 0x0b141 <= r
											if r < 0x0b15c {
												return LVT
											} else { // 0x0b15c <= r
												return LV
											}
										}
									}
								} else { // 0x0b15d <= r
									if r < 0x0b194 {
										if r < 0x0b178 {
											return LVT
										} else { // 0x0b178 <= r
											if r < 0x0b179 {
												return LV
											} else { // 0x0b179 <= r
												return LVT
											}
										}
									} else { // 0x0b194 <= r
										if r < 0x0b1b0 {
											if r < 0x0b195 {
												return LV
											} else { // 0x0b195 <= r
												return LVT
											}
										} else { // 0x0b1b0 <= r
											if r < 0x0b1b1 {
												return LV
											} else { // 0x0b1b1 <= r
												return LVT
											}
										}
									}
								}
							}
						} else { // 0x0b1cc <= r
							if r < 0x0b274 {
								if r < 0x0b220 {
									if r < 0x0b1e9 {
										if r < 0x0b1cd {
											return LV
										} else { // 0x0b1cd <= r
											if r < 0x0b1e8 {
												return LVT
											} else { // 0x0b1e8 <= r
												return LV
											}
										}
									} else { // 0x0b1e9 <= r
										if r < 0x0b204 {
											return LVT
										} else { // 0x0b204 <= r
											if r < 0x0b205 {
												return LV
											} else { // 0x0b205 <= r
												return LVT
											}
										}
									}
								} else { // 0x0b220 <= r
									if r < 0x0b23d {
										if r < 0x0b221 {
											return LV
										} else { // 0x0b221 <= r
											if r < 0x0b23c {
												return LVT
											} else { // 0x0b23c <= r
												return LV
											}
										}
									} else { // 0x0b23d <= r
										if r < 0x0b258 {
											return LVT
										} else { // 0x0b258 <= r
											if r < 0x0b259 {
												return LV
											} else { // 0x0b259 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0b274 <= r
								if r < 0x0b2c8 {
									if r < 0x0b291 {
										if r < 0x0b275 {
											return LV
										} else { // 0x0b275 <= r
											if r < 0x0b290 {
												return LVT
											} else { // 0x0b290 <= r
												return LV
											}
										}
									} else { // 0x0b291 <= r
										if r < 0x0b2ac {
											return LVT
										} else { // 0x0b2ac <= r
											if r < 0x0b2ad {
												return LV
											} else { // 0x0b2ad <= r
												return LVT
											}
										}
									}
								} else { // 0x0b2c8 <= r
									if r < 0x0b2e5 {
										if r < 0x0b2c9 {
											return LV
										} else { // 0x0b2c9 <= r
											if r < 0x0b2e4 {
												return LVT
											} else { // 0x0b2e4 <= r
												return LV
											}
										}
									} else { // 0x0b2e5 <= r
										if r < 0x0b301 {
											if r < 0x0b300 {
												return LVT
											} else { // 0x0b300 <= r
												return LV
											}
										} else { // 0x0b301 <= r
											if r < 0x0b31c {
												return LVT
											} else { // 0x0b31c <= r
												return LV
											}
										}
									}
								}
							}
						}
					} else { // 0x0b31d <= r
						if r < 0x0b488 {
							if r < 0x0b3c5 {
								if r < 0x0b371 {
									if r < 0x0b354 {
										if r < 0x0b338 {
											return LVT
										} else { // 0x0b338 <= r
											if r < 0x0b339 {
												return LV
											} else { // 0x0b339 <= r
												return LVT
											}
										}
									} else { // 0x0b354 <= r
										if r < 0x0b355 {
											return LV
										} else { // 0x0b355 <= r
											if r < 0x0b370 {
												return LVT
											} else { // 0x0b370 <= r
												return LV
											}
										}
									}
								} else { // 0x0b371 <= r
									if r < 0x0b3a8 {
										if r < 0x0b38c {
											return LVT
										} else { // 0x0b38c <= r
											if r < 0x0b38d {
												return LV
											} else { // 0x0b38d <= r
												return LVT
											}
										}
									} else { // 0x0b3a8 <= r
										if r < 0x0b3a9 {
											return LV
										} else { // 0x0b3a9 <= r
											if r < 0x0b3c4 {
												return LVT
											} else { // 0x0b3c4 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0b3c5 <= r
								if r < 0x0b419 {
									if r < 0x0b3fc {
										if r < 0x0b3e0 {
											return LVT
										} else { // 0x0b3e0 <= r
											if r < 0x0b3e1 {
												return LV
											} else { // 0x0b3e1 <= r
												return LVT
											}
										}
									} else { // 0x0b3fc <= r
										if r < 0x0b3fd {
											return LV
										} else { // 0x0b3fd <= r
											if r < 0x0b418 {
												return LVT
											} else { // 0x0b418 <= r
												return LV
											}
										}
									}
								} else { // 0x0b419 <= r
									if r < 0x0b450 {
										if r < 0x0b434 {
											return LVT
										} else { // 0x0b434 <= r
											if r < 0x0b435 {
												return LV
											} else { // 0x0b435 <= r
												return LVT
											}
										}
									} else { // 0x0b450 <= r
										if r < 0x0b46c {
											if r < 0x0b451 {
												return LV
											} else { // 0x0b451 <= r
												return LVT
											}
										} else { // 0x0b46c <= r
											if r < 0x0b46d {
												return LV
											} else { // 0x0b46d <= r
												return LVT
											}
										}
									}
								}
							}
						} else { // 0x0b488 <= r
							if r < 0x0b531 {
								if r < 0x0b4dc {
									if r < 0x0b4a5 {
										if r < 0x0b489 {
											return LV
										} else { // 0x0b489 <= r
											if r < 0x0b4a4 {
												return LVT
											} else { // 0x0b4a4 <= r
												return LV
											}
										}
									} else { // 0x0b4a5 <= r
										if r < 0x0b4c0 {
											return LVT
										} else { // 0x0b4c0 <= r
											if r < 0x0b4c1 {
												return LV
											} else { // 0x0b4c1 <= r
												return LVT
											}
										}
									}
								} else { // 0x0b4dc <= r
									if r < 0x0b4f9 {
										if r < 0x0b4dd {
											return LV
										} else { // 0x0b4dd <= r
											if r < 0x0b4f8 {
												return LVT
											} else { // 0x0b4f8 <= r
												return LV
											}
										}
									} else { // 0x0b4f9 <= r
										if r < 0x0b515 {
											if r < 0x0b514 {
												return LVT
											} else { // 0x0b514 <= r
												return LV
											}
										} else { // 0x0b515 <= r
											if r < 0x0b530 {
												return LVT
											} else { // 0x0b530 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0b531 <= r
								if r < 0x0b585 {
									if r < 0x0b568 {
										if r < 0x0b54c {
											return LVT
										} else { // 0x0b54c <= r
											if r < 0x0b54d {
												return LV
											} else { // 0x0b54d <= r
												return LVT
											}
										}
									} else { // 0x0b568 <= r
										if r < 0x0b569 {
											return LV
										} else { // 0x0b569 <= r
											if r < 0x0b584 {
												return LVT
											} else { // 0x0b584 <= r
												return LV
											}
										}
									}
								} else { // 0x0b585 <= r
									if r < 0x0b5bc {
										if r < 0x0b5a0 {
											return LVT
										} else { // 0x0b5a0 <= r
											if r < 0x0b5a1 {
												return LV
											} else { // 0x0b5a1 <= r
												return LVT
											}
										}
									} else { // 0x0b5bc <= r
										if r < 0x0b5d8 {
											if r < 0x0b5bd {
												return LV
											} else { // 0x0b5bd <= r
												return LVT
											}
										} else { // 0x0b5d8 <= r
											if r < 0x0b5d9 {
												return LV
											} else { // 0x0b5d9 <= r
												return LVT
											}
										}
									}
								}
							}
						}
					}
				} else { // 0x0b5f4 <= r
					if r < 0x0b8b0 {
						if r < 0x0b745 {
							if r < 0x0b69c {
								if r < 0x0b648 {
									if r < 0x0b611 {
										if r < 0x0b5f5 {
											return LV
										} else { // 0x0b5f5 <= r
											if r < 0x0b610 {
												return LVT
											} else { // 0x0b610 <= r
												return LV
											}
										}
									} else { // 0x0b611 <= r
										if r < 0x0b62c {
											return LVT
										} else { // 0x0b62c <= r
											if r < 0x0b62d {
												return LV
											} else { // 0x0b62d <= r
												return LVT
											}
										}
									}
								} else { // 0x0b648 <= r
									if r < 0x0b665 {
										if r < 0x0b649 {
											return LV
										} else { // 0x0b649 <= r
											if r < 0x0b664 {
												return LVT
											} else { // 0x0b664 <= r
												return LV
											}
										}
									} else { // 0x0b665 <= r
										if r < 0x0b680 {
											return LVT
										} else { // 0x0b680 <= r
											if r < 0x0b681 {
												return LV
											} else { // 0x0b681 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0b69c <= r
								if r < 0x0b6f0 {
									if r < 0x0b6b9 {
										if r < 0x0b69d {
											return LV
										} else { // 0x0b69d <= r
											if r < 0x0b6b8 {
												return LVT
											} else { // 0x0b6b8 <= r
												return LV
											}
										}
									} else { // 0x0b6b9 <= r
										if r < 0x0b6d4 {
											return LVT
										} else { // 0x0b6d4 <= r
											if r < 0x0b6d5 {
												return LV
											} else { // 0x0b6d5 <= r
												return LVT
											}
										}
									}
								} else { // 0x0b6f0 <= r
									if r < 0x0b70d {
										if r < 0x0b6f1 {
											return LV
										} else { // 0x0b6f1 <= r
											if r < 0x0b70c {
												return LVT
											} else { // 0x0b70c <= r
												return LV
											}
										}
									} else { // 0x0b70d <= r
										if r < 0x0b729 {
											if r < 0x0b728 {
												return LVT
											} else { // 0x0b728 <= r
												return LV
											}
										} else { // 0x0b729 <= r
											if r < 0x0b744 {
												return LVT
											} else { // 0x0b744 <= r
												return LV
											}
										}
									}
								}
							}
						} else { // 0x0b745 <= r
							if r < 0x0b7ed {
								if r < 0x0b799 {
									if r < 0x0b77c {
										if r < 0x0b760 {
											return LVT
										} else { // 0x0b760 <= r
											if r < 0x0b761 {
												return LV
											} else { // 0x0b761 <= r
												return LVT
											}
										}
									} else { // 0x0b77c <= r
										if r < 0x0b77d {
											return LV
										} else { // 0x0b77d <= r
											if r < 0x0b798 {
												return LVT
											} else { // 0x0b798 <= r
												return LV
											}
										}
									}
								} else { // 0x0b799 <= r
									if r < 0x0b7d0 {
										if r < 0x0b7b4 {
											return LVT
										} else { // 0x0b7b4 <= r
											if r < 0x0b7b5 {
												return LV
											} else { // 0x0b7b5 <= r
												return LVT
											}
										}
									} else { // 0x0b7d0 <= r
										if r < 0x0b7d1 {
											return LV
										} else { // 0x0b7d1 <= r
											if r < 0x0b7ec {
												return LVT
											} else { // 0x0b7ec <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0b7ed <= r
								if r < 0x0b841 {
									if r < 0x0b824 {
										if r < 0x0b808 {
											return LVT
										} else { // 0x0b808 <= r
											if r < 0x0b809 {
												return LV
											} else { // 0x0b809 <= r
												return LVT
											}
										}
									} else { // 0x0b824 <= r
										if r < 0x0b825 {
											return LV
										} else { // 0x0b825 <= r
											if r < 0x0b840 {
												return LVT
											} else { // 0x0b840 <= r
												return LV
											}
										}
									}
								} else { // 0x0b841 <= r
									if r < 0x0b878 {
										if r < 0x0b85c {
											return LVT
										} else { // 0x0b85c <= r
											if r < 0x0b85d {
												return LV
											} else { // 0x0b85d <= r
												return LVT
											}
										}
									} else { // 0x0b878 <= r
										if r < 0x0b894 {
											if r < 0x0b879 {
												return LV
											} else { // 0x0b879 <= r
												return LVT
											}
										} else { // 0x0b894 <= r
											if r < 0x0b895 {
												return LV
											} else { // 0x0b895 <= r
												return LVT
											}
										}
									}
								}
							}
						}
					} else { // 0x0b8b0 <= r
						if r < 0x0ba01 {
							if r < 0x0b958 {
								if r < 0x0b904 {
									if r < 0x0b8cd {
										if r < 0x0b8b1 {
											return LV
										} else { // 0x0b8b1 <= r
											if r < 0x0b8cc {
												return LVT
											} else { // 0x0b8cc <= r
												return LV
											}
										}
									} else { // 0x0b8cd <= r
										if r < 0x0b8e8 {
											return LVT
										} else { // 0x0b8e8 <= r
											if r < 0x0b8e9 {
												return LV
											} else { // 0x0b8e9 <= r
												return LVT
											}
										}
									}
								} else { // 0x0b904 <= r
									if r < 0x0b921 {
										if r < 0x0b905 {
											return LV
										} else { // 0x0b905 <= r
											if r < 0x0b920 {
												return LVT
											} else { // 0x0b920 <= r
												return LV
											}
										}
									} else { // 0x0b921 <= r
										if r < 0x0b93c {
											return LVT
										} else { // 0x0b93c <= r
											if r < 0x0b93d {
												return LV
											} else { // 0x0b93d <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0b958 <= r
								if r < 0x0b9ac {
									if r < 0x0b975 {
										if r < 0x0b959 {
											return LV
										} else { // 0x0b959 <= r
											if r < 0x0b974 {
												return LVT
											} else { // 0x0b974 <= r
												return LV
											}
										}
									} else { // 0x0b975 <= r
										if r < 0x0b990 {
											return LVT
										} else { // 0x0b990 <= r
											if r < 0x0b991 {
												return LV
											} else { // 0x0b991 <= r
												return LVT
											}
										}
									}
								} else { // 0x0b9ac <= r
									if r < 0x0b9c9 {
										if r < 0x0b9ad {
											return LV
										} else { // 0x0b9ad <= r
											if r < 0x0b9c8 {
												return LVT
											} else { // 0x0b9c8 <= r
												return LV
											}
										}
									} else { // 0x0b9c9 <= r
										if r < 0x0b9e5 {
											if r < 0x0b9e4 {
												return LVT
											} else { // 0x0b9e4 <= r
												return LV
											}
										} else { // 0x0b9e5 <= r
											if r < 0x0ba00 {
												return LVT
											} else { // 0x0ba00 <= r
												return LV
											}
										}
									}
								}
							}
						} else { // 0x0ba01 <= r
							if r < 0x0bac4 {
								if r < 0x0ba55 {
									if r < 0x0ba38 {
										if r < 0x0ba1c {
											return LVT
										} else { // 0x0ba1c <= r
											if r < 0x0ba1d {
												return LV
											} else { // 0x0ba1d <= r
												return LVT
											}
										}
									} else { // 0x0ba38 <= r
										if r < 0x0ba39 {
											return LV
										} else { // 0x0ba39 <= r
											if r < 0x0ba54 {
												return LVT
											} else { // 0x0ba54 <= r
												return LV
											}
										}
									}
								} else { // 0x0ba55 <= r
									if r < 0x0ba8c {
										if r < 0x0ba70 {
											return LVT
										} else { // 0x0ba70 <= r
											if r < 0x0ba71 {
												return LV
											} else { // 0x0ba71 <= r
												return LVT
											}
										}
									} else { // 0x0ba8c <= r
										if r < 0x0baa8 {
											if r < 0x0ba8d {
												return LV
											} else { // 0x0ba8d <= r
												return LVT
											}
										} else { // 0x0baa8 <= r
											if r < 0x0baa9 {
												return LV
											} else { // 0x0baa9 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0bac4 <= r
								if r < 0x0bb18 {
									if r < 0x0bae1 {
										if r < 0x0bac5 {
											return LV
										} else { // 0x0bac5 <= r
											if r < 0x0bae0 {
												return LVT
											} else { // 0x0bae0 <= r
												return LV
											}
										}
									} else { // 0x0bae1 <= r
										if r < 0x0bafc {
											return LVT
										} else { // 0x0bafc <= r
											if r < 0x0bafd {
												return LV
											} else { // 0x0bafd <= r
												return LVT
											}
										}
									}
								} else { // 0x0bb18 <= r
									if r < 0x0bb35 {
										if r < 0x0bb19 {
											return LV
										} else { // 0x0bb19 <= r
											if r < 0x0bb34 {
												return LVT
											} else { // 0x0bb34 <= r
												return LV
											}
										}
									} else { // 0x0bb35 <= r
										if r < 0x0bb51 {
											if r < 0x0bb50 {
												return LVT
											} else { // 0x0bb50 <= r
												return LV
											}
										} else { // 0x0bb51 <= r
											if r < 0x0bb6c {
												return LVT
											} else { // 0x0bb6c <= r
												return LV
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else { // 0x0bb6d <= r
		if r < 0x0d185 {
			if r < 0x0c679 {
				if r < 0x0c100 {
					if r < 0x0be29 {
						if r < 0x0bcd8 {
							if r < 0x0bc15 {
								if r < 0x0bbc1 {
									if r < 0x0bba4 {
										if r < 0x0bb88 {
											return LVT
										} else { // 0x0bb88 <= r
											if r < 0x0bb89 {
												return LV
											} else { // 0x0bb89 <= r
												return LVT
											}
										}
									} else { // 0x0bba4 <= r
										if r < 0x0bba5 {
											return LV
										} else { // 0x0bba5 <= r
											if r < 0x0bbc0 {
												return LVT
											} else { // 0x0bbc0 <= r
												return LV
											}
										}
									}
								} else { // 0x0bbc1 <= r
									if r < 0x0bbf8 {
										if r < 0x0bbdc {
											return LVT
										} else { // 0x0bbdc <= r
											if r < 0x0bbdd {
												return LV
											} else { // 0x0bbdd <= r
												return LVT
											}
										}
									} else { // 0x0bbf8 <= r
										if r < 0x0bbf9 {
											return LV
										} else { // 0x0bbf9 <= r
											if r < 0x0bc14 {
												return LVT
											} else { // 0x0bc14 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0bc15 <= r
								if r < 0x0bc69 {
									if r < 0x0bc4c {
										if r < 0x0bc30 {
											return LVT
										} else { // 0x0bc30 <= r
											if r < 0x0bc31 {
												return LV
											} else { // 0x0bc31 <= r
												return LVT
											}
										}
									} else { // 0x0bc4c <= r
										if r < 0x0bc4d {
											return LV
										} else { // 0x0bc4d <= r
											if r < 0x0bc68 {
												return LVT
											} else { // 0x0bc68 <= r
												return LV
											}
										}
									}
								} else { // 0x0bc69 <= r
									if r < 0x0bca0 {
										if r < 0x0bc84 {
											return LVT
										} else { // 0x0bc84 <= r
											if r < 0x0bc85 {
												return LV
											} else { // 0x0bc85 <= r
												return LVT
											}
										}
									} else { // 0x0bca0 <= r
										if r < 0x0bcbc {
											if r < 0x0bca1 {
												return LV
											} else { // 0x0bca1 <= r
												return LVT
											}
										} else { // 0x0bcbc <= r
											if r < 0x0bcbd {
												return LV
											} else { // 0x0bcbd <= r
												return LVT
											}
										}
									}
								}
							}
						} else { // 0x0bcd8 <= r
							if r < 0x0bd80 {
								if r < 0x0bd2c {
									if r < 0x0bcf5 {
										if r < 0x0bcd9 {
											return LV
										} else { // 0x0bcd9 <= r
											if r < 0x0bcf4 {
												return LVT
											} else { // 0x0bcf4 <= r
												return LV
											}
										}
									} else { // 0x0bcf5 <= r
										if r < 0x0bd10 {
											return LVT
										} else { // 0x0bd10 <= r
											if r < 0x0bd11 {
												return LV
											} else { // 0x0bd11 <= r
												return LVT
											}
										}
									}
								} else { // 0x0bd2c <= r
									if r < 0x0bd49 {
										if r < 0x0bd2d {
											return LV
										} else { // 0x0bd2d <= r
											if r < 0x0bd48 {
												return LVT
											} else { // 0x0bd48 <= r
												return LV
											}
										}
									} else { // 0x0bd49 <= r
										if r < 0x0bd64 {
											return LVT
										} else { // 0x0bd64 <= r
											if r < 0x0bd65 {
												return LV
											} else { // 0x0bd65 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0bd80 <= r
								if r < 0x0bdd4 {
									if r < 0x0bd9d {
										if r < 0x0bd81 {
											return LV
										} else { // 0x0bd81 <= r
											if r < 0x0bd9c {
												return LVT
											} else { // 0x0bd9c <= r
												return LV
											}
										}
									} else { // 0x0bd9d <= r
										if r < 0x0bdb8 {
											return LVT
										} else { // 0x0bdb8 <= r
											if r < 0x0bdb9 {
												return LV
											} else { // 0x0bdb9 <= r
												return LVT
											}
										}
									}
								} else { // 0x0bdd4 <= r
									if r < 0x0bdf1 {
										if r < 0x0bdd5 {
											return LV
										} else { // 0x0bdd5 <= r
											if r < 0x0bdf0 {
												return LVT
											} else { // 0x0bdf0 <= r
												return LV
											}
										}
									} else { // 0x0bdf1 <= r
										if r < 0x0be0d {
											if r < 0x0be0c {
												return LVT
											} else { // 0x0be0c <= r
												return LV
											}
										} else { // 0x0be0d <= r
											if r < 0x0be28 {
												return LVT
											} else { // 0x0be28 <= r
												return LV
											}
										}
									}
								}
							}
						}
					} else { // 0x0be29 <= r
						if r < 0x0bf94 {
							if r < 0x0bed1 {
								if r < 0x0be7d {
									if r < 0x0be60 {
										if r < 0x0be44 {
											return LVT
										} else { // 0x0be44 <= r
											if r < 0x0be45 {
												return LV
											} else { // 0x0be45 <= r
												return LVT
											}
										}
									} else { // 0x0be60 <= r
										if r < 0x0be61 {
											return LV
										} else { // 0x0be61 <= r
											if r < 0x0be7c {
												return LVT
											} else { // 0x0be7c <= r
												return LV
											}
										}
									}
								} else { // 0x0be7d <= r
									if r < 0x0beb4 {
										if r < 0x0be98 {
											return LVT
										} else { // 0x0be98 <= r
											if r < 0x0be99 {
												return LV
											} else { // 0x0be99 <= r
												return LVT
											}
										}
									} else { // 0x0beb4 <= r
										if r < 0x0beb5 {
											return LV
										} else { // 0x0beb5 <= r
											if r < 0x0bed0 {
												return LVT
											} else { // 0x0bed0 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0bed1 <= r
								if r < 0x0bf25 {
									if r < 0x0bf08 {
										if r < 0x0beec {
											return LVT
										} else { // 0x0beec <= r
											if r < 0x0beed {
												return LV
											} else { // 0x0beed <= r
												return LVT
											}
										}
									} else { // 0x0bf08 <= r
										if r < 0x0bf09 {
											return LV
										} else { // 0x0bf09 <= r
											if r < 0x0bf24 {
												return LVT
											} else { // 0x0bf24 <= r
												return LV
											}
										}
									}
								} else { // 0x0bf25 <= r
									if r < 0x0bf5c {
										if r < 0x0bf40 {
											return LVT
										} else { // 0x0bf40 <= r
											if r < 0x0bf41 {
												return LV
											} else { // 0x0bf41 <= r
												return LVT
											}
										}
									} else { // 0x0bf5c <= r
										if r < 0x0bf78 {
											if r < 0x0bf5d {
												return LV
											} else { // 0x0bf5d <= r
												return LVT
											}
										} else { // 0x0bf78 <= r
											if r < 0x0bf79 {
												return LV
											} else { // 0x0bf79 <= r
												return LVT
											}
										}
									}
								}
							}
						} else { // 0x0bf94 <= r
							if r < 0x0c03d {
								if r < 0x0bfe8 {
									if r < 0x0bfb1 {
										if r < 0x0bf95 {
											return LV
										} else { // 0x0bf95 <= r
											if r < 0x0bfb0 {
												return LVT
											} else { // 0x0bfb0 <= r
												return LV
											}
										}
									} else { // 0x0bfb1 <= r
										if r < 0x0bfcc {
											return LVT
										} else { // 0x0bfcc <= r
											if r < 0x0bfcd {
												return LV
											} else { // 0x0bfcd <= r
												return LVT
											}
										}
									}
								} else { // 0x0bfe8 <= r
									if r < 0x0c005 {
										if r < 0x0bfe9 {
											return LV
										} else { // 0x0bfe9 <= r
											if r < 0x0c004 {
												return LVT
											} else { // 0x0c004 <= r
												return LV
											}
										}
									} else { // 0x0c005 <= r
										if r < 0x0c021 {
											if r < 0x0c020 {
												return LVT
											} else { // 0x0c020 <= r
												return LV
											}
										} else { // 0x0c021 <= r
											if r < 0x0c03c {
												return LVT
											} else { // 0x0c03c <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0c03d <= r
								if r < 0x0c091 {
									if r < 0x0c074 {
										if r < 0x0c058 {
											return LVT
										} else { // 0x0c058 <= r
											if r < 0x0c059 {
												return LV
											} else { // 0x0c059 <= r
												return LVT
											}
										}
									} else { // 0x0c074 <= r
										if r < 0x0c075 {
											return LV
										} else { // 0x0c075 <= r
											if r < 0x0c090 {
												return LVT
											} else { // 0x0c090 <= r
												return LV
											}
										}
									}
								} else { // 0x0c091 <= r
									if r < 0x0c0c8 {
										if r < 0x0c0ac {
											return LVT
										} else { // 0x0c0ac <= r
											if r < 0x0c0ad {
												return LV
											} else { // 0x0c0ad <= r
												return LVT
											}
										}
									} else { // 0x0c0c8 <= r
										if r < 0x0c0e4 {
											if r < 0x0c0c9 {
												return LV
											} else { // 0x0c0c9 <= r
												return LVT
											}
										} else { // 0x0c0e4 <= r
											if r < 0x0c0e5 {
												return LV
											} else { // 0x0c0e5 <= r
												return LVT
											}
										}
									}
								}
							}
						}
					}
				} else { // 0x0c100 <= r
					if r < 0x0c3bc {
						if r < 0x0c251 {
							if r < 0x0c1a8 {
								if r < 0x0c154 {
									if r < 0x0c11d {
										if r < 0x0c101 {
											return LV
										} else { // 0x0c101 <= r
											if r < 0x0c11c {
												return LVT
											} else { // 0x0c11c <= r
												return LV
											}
										}
									} else { // 0x0c11d <= r
										if r < 0x0c138 {
											return LVT
										} else { // 0x0c138 <= r
											if r < 0x0c139 {
												return LV
											} else { // 0x0c139 <= r
												return LVT
											}
										}
									}
								} else { // 0x0c154 <= r
									if r < 0x0c171 {
										if r < 0x0c155 {
											return LV
										} else { // 0x0c155 <= r
											if r < 0x0c170 {
												return LVT
											} else { // 0x0c170 <= r
												return LV
											}
										}
									} else { // 0x0c171 <= r
										if r < 0x0c18c {
											return LVT
										} else { // 0x0c18c <= r
											if r < 0x0c18d {
												return LV
											} else { // 0x0c18d <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0c1a8 <= r
								if r < 0x0c1fc {
									if r < 0x0c1c5 {
										if r < 0x0c1a9 {
											return LV
										} else { // 0x0c1a9 <= r
											if r < 0x0c1c4 {
												return LVT
											} else { // 0x0c1c4 <= r
												return LV
											}
										}
									} else { // 0x0c1c5 <= r
										if r < 0x0c1e0 {
											return LVT
										} else { // 0x0c1e0 <= r
											if r < 0x0c1e1 {
												return LV
											} else { // 0x0c1e1 <= r
												return LVT
											}
										}
									}
								} else { // 0x0c1fc <= r
									if r < 0x0c219 {
										if r < 0x0c1fd {
											return LV
										} else { // 0x0c1fd <= r
											if r < 0x0c218 {
												return LVT
											} else { // 0x0c218 <= r
												return LV
											}
										}
									} else { // 0x0c219 <= r
										if r < 0x0c235 {
											if r < 0x0c234 {
												return LVT
											} else { // 0x0c234 <= r
												return LV
											}
										} else { // 0x0c235 <= r
											if r < 0x0c250 {
												return LVT
											} else { // 0x0c250 <= r
												return LV
											}
										}
									}
								}
							}
						} else { // 0x0c251 <= r
							if r < 0x0c2f9 {
								if r < 0x0c2a5 {
									if r < 0x0c288 {
										if r < 0x0c26c {
											return LVT
										} else { // 0x0c26c <= r
											if r < 0x0c26d {
												return LV
											} else { // 0x0c26d <= r
												return LVT
											}
										}
									} else { // 0x0c288 <= r
										if r < 0x0c289 {
											return LV
										} else { // 0x0c289 <= r
											if r < 0x0c2a4 {
												return LVT
											} else { // 0x0c2a4 <= r
												return LV
											}
										}
									}
								} else { // 0x0c2a5 <= r
									if r < 0x0c2dc {
										if r < 0x0c2c0 {
											return LVT
										} else { // 0x0c2c0 <= r
											if r < 0x0c2c1 {
												return LV
											} else { // 0x0c2c1 <= r
												return LVT
											}
										}
									} else { // 0x0c2dc <= r
										if r < 0x0c2dd {
											return LV
										} else { // 0x0c2dd <= r
											if r < 0x0c2f8 {
												return LVT
											} else { // 0x0c2f8 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0c2f9 <= r
								if r < 0x0c34d {
									if r < 0x0c330 {
										if r < 0x0c314 {
											return LVT
										} else { // 0x0c314 <= r
											if r < 0x0c315 {
												return LV
											} else { // 0x0c315 <= r
												return LVT
											}
										}
									} else { // 0x0c330 <= r
										if r < 0x0c331 {
											return LV
										} else { // 0x0c331 <= r
											if r < 0x0c34c {
												return LVT
											} else { // 0x0c34c <= r
												return LV
											}
										}
									}
								} else { // 0x0c34d <= r
									if r < 0x0c384 {
										if r < 0x0c368 {
											return LVT
										} else { // 0x0c368 <= r
											if r < 0x0c369 {
												return LV
											} else { // 0x0c369 <= r
												return LVT
											}
										}
									} else { // 0x0c384 <= r
										if r < 0x0c3a0 {
											if r < 0x0c385 {
												return LV
											} else { // 0x0c385 <= r
												return LVT
											}
										} else { // 0x0c3a0 <= r
											if r < 0x0c3a1 {
												return LV
											} else { // 0x0c3a1 <= r
												return LVT
											}
										}
									}
								}
							}
						}
					} else { // 0x0c3bc <= r
						if r < 0x0c50d {
							if r < 0x0c464 {
								if r < 0x0c410 {
									if r < 0x0c3d9 {
										if r < 0x0c3bd {
											return LV
										} else { // 0x0c3bd <= r
											if r < 0x0c3d8 {
												return LVT
											} else { // 0x0c3d8 <= r
												return LV
											}
										}
									} else { // 0x0c3d9 <= r
										if r < 0x0c3f4 {
											return LVT
										} else { // 0x0c3f4 <= r
											if r < 0x0c3f5 {
												return LV
											} else { // 0x0c3f5 <= r
												return LVT
											}
										}
									}
								} else { // 0x0c410 <= r
									if r < 0x0c42d {
										if r < 0x0c411 {
											return LV
										} else { // 0x0c411 <= r
											if r < 0x0c42c {
												return LVT
											} else { // 0x0c42c <= r
												return LV
											}
										}
									} else { // 0x0c42d <= r
										if r < 0x0c448 {
											return LVT
										} else { // 0x0c448 <= r
											if r < 0x0c449 {
												return LV
											} else { // 0x0c449 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0c464 <= r
								if r < 0x0c4b8 {
									if r < 0x0c481 {
										if r < 0x0c465 {
											return LV
										} else { // 0x0c465 <= r
											if r < 0x0c480 {
												return LVT
											} else { // 0x0c480 <= r
												return LV
											}
										}
									} else { // 0x0c481 <= r
										if r < 0x0c49c {
											return LVT
										} else { // 0x0c49c <= r
											if r < 0x0c49d {
												return LV
											} else { // 0x0c49d <= r
												return LVT
											}
										}
									}
								} else { // 0x0c4b8 <= r
									if r < 0x0c4d5 {
										if r < 0x0c4b9 {
											return LV
										} else { // 0x0c4b9 <= r
											if r < 0x0c4d4 {
												return LVT
											} else { // 0x0c4d4 <= r
												return LV
											}
										}
									} else { // 0x0c4d5 <= r
										if r < 0x0c4f1 {
											if r < 0x0c4f0 {
												return LVT
											} else { // 0x0c4f0 <= r
												return LV
											}
										} else { // 0x0c4f1 <= r
											if r < 0x0c50c {
												return LVT
											} else { // 0x0c50c <= r
												return LV
											}
										}
									}
								}
							}
						} else { // 0x0c50d <= r
							if r < 0x0c5d0 {
								if r < 0x0c561 {
									if r < 0x0c544 {
										if r < 0x0c528 {
											return LVT
										} else { // 0x0c528 <= r
											if r < 0x0c529 {
												return LV
											} else { // 0x0c529 <= r
												return LVT
											}
										}
									} else { // 0x0c544 <= r
										if r < 0x0c545 {
											return LV
										} else { // 0x0c545 <= r
											if r < 0x0c560 {
												return LVT
											} else { // 0x0c560 <= r
												return LV
											}
										}
									}
								} else { // 0x0c561 <= r
									if r < 0x0c598 {
										if r < 0x0c57c {
											return LVT
										} else { // 0x0c57c <= r
											if r < 0x0c57d {
												return LV
											} else { // 0x0c57d <= r
												return LVT
											}
										}
									} else { // 0x0c598 <= r
										if r < 0x0c5b4 {
											if r < 0x0c599 {
												return LV
											} else { // 0x0c599 <= r
												return LVT
											}
										} else { // 0x0c5b4 <= r
											if r < 0x0c5b5 {
												return LV
											} else { // 0x0c5b5 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0c5d0 <= r
								if r < 0x0c624 {
									if r < 0x0c5ed {
										if r < 0x0c5d1 {
											return LV
										} else { // 0x0c5d1 <= r
											if r < 0x0c5ec {
												return LVT
											} else { // 0x0c5ec <= r
												return LV
											}
										}
									} else { // 0x0c5ed <= r
										if r < 0x0c608 {
											return LVT
										} else { // 0x0c608 <= r
											if r < 0x0c609 {
												return LV
											} else { // 0x0c609 <= r
												return LVT
											}
										}
									}
								} else { // 0x0c624 <= r
									if r < 0x0c641 {
										if r < 0x0c625 {
											return LV
										} else { // 0x0c625 <= r
											if r < 0x0c640 {
												return LVT
											} else { // 0x0c640 <= r
												return LV
											}
										}
									} else { // 0x0c641 <= r
										if r < 0x0c65d {
											if r < 0x0c65c {
												return LVT
											} else { // 0x0c65c <= r
												return LV
											}
										} else { // 0x0c65d <= r
											if r < 0x0c678 {
												return LVT
											} else { // 0x0c678 <= r
												return LV
											}
										}
									}
								}
							}
						}
					}
				}
			} else { // 0x0c679 <= r
				if r < 0x0cc0c {
					if r < 0x0c935 {
						if r < 0x0c7e4 {
							if r < 0x0c721 {
								if r < 0x0c6cd {
									if r < 0x0c6b0 {
										if r < 0x0c694 {
											return LVT
										} else { // 0x0c694 <= r
											if r < 0x0c695 {
												return LV
											} else { // 0x0c695 <= r
												return LVT
											}
										}
									} else { // 0x0c6b0 <= r
										if r < 0x0c6b1 {
											return LV
										} else { // 0x0c6b1 <= r
											if r < 0x0c6cc {
												return LVT
											} else { // 0x0c6cc <= r
												return LV
											}
										}
									}
								} else { // 0x0c6cd <= r
									if r < 0x0c704 {
										if r < 0x0c6e8 {
											return LVT
										} else { // 0x0c6e8 <= r
											if r < 0x0c6e9 {
												return LV
											} else { // 0x0c6e9 <= r
												return LVT
											}
										}
									} else { // 0x0c704 <= r
										if r < 0x0c705 {
											return LV
										} else { // 0x0c705 <= r
											if r < 0x0c720 {
												return LVT
											} else { // 0x0c720 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0c721 <= r
								if r < 0x0c775 {
									if r < 0x0c758 {
										if r < 0x0c73c {
											return LVT
										} else { // 0x0c73c <= r
											if r < 0x0c73d {
												return LV
											} else { // 0x0c73d <= r
												return LVT
											}
										}
									} else { // 0x0c758 <= r
										if r < 0x0c759 {
											return LV
										} else { // 0x0c759 <= r
											if r < 0x0c774 {
												return LVT
											} else { // 0x0c774 <= r
												return LV
											}
										}
									}
								} else { // 0x0c775 <= r
									if r < 0x0c7ac {
										if r < 0x0c790 {
											return LVT
										} else { // 0x0c790 <= r
											if r < 0x0c791 {
												return LV
											} else { // 0x0c791 <= r
												return LVT
											}
										}
									} else { // 0x0c7ac <= r
										if r < 0x0c7c8 {
											if r < 0x0c7ad {
												return LV
											} else { // 0x0c7ad <= r
												return LVT
											}
										} else { // 0x0c7c8 <= r
											if r < 0x0c7c9 {
												return LV
											} else { // 0x0c7c9 <= r
												return LVT
											}
										}
									}
								}
							}
						} else { // 0x0c7e4 <= r
							if r < 0x0c88c {
								if r < 0x0c838 {
									if r < 0x0c801 {
										if r < 0x0c7e5 {
											return LV
										} else { // 0x0c7e5 <= r
											if r < 0x0c800 {
												return LVT
											} else { // 0x0c800 <= r
												return LV
											}
										}
									} else { // 0x0c801 <= r
										if r < 0x0c81c {
											return LVT
										} else { // 0x0c81c <= r
											if r < 0x0c81d {
												return LV
											} else { // 0x0c81d <= r
												return LVT
											}
										}
									}
								} else { // 0x0c838 <= r
									if r < 0x0c855 {
										if r < 0x0c839 {
											return LV
										} else { // 0x0c839 <= r
											if r < 0x0c854 {
												return LVT
											} else { // 0x0c854 <= r
												return LV
											}
										}
									} else { // 0x0c855 <= r
										if r < 0x0c870 {
											return LVT
										} else { // 0x0c870 <= r
											if r < 0x0c871 {
												return LV
											} else { // 0x0c871 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0c88c <= r
								if r < 0x0c8e0 {
									if r < 0x0c8a9 {
										if r < 0x0c88d {
											return LV
										} else { // 0x0c88d <= r
											if r < 0x0c8a8 {
												return LVT
											} else { // 0x0c8a8 <= r
												return LV
											}
										}
									} else { // 0x0c8a9 <= r
										if r < 0x0c8c4 {
											return LVT
										} else { // 0x0c8c4 <= r
											if r < 0x0c8c5 {
												return LV
											} else { // 0x0c8c5 <= r
												return LVT
											}
										}
									}
								} else { // 0x0c8e0 <= r
									if r < 0x0c8fd {
										if r < 0x0c8e1 {
											return LV
										} else { // 0x0c8e1 <= r
											if r < 0x0c8fc {
												return LVT
											} else { // 0x0c8fc <= r
												return LV
											}
										}
									} else { // 0x0c8fd <= r
										if r < 0x0c919 {
											if r < 0x0c918 {
												return LVT
											} else { // 0x0c918 <= r
												return LV
											}
										} else { // 0x0c919 <= r
											if r < 0x0c934 {
												return LVT
											} else { // 0x0c934 <= r
												return LV
											}
										}
									}
								}
							}
						}
					} else { // 0x0c935 <= r
						if r < 0x0caa0 {
							if r < 0x0c9dd {
								if r < 0x0c989 {
									if r < 0x0c96c {
										if r < 0x0c950 {
											return LVT
										} else { // 0x0c950 <= r
											if r < 0x0c951 {
												return LV
											} else { // 0x0c951 <= r
												return LVT
											}
										}
									} else { // 0x0c96c <= r
										if r < 0x0c96d {
											return LV
										} else { // 0x0c96d <= r
											if r < 0x0c988 {
												return LVT
											} else { // 0x0c988 <= r
												return LV
											}
										}
									}
								} else { // 0x0c989 <= r
									if r < 0x0c9c0 {
										if r < 0x0c9a4 {
											return LVT
										} else { // 0x0c9a4 <= r
											if r < 0x0c9a5 {
												return LV
											} else { // 0x0c9a5 <= r
												return LVT
											}
										}
									} else { // 0x0c9c0 <= r
										if r < 0x0c9c1 {
											return LV
										} else { // 0x0c9c1 <= r
											if r < 0x0c9dc {
												return LVT
											} else { // 0x0c9dc <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0c9dd <= r
								if r < 0x0ca31 {
									if r < 0x0ca14 {
										if r < 0x0c9f8 {
											return LVT
										} else { // 0x0c9f8 <= r
											if r < 0x0c9f9 {
												return LV
											} else { // 0x0c9f9 <= r
												return LVT
											}
										}
									} else { // 0x0ca14 <= r
										if r < 0x0ca15 {
											return LV
										} else { // 0x0ca15 <= r
											if r < 0x0ca30 {
												return LVT
											} else { // 0x0ca30 <= r
												return LV
											}
										}
									}
								} else { // 0x0ca31 <= r
									if r < 0x0ca68 {
										if r < 0x0ca4c {
											return LVT
										} else { // 0x0ca4c <= r
											if r < 0x0ca4d {
												return LV
											} else { // 0x0ca4d <= r
												return LVT
											}
										}
									} else { // 0x0ca68 <= r
										if r < 0x0ca84 {
											if r < 0x0ca69 {
												return LV
											} else { // 0x0ca69 <= r
												return LVT
											}
										} else { // 0x0ca84 <= r
											if r < 0x0ca85 {
												return LV
											} else { // 0x0ca85 <= r
												return LVT
											}
										}
									}
								}
							}
						} else { // 0x0caa0 <= r
							if r < 0x0cb49 {
								if r < 0x0caf4 {
									if r < 0x0cabd {
										if r < 0x0caa1 {
											return LV
										} else { // 0x0caa1 <= r
											if r < 0x0cabc {
												return LVT
											} else { // 0x0cabc <= r
												return LV
											}
										}
									} else { // 0x0cabd <= r
										if r < 0x0cad8 {
											return LVT
										} else { // 0x0cad8 <= r
											if r < 0x0cad9 {
												return LV
											} else { // 0x0cad9 <= r
												return LVT
											}
										}
									}
								} else { // 0x0caf4 <= r
									if r < 0x0cb11 {
										if r < 0x0caf5 {
											return LV
										} else { // 0x0caf5 <= r
											if r < 0x0cb10 {
												return LVT
											} else { // 0x0cb10 <= r
												return LV
											}
										}
									} else { // 0x0cb11 <= r
										if r < 0x0cb2d {
											if r < 0x0cb2c {
												return LVT
											} else { // 0x0cb2c <= r
												return LV
											}
										} else { // 0x0cb2d <= r
											if r < 0x0cb48 {
												return LVT
											} else { // 0x0cb48 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0cb49 <= r
								if r < 0x0cb9d {
									if r < 0x0cb80 {
										if r < 0x0cb64 {
											return LVT
										} else { // 0x0cb64 <= r
											if r < 0x0cb65 {
												return LV
											} else { // 0x0cb65 <= r
												return LVT
											}
										}
									} else { // 0x0cb80 <= r
										if r < 0x0cb81 {
											return LV
										} else { // 0x0cb81 <= r
											if r < 0x0cb9c {
												return LVT
											} else { // 0x0cb9c <= r
												return LV
											}
										}
									}
								} else { // 0x0cb9d <= r
									if r < 0x0cbd4 {
										if r < 0x0cbb8 {
											return LVT
										} else { // 0x0cbb8 <= r
											if r < 0x0cbb9 {
												return LV
											} else { // 0x0cbb9 <= r
												return LVT
											}
										}
									} else { // 0x0cbd4 <= r
										if r < 0x0cbf0 {
											if r < 0x0cbd5 {
												return LV
											} else { // 0x0cbd5 <= r
												return LVT
											}
										} else { // 0x0cbf0 <= r
											if r < 0x0cbf1 {
												return LV
											} else { // 0x0cbf1 <= r
												return LVT
											}
										}
									}
								}
							}
						}
					}
				} else { // 0x0cc0c <= r
					if r < 0x0cec8 {
						if r < 0x0cd5d {
							if r < 0x0ccb4 {
								if r < 0x0cc60 {
									if r < 0x0cc29 {
										if r < 0x0cc0d {
											return LV
										} else { // 0x0cc0d <= r
											if r < 0x0cc28 {
												return LVT
											} else { // 0x0cc28 <= r
												return LV
											}
										}
									} else { // 0x0cc29 <= r
										if r < 0x0cc44 {
											return LVT
										} else { // 0x0cc44 <= r
											if r < 0x0cc45 {
												return LV
											} else { // 0x0cc45 <= r
												return LVT
											}
										}
									}
								} else { // 0x0cc60 <= r
									if r < 0x0cc7d {
										if r < 0x0cc61 {
											return LV
										} else { // 0x0cc61 <= r
											if r < 0x0cc7c {
												return LVT
											} else { // 0x0cc7c <= r
												return LV
											}
										}
									} else { // 0x0cc7d <= r
										if r < 0x0cc98 {
											return LVT
										} else { // 0x0cc98 <= r
											if r < 0x0cc99 {
												return LV
											} else { // 0x0cc99 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0ccb4 <= r
								if r < 0x0cd08 {
									if r < 0x0ccd1 {
										if r < 0x0ccb5 {
											return LV
										} else { // 0x0ccb5 <= r
											if r < 0x0ccd0 {
												return LVT
											} else { // 0x0ccd0 <= r
												return LV
											}
										}
									} else { // 0x0ccd1 <= r
										if r < 0x0ccec {
											return LVT
										} else { // 0x0ccec <= r
											if r < 0x0cced {
												return LV
											} else { // 0x0cced <= r
												return LVT
											}
										}
									}
								} else { // 0x0cd08 <= r
									if r < 0x0cd25 {
										if r < 0x0cd09 {
											return LV
										} else { // 0x0cd09 <= r
											if r < 0x0cd24 {
												return LVT
											} else { // 0x0cd24 <= r
												return LV
											}
										}
									} else { // 0x0cd25 <= r
										if r < 0x0cd41 {
											if r < 0x0cd40 {
												return LVT
											} else { // 0x0cd40 <= r
												return LV
											}
										} else { // 0x0cd41 <= r
											if r < 0x0cd5c {
												return LVT
											} else { // 0x0cd5c <= r
												return LV
											}
										}
									}
								}
							}
						} else { // 0x0cd5d <= r
							if r < 0x0ce05 {
								if r < 0x0cdb1 {
									if r < 0x0cd94 {
										if r < 0x0cd78 {
											return LVT
										} else { // 0x0cd78 <= r
											if r < 0x0cd79 {
												return LV
											} else { // 0x0cd79 <= r
												return LVT
											}
										}
									} else { // 0x0cd94 <= r
										if r < 0x0cd95 {
											return LV
										} else { // 0x0cd95 <= r
											if r < 0x0cdb0 {
												return LVT
											} else { // 0x0cdb0 <= r
												return LV
											}
										}
									}
								} else { // 0x0cdb1 <= r
									if r < 0x0cde8 {
										if r < 0x0cdcc {
											return LVT
										} else { // 0x0cdcc <= r
											if r < 0x0cdcd {
												return LV
											} else { // 0x0cdcd <= r
												return LVT
											}
										}
									} else { // 0x0cde8 <= r
										if r < 0x0cde9 {
											return LV
										} else { // 0x0cde9 <= r
											if r < 0x0ce04 {
												return LVT
											} else { // 0x0ce04 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0ce05 <= r
								if r < 0x0ce59 {
									if r < 0x0ce3c {
										if r < 0x0ce20 {
											return LVT
										} else { // 0x0ce20 <= r
											if r < 0x0ce21 {
												return LV
											} else { // 0x0ce21 <= r
												return LVT
											}
										}
									} else { // 0x0ce3c <= r
										if r < 0x0ce3d {
											return LV
										} else { // 0x0ce3d <= r
											if r < 0x0ce58 {
												return LVT
											} else { // 0x0ce58 <= r
												return LV
											}
										}
									}
								} else { // 0x0ce59 <= r
									if r < 0x0ce90 {
										if r < 0x0ce74 {
											return LVT
										} else { // 0x0ce74 <= r
											if r < 0x0ce75 {
												return LV
											} else { // 0x0ce75 <= r
												return LVT
											}
										}
									} else { // 0x0ce90 <= r
										if r < 0x0ceac {
											if r < 0x0ce91 {
												return LV
											} else { // 0x0ce91 <= r
												return LVT
											}
										} else { // 0x0ceac <= r
											if r < 0x0cead {
												return LV
											} else { // 0x0cead <= r
												return LVT
											}
										}
									}
								}
							}
						}
					} else { // 0x0cec8 <= r
						if r < 0x0d019 {
							if r < 0x0cf70 {
								if r < 0x0cf1c {
									if r < 0x0cee5 {
										if r < 0x0cec9 {
											return LV
										} else { // 0x0cec9 <= r
											if r < 0x0cee4 {
												return LVT
											} else { // 0x0cee4 <= r
												return LV
											}
										}
									} else { // 0x0cee5 <= r
										if r < 0x0cf00 {
											return LVT
										} else { // 0x0cf00 <= r
											if r < 0x0cf01 {
												return LV
											} else { // 0x0cf01 <= r
												return LVT
											}
										}
									}
								} else { // 0x0cf1c <= r
									if r < 0x0cf39 {
										if r < 0x0cf1d {
											return LV
										} else { // 0x0cf1d <= r
											if r < 0x0cf38 {
												return LVT
											} else { // 0x0cf38 <= r
												return LV
											}
										}
									} else { // 0x0cf39 <= r
										if r < 0x0cf54 {
											return LVT
										} else { // 0x0cf54 <= r
											if r < 0x0cf55 {
												return LV
											} else { // 0x0cf55 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0cf70 <= r
								if r < 0x0cfc4 {
									if r < 0x0cf8d {
										if r < 0x0cf71 {
											return LV
										} else { // 0x0cf71 <= r
											if r < 0x0cf8c {
												return LVT
											} else { // 0x0cf8c <= r
												return LV
											}
										}
									} else { // 0x0cf8d <= r
										if r < 0x0cfa8 {
											return LVT
										} else { // 0x0cfa8 <= r
											if r < 0x0cfa9 {
												return LV
											} else { // 0x0cfa9 <= r
												return LVT
											}
										}
									}
								} else { // 0x0cfc4 <= r
									if r < 0x0cfe1 {
										if r < 0x0cfc5 {
											return LV
										} else { // 0x0cfc5 <= r
											if r < 0x0cfe0 {
												return LVT
											} else { // 0x0cfe0 <= r
												return LV
											}
										}
									} else { // 0x0cfe1 <= r
										if r < 0x0cffd {
											if r < 0x0cffc {
												return LVT
											} else { // 0x0cffc <= r
												return LV
											}
										} else { // 0x0cffd <= r
											if r < 0x0d018 {
												return LVT
											} else { // 0x0d018 <= r
												return LV
											}
										}
									}
								}
							}
						} else { // 0x0d019 <= r
							if r < 0x0d0dc {
								if r < 0x0d06d {
									if r < 0x0d050 {
										if r < 0x0d034 {
											return LVT
										} else { // 0x0d034 <= r
											if r < 0x0d035 {
												return LV
											} else { // 0x0d035 <= r
												return LVT
											}
										}
									} else { // 0x0d050 <= r
										if r < 0x0d051 {
											return LV
										} else { // 0x0d051 <= r
											if r < 0x0d06c {
												return LVT
											} else { // 0x0d06c <= r
												return LV
											}
										}
									}
								} else { // 0x0d06d <= r
									if r < 0x0d0a4 {
										if r < 0x0d088 {
											return LVT
										} else { // 0x0d088 <= r
											if r < 0x0d089 {
												return LV
											} else { // 0x0d089 <= r
												return LVT
											}
										}
									} else { // 0x0d0a4 <= r
										if r < 0x0d0c0 {
											if r < 0x0d0a5 {
												return LV
											} else { // 0x0d0a5 <= r
												return LVT
											}
										} else { // 0x0d0c0 <= r
											if r < 0x0d0c1 {
												return LV
											} else { // 0x0d0c1 <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0d0dc <= r
								if r < 0x0d130 {
									if r < 0x0d0f9 {
										if r < 0x0d0dd {
											return LV
										} else { // 0x0d0dd <= r
											if r < 0x0d0f8 {
												return LVT
											} else { // 0x0d0f8 <= r
												return LV
											}
										}
									} else { // 0x0d0f9 <= r
										if r < 0x0d114 {
											return LVT
										} else { // 0x0d114 <= r
											if r < 0x0d115 {
												return LV
											} else { // 0x0d115 <= r
												return LVT
											}
										}
									}
								} else { // 0x0d130 <= r
									if r < 0x0d14d {
										if r < 0x0d131 {
											return LV
										} else { // 0x0d131 <= r
											if r < 0x0d14c {
												return LVT
											} else { // 0x0d14c <= r
												return LV
											}
										}
									} else { // 0x0d14d <= r
										if r < 0x0d169 {
											if r < 0x0d168 {
												return LVT
											} else { // 0x0d168 <= r
												return LV
											}
										} else { // 0x0d169 <= r
											if r < 0x0d184 {
												return LVT
											} else { // 0x0d184 <= r
												return LV
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else { // 0x0d185 <= r
			if r < 0x1133f {
				if r < 0x0d718 {
					if r < 0x0d441 {
						if r < 0x0d2f0 {
							if r < 0x0d22d {
								if r < 0x0d1d9 {
									if r < 0x0d1bc {
										if r < 0x0d1a0 {
											return LVT
										} else { // 0x0d1a0 <= r
											if r < 0x0d1a1 {
												return LV
											} else { // 0x0d1a1 <= r
												return LVT
											}
										}
									} else { // 0x0d1bc <= r
										if r < 0x0d1bd {
											return LV
										} else { // 0x0d1bd <= r
											if r < 0x0d1d8 {
												return LVT
											} else { // 0x0d1d8 <= r
												return LV
											}
										}
									}
								} else { // 0x0d1d9 <= r
									if r < 0x0d210 {
										if r < 0x0d1f4 {
											return LVT
										} else { // 0x0d1f4 <= r
											if r < 0x0d1f5 {
												return LV
											} else { // 0x0d1f5 <= r
												return LVT
											}
										}
									} else { // 0x0d210 <= r
										if r < 0x0d211 {
											return LV
										} else { // 0x0d211 <= r
											if r < 0x0d22c {
												return LVT
											} else { // 0x0d22c <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0d22d <= r
								if r < 0x0d281 {
									if r < 0x0d264 {
										if r < 0x0d248 {
											return LVT
										} else { // 0x0d248 <= r
											if r < 0x0d249 {
												return LV
											} else { // 0x0d249 <= r
												return LVT
											}
										}
									} else { // 0x0d264 <= r
										if r < 0x0d265 {
											return LV
										} else { // 0x0d265 <= r
											if r < 0x0d280 {
												return LVT
											} else { // 0x0d280 <= r
												return LV
											}
										}
									}
								} else { // 0x0d281 <= r
									if r < 0x0d2b8 {
										if r < 0x0d29c {
											return LVT
										} else { // 0x0d29c <= r
											if r < 0x0d29d {
												return LV
											} else { // 0x0d29d <= r
												return LVT
											}
										}
									} else { // 0x0d2b8 <= r
										if r < 0x0d2d4 {
											if r < 0x0d2b9 {
												return LV
											} else { // 0x0d2b9 <= r
												return LVT
											}
										} else { // 0x0d2d4 <= r
											if r < 0x0d2d5 {
												return LV
											} else { // 0x0d2d5 <= r
												return LVT
											}
										}
									}
								}
							}
						} else { // 0x0d2f0 <= r
							if r < 0x0d398 {
								if r < 0x0d344 {
									if r < 0x0d30d {
										if r < 0x0d2f1 {
											return LV
										} else { // 0x0d2f1 <= r
											if r < 0x0d30c {
												return LVT
											} else { // 0x0d30c <= r
												return LV
											}
										}
									} else { // 0x0d30d <= r
										if r < 0x0d328 {
											return LVT
										} else { // 0x0d328 <= r
											if r < 0x0d329 {
												return LV
											} else { // 0x0d329 <= r
												return LVT
											}
										}
									}
								} else { // 0x0d344 <= r
									if r < 0x0d361 {
										if r < 0x0d345 {
											return LV
										} else { // 0x0d345 <= r
											if r < 0x0d360 {
												return LVT
											} else { // 0x0d360 <= r
												return LV
											}
										}
									} else { // 0x0d361 <= r
										if r < 0x0d37c {
											return LVT
										} else { // 0x0d37c <= r
											if r < 0x0d37d {
												return LV
											} else { // 0x0d37d <= r
												return LVT
											}
										}
									}
								}
							} else { // 0x0d398 <= r
								if r < 0x0d3ec {
									if r < 0x0d3b5 {
										if r < 0x0d399 {
											return LV
										} else { // 0x0d399 <= r
											if r < 0x0d3b4 {
												return LVT
											} else { // 0x0d3b4 <= r
												return LV
											}
										}
									} else { // 0x0d3b5 <= r
										if r < 0x0d3d0 {
											return LVT
										} else { // 0x0d3d0 <= r
											if r < 0x0d3d1 {
												return LV
											} else { // 0x0d3d1 <= r
												return LVT
											}
										}
									}
								} else { // 0x0d3ec <= r
									if r < 0x0d409 {
										if r < 0x0d3ed {
											return LV
										} else { // 0x0d3ed <= r
											if r < 0x0d408 {
												return LVT
											} else { // 0x0d408 <= r
												return LV
											}
										}
									} else { // 0x0d409 <= r
										if r < 0x0d425 {
											if r < 0x0d424 {
												return LVT
											} else { // 0x0d424 <= r
												return LV
											}
										} else { // 0x0d425 <= r
											if r < 0x0d440 {
												return LVT
											} else { // 0x0d440 <= r
												return LV
											}
										}
									}
								}
							}
						}
					} else { // 0x0d441 <= r
						if r < 0x0d5ac {
							if r < 0x0d4e9 {
								if r < 0x0d495 {
									if r < 0x0d478 {
										if r < 0x0d45c {
											return LVT
										} else { // 0x0d45c <= r
											if r < 0x0d45d {
												return LV
											} else { // 0x0d45d <= r
												return LVT
											}
										}
									} else { // 0x0d478 <= r
										if r < 0x0d479 {
											return LV
										} else { // 0x0d479 <= r
											if r < 0x0d494 {
												return LVT
											} else { // 0x0d494 <= r
												return LV
											}
										}
									}
								} else { // 0x0d495 <= r
									if r < 0x0d4cc {
										if r < 0x0d4b0 {
											return LVT
										} else { // 0x0d4b0 <= r
											if r < 0x0d4b1 {
												return LV
											} else { // 0x0d4b1 <= r
												return LVT
											}
										}
									} else { // 0x0d4cc <= r
										if r < 0x0d4cd {
											return LV
										} else { // 0x0d4cd <= r
											if r < 0x0d4e8 {
												return LVT
											} else { // 0x0d4e8 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0d4e9 <= r
								if r < 0x0d53d {
									if r < 0x0d520 {
										if r < 0x0d504 {
											return LVT
										} else { // 0x0d504 <= r
											if r < 0x0d505 {
												return LV
											} else { // 0x0d505 <= r
												return LVT
											}
										}
									} else { // 0x0d520 <= r
										if r < 0x0d521 {
											return LV
										} else { // 0x0d521 <= r
											if r < 0x0d53c {
												return LVT
											} else { // 0x0d53c <= r
												return LV
											}
										}
									}
								} else { // 0x0d53d <= r
									if r < 0x0d574 {
										if r < 0x0d558 {
											return LVT
										} else { // 0x0d558 <= r
											if r < 0x0d559 {
												return LV
											} else { // 0x0d559 <= r
												return LVT
											}
										}
									} else { // 0x0d574 <= r
										if r < 0x0d590 {
											if r < 0x0d575 {
												return LV
											} else { // 0x0d575 <= r
												return LVT
											}
										} else { // 0x0d590 <= r
											if r < 0x0d591 {
												return LV
											} else { // 0x0d591 <= r
												return LVT
											}
										}
									}
								}
							}
						} else { // 0x0d5ac <= r
							if r < 0x0d655 {
								if r < 0x0d600 {
									if r < 0x0d5c9 {
										if r < 0x0d5ad {
											return LV
										} else { // 0x0d5ad <= r
											if r < 0x0d5c8 {
												return LVT
											} else { // 0x0d5c8 <= r
												return LV
											}
										}
									} else { // 0x0d5c9 <= r
										if r < 0x0d5e4 {
											return LVT
										} else { // 0x0d5e4 <= r
											if r < 0x0d5e5 {
												return LV
											} else { // 0x0d5e5 <= r
												return LVT
											}
										}
									}
								} else { // 0x0d600 <= r
									if r < 0x0d61d {
										if r < 0x0d601 {
											return LV
										} else { // 0x0d601 <= r
											if r < 0x0d61c {
												return LVT
											} else { // 0x0d61c <= r
												return LV
											}
										}
									} else { // 0x0d61d <= r
										if r < 0x0d639 {
											if r < 0x0d638 {
												return LVT
											} else { // 0x0d638 <= r
												return LV
											}
										} else { // 0x0d639 <= r
											if r < 0x0d654 {
												return LVT
											} else { // 0x0d654 <= r
												return LV
											}
										}
									}
								}
							} else { // 0x0d655 <= r
								if r < 0x0d6a9 {
									if r < 0x0d68c {
										if r < 0x0d670 {
											return LVT
										} else { // 0x0d670 <= r
											if r < 0x0d671 {
												return LV
											} else { // 0x0d671 <= r
												return LVT
											}
										}
									} else { // 0x0d68c <= r
										if r < 0x0d68d {
											return LV
										} else { // 0x0d68d <= r
											if r < 0x0d6a8 {
												return LVT
											} else { // 0x0d6a8 <= r
												return LV
											}
										}
									}
								} else { // 0x0d6a9 <= r
									if r < 0x0d6e0 {
										if r < 0x0d6c4 {
											return LVT
										} else { // 0x0d6c4 <= r
											if r < 0x0d6c5 {
												return LV
											} else { // 0x0d6c5 <= r
												return LVT
											}
										}
									} else { // 0x0d6e0 <= r
										if r < 0x0d6fc {
											if r < 0x0d6e1 {
												return LV
											} else { // 0x0d6e1 <= r
												return LVT
											}
										} else { // 0x0d6fc <= r
											if r < 0x0d6fd {
												return LV
											} else { // 0x0d6fd <= r
												return LVT
											}
										}
									}
								}
							}
						}
					}
				} else { // 0x0d718 <= r
					if r < 0x11003 {
						if r < 0x0ff9e {
							if r < 0x0d7c7 {
								if r < 0x0d76c {
									if r < 0x0d735 {
										if r < 0x0d719 {
											return LV
										} else { // 0x0d719 <= r
											if r < 0x0d734 {
												return LVT
											} else { // 0x0d734 <= r
												return LV
											}
										}
									} else { // 0x0d735 <= r
										if r < 0x0d750 {
											return LVT
										} else { // 0x0d750 <= r
											if r < 0x0d751 {
												return LV
											} else { // 0x0d751 <= r
												return LVT
											}
										}
									}
								} else { // 0x0d76c <= r
									if r < 0x0d789 {
										if r < 0x0d76d {
											return LV
										} else { // 0x0d76d <= r
											if r < 0x0d788 {
												return LVT
											} else { // 0x0d788 <= r
												return LV
											}
										}
									} else { // 0x0d789 <= r
										if r < 0x0d7a4 {
											return LVT
										} else { // 0x0d7a4 <= r
											if r < 0x0d7b0 {
												return Any
											} else { // 0x0d7b0 <= r
												return V
											}
										}
									}
								}
							} else { // 0x0d7c7 <= r
								if r < 0x0fb1f {
									if r < 0x0d800 {
										if r < 0x0d7cb {
											return Any
										} else { // 0x0d7cb <= r
											if r < 0x0d7fc {
												return T
											} else { // 0x0d7fc <= r
												return Any
											}
										}
									} else { // 0x0d800 <= r
										if r < 0x0e000 {
											return Control
										} else { // 0x0e000 <= r
											if r < 0x0fb1e {
												return Any
											} else { // 0x0fb1e <= r
												return Extend
											}
										}
									}
								} else { // 0x0fb1f <= r
									if r < 0x0fe20 {
										if r < 0x0fe00 {
											return Any
										} else { // 0x0fe00 <= r
											if r < 0x0fe10 {
												return Extend
											} else { // 0x0fe10 <= r
												return Any
											}
										}
									} else { // 0x0fe20 <= r
										if r < 0x0feff {
											if r < 0x0fe30 {
												return Extend
											} else { // 0x0fe30 <= r
												return Any
											}
										} else { // 0x0feff <= r
											if r < 0x0ff00 {
												return Control
											} else { // 0x0ff00 <= r
												return Any
											}
										}
									}
								}
							}
						} else { // 0x0ff9e <= r
							if r < 0x10a05 {
								if r < 0x102e0 {
									if r < 0x0fffc {
										if r < 0x0ffa0 {
											return Extend
										} else { // 0x0ffa0 <= r
											if r < 0x0fff0 {
												return Any
											} else { // 0x0fff0 <= r
												return Control
											}
										}
									} else { // 0x0fffc <= r
										if r < 0x101fd {
											return Any
										} else { // 0x101fd <= r
											if r < 0x101fe {
												return Extend
											} else { // 0x101fe <= r
												return Any
											}
										}
									}
								} else { // 0x102e0 <= r
									if r < 0x1037b {
										if r < 0x102e1 {
											return Extend
										} else { // 0x102e1 <= r
											if r < 0x10376 {
												return Any
											} else { // 0x10376 <= r
												return Extend
											}
										}
									} else { // 0x1037b <= r
										if r < 0x10a01 {
											return Any
										} else { // 0x10a01 <= r
											if r < 0x10a04 {
												return Extend
											} else { // 0x10a04 <= r
												return Any
											}
										}
									}
								}
							} else { // 0x10a05 <= r
								if r < 0x10a3f {
									if r < 0x10a10 {
										if r < 0x10a07 {
											return Extend
										} else { // 0x10a07 <= r
											if r < 0x10a0c {
												return Any
											} else { // 0x10a0c <= r
												return Extend
											}
										}
									} else { // 0x10a10 <= r
										if r < 0x10a38 {
											return Any
										} else { // 0x10a38 <= r
											if r < 0x10a3b {
												return Extend
											} else { // 0x10a3b <= r
												return Any
											}
										}
									}
								} else { // 0x10a3f <= r
									if r < 0x10ae7 {
										if r < 0x10a40 {
											return Extend
										} else { // 0x10a40 <= r
											if r < 0x10ae5 {
												return Any
											} else { // 0x10ae5 <= r
												return Extend
											}
										}
									} else { // 0x10ae7 <= r
										if r < 0x11001 {
											if r < 0x11000 {
												return Any
											} else { // 0x11000 <= r
												return SpacingMark
											}
										} else { // 0x11001 <= r
											if r < 0x11002 {
												return Extend
											} else { // 0x11002 <= r
												return SpacingMark
											}
										}
									}
								}
							}
						}
					} else { // 0x11003 <= r
						if r < 0x111b6 {
							if r < 0x110be {
								if r < 0x110b0 {
									if r < 0x1107f {
										if r < 0x11038 {
											return Any
										} else { // 0x11038 <= r
											if r < 0x11047 {
												return Extend
											} else { // 0x11047 <= r
												return Any
											}
										}
									} else { // 0x1107f <= r
										if r < 0x11082 {
											return Extend
										} else { // 0x11082 <= r
											if r < 0x11083 {
												return SpacingMark
											} else { // 0x11083 <= r
												return Any
											}
										}
									}
								} else { // 0x110b0 <= r
									if r < 0x110b9 {
										if r < 0x110b3 {
											return SpacingMark
										} else { // 0x110b3 <= r
											if r < 0x110b7 {
												return Extend
											} else { // 0x110b7 <= r
												return SpacingMark
											}
										}
									} else { // 0x110b9 <= r
										if r < 0x110bb {
											return Extend
										} else { // 0x110bb <= r
											if r < 0x110bd {
												return Any
											} else { // 0x110bd <= r
												return Prepend
											}
										}
									}
								}
							} else { // 0x110be <= r
								if r < 0x11135 {
									if r < 0x11127 {
										if r < 0x11100 {
											return Any
										} else { // 0x11100 <= r
											if r < 0x11103 {
												return Extend
											} else { // 0x11103 <= r
												return Any
											}
										}
									} else { // 0x11127 <= r
										if r < 0x1112c {
											return Extend
										} else { // 0x1112c <= r
											if r < 0x1112d {
												return SpacingMark
											} else { // 0x1112d <= r
												return Extend
											}
										}
									}
								} else { // 0x11135 <= r
									if r < 0x11180 {
										if r < 0x11173 {
											return Any
										} else { // 0x11173 <= r
											if r < 0x11174 {
												return Extend
											} else { // 0x11174 <= r
												return Any
											}
										}
									} else { // 0x11180 <= r
										if r < 0x11183 {
											if r < 0x11182 {
												return Extend
											} else { // 0x11182 <= r
												return SpacingMark
											}
										} else { // 0x11183 <= r
											if r < 0x111b3 {
												return Any
											} else { // 0x111b3 <= r
												return SpacingMark
											}
										}
									}
								}
							}
						} else { // 0x111b6 <= r
							if r < 0x11238 {
								if r < 0x111cd {
									if r < 0x111c2 {
										if r < 0x111bf {
											return Extend
										} else { // 0x111bf <= r
											if r < 0x111c1 {
												return SpacingMark
											} else { // 0x111c1 <= r
												return Any
											}
										}
									} else { // 0x111c2 <= r
										if r < 0x111c4 {
											return Prepend
										} else { // 0x111c4 <= r
											if r < 0x111ca {
												return Any
											} else { // 0x111ca <= r
												return Extend
											}
										}
									}
								} else { // 0x111cd <= r
									if r < 0x11232 {
										if r < 0x1122c {
											return Any
										} else { // 0x1122c <= r
											if r < 0x1122f {
												return SpacingMark
											} else { // 0x1122f <= r
												return Extend
											}
										}
									} else { // 0x11232 <= r
										if r < 0x11235 {
											if r < 0x11234 {
												return SpacingMark
											} else { // 0x11234 <= r
												return Extend
											}
										} else { // 0x11235 <= r
											if r < 0x11236 {
												return SpacingMark
											} else { // 0x11236 <= r
												return Extend
											}
										}
									}
								}
							} else { // 0x11238 <= r
								if r < 0x112eb {
									if r < 0x112df {
										if r < 0x1123e {
											return Any
										} else { // 0x1123e <= r
											if r < 0x1123f {
												return Extend
											} else { // 0x1123f <= r
												return Any
											}
										}
									} else { // 0x112df <= r
										if r < 0x112e0 {
											return Extend
										} else { // 0x112e0 <= r
											if r < 0x112e3 {
												return SpacingMark
											} else { // 0x112e3 <= r
												return Extend
											}
										}
									}
								} else { // 0x112eb <= r
									if r < 0x11304 {
										if r < 0x11300 {
											return Any
										} else { // 0x11300 <= r
											if r < 0x11302 {
												return Extend
											} else { // 0x11302 <= r
												return SpacingMark
											}
										}
									} else { // 0x11304 <= r
										if r < 0x1133d {
											if r < 0x1133c {
												return Any
											} else { // 0x1133c <= r
												return Extend
											}
										} else { // 0x1133d <= r
											if r < 0x1133e {
												return Any
											} else { // 0x1133e <= r
												return Extend
											}
										}
									}
								}
							}
						}
					}
				}
			} else { // 0x1133f <= r
				if r < 0x1d173 {
					if r < 0x1163e {
						if r < 0x114b3 {
							if r < 0x11366 {
								if r < 0x1134b {
									if r < 0x11345 {
										if r < 0x11340 {
											return SpacingMark
										} else { // 0x11340 <= r
											if r < 0x11341 {
												return Extend
											} else { // 0x11341 <= r
												return SpacingMark
											}
										}
									} else { // 0x11345 <= r
										if r < 0x11347 {
											return Any
										} else { // 0x11347 <= r
											if r < 0x11349 {
												return SpacingMark
											} else { // 0x11349 <= r
												return Any
											}
										}
									}
								} else { // 0x1134b <= r
									if r < 0x11358 {
										if r < 0x1134e {
											return SpacingMark
										} else { // 0x1134e <= r
											if r < 0x11357 {
												return Any
											} else { // 0x11357 <= r
												return Extend
											}
										}
									} else { // 0x11358 <= r
										if r < 0x11362 {
											return Any
										} else { // 0x11362 <= r
											if r < 0x11364 {
												return SpacingMark
											} else { // 0x11364 <= r
												return Any
											}
										}
									}
								}
							} else { // 0x11366 <= r
								if r < 0x11440 {
									if r < 0x11375 {
										if r < 0x1136d {
											return Extend
										} else { // 0x1136d <= r
											if r < 0x11370 {
												return Any
											} else { // 0x11370 <= r
												return Extend
											}
										}
									} else { // 0x11375 <= r
										if r < 0x11435 {
											return Any
										} else { // 0x11435 <= r
											if r < 0x11438 {
												return SpacingMark
											} else { // 0x11438 <= r
												return Extend
											}
										}
									}
								} else { // 0x11440 <= r
									if r < 0x11446 {
										if r < 0x11442 {
											return SpacingMark
										} else { // 0x11442 <= r
											if r < 0x11445 {
												return Extend
											} else { // 0x11445 <= r
												return SpacingMark
											}
										}
									} else { // 0x11446 <= r
										if r < 0x114b0 {
											if r < 0x11447 {
												return Extend
											} else { // 0x11447 <= r
												return Any
											}
										} else { // 0x114b0 <= r
											if r < 0x114b1 {
												return Extend
											} else { // 0x114b1 <= r
												return SpacingMark
											}
										}
									}
								}
							}
						} else { // 0x114b3 <= r
							if r < 0x115b2 {
								if r < 0x114bf {
									if r < 0x114bb {
										if r < 0x114b9 {
											return Extend
										} else { // 0x114b9 <= r
											if r < 0x114ba {
												return SpacingMark
											} else { // 0x114ba <= r
												return Extend
											}
										}
									} else { // 0x114bb <= r
										if r < 0x114bd {
											return SpacingMark
										} else { // 0x114bd <= r
											if r < 0x114be {
												return Extend
											} else { // 0x114be <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x114bf <= r
									if r < 0x114c4 {
										if r < 0x114c1 {
											return Extend
										} else { // 0x114c1 <= r
											if r < 0x114c2 {
												return SpacingMark
											} else { // 0x114c2 <= r
												return Extend
											}
										}
									} else { // 0x114c4 <= r
										if r < 0x115af {
											return Any
										} else { // 0x115af <= r
											if r < 0x115b0 {
												return Extend
											} else { // 0x115b0 <= r
												return SpacingMark
											}
										}
									}
								}
							} else { // 0x115b2 <= r
								if r < 0x115c1 {
									if r < 0x115bc {
										if r < 0x115b6 {
											return Extend
										} else { // 0x115b6 <= r
											if r < 0x115b8 {
												return Any
											} else { // 0x115b8 <= r
												return SpacingMark
											}
										}
									} else { // 0x115bc <= r
										if r < 0x115be {
											return Extend
										} else { // 0x115be <= r
											if r < 0x115bf {
												return SpacingMark
											} else { // 0x115bf <= r
												return Extend
											}
										}
									}
								} else { // 0x115c1 <= r
									if r < 0x11630 {
										if r < 0x115dc {
											return Any
										} else { // 0x115dc <= r
											if r < 0x115de {
												return Extend
											} else { // 0x115de <= r
												return Any
											}
										}
									} else { // 0x11630 <= r
										if r < 0x1163b {
											if r < 0x11633 {
												return SpacingMark
											} else { // 0x11633 <= r
												return Extend
											}
										} else { // 0x1163b <= r
											if r < 0x1163d {
												return SpacingMark
											} else { // 0x1163d <= r
												return Extend
											}
										}
									}
								}
							}
						}
					} else { // 0x1163e <= r
						if r < 0x11ca8 {
							if r < 0x11720 {
								if r < 0x116ae {
									if r < 0x116ab {
										if r < 0x1163f {
											return SpacingMark
										} else { // 0x1163f <= r
											if r < 0x11641 {
												return Extend
											} else { // 0x11641 <= r
												return Any
											}
										}
									} else { // 0x116ab <= r
										if r < 0x116ac {
											return Extend
										} else { // 0x116ac <= r
											if r < 0x116ad {
												return SpacingMark
											} else { // 0x116ad <= r
												return Extend
											}
										}
									}
								} else { // 0x116ae <= r
									if r < 0x116b7 {
										if r < 0x116b0 {
											return SpacingMark
										} else { // 0x116b0 <= r
											if r < 0x116b6 {
												return Extend
											} else { // 0x116b6 <= r
												return SpacingMark
											}
										}
									} else { // 0x116b7 <= r
										if r < 0x116b8 {
											return Extend
										} else { // 0x116b8 <= r
											if r < 0x1171d {
												return Any
											} else { // 0x1171d <= r
												return Extend
											}
										}
									}
								}
							} else { // 0x11720 <= r
								if r < 0x11c30 {
									if r < 0x11727 {
										if r < 0x11722 {
											return SpacingMark
										} else { // 0x11722 <= r
											if r < 0x11726 {
												return Extend
											} else { // 0x11726 <= r
												return SpacingMark
											}
										}
									} else { // 0x11727 <= r
										if r < 0x1172c {
											return Extend
										} else { // 0x1172c <= r
											if r < 0x11c2f {
												return Any
											} else { // 0x11c2f <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x11c30 <= r
									if r < 0x11c3e {
										if r < 0x11c37 {
											return Extend
										} else { // 0x11c37 <= r
											if r < 0x11c38 {
												return Any
											} else { // 0x11c38 <= r
												return Extend
											}
										}
									} else { // 0x11c3e <= r
										if r < 0x11c40 {
											if r < 0x11c3f {
												return SpacingMark
											} else { // 0x11c3f <= r
												return Extend
											}
										} else { // 0x11c40 <= r
											if r < 0x11c92 {
												return Any
											} else { // 0x11c92 <= r
												return Extend
											}
										}
									}
								}
							}
						} else { // 0x11ca8 <= r
							if r < 0x16f7f {
								if r < 0x11cb5 {
									if r < 0x11cb1 {
										if r < 0x11ca9 {
											return Any
										} else { // 0x11ca9 <= r
											if r < 0x11caa {
												return SpacingMark
											} else { // 0x11caa <= r
												return Extend
											}
										}
									} else { // 0x11cb1 <= r
										if r < 0x11cb2 {
											return SpacingMark
										} else { // 0x11cb2 <= r
											if r < 0x11cb4 {
												return Extend
											} else { // 0x11cb4 <= r
												return SpacingMark
											}
										}
									}
								} else { // 0x11cb5 <= r
									if r < 0x16af5 {
										if r < 0x11cb7 {
											return Extend
										} else { // 0x11cb7 <= r
											if r < 0x16af0 {
												return Any
											} else { // 0x16af0 <= r
												return Extend
											}
										}
									} else { // 0x16af5 <= r
										if r < 0x16b37 {
											if r < 0x16b30 {
												return Any
											} else { // 0x16b30 <= r
												return Extend
											}
										} else { // 0x16b37 <= r
											if r < 0x16f51 {
												return Any
											} else { // 0x16f51 <= r
												return SpacingMark
											}
										}
									}
								}
							} else { // 0x16f7f <= r
								if r < 0x1bca4 {
									if r < 0x1bc9d {
										if r < 0x16f8f {
											return Any
										} else { // 0x16f8f <= r
											if r < 0x16f93 {
												return Extend
											} else { // 0x16f93 <= r
												return Any
											}
										}
									} else { // 0x1bc9d <= r
										if r < 0x1bc9f {
											return Extend
										} else { // 0x1bc9f <= r
											if r < 0x1bca0 {
												return Any
											} else { // 0x1bca0 <= r
												return Control
											}
										}
									}
								} else { // 0x1bca4 <= r
									if r < 0x1d167 {
										if r < 0x1d165 {
											return Any
										} else { // 0x1d165 <= r
											if r < 0x1d166 {
												return Extend
											} else { // 0x1d166 <= r
												return SpacingMark
											}
										}
									} else { // 0x1d167 <= r
										if r < 0x1d16d {
											if r < 0x1d16a {
												return Extend
											} else { // 0x1d16a <= r
												return Any
											}
										} else { // 0x1d16d <= r
											if r < 0x1d16e {
												return SpacingMark
											} else { // 0x1d16e <= r
												return Extend
											}
										}
									}
								}
							}
						}
					}
				} else { // 0x1d173 <= r
					if r < 0x1f46a {
						if r < 0x1e01b {
							if r < 0x1da6d {
								if r < 0x1d1ae {
									if r < 0x1d185 {
										if r < 0x1d17b {
											return Control
										} else { // 0x1d17b <= r
											if r < 0x1d183 {
												return Extend
											} else { // 0x1d183 <= r
												return Any
											}
										}
									} else { // 0x1d185 <= r
										if r < 0x1d18c {
											return Extend
										} else { // 0x1d18c <= r
											if r < 0x1d1aa {
												return Any
											} else { // 0x1d1aa <= r
												return Extend
											}
										}
									}
								} else { // 0x1d1ae <= r
									if r < 0x1da00 {
										if r < 0x1d242 {
											return Any
										} else { // 0x1d242 <= r
											if r < 0x1d245 {
												return Extend
											} else { // 0x1d245 <= r
												return Any
											}
										}
									} else { // 0x1da00 <= r
										if r < 0x1da37 {
											return Extend
										} else { // 0x1da37 <= r
											if r < 0x1da3b {
												return Any
											} else { // 0x1da3b <= r
												return Extend
											}
										}
									}
								}
							} else { // 0x1da6d <= r
								if r < 0x1daa0 {
									if r < 0x1da84 {
										if r < 0x1da75 {
											return Any
										} else { // 0x1da75 <= r
											if r < 0x1da76 {
												return Extend
											} else { // 0x1da76 <= r
												return Any
											}
										}
									} else { // 0x1da84 <= r
										if r < 0x1da85 {
											return Extend
										} else { // 0x1da85 <= r
											if r < 0x1da9b {
												return Any
											} else { // 0x1da9b <= r
												return Extend
											}
										}
									}
								} else { // 0x1daa0 <= r
									if r < 0x1e000 {
										if r < 0x1daa1 {
											return Any
										} else { // 0x1daa1 <= r
											if r < 0x1dab0 {
												return Extend
											} else { // 0x1dab0 <= r
												return Any
											}
										}
									} else { // 0x1e000 <= r
										if r < 0x1e008 {
											if r < 0x1e007 {
												return Extend
											} else { // 0x1e007 <= r
												return Any
											}
										} else { // 0x1e008 <= r
											if r < 0x1e019 {
												return Extend
											} else { // 0x1e019 <= r
												return Any
											}
										}
									}
								}
							}
						} else { // 0x1e01b <= r
							if r < 0x1f385 {
								if r < 0x1e8d0 {
									if r < 0x1e025 {
										if r < 0x1e022 {
											return Extend
										} else { // 0x1e022 <= r
											if r < 0x1e023 {
												return Any
											} else { // 0x1e023 <= r
												return Extend
											}
										}
									} else { // 0x1e025 <= r
										if r < 0x1e026 {
											return Any
										} else { // 0x1e026 <= r
											if r < 0x1e02b {
												return Extend
											} else { // 0x1e02b <= r
												return Any
											}
										}
									}
								} else { // 0x1e8d0 <= r
									if r < 0x1e94b {
										if r < 0x1e8d7 {
											return Extend
										} else { // 0x1e8d7 <= r
											if r < 0x1e944 {
												return Any
											} else { // 0x1e944 <= r
												return Extend
											}
										}
									} else { // 0x1e94b <= r
										if r < 0x1f1e6 {
											return Any
										} else { // 0x1f1e6 <= r
											if r < 0x1f200 {
												return Regional_Indicator
											} else { // 0x1f200 <= r
												return Any
											}
										}
									}
								}
							} else { // 0x1f385 <= r
								if r < 0x1f3fb {
									if r < 0x1f3c5 {
										if r < 0x1f386 {
											return E_Base
										} else { // 0x1f386 <= r
											if r < 0x1f3c3 {
												return Any
											} else { // 0x1f3c3 <= r
												return E_Base
											}
										}
									} else { // 0x1f3c5 <= r
										if r < 0x1f3ca {
											return Any
										} else { // 0x1f3ca <= r
											if r < 0x1f3cc {
												return E_Base
											} else { // 0x1f3cc <= r
												return Any
											}
										}
									}
								} else { // 0x1f3fb <= r
									if r < 0x1f444 {
										if r < 0x1f400 {
											return E_Modifier
										} else { // 0x1f400 <= r
											if r < 0x1f442 {
												return Any
											} else { // 0x1f442 <= r
												return E_Base
											}
										}
									} else { // 0x1f444 <= r
										if r < 0x1f451 {
											if r < 0x1f446 {
												return Any
											} else { // 0x1f446 <= r
												return E_Base
											}
										} else { // 0x1f451 <= r
											if r < 0x1f466 {
												return Any
											} else { // 0x1f466 <= r
												return E_Base_GAZ
											}
										}
									}
								}
							}
						}
					} else { // 0x1f46a <= r
						if r < 0x1f645 {
							if r < 0x1f48c {
								if r < 0x1f47d {
									if r < 0x1f470 {
										if r < 0x1f46e {
											return Any
										} else { // 0x1f46e <= r
											if r < 0x1f46f {
												return E_Base
											} else { // 0x1f46f <= r
												return Any
											}
										}
									} else { // 0x1f470 <= r
										if r < 0x1f479 {
											return E_Base
										} else { // 0x1f479 <= r
											if r < 0x1f47c {
												return Any
											} else { // 0x1f47c <= r
												return E_Base
											}
										}
									}
								} else { // 0x1f47d <= r
									if r < 0x1f485 {
										if r < 0x1f481 {
											return Any
										} else { // 0x1f481 <= r
											if r < 0x1f484 {
												return E_Base
											} else { // 0x1f484 <= r
												return Any
											}
										}
									} else { // 0x1f485 <= r
										if r < 0x1f488 {
											return E_Base
										} else { // 0x1f488 <= r
											if r < 0x1f48b {
												return Any
											} else { // 0x1f48b <= r
												return Glue_After_Zwj
											}
										}
									}
								}
							} else { // 0x1f48c <= r
								if r < 0x1f57b {
									if r < 0x1f575 {
										if r < 0x1f4aa {
											return Any
										} else { // 0x1f4aa <= r
											if r < 0x1f4ab {
												return E_Base
											} else { // 0x1f4ab <= r
												return Any
											}
										}
									} else { // 0x1f575 <= r
										if r < 0x1f576 {
											return E_Base
										} else { // 0x1f576 <= r
											if r < 0x1f57a {
												return Any
											} else { // 0x1f57a <= r
												return E_Base
											}
										}
									}
								} else { // 0x1f57b <= r
									if r < 0x1f595 {
										if r < 0x1f590 {
											return Any
										} else { // 0x1f590 <= r
											if r < 0x1f591 {
												return E_Base
											} else { // 0x1f591 <= r
												return Any
											}
										}
									} else { // 0x1f595 <= r
										if r < 0x1f5e8 {
											if r < 0x1f597 {
												return E_Base
											} else { // 0x1f597 <= r
												return Any
											}
										} else { // 0x1f5e8 <= r
											if r < 0x1f5e9 {
												return Glue_After_Zwj
											} else { // 0x1f5e9 <= r
												return Any
											}
										}
									}
								}
							}
						} else { // 0x1f645 <= r
							if r < 0x1f927 {
								if r < 0x1f6b4 {
									if r < 0x1f650 {
										if r < 0x1f648 {
											return E_Base
										} else { // 0x1f648 <= r
											if r < 0x1f64b {
												return Any
											} else { // 0x1f64b <= r
												return E_Base
											}
										}
									} else { // 0x1f650 <= r
										if r < 0x1f6a3 {
											return Any
										} else { // 0x1f6a3 <= r
											if r < 0x1f6a4 {
												return E_Base
											} else { // 0x1f6a4 <= r
												return Any
											}
										}
									}
								} else { // 0x1f6b4 <= r
									if r < 0x1f6c1 {
										if r < 0x1f6b7 {
											return E_Base
										} else { // 0x1f6b7 <= r
											if r < 0x1f6c0 {
												return Any
											} else { // 0x1f6c0 <= r
												return E_Base
											}
										}
									} else { // 0x1f6c1 <= r
										if r < 0x1f91f {
											if r < 0x1f918 {
												return Any
											} else { // 0x1f918 <= r
												return E_Base
											}
										} else { // 0x1f91f <= r
											if r < 0x1f926 {
												return Any
											} else { // 0x1f926 <= r
												return E_Base
											}
										}
									}
								}
							} else { // 0x1f927 <= r
								if r < 0x1f93f {
									if r < 0x1f933 {
										if r < 0x1f930 {
											return Any
										} else { // 0x1f930 <= r
											if r < 0x1f931 {
												return E_Base
											} else { // 0x1f931 <= r
												return Any
											}
										}
									} else { // 0x1f933 <= r
										if r < 0x1f93a {
											return E_Base
										} else { // 0x1f93a <= r
											if r < 0x1f93c {
												return Any
											} else { // 0x1f93c <= r
												return E_Base
											}
										}
									}
								} else { // 0x1f93f <= r
									if r < 0xe0080 {
										if r < 0xe0000 {
											return Any
										} else { // 0xe0000 <= r
											if r < 0xe0020 {
												return Control
											} else { // 0xe0020 <= r
												return Extend
											}
										}
									} else { // 0xe0080 <= r
										if r < 0xe01f0 {
											if r < 0xe0100 {
												return Control
											} else { // 0xe0100 <= r
												return Extend
											}
										} else { // 0xe01f0 <= r
											if r < 0x10ffff {
												return Control
											} else { // 0x10ffff <= r
												return OutOfRange
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

var keys = [...]rune{
	0x0000,
	0x000a,
	0x000b,
	0x000d,
	0x000e,
	0x0020,
	0x007f,
	0x00a0,
	0x00ad,
	0x00ae,
	0x0300,
	0x0370,
	0x0483,
	0x048a,
	0x0591,
	0x05be,
	0x05bf,
	0x05c0,
	0x05c1,
	0x05c3,
	0x05c4,
	0x05c6,
	0x05c7,
	0x05c8,
	0x0600,
	0x0606,
	0x0610,
	0x061b,
	0x061c,
	0x061d,
	0x064b,
	0x0660,
	0x0670,
	0x0671,
	0x06d6,
	0x06dd,
	0x06de,
	0x06df,
	0x06e5,
	0x06e7,
	0x06e9,
	0x06ea,
	0x06ee,
	0x070f,
	0x0710,
	0x0711,
	0x0712,
	0x0730,
	0x074b,
	0x07a6,
	0x07b1,
	0x07eb,
	0x07f4,
	0x0816,
	0x081a,
	0x081b,
	0x0824,
	0x0825,
	0x0828,
	0x0829,
	0x082e,
	0x0859,
	0x085c,
	0x08d4,
	0x08e2,
	0x08e3,
	0x0903,
	0x0904,
	0x093a,
	0x093b,
	0x093c,
	0x093d,
	0x093e,
	0x0941,
	0x0949,
	0x094d,
	0x094e,
	0x0950,
	0x0951,
	0x0958,
	0x0962,
	0x0964,
	0x0981,
	0x0982,
	0x0984,
	0x09bc,
	0x09bd,
	0x09be,
	0x09bf,
	0x09c1,
	0x09c5,
	0x09c7,
	0x09c9,
	0x09cb,
	0x09cd,
	0x09ce,
	0x09d7,
	0x09d8,
	0x09e2,
	0x09e4,
	0x0a01,
	0x0a03,
	0x0a04,
	0x0a3c,
	0x0a3d,
	0x0a3e,
	0x0a41,
	0x0a43,
	0x0a47,
	0x0a49,
	0x0a4b,
	0x0a4e,
	0x0a51,
	0x0a52,
	0x0a70,
	0x0a72,
	0x0a75,
	0x0a76,
	0x0a81,
	0x0a83,
	0x0a84,
	0x0abc,
	0x0abd,
	0x0abe,
	0x0ac1,
	0x0ac6,
	0x0ac7,
	0x0ac9,
	0x0aca,
	0x0acb,
	0x0acd,
	0x0ace,
	0x0ae2,
	0x0ae4,
	0x0b01,
	0x0b02,
	0x0b04,
	0x0b3c,
	0x0b3d,
	0x0b3e,
	0x0b40,
	0x0b41,
	0x0b45,
	0x0b47,
	0x0b49,
	0x0b4b,
	0x0b4d,
	0x0b4e,
	0x0b56,
	0x0b58,
	0x0b62,
	0x0b64,
	0x0b82,
	0x0b83,
	0x0bbe,
	0x0bbf,
	0x0bc0,
	0x0bc1,
	0x0bc3,
	0x0bc6,
	0x0bc9,
	0x0bca,
	0x0bcd,
	0x0bce,
	0x0bd7,
	0x0bd8,
	0x0c00,
	0x0c01,
	0x0c04,
	0x0c3e,
	0x0c41,
	0x0c45,
	0x0c46,
	0x0c49,
	0x0c4a,
	0x0c4e,
	0x0c55,
	0x0c57,
	0x0c62,
	0x0c64,
	0x0c81,
	0x0c82,
	0x0c84,
	0x0cbc,
	0x0cbd,
	0x0cbe,
	0x0cbf,
	0x0cc0,
	0x0cc2,
	0x0cc3,
	0x0cc5,
	0x0cc6,
	0x0cc7,
	0x0cc9,
	0x0cca,
	0x0ccc,
	0x0cce,
	0x0cd5,
	0x0cd7,
	0x0ce2,
	0x0ce4,
	0x0d01,
	0x0d02,
	0x0d04,
	0x0d3e,
	0x0d3f,
	0x0d41,
	0x0d45,
	0x0d46,
	0x0d49,
	0x0d4a,
	0x0d4d,
	0x0d4e,
	0x0d4f,
	0x0d57,
	0x0d58,
	0x0d62,
	0x0d64,
	0x0d82,
	0x0d84,
	0x0dca,
	0x0dcb,
	0x0dcf,
	0x0dd0,
	0x0dd2,
	0x0dd5,
	0x0dd6,
	0x0dd7,
	0x0dd8,
	0x0ddf,
	0x0de0,
	0x0df2,
	0x0df4,
	0x0e31,
	0x0e32,
	0x0e33,
	0x0e34,
	0x0e3b,
	0x0e47,
	0x0e4f,
	0x0eb1,
	0x0eb2,
	0x0eb3,
	0x0eb4,
	0x0eba,
	0x0ebb,
	0x0ebd,
	0x0ec8,
	0x0ece,
	0x0f18,
	0x0f1a,
	0x0f35,
	0x0f36,
	0x0f37,
	0x0f38,
	0x0f39,
	0x0f3a,
	0x0f3e,
	0x0f40,
	0x0f71,
	0x0f7f,
	0x0f80,
	0x0f85,
	0x0f86,
	0x0f88,
	0x0f8d,
	0x0f98,
	0x0f99,
	0x0fbd,
	0x0fc6,
	0x0fc7,
	0x102d,
	0x1031,
	0x1032,
	0x1038,
	0x1039,
	0x103b,
	0x103d,
	0x103f,
	0x1056,
	0x1058,
	0x105a,
	0x105e,
	0x1061,
	0x1071,
	0x1075,
	0x1082,
	0x1083,
	0x1084,
	0x1085,
	0x1087,
	0x108d,
	0x108e,
	0x109d,
	0x109e,
	0x1100,
	0x1160,
	0x11a8,
	0x1200,
	0x135d,
	0x1360,
	0x1712,
	0x1715,
	0x1732,
	0x1735,
	0x1752,
	0x1754,
	0x1772,
	0x1774,
	0x17b4,
	0x17b6,
	0x17b7,
	0x17be,
	0x17c6,
	0x17c7,
	0x17c9,
	0x17d4,
	0x17dd,
	0x17de,
	0x180b,
	0x180e,
	0x180f,
	0x1885,
	0x1887,
	0x18a9,
	0x18aa,
	0x1920,
	0x1923,
	0x1927,
	0x1929,
	0x192c,
	0x1930,
	0x1932,
	0x1933,
	0x1939,
	0x193c,
	0x1a17,
	0x1a19,
	0x1a1b,
	0x1a1c,
	0x1a55,
	0x1a56,
	0x1a57,
	0x1a58,
	0x1a5f,
	0x1a60,
	0x1a61,
	0x1a62,
	0x1a63,
	0x1a65,
	0x1a6d,
	0x1a73,
	0x1a7d,
	0x1a7f,
	0x1a80,
	0x1ab0,
	0x1abf,
	0x1b00,
	0x1b04,
	0x1b05,
	0x1b34,
	0x1b35,
	0x1b36,
	0x1b3b,
	0x1b3c,
	0x1b3d,
	0x1b42,
	0x1b43,
	0x1b45,
	0x1b6b,
	0x1b74,
	0x1b80,
	0x1b82,
	0x1b83,
	0x1ba1,
	0x1ba2,
	0x1ba6,
	0x1ba8,
	0x1baa,
	0x1bab,
	0x1bae,
	0x1be6,
	0x1be7,
	0x1be8,
	0x1bea,
	0x1bed,
	0x1bee,
	0x1bef,
	0x1bf2,
	0x1bf4,
	0x1c24,
	0x1c2c,
	0x1c34,
	0x1c36,
	0x1c38,
	0x1cd0,
	0x1cd3,
	0x1cd4,
	0x1ce1,
	0x1ce2,
	0x1ce9,
	0x1ced,
	0x1cee,
	0x1cf2,
	0x1cf4,
	0x1cf5,
	0x1cf8,
	0x1cfa,
	0x1dc0,
	0x1df6,
	0x1dfb,
	0x1e00,
	0x200b,
	0x200c,
	0x200d,
	0x200e,
	0x2010,
	0x2028,
	0x202f,
	0x2060,
	0x2070,
	0x20d0,
	0x20f1,
	0x261d,
	0x261e,
	0x26f9,
	0x26fa,
	0x270a,
	0x270e,
	0x2764,
	0x2765,
	0x2cef,
	0x2cf2,
	0x2d7f,
	0x2d80,
	0x2de0,
	0x2e00,
	0x302a,
	0x3030,
	0x3099,
	0x309b,
	0xa66f,
	0xa673,
	0xa674,
	0xa67e,
	0xa69e,
	0xa6a0,
	0xa6f0,
	0xa6f2,
	0xa802,
	0xa803,
	0xa806,
	0xa807,
	0xa80b,
	0xa80c,
	0xa823,
	0xa825,
	0xa827,
	0xa828,
	0xa880,
	0xa882,
	0xa8b4,
	0xa8c4,
	0xa8c6,
	0xa8e0,
	0xa8f2,
	0xa926,
	0xa92e,
	0xa947,
	0xa952,
	0xa954,
	0xa960,
	0xa97d,
	0xa980,
	0xa983,
	0xa984,
	0xa9b3,
	0xa9b4,
	0xa9b6,
	0xa9ba,
	0xa9bc,
	0xa9bd,
	0xa9c1,
	0xa9e5,
	0xa9e6,
	0xaa29,
	0xaa2f,
	0xaa31,
	0xaa33,
	0xaa35,
	0xaa37,
	0xaa43,
	0xaa44,
	0xaa4c,
	0xaa4d,
	0xaa4e,
	0xaa7c,
	0xaa7d,
	0xaab0,
	0xaab1,
	0xaab2,
	0xaab5,
	0xaab7,
	0xaab9,
	0xaabe,
	0xaac0,
	0xaac1,
	0xaac2,
	0xaaeb,
	0xaaec,
	0xaaee,
	0xaaf0,
	0xaaf5,
	0xaaf6,
	0xaaf7,
	0xabe3,
	0xabe5,
	0xabe6,
	0xabe8,
	0xabe9,
	0xabeb,
	0xabec,
	0xabed,
	0xabee,
	0xac00,
	0xac01,
	0xac1c,
	0xac1d,
	0xac38,
	0xac39,
	0xac54,
	0xac55,
	0xac70,
	0xac71,
	0xac8c,
	0xac8d,
	0xaca8,
	0xaca9,
	0xacc4,
	0xacc5,
	0xace0,
	0xace1,
	0xacfc,
	0xacfd,
	0xad18,
	0xad19,
	0xad34,
	0xad35,
	0xad50,
	0xad51,
	0xad6c,
	0xad6d,
	0xad88,
	0xad89,
	0xada4,
	0xada5,
	0xadc0,
	0xadc1,
	0xaddc,
	0xaddd,
	0xadf8,
	0xadf9,
	0xae14,
	0xae15,
	0xae30,
	0xae31,
	0xae4c,
	0xae4d,
	0xae68,
	0xae69,
	0xae84,
	0xae85,
	0xaea0,
	0xaea1,
	0xaebc,
	0xaebd,
	0xaed8,
	0xaed9,
	0xaef4,
	0xaef5,
	0xaf10,
	0xaf11,
	0xaf2c,
	0xaf2d,
	0xaf48,
	0xaf49,
	0xaf64,
	0xaf65,
	0xaf80,
	0xaf81,
	0xaf9c,
	0xaf9d,
	0xafb8,
	0xafb9,
	0xafd4,
	0xafd5,
	0xaff0,
	0xaff1,
	0xb00c,
	0xb00d,
	0xb028,
	0xb029,
	0xb044,
	0xb045,
	0xb060,
	0xb061,
	0xb07c,
	0xb07d,
	0xb098,
	0xb099,
	0xb0b4,
	0xb0b5,
	0xb0d0,
	0xb0d1,
	0xb0ec,
	0xb0ed,
	0xb108,
	0xb109,
	0xb124,
	0xb125,
	0xb140,
	0xb141,
	0xb15c,
	0xb15d,
	0xb178,
	0xb179,
	0xb194,
	0xb195,
	0xb1b0,
	0xb1b1,
	0xb1cc,
	0xb1cd,
	0xb1e8,
	0xb1e9,
	0xb204,
	0xb205,
	0xb220,
	0xb221,
	0xb23c,
	0xb23d,
	0xb258,
	0xb259,
	0xb274,
	0xb275,
	0xb290,
	0xb291,
	0xb2ac,
	0xb2ad,
	0xb2c8,
	0xb2c9,
	0xb2e4,
	0xb2e5,
	0xb300,
	0xb301,
	0xb31c,
	0xb31d,
	0xb338,
	0xb339,
	0xb354,
	0xb355,
	0xb370,
	0xb371,
	0xb38c,
	0xb38d,
	0xb3a8,
	0xb3a9,
	0xb3c4,
	0xb3c5,
	0xb3e0,
	0xb3e1,
	0xb3fc,
	0xb3fd,
	0xb418,
	0xb419,
	0xb434,
	0xb435,
	0xb450,
	0xb451,
	0xb46c,
	0xb46d,
	0xb488,
	0xb489,
	0xb4a4,
	0xb4a5,
	0xb4c0,
	0xb4c1,
	0xb4dc,
	0xb4dd,
	0xb4f8,
	0xb4f9,
	0xb514,
	0xb515,
	0xb530,
	0xb531,
	0xb54c,
	0xb54d,
	0xb568,
	0xb569,
	0xb584,
	0xb585,
	0xb5a0,
	0xb5a1,
	0xb5bc,
	0xb5bd,
	0xb5d8,
	0xb5d9,
	0xb5f4,
	0xb5f5,
	0xb610,
	0xb611,
	0xb62c,
	0xb62d,
	0xb648,
	0xb649,
	0xb664,
	0xb665,
	0xb680,
	0xb681,
	0xb69c,
	0xb69d,
	0xb6b8,
	0xb6b9,
	0xb6d4,
	0xb6d5,
	0xb6f0,
	0xb6f1,
	0xb70c,
	0xb70d,
	0xb728,
	0xb729,
	0xb744,
	0xb745,
	0xb760,
	0xb761,
	0xb77c,
	0xb77d,
	0xb798,
	0xb799,
	0xb7b4,
	0xb7b5,
	0xb7d0,
	0xb7d1,
	0xb7ec,
	0xb7ed,
	0xb808,
	0xb809,
	0xb824,
	0xb825,
	0xb840,
	0xb841,
	0xb85c,
	0xb85d,
	0xb878,
	0xb879,
	0xb894,
	0xb895,
	0xb8b0,
	0xb8b1,
	0xb8cc,
	0xb8cd,
	0xb8e8,
	0xb8e9,
	0xb904,
	0xb905,
	0xb920,
	0xb921,
	0xb93c,
	0xb93d,
	0xb958,
	0xb959,
	0xb974,
	0xb975,
	0xb990,
	0xb991,
	0xb9ac,
	0xb9ad,
	0xb9c8,
	0xb9c9,
	0xb9e4,
	0xb9e5,
	0xba00,
	0xba01,
	0xba1c,
	0xba1d,
	0xba38,
	0xba39,
	0xba54,
	0xba55,
	0xba70,
	0xba71,
	0xba8c,
	0xba8d,
	0xbaa8,
	0xbaa9,
	0xbac4,
	0xbac5,
	0xbae0,
	0xbae1,
	0xbafc,
	0xbafd,
	0xbb18,
	0xbb19,
	0xbb34,
	0xbb35,
	0xbb50,
	0xbb51,
	0xbb6c,
	0xbb6d,
	0xbb88,
	0xbb89,
	0xbba4,
	0xbba5,
	0xbbc0,
	0xbbc1,
	0xbbdc,
	0xbbdd,
	0xbbf8,
	0xbbf9,
	0xbc14,
	0xbc15,
	0xbc30,
	0xbc31,
	0xbc4c,
	0xbc4d,
	0xbc68,
	0xbc69,
	0xbc84,
	0xbc85,
	0xbca0,
	0xbca1,
	0xbcbc,
	0xbcbd,
	0xbcd8,
	0xbcd9,
	0xbcf4,
	0xbcf5,
	0xbd10,
	0xbd11,
	0xbd2c,
	0xbd2d,
	0xbd48,
	0xbd49,
	0xbd64,
	0xbd65,
	0xbd80,
	0xbd81,
	0xbd9c,
	0xbd9d,
	0xbdb8,
	0xbdb9,
	0xbdd4,
	0xbdd5,
	0xbdf0,
	0xbdf1,
	0xbe0c,
	0xbe0d,
	0xbe28,
	0xbe29,
	0xbe44,
	0xbe45,
	0xbe60,
	0xbe61,
	0xbe7c,
	0xbe7d,
	0xbe98,
	0xbe99,
	0xbeb4,
	0xbeb5,
	0xbed0,
	0xbed1,
	0xbeec,
	0xbeed,
	0xbf08,
	0xbf09,
	0xbf24,
	0xbf25,
	0xbf40,
	0xbf41,
	0xbf5c,
	0xbf5d,
	0xbf78,
	0xbf79,
	0xbf94,
	0xbf95,
	0xbfb0,
	0xbfb1,
	0xbfcc,
	0xbfcd,
	0xbfe8,
	0xbfe9,
	0xc004,
	0xc005,
	0xc020,
	0xc021,
	0xc03c,
	0xc03d,
	0xc058,
	0xc059,
	0xc074,
	0xc075,
	0xc090,
	0xc091,
	0xc0ac,
	0xc0ad,
	0xc0c8,
	0xc0c9,
	0xc0e4,
	0xc0e5,
	0xc100,
	0xc101,
	0xc11c,
	0xc11d,
	0xc138,
	0xc139,
	0xc154,
	0xc155,
	0xc170,
	0xc171,
	0xc18c,
	0xc18d,
	0xc1a8,
	0xc1a9,
	0xc1c4,
	0xc1c5,
	0xc1e0,
	0xc1e1,
	0xc1fc,
	0xc1fd,
	0xc218,
	0xc219,
	0xc234,
	0xc235,
	0xc250,
	0xc251,
	0xc26c,
	0xc26d,
	0xc288,
	0xc289,
	0xc2a4,
	0xc2a5,
	0xc2c0,
	0xc2c1,
	0xc2dc,
	0xc2dd,
	0xc2f8,
	0xc2f9,
	0xc314,
	0xc315,
	0xc330,
	0xc331,
	0xc34c,
	0xc34d,
	0xc368,
	0xc369,
	0xc384,
	0xc385,
	0xc3a0,
	0xc3a1,
	0xc3bc,
	0xc3bd,
	0xc3d8,
	0xc3d9,
	0xc3f4,
	0xc3f5,
	0xc410,
	0xc411,
	0xc42c,
	0xc42d,
	0xc448,
	0xc449,
	0xc464,
	0xc465,
	0xc480,
	0xc481,
	0xc49c,
	0xc49d,
	0xc4b8,
	0xc4b9,
	0xc4d4,
	0xc4d5,
	0xc4f0,
	0xc4f1,
	0xc50c,
	0xc50d,
	0xc528,
	0xc529,
	0xc544,
	0xc545,
	0xc560,
	0xc561,
	0xc57c,
	0xc57d,
	0xc598,
	0xc599,
	0xc5b4,
	0xc5b5,
	0xc5d0,
	0xc5d1,
	0xc5ec,
	0xc5ed,
	0xc608,
	0xc609,
	0xc624,
	0xc625,
	0xc640,
	0xc641,
	0xc65c,
	0xc65d,
	0xc678,
	0xc679,
	0xc694,
	0xc695,
	0xc6b0,
	0xc6b1,
	0xc6cc,
	0xc6cd,
	0xc6e8,
	0xc6e9,
	0xc704,
	0xc705,
	0xc720,
	0xc721,
	0xc73c,
	0xc73d,
	0xc758,
	0xc759,
	0xc774,
	0xc775,
	0xc790,
	0xc791,
	0xc7ac,
	0xc7ad,
	0xc7c8,
	0xc7c9,
	0xc7e4,
	0xc7e5,
	0xc800,
	0xc801,
	0xc81c,
	0xc81d,
	0xc838,
	0xc839,
	0xc854,
	0xc855,
	0xc870,
	0xc871,
	0xc88c,
	0xc88d,
	0xc8a8,
	0xc8a9,
	0xc8c4,
	0xc8c5,
	0xc8e0,
	0xc8e1,
	0xc8fc,
	0xc8fd,
	0xc918,
	0xc919,
	0xc934,
	0xc935,
	0xc950,
	0xc951,
	0xc96c,
	0xc96d,
	0xc988,
	0xc989,
	0xc9a4,
	0xc9a5,
	0xc9c0,
	0xc9c1,
	0xc9dc,
	0xc9dd,
	0xc9f8,
	0xc9f9,
	0xca14,
	0xca15,
	0xca30,
	0xca31,
	0xca4c,
	0xca4d,
	0xca68,
	0xca69,
	0xca84,
	0xca85,
	0xcaa0,
	0xcaa1,
	0xcabc,
	0xcabd,
	0xcad8,
	0xcad9,
	0xcaf4,
	0xcaf5,
	0xcb10,
	0xcb11,
	0xcb2c,
	0xcb2d,
	0xcb48,
	0xcb49,
	0xcb64,
	0xcb65,
	0xcb80,
	0xcb81,
	0xcb9c,
	0xcb9d,
	0xcbb8,
	0xcbb9,
	0xcbd4,
	0xcbd5,
	0xcbf0,
	0xcbf1,
	0xcc0c,
	0xcc0d,
	0xcc28,
	0xcc29,
	0xcc44,
	0xcc45,
	0xcc60,
	0xcc61,
	0xcc7c,
	0xcc7d,
	0xcc98,
	0xcc99,
	0xccb4,
	0xccb5,
	0xccd0,
	0xccd1,
	0xccec,
	0xcced,
	0xcd08,
	0xcd09,
	0xcd24,
	0xcd25,
	0xcd40,
	0xcd41,
	0xcd5c,
	0xcd5d,
	0xcd78,
	0xcd79,
	0xcd94,
	0xcd95,
	0xcdb0,
	0xcdb1,
	0xcdcc,
	0xcdcd,
	0xcde8,
	0xcde9,
	0xce04,
	0xce05,
	0xce20,
	0xce21,
	0xce3c,
	0xce3d,
	0xce58,
	0xce59,
	0xce74,
	0xce75,
	0xce90,
	0xce91,
	0xceac,
	0xcead,
	0xcec8,
	0xcec9,
	0xcee4,
	0xcee5,
	0xcf00,
	0xcf01,
	0xcf1c,
	0xcf1d,
	0xcf38,
	0xcf39,
	0xcf54,
	0xcf55,
	0xcf70,
	0xcf71,
	0xcf8c,
	0xcf8d,
	0xcfa8,
	0xcfa9,
	0xcfc4,
	0xcfc5,
	0xcfe0,
	0xcfe1,
	0xcffc,
	0xcffd,
	0xd018,
	0xd019,
	0xd034,
	0xd035,
	0xd050,
	0xd051,
	0xd06c,
	0xd06d,
	0xd088,
	0xd089,
	0xd0a4,
	0xd0a5,
	0xd0c0,
	0xd0c1,
	0xd0dc,
	0xd0dd,
	0xd0f8,
	0xd0f9,
	0xd114,
	0xd115,
	0xd130,
	0xd131,
	0xd14c,
	0xd14d,
	0xd168,
	0xd169,
	0xd184,
	0xd185,
	0xd1a0,
	0xd1a1,
	0xd1bc,
	0xd1bd,
	0xd1d8,
	0xd1d9,
	0xd1f4,
	0xd1f5,
	0xd210,
	0xd211,
	0xd22c,
	0xd22d,
	0xd248,
	0xd249,
	0xd264,
	0xd265,
	0xd280,
	0xd281,
	0xd29c,
	0xd29d,
	0xd2b8,
	0xd2b9,
	0xd2d4,
	0xd2d5,
	0xd2f0,
	0xd2f1,
	0xd30c,
	0xd30d,
	0xd328,
	0xd329,
	0xd344,
	0xd345,
	0xd360,
	0xd361,
	0xd37c,
	0xd37d,
	0xd398,
	0xd399,
	0xd3b4,
	0xd3b5,
	0xd3d0,
	0xd3d1,
	0xd3ec,
	0xd3ed,
	0xd408,
	0xd409,
	0xd424,
	0xd425,
	0xd440,
	0xd441,
	0xd45c,
	0xd45d,
	0xd478,
	0xd479,
	0xd494,
	0xd495,
	0xd4b0,
	0xd4b1,
	0xd4cc,
	0xd4cd,
	0xd4e8,
	0xd4e9,
	0xd504,
	0xd505,
	0xd520,
	0xd521,
	0xd53c,
	0xd53d,
	0xd558,
	0xd559,
	0xd574,
	0xd575,
	0xd590,
	0xd591,
	0xd5ac,
	0xd5ad,
	0xd5c8,
	0xd5c9,
	0xd5e4,
	0xd5e5,
	0xd600,
	0xd601,
	0xd61c,
	0xd61d,
	0xd638,
	0xd639,
	0xd654,
	0xd655,
	0xd670,
	0xd671,
	0xd68c,
	0xd68d,
	0xd6a8,
	0xd6a9,
	0xd6c4,
	0xd6c5,
	0xd6e0,
	0xd6e1,
	0xd6fc,
	0xd6fd,
	0xd718,
	0xd719,
	0xd734,
	0xd735,
	0xd750,
	0xd751,
	0xd76c,
	0xd76d,
	0xd788,
	0xd789,
	0xd7a4,
	0xd7b0,
	0xd7c7,
	0xd7cb,
	0xd7fc,
	0xd800,
	0xe000,
	0xfb1e,
	0xfb1f,
	0xfe00,
	0xfe10,
	0xfe20,
	0xfe30,
	0xfeff,
	0xff00,
	0xff9e,
	0xffa0,
	0xfff0,
	0xfffc,
	0x101fd,
	0x101fe,
	0x102e0,
	0x102e1,
	0x10376,
	0x1037b,
	0x10a01,
	0x10a04,
	0x10a05,
	0x10a07,
	0x10a0c,
	0x10a10,
	0x10a38,
	0x10a3b,
	0x10a3f,
	0x10a40,
	0x10ae5,
	0x10ae7,
	0x11000,
	0x11001,
	0x11002,
	0x11003,
	0x11038,
	0x11047,
	0x1107f,
	0x11082,
	0x11083,
	0x110b0,
	0x110b3,
	0x110b7,
	0x110b9,
	0x110bb,
	0x110bd,
	0x110be,
	0x11100,
	0x11103,
	0x11127,
	0x1112c,
	0x1112d,
	0x11135,
	0x11173,
	0x11174,
	0x11180,
	0x11182,
	0x11183,
	0x111b3,
	0x111b6,
	0x111bf,
	0x111c1,
	0x111c2,
	0x111c4,
	0x111ca,
	0x111cd,
	0x1122c,
	0x1122f,
	0x11232,
	0x11234,
	0x11235,
	0x11236,
	0x11238,
	0x1123e,
	0x1123f,
	0x112df,
	0x112e0,
	0x112e3,
	0x112eb,
	0x11300,
	0x11302,
	0x11304,
	0x1133c,
	0x1133d,
	0x1133e,
	0x1133f,
	0x11340,
	0x11341,
	0x11345,
	0x11347,
	0x11349,
	0x1134b,
	0x1134e,
	0x11357,
	0x11358,
	0x11362,
	0x11364,
	0x11366,
	0x1136d,
	0x11370,
	0x11375,
	0x11435,
	0x11438,
	0x11440,
	0x11442,
	0x11445,
	0x11446,
	0x11447,
	0x114b0,
	0x114b1,
	0x114b3,
	0x114b9,
	0x114ba,
	0x114bb,
	0x114bd,
	0x114be,
	0x114bf,
	0x114c1,
	0x114c2,
	0x114c4,
	0x115af,
	0x115b0,
	0x115b2,
	0x115b6,
	0x115b8,
	0x115bc,
	0x115be,
	0x115bf,
	0x115c1,
	0x115dc,
	0x115de,
	0x11630,
	0x11633,
	0x1163b,
	0x1163d,
	0x1163e,
	0x1163f,
	0x11641,
	0x116ab,
	0x116ac,
	0x116ad,
	0x116ae,
	0x116b0,
	0x116b6,
	0x116b7,
	0x116b8,
	0x1171d,
	0x11720,
	0x11722,
	0x11726,
	0x11727,
	0x1172c,
	0x11c2f,
	0x11c30,
	0x11c37,
	0x11c38,
	0x11c3e,
	0x11c3f,
	0x11c40,
	0x11c92,
	0x11ca8,
	0x11ca9,
	0x11caa,
	0x11cb1,
	0x11cb2,
	0x11cb4,
	0x11cb5,
	0x11cb7,
	0x16af0,
	0x16af5,
	0x16b30,
	0x16b37,
	0x16f51,
	0x16f7f,
	0x16f8f,
	0x16f93,
	0x1bc9d,
	0x1bc9f,
	0x1bca0,
	0x1bca4,
	0x1d165,
	0x1d166,
	0x1d167,
	0x1d16a,
	0x1d16d,
	0x1d16e,
	0x1d173,
	0x1d17b,
	0x1d183,
	0x1d185,
	0x1d18c,
	0x1d1aa,
	0x1d1ae,
	0x1d242,
	0x1d245,
	0x1da00,
	0x1da37,
	0x1da3b,
	0x1da6d,
	0x1da75,
	0x1da76,
	0x1da84,
	0x1da85,
	0x1da9b,
	0x1daa0,
	0x1daa1,
	0x1dab0,
	0x1e000,
	0x1e007,
	0x1e008,
	0x1e019,
	0x1e01b,
	0x1e022,
	0x1e023,
	0x1e025,
	0x1e026,
	0x1e02b,
	0x1e8d0,
	0x1e8d7,
	0x1e944,
	0x1e94b,
	0x1f1e6,
	0x1f200,
	0x1f385,
	0x1f386,
	0x1f3c3,
	0x1f3c5,
	0x1f3ca,
	0x1f3cc,
	0x1f3fb,
	0x1f400,
	0x1f442,
	0x1f444,
	0x1f446,
	0x1f451,
	0x1f466,
	0x1f46a,
	0x1f46e,
	0x1f46f,
	0x1f470,
	0x1f479,
	0x1f47c,
	0x1f47d,
	0x1f481,
	0x1f484,
	0x1f485,
	0x1f488,
	0x1f48b,
	0x1f48c,
	0x1f4aa,
	0x1f4ab,
	0x1f575,
	0x1f576,
	0x1f57a,
	0x1f57b,
	0x1f590,
	0x1f591,
	0x1f595,
	0x1f597,
	0x1f5e8,
	0x1f5e9,
	0x1f645,
	0x1f648,
	0x1f64b,
	0x1f650,
	0x1f6a3,
	0x1f6a4,
	0x1f6b4,
	0x1f6b7,
	0x1f6c0,
	0x1f6c1,
	0x1f918,
	0x1f91f,
	0x1f926,
	0x1f927,
	0x1f930,
	0x1f931,
	0x1f933,
	0x1f93a,
	0x1f93c,
	0x1f93f,
	0xe0000,
	0xe0020,
	0xe0080,
	0xe0100,
	0xe01f0,
	0x10ffff,
}
var klen = len(keys)

var vals = [...]Prop{
	Control,
	LF,
	Control,
	CR,
	Control,
	Any,
	Control,
	Any,
	Control,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Prepend,
	Any,
	Extend,
	Any,
	Control,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Prepend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Prepend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Prepend,
	Extend,
	SpacingMark,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	Prepend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	L,
	V,
	T,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Control,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Control,
	Extend,
	ZWJ,
	Control,
	Any,
	Control,
	Any,
	Control,
	Any,
	Extend,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	Glue_After_Zwj,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	L,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	Any,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	LV,
	LVT,
	Any,
	V,
	Any,
	T,
	Any,
	Control,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Control,
	Any,
	Extend,
	Any,
	Control,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Prepend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	Prepend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Any,
	SpacingMark,
	Any,
	SpacingMark,
	Any,
	Extend,
	Any,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	SpacingMark,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	SpacingMark,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Control,
	Any,
	Extend,
	SpacingMark,
	Extend,
	Any,
	SpacingMark,
	Extend,
	Control,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Extend,
	Any,
	Regional_Indicator,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Modifier,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base_GAZ,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	Glue_After_Zwj,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	Glue_After_Zwj,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	E_Base,
	Any,
	Control,
	Extend,
	Control,
	Extend,
	Control,
	OutOfRange,
}

func Property2(r rune) Prop {
	a, b := 0, klen
	for {
		c := b - a
		if c <= 1 {
			break
		}
		m := a + c/2
		if r < keys[m] {
			b = m
		} else {
			a = m
		}
	}
	return vals[a]
}
