package main

import "fmt"

func main() {
	var a int8 = 15                      // 2^8 = 256, range: -128 to 127
	var b int16 = 30000                  // 2^16 = 65536, range: -32768 to 32767
	var c int32 = 2000000000             // 2^32 = 4294967296, range: -2147483648 to 2147483647
	var d int64 = 9000000000000000000    // 2^64 = 18446744073709551616, range: -9223372036854775808 to 9223372036854775807
	var f float32 = 20.5                 // 2^32 = 4294967296, range: -3.402823466e+38 to 3.402823466e+38
	var e float64 = 1000000000.123456789 // 2^64 = 18446744073709551616, range: -1.7976931348623157e+308 to 1.7976931348623157e+308
	love := "♡"
	fmt.Printf("Integer: %d, Float: %.2f, unicode: %c", a, b, love)
	fmt.Printf("\nInt8: %d, Int16: %d, Int32: %d, Int64: %d, Float32: %.2f, Float64: %.9f\n", a, b, c, d, f, e)
}
