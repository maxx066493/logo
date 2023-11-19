// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package fr

// expBySqrtExp is equivalent to z.Exp(x, 221fc8bf5346d7e168584bf946c1e6a48e68f3c8cb5f873d7)
//
// uses github.com/mmcloughlin/addchain v0.4.0 to generate a shorter addition chain
func (z *Element) expBySqrtExp(x Element) *Element {
	// addition chain:
	//
	//	_10       = 2*1
	//	_100      = 2*_10
	//	_101      = 1 + _100
	//	_110      = 1 + _101
	//	_111      = 1 + _110
	//	_1001     = _10 + _111
	//	_1011     = _10 + _1001
	//	_1101     = _10 + _1011
	//	_1111     = _10 + _1101
	//	_10001    = _10 + _1111
	//	_10111    = _110 + _10001
	//	_11001    = _10 + _10111
	//	_11011    = _10 + _11001
	//	_110110   = 2*_11011
	//	_111111   = _1001 + _110110
	//	_1111110  = 2*_111111
	//	_1111111  = 1 + _1111110
	//	_10001000 = _1001 + _1111111
	//	i42       = ((_10001000 << 8 + _1111111) << 7 + _10001) << 7
	//	i55       = ((_111111 + i42) << 4 + _101) << 6 + _1101
	//	i74       = ((i55 << 8 + _11011) << 2 + 1) << 7
	//	i87       = ((_111111 + i74) << 8 + _1011) << 2 + 1
	//	i113      = ((i87 << 8 + _1011) << 8 + _1001) << 8
	//	i129      = ((_1111111 + i113) << 5 + _101) << 8 + _11011
	//	i151      = ((i129 << 9 + _1111) << 6 + _1101) << 5
	//	i166      = ((_1001 + i151) << 7 + _10001) << 5 + _11001
	//	i184      = ((i166 << 3 + _101) << 7 + _1111) << 6
	//	i198      = ((_1111 + i184) << 7 + _10001) << 4 + _1001
	//	i219      = ((i198 << 5 + _1101) << 7 + _111111) << 7
	//	return      ((_111 + i219) << 6 + _1111) << 6 + _10111
	//
	// Operations: 190 squares 44 multiplies

	// Allocate Temporaries.
	var (
		t0  = new(Element)
		t1  = new(Element)
		t2  = new(Element)
		t3  = new(Element)
		t4  = new(Element)
		t5  = new(Element)
		t6  = new(Element)
		t7  = new(Element)
		t8  = new(Element)
		t9  = new(Element)
		t10 = new(Element)
		t11 = new(Element)
	)

	// var t0,t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11 Element
	// Step 1: t2 = x^0x2
	t2.Square(&x)

	// Step 2: z = x^0x4
	z.Square(t2)

	// Step 3: t6 = x^0x5
	t6.Mul(&x, z)

	// Step 4: z = x^0x6
	z.Mul(&x, t6)

	// Step 5: t1 = x^0x7
	t1.Mul(&x, z)

	// Step 6: t4 = x^0x9
	t4.Mul(t2, t1)

	// Step 7: t10 = x^0xb
	t10.Mul(t2, t4)

	// Step 8: t3 = x^0xd
	t3.Mul(t2, t10)

	// Step 9: t0 = x^0xf
	t0.Mul(t2, t3)

	// Step 10: t5 = x^0x11
	t5.Mul(t2, t0)

	// Step 11: z = x^0x17
	z.Mul(z, t5)

	// Step 12: t7 = x^0x19
	t7.Mul(t2, z)

	// Step 13: t8 = x^0x1b
	t8.Mul(t2, t7)

	// Step 14: t2 = x^0x36
	t2.Square(t8)

	// Step 15: t2 = x^0x3f
	t2.Mul(t4, t2)

	// Step 16: t9 = x^0x7e
	t9.Square(t2)

	// Step 17: t9 = x^0x7f
	t9.Mul(&x, t9)

	// Step 18: t11 = x^0x88
	t11.Mul(t4, t9)

	// Step 26: t11 = x^0x8800
	for s := 0; s < 8; s++ {
		t11.Square(t11)
	}

	// Step 27: t11 = x^0x887f
	t11.Mul(t9, t11)

	// Step 34: t11 = x^0x443f80
	for s := 0; s < 7; s++ {
		t11.Square(t11)
	}

	// Step 35: t11 = x^0x443f91
	t11.Mul(t5, t11)

	// Step 42: t11 = x^0x221fc880
	for s := 0; s < 7; s++ {
		t11.Square(t11)
	}

	// Step 43: t11 = x^0x221fc8bf
	t11.Mul(t2, t11)

	// Step 47: t11 = x^0x221fc8bf0
	for s := 0; s < 4; s++ {
		t11.Square(t11)
	}

	// Step 48: t11 = x^0x221fc8bf5
	t11.Mul(t6, t11)

	// Step 54: t11 = x^0x887f22fd40
	for s := 0; s < 6; s++ {
		t11.Square(t11)
	}

	// Step 55: t11 = x^0x887f22fd4d
	t11.Mul(t3, t11)

	// Step 63: t11 = x^0x887f22fd4d00
	for s := 0; s < 8; s++ {
		t11.Square(t11)
	}

	// Step 64: t11 = x^0x887f22fd4d1b
	t11.Mul(t8, t11)

	// Step 66: t11 = x^0x221fc8bf5346c
	for s := 0; s < 2; s++ {
		t11.Square(t11)
	}

	// Step 67: t11 = x^0x221fc8bf5346d
	t11.Mul(&x, t11)

	// Step 74: t11 = x^0x110fe45fa9a3680
	for s := 0; s < 7; s++ {
		t11.Square(t11)
	}

	// Step 75: t11 = x^0x110fe45fa9a36bf
	t11.Mul(t2, t11)

	// Step 83: t11 = x^0x110fe45fa9a36bf00
	for s := 0; s < 8; s++ {
		t11.Square(t11)
	}

	// Step 84: t11 = x^0x110fe45fa9a36bf0b
	t11.Mul(t10, t11)

	// Step 86: t11 = x^0x443f917ea68dafc2c
	for s := 0; s < 2; s++ {
		t11.Square(t11)
	}

	// Step 87: t11 = x^0x443f917ea68dafc2d
	t11.Mul(&x, t11)

	// Step 95: t11 = x^0x443f917ea68dafc2d00
	for s := 0; s < 8; s++ {
		t11.Square(t11)
	}

	// Step 96: t10 = x^0x443f917ea68dafc2d0b
	t10.Mul(t10, t11)

	// Step 104: t10 = x^0x443f917ea68dafc2d0b00
	for s := 0; s < 8; s++ {
		t10.Square(t10)
	}

	// Step 105: t10 = x^0x443f917ea68dafc2d0b09
	t10.Mul(t4, t10)

	// Step 113: t10 = x^0x443f917ea68dafc2d0b0900
	for s := 0; s < 8; s++ {
		t10.Square(t10)
	}

	// Step 114: t9 = x^0x443f917ea68dafc2d0b097f
	t9.Mul(t9, t10)

	// Step 119: t9 = x^0x887f22fd4d1b5f85a1612fe0
	for s := 0; s < 5; s++ {
		t9.Square(t9)
	}

	// Step 120: t9 = x^0x887f22fd4d1b5f85a1612fe5
	t9.Mul(t6, t9)

	// Step 128: t9 = x^0x887f22fd4d1b5f85a1612fe500
	for s := 0; s < 8; s++ {
		t9.Square(t9)
	}

	// Step 129: t8 = x^0x887f22fd4d1b5f85a1612fe51b
	t8.Mul(t8, t9)

	// Step 138: t8 = x^0x110fe45fa9a36bf0b42c25fca3600
	for s := 0; s < 9; s++ {
		t8.Square(t8)
	}

	// Step 139: t8 = x^0x110fe45fa9a36bf0b42c25fca360f
	t8.Mul(t0, t8)

	// Step 145: t8 = x^0x443f917ea68dafc2d0b097f28d83c0
	for s := 0; s < 6; s++ {
		t8.Square(t8)
	}

	// Step 146: t8 = x^0x443f917ea68dafc2d0b097f28d83cd
	t8.Mul(t3, t8)

	// Step 151: t8 = x^0x887f22fd4d1b5f85a1612fe51b079a0
	for s := 0; s < 5; s++ {
		t8.Square(t8)
	}

	// Step 152: t8 = x^0x887f22fd4d1b5f85a1612fe51b079a9
	t8.Mul(t4, t8)

	// Step 159: t8 = x^0x443f917ea68dafc2d0b097f28d83cd480
	for s := 0; s < 7; s++ {
		t8.Square(t8)
	}

	// Step 160: t8 = x^0x443f917ea68dafc2d0b097f28d83cd491
	t8.Mul(t5, t8)

	// Step 165: t8 = x^0x887f22fd4d1b5f85a1612fe51b079a9220
	for s := 0; s < 5; s++ {
		t8.Square(t8)
	}

	// Step 166: t7 = x^0x887f22fd4d1b5f85a1612fe51b079a9239
	t7.Mul(t7, t8)

	// Step 169: t7 = x^0x443f917ea68dafc2d0b097f28d83cd491c8
	for s := 0; s < 3; s++ {
		t7.Square(t7)
	}

	// Step 170: t6 = x^0x443f917ea68dafc2d0b097f28d83cd491cd
	t6.Mul(t6, t7)

	// Step 177: t6 = x^0x221fc8bf5346d7e168584bf946c1e6a48e680
	for s := 0; s < 7; s++ {
		t6.Square(t6)
	}

	// Step 178: t6 = x^0x221fc8bf5346d7e168584bf946c1e6a48e68f
	t6.Mul(t0, t6)

	// Step 184: t6 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3c0
	for s := 0; s < 6; s++ {
		t6.Square(t6)
	}

	// Step 185: t6 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf
	t6.Mul(t0, t6)

	// Step 192: t6 = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e780
	for s := 0; s < 7; s++ {
		t6.Square(t6)
	}

	// Step 193: t5 = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e791
	t5.Mul(t5, t6)

	// Step 197: t5 = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e7910
	for s := 0; s < 4; s++ {
		t5.Square(t5)
	}

	// Step 198: t4 = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e7919
	t4.Mul(t4, t5)

	// Step 203: t4 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf2320
	for s := 0; s < 5; s++ {
		t4.Square(t4)
	}

	// Step 204: t3 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf232d
	t3.Mul(t3, t4)

	// Step 211: t3 = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e7919680
	for s := 0; s < 7; s++ {
		t3.Square(t3)
	}

	// Step 212: t2 = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e79196bf
	t2.Mul(t2, t3)

	// Step 219: t2 = x^0x221fc8bf5346d7e168584bf946c1e6a48e68f3c8cb5f80
	for s := 0; s < 7; s++ {
		t2.Square(t2)
	}

	// Step 220: t1 = x^0x221fc8bf5346d7e168584bf946c1e6a48e68f3c8cb5f87
	t1.Mul(t1, t2)

	// Step 226: t1 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf232d7e1c0
	for s := 0; s < 6; s++ {
		t1.Square(t1)
	}

	// Step 227: t0 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf232d7e1cf
	t0.Mul(t0, t1)

	// Step 233: t0 = x^0x221fc8bf5346d7e168584bf946c1e6a48e68f3c8cb5f873c0
	for s := 0; s < 6; s++ {
		t0.Square(t0)
	}

	// Step 234: z = x^0x221fc8bf5346d7e168584bf946c1e6a48e68f3c8cb5f873d7
	z.Mul(z, t0)

	return z
}

