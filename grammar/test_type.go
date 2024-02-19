package grammar

import (
	"fmt"
	"math"
	"unsafe"
)

func TestDataType() {
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)

	fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)
	fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16)
	fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32)
	fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))

	var f32 float32
	var f64 float64

	fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64)

	var ui uint
	ui = uint(math.MaxUint64) // 再+1会导致overflows错误
	fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, ui)

	var iMax, iMin int
	iMax = math.MaxInt64 // 再+1会导致overflows错误
	iMin = math.MinInt64 // 再-1会导致overflows错误

	fmt.Printf("%T %dB %v~%v\n", iMax, unsafe.Sizeof(iMax), iMin, iMax)
}

func TestBoolType(age int) {
	if age >= 18 {
		fmt.Println("你已经成年")
	} else {
		fmt.Println("你还未成年")
	}
}
