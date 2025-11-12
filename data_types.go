package main

import (
	"fmt"
)

func main() {
	// 1. Bit Manipulation (using int)
	// Bits are the fundamental 0s and 1s. We use bitwise operators
	// to manipulate them directly.
	fmt.Println("## 1. Bit Manipulation ##")
	a := 12 // Binary: 0000 1100
	b := 10 // Binary: 0000 1010

	// AND (&): 1 only if both bits are 1
	//   0000 1100 (12)
	// & 0000 1010 (10)
	//   ----------
	//   0000 1000 (8)
	fmt.Printf("Bitwise AND (a & b): %d\n", a&b) // Output: 8

	// OR (|): 1 if at least one bit is 1
	//   0000 1100 (12)
	// | 0000 1010 (10)
	//   ----------
	//   0000 1110 (14)
	fmt.Printf("Bitwise OR (a | b):  %d\n", a|b) // Output: 14

	// XOR (^): 1 only if bits are different
	//   0000 1100 (12)
	// ^ 0000 1010 (10)
	//   ----------
	//   0000 0110 (6)
	fmt.Printf("Bitwise XOR (a ^ b): %d\n", a^b) // Output: 6

	// Left Shift (<<): Shifts bits left, multiplies by 2
	// 0000 1100 (12) << 1  ->  0001 1000 (24)
	fmt.Printf("Left Shift (a << 1): %d\n", a<<1) // Output: 24

	// Right Shift (>>): Shifts bits right, divides by 2
	// 0000 1100 (12) >> 1  ->  0000 0110 (6)
	fmt.Printf("Right Shift (a >> 1): %d\n", a>>1) // Output: 6

	// 2. byte (alias for uint8)
	// A byte is just an 8-bit unsigned integer.
	// It's often used to represent raw data or ASCII characters.
	fmt.Println("\n## 2. byte (uint8) ##")
	var myByte byte = 'A' // 'A' is 65 in the ASCII table
	fmt.Printf("Value: %d, Character: %c, Type: %T\n", myByte, myByte, myByte)

	var myUint8 uint8 = 255 // uint8 range: 0 to 255
	fmt.Printf("uint8 Value: %d, Type: %T\n", myUint8, myUint8)

	// Overflow example: 255 + 1 wraps back to 0
	// myUint8 = myUint8 + 1 // This would cause an overflow error at compile time
	// fmt.Println(myUint8)

	// 3. Integer Types (int, uint, uint16)
	// Integers are whole numbers.
	// 'int' vs 'uint': 'int' can be negative (signed), 'uint' cannot (unsigned).
	// '16': Specifies the bit-size.
	fmt.Println("\n## 3. Integer Types ##")

	// int: Default signed integer. Size (32 or 64 bit) depends on your system.
	var myInt int = -100
	fmt.Printf("int Value: %d, Type: %T\n", myInt, myInt)

	// uint: Default unsigned integer. Size (32 or 64 bit) depends on your system.
	var myUint uint = 100
	fmt.Printf("uint Value: %d, Type: %T\n", myUint, myUint)

	// uint16: A 16-bit unsigned integer. Range: 0 to 65,535
	var myUint16 uint16 = 65000
	fmt.Printf("uint16 Value: %d, Type: %T\n", myUint16, myUint16)

	// Type Conversion: Go requires explicit conversion.
	// You can't add an int and a uint16 directly.
	// sum := myInt + myUint16 // This is a compiler error!

	// Correct way:
	sum := myInt + int(myUint16)
	fmt.Printf("Explicit Conversion Sum: %d\n", sum)

	// 4. Floating Point Types (float64, float32)
	// Used for numbers with decimal points.
	// 'double' in other languages is 'float64' in Go.
	fmt.Println("\n## 4. Floating Point Types ##")

	// float64 (double precision): Default and most common float type.
	var myDouble float64 = 123.456789012345
	fmt.Printf("float64 (double): %f, Type: %T\n", myDouble, myDouble)

	// float32 (single precision): Less precise, uses half the memory.
	var myFloat float32 = 9.87654321
	fmt.Printf("float32:          %f, Type: %T\n", myFloat, myFloat)

	// Note the precision loss when printing float32:
	fmt.Printf("float32 (less precise): %.10f\n", myFloat)
	// Note the better precision with float64:
	fmt.Printf("float64 (more precise): %.10f\n", myDouble)

	// Arithmetic (Note: 0.1 + 0.2 is not exactly 0.3)
	f1 := 0.1
	f2 := 0.2
	fmt.Printf("Float arithmetic (0.1 + 0.2): %.60f\n", f1+f2)
}