// expByLegendreExp is equivalent to z.Exp(x, 221fc8bf5346d7e168584bf946c1e6a48e68f3c8cb5f873d7800000000000000)
//
// uses github.com/mmcloughlin/addchain v0.4.0 to generate a shorter addition chain
func (z *Element) expByLegendreExp(x Element) *Element {
	// addition chain:
	//
	//	_10       = 2*1
	//	_11       = 1 + _10
	//	_101      = _10 + _11
	//	_110      = 1 + _101
	//	_1001     = _11 + _110
	//	_1010     = 1 + _1001
	//	_1011     = 1 + _1010
	//	_1111     = _101 + _1010
	//	_10001    = _10 + _1111
	//	_11011    = _1010 + _10001
	//	_11101    = _10 + _11011
	//	_100011   = _110 + _11101
	//	_101001   = _110 + _100011
	//	_101101   = _1010 + _100011
	//	_101111   = _10 + _101101
	//	_110101   = _110 + _101111
	//	_111001   = _1010 + _101111
	//	_111111   = _110 + _111001
	//	_1111110  = 2*_111111
	//	_1111111  = 1 + _1111110
	//	_10001000 = _1001 + _1111111
	//	i45       = ((_10001000 << 8 + _1111111) << 7 + _10001) << 7
	//	i58       = ((_111111 + i45) << 7 + _101001) << 3 + _101
	//	i77       = ((i58 << 8 + _11011) << 7 + _101111) << 2
	//	i98       = ((_11 + i77) << 10 + _101101) << 8 + _1011
	//	i121      = ((i98 << 8 + _1001) << 8 + _1111111) << 5
	//	i141      = ((_101 + i121) << 8 + _11011) << 9 + _1111
	//	i166      = ((i141 << 8 + _110101) << 6 + _1001) << 9
	//	i179      = ((_111001 + i166) << 3 + _101) << 7 + _1111
	//	i203      = ((i179 << 6 + _1111) << 8 + _100011) << 8
	//	i223      = ((_101101 + i203) << 7 + _111111) << 10 + _111001
	//	return      ((i223 << 5 + _11101) << 5 + _1111) << 59
	//
	// Operations: 248 squares 46 multiplies

	// Allocate Temporaries.
	var (
		t0  = new(Element)
		t1  = new(Element)
		t2  = new(Element)
		t3  = new(Element)
		t4  = new(Element)
		t5  = new(Element)
		t6  = new(Element)
		t7  = new(Element)
		t8  = new(Element)
		t9  = new(Element)
		t10 = new(Element)
		t11 = new(Element)
		t12 = new(Element)
		t13 = new(Element)
		t14 = new(Element)
		t15 = new(Element)
	)

	// var t0,t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11,t12,t13,t14,t15 Element
	// Step 1: t7 = x^0x2
	t7.Square(&x)

	// Step 2: t11 = x^0x3
	t11.Mul(&x, t7)

	// Step 3: t5 = x^0x5
	t5.Mul(t7, t11)

	// Step 4: t2 = x^0x6
	t2.Mul(&x, t5)

	// Step 5: t6 = x^0x9
	t6.Mul(t11, t2)

	// Step 6: t1 = x^0xa
	t1.Mul(&x, t6)

	// Step 7: t10 = x^0xb
	t10.Mul(&x, t1)

	// Step 8: z = x^0xf
	z.Mul(t5, t1)

	// Step 9: t14 = x^0x11
	t14.Mul(t7, z)

	// Step 10: t8 = x^0x1b
	t8.Mul(t1, t14)

	// Step 11: t0 = x^0x1d
	t0.Mul(t7, t8)

	// Step 12: t4 = x^0x23
	t4.Mul(t2, t0)

	// Step 13: t13 = x^0x29
	t13.Mul(t2, t4)

	// Step 14: t3 = x^0x2d
	t3.Mul(t1, t4)

	// Step 15: t12 = x^0x2f
	t12.Mul(t7, t3)

	// Step 16: t7 = x^0x35
	t7.Mul(t2, t12)

	// Step 17: t1 = x^0x39
	t1.Mul(t1, t12)

	// Step 18: t2 = x^0x3f
	t2.Mul(t2, t1)

	// Step 19: t9 = x^0x7e
	t9.Square(t2)

	// Step 20: t9 = x^0x7f
	t9.Mul(&x, t9)

	// Step 21: t15 = x^0x88
	t15.Mul(t6, t9)

	// Step 29: t15 = x^0x8800
	for s := 0; s < 8; s++ {
		t15.Square(t15)
	}

	// Step 30: t15 = x^0x887f
	t15.Mul(t9, t15)

	// Step 37: t15 = x^0x443f80
	for s := 0; s < 7; s++ {
		t15.Square(t15)
	}

	// Step 38: t14 = x^0x443f91
	t14.Mul(t14, t15)

	// Step 45: t14 = x^0x221fc880
	for s := 0; s < 7; s++ {
		t14.Square(t14)
	}

	// Step 46: t14 = x^0x221fc8bf
	t14.Mul(t2, t14)

	// Step 53: t14 = x^0x110fe45f80
	for s := 0; s < 7; s++ {
		t14.Square(t14)
	}

	// Step 54: t13 = x^0x110fe45fa9
	t13.Mul(t13, t14)

	// Step 57: t13 = x^0x887f22fd48
	for s := 0; s < 3; s++ {
		t13.Square(t13)
	}

	// Step 58: t13 = x^0x887f22fd4d
	t13.Mul(t5, t13)

	// Step 66: t13 = x^0x887f22fd4d00
	for s := 0; s < 8; s++ {
		t13.Square(t13)
	}

	// Step 67: t13 = x^0x887f22fd4d1b
	t13.Mul(t8, t13)

	// Step 74: t13 = x^0x443f917ea68d80
	for s := 0; s < 7; s++ {
		t13.Square(t13)
	}

	// Step 75: t12 = x^0x443f917ea68daf
	t12.Mul(t12, t13)

	// Step 77: t12 = x^0x110fe45fa9a36bc
	for s := 0; s < 2; s++ {
		t12.Square(t12)
	}

	// Step 78: t11 = x^0x110fe45fa9a36bf
	t11.Mul(t11, t12)

	// Step 88: t11 = x^0x443f917ea68dafc00
	for s := 0; s < 10; s++ {
		t11.Square(t11)
	}

	// Step 89: t11 = x^0x443f917ea68dafc2d
	t11.Mul(t3, t11)

	// Step 97: t11 = x^0x443f917ea68dafc2d00
	for s := 0; s < 8; s++ {
		t11.Square(t11)
	}

	// Step 98: t10 = x^0x443f917ea68dafc2d0b
	t10.Mul(t10, t11)

	// Step 106: t10 = x^0x443f917ea68dafc2d0b00
	for s := 0; s < 8; s++ {
		t10.Square(t10)
	}

	// Step 107: t10 = x^0x443f917ea68dafc2d0b09
	t10.Mul(t6, t10)

	// Step 115: t10 = x^0x443f917ea68dafc2d0b0900
	for s := 0; s < 8; s++ {
		t10.Square(t10)
	}

	// Step 116: t9 = x^0x443f917ea68dafc2d0b097f
	t9.Mul(t9, t10)

	// Step 121: t9 = x^0x887f22fd4d1b5f85a1612fe0
	for s := 0; s < 5; s++ {
		t9.Square(t9)
	}

	// Step 122: t9 = x^0x887f22fd4d1b5f85a1612fe5
	t9.Mul(t5, t9)

	// Step 130: t9 = x^0x887f22fd4d1b5f85a1612fe500
	for s := 0; s < 8; s++ {
		t9.Square(t9)
	}

	// Step 131: t8 = x^0x887f22fd4d1b5f85a1612fe51b
	t8.Mul(t8, t9)

	// Step 140: t8 = x^0x110fe45fa9a36bf0b42c25fca3600
	for s := 0; s < 9; s++ {
		t8.Square(t8)
	}

	// Step 141: t8 = x^0x110fe45fa9a36bf0b42c25fca360f
	t8.Mul(z, t8)

	// Step 149: t8 = x^0x110fe45fa9a36bf0b42c25fca360f00
	for s := 0; s < 8; s++ {
		t8.Square(t8)
	}

	// Step 150: t7 = x^0x110fe45fa9a36bf0b42c25fca360f35
	t7.Mul(t7, t8)

	// Step 156: t7 = x^0x443f917ea68dafc2d0b097f28d83cd40
	for s := 0; s < 6; s++ {
		t7.Square(t7)
	}

	// Step 157: t6 = x^0x443f917ea68dafc2d0b097f28d83cd49
	t6.Mul(t6, t7)

	// Step 166: t6 = x^0x887f22fd4d1b5f85a1612fe51b079a9200
	for s := 0; s < 9; s++ {
		t6.Square(t6)
	}

	// Step 167: t6 = x^0x887f22fd4d1b5f85a1612fe51b079a9239
	t6.Mul(t1, t6)

	// Step 170: t6 = x^0x443f917ea68dafc2d0b097f28d83cd491c8
	for s := 0; s < 3; s++ {
		t6.Square(t6)
	}

	// Step 171: t5 = x^0x443f917ea68dafc2d0b097f28d83cd491cd
	t5.Mul(t5, t6)

	// Step 178: t5 = x^0x221fc8bf5346d7e168584bf946c1e6a48e680
	for s := 0; s < 7; s++ {
		t5.Square(t5)
	}

	// Step 179: t5 = x^0x221fc8bf5346d7e168584bf946c1e6a48e68f
	t5.Mul(z, t5)

	// Step 185: t5 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3c0
	for s := 0; s < 6; s++ {
		t5.Square(t5)
	}

	// Step 186: t5 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf
	t5.Mul(z, t5)

	// Step 194: t5 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf00
	for s := 0; s < 8; s++ {
		t5.Square(t5)
	}

	// Step 195: t4 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf23
	t4.Mul(t4, t5)

	// Step 203: t4 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf2300
	for s := 0; s < 8; s++ {
		t4.Square(t4)
	}

	// Step 204: t3 = x^0x887f22fd4d1b5f85a1612fe51b079a9239a3cf232d
	t3.Mul(t3, t4)

	// Step 211: t3 = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e7919680
	for s := 0; s < 7; s++ {
		t3.Square(t3)
	}

	// Step 212: t2 = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e79196bf
	t2.Mul(t2, t3)

	// Step 222: t2 = x^0x110fe45fa9a36bf0b42c25fca360f352473479e465afc00
	for s := 0; s < 10; s++ {
		t2.Square(t2)
	}

	// Step 223: t1 = x^0x110fe45fa9a36bf0b42c25fca360f352473479e465afc39
	t1.Mul(t1, t2)

	// Step 228: t1 = x^0x221fc8bf5346d7e168584bf946c1e6a48e68f3c8cb5f8720
	for s := 0; s < 5; s++ {
		t1.Square(t1)
	}

	// Step 229: t0 = x^0x221fc8bf5346d7e168584bf946c1e6a48e68f3c8cb5f873d
	t0.Mul(t0, t1)

	// Step 234: t0 = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e79196bf0e7a0
	for s := 0; s < 5; s++ {
		t0.Square(t0)
	}

	// Step 235: z = x^0x443f917ea68dafc2d0b097f28d83cd491cd1e79196bf0e7af
	z.Mul(z, t0)

	// Step 294: z = x^0x221fc8bf5346d7e168584bf946c1e6a48e68f3c8cb5f873d7800000000000000
	for s := 0; s < 59; s++ {
		z.Square(z)
	}

	return z
}
