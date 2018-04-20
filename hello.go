package main
import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Variables
var numberOfTimes int

var var1, var2, var3 int

// variable numoftimes type integer set equal to 3
var numberOfTimes int = 3

var var1, var2, var3 int = 1, 2, 3


vname1, vname2, vname3 := v1, v2, v3
// := only use inside functions -- for global variables we should only 
// use var


// Blank variable
// "_" or _ is a blank variable and it is used to ignore a value
_, remainder := divide(10, 3)
// function divide which accepts two arguments, arg1 / arg2 
// IF we only want the remainder, we ignore the quotient(arg1) using 
// _



// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Example of program that will not run 

package main

func main(){
	var i int
	// if nothing is assigned to variable i then there will be a compilation error
}


// CONSTANTS
// constants are defined during compile time, cannot be changed during 
// runtime
const constantName = value
const Pi = float32 = 3.1415926
const i = 10000
const MaxThread = 10
const prefix = "astaxie_" // Elementary types


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// BOOLEANS

// true or false
// Go does not allow to convert boolean into number

var isActive bool // global variable

var enabled, disabled = true, false

func test(){
	var available bool // local variable
	valid := false // brief statement of variable
	available = true // assigned value to variable
}


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// INTEGERS

// Integers include signed and unsigned, int and uint
// Both are same length, the specific length depends on operating system
// 32-bit in 32-bit OS's and 64-bit in 64-bit OS's 

Specific length integers: {
rune, int8, int16, int64, byte, uint8, uint16, uint32, uint64
rune is = to int32

byte is = to uint8
}

// GO DOES NOT ALLOW ASSIGNING VARIABLES BETWEEN DATA TYPES
var a int8
var b int32

c := a + b
// THIS WILL CAUSE compilation errors


// FRACTIONS
float32, float64
// no type called float is present.


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// COMPLEX NUMBERS

// 64-bit real, 64-bit imaginary part
complex128 // default type

// There exists a smaller complex number 

// 32-bit real, 32-bit imaginary
complex64

// The form for these complex numbers is RE + IMi
// real + imaginary


// Example: 

var c complex64 = 5+5i
// output: (5+5i)
fmt.Printf("Value is: %v",c)
// interpolation - c is substituded for v


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// STRINGS

// Strings are represented in double quotes "" or `` back ticks

var frenchHello string // defining a string
var emptyString string = "" // empty string

func test(){
	no, yes, maybe := "no", "yes", "maybe"
	japaneseHello := "Ohaiu"
	frenchHello := "Bonjour"
}

// String objects do not allow value change. Will receive errors when 
// compiling 

var s string = "hello"
s[0] = 'c' // does not work

// to change a string, a new one has to be created

s := "hello"
c := []byte(s) // converts string to byte size

c[0] = 'c'

s2 := string(c) // converts back to string type

fmt.Printf("%s\n", s2)


// + operator can be used to concatenate two strings

s := "hello,"
m := " world"
a := s + m

fmt.Printf("%s\n", a)

s := "hello"
s = "c" + s[1:] // returns substring from index 1 till the end

fmt.Printf("%s\n", s)


// multiple line string

m := `hello 
world`

// backticks will not escape characters in a string

// ERROR TYPES
// There is one error type for the purpose of dealing with errors in the package 
// called errors

// Go requires us to explicitly handle our errors or ignore them
// Typical error handling forms a 
if err != nil

err := errors.New("emit macho dwarf: elf header corrupted")

if err != nil {
	fmt.Printf(err)
}





// Variables and Data structures
i := 1234 // type: int

j := int32(1) // type: int32

f := float32(3.14) // type: float32

bytes := [5]byte{'h', 'e', 'l', 'l', 'o'} // type: [5]byte

primes := [4]int{2, 3, 5, 7} // type: [4]int


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// define by groups 

// if you want to define multiple constants 
// variables or import packages, you can use the group form

import 'fmt'
import 'os'

const i = 100
const pi = 3.1415
const prefix = "Go_"

var i int 
var pi float32
var prefix string


// if you have imports from both standard library and custom
// imports, the standard library imports go first, followed 
// by others

import(
	"fmt", 
	"os",
	"github.com/thewhitetulip/Tasks/views"
)

// value of first const will be 0 unless it is assigned to iota
// if there are not assigned values for the elements except the last one, 
// all constants will have the same value as the last ones

const(
	i = 100
	pi = 3.1415
	prefix = "Go_"
)

var(
	i int
	pi float32
	prefix string
)


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// iota enumerate

// Go has a keyword called iota, this keyword is to make enum, 
// it begins with 0, increased by 1.

const(
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w 
	// if there is no expression after the constants name, the last 
	// expression is used
	// last line reads as w = iota implicitly 
	// w == 3
)


const v = iota // once iota meets 'const', it resets to 0, v == 0

const(
	e, f, g = iota, iota, iota
	// e==o, f==0, g==0 values of iota are the same in one line
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// SOME GO RULES BY DEFAULT

// Ant variable that begins with a capital letter means it will be exported
// private otherwise.

// Same rule applies for functions and constants, no public or private keyword
// exists in Go


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// array, slice, map

 // Arrays

 var arr [n]type

 // in [n]type, n is the length of the array, type is the type of its 
// elements. Use [] to get or set array values

var arr [10]int // array of type [10]int

arr[0] = 42 // array is 0-based
arr[1] = 13 // assign value to element

fmt.Printf("The first element is %d\n", arr[0])
// gets element value, returns 42

fmt.Printf("The last element is %d\n", arr[9])

// returns value of the tenth value in the array

// its possible to use := when defining arrays

a := [3]int{1, 2, 3} // define an int array with 3 elements

b := [10]int{1, 2, 3} // defines array with 10 elements
// first three are assigned 1, 2, 3

// The rest are a default value of 0

c := [...]int{4, 5, 6} // use `...` to replace the length parameter and Go will 
// calculate it for you 




// You may want to use arrays as arrays' elements. let's see  how to do this.
// define a two-dimensional array with 2 elements, and each element
// has 4 elements 
doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// SLICE

// The main disadvantage of arrays is that we need to know the size prehand. At times this isn't 
// possible.

// Instead, we can use a `dynamic array`. This is called a 'slice' in Go

// Slice is not really a dynamic array. It's a reference type. 
// slice points to a an underlying array whose declaration is similar to array,
// but doesnt need length

// Just like defining an array, except we exclude the length.

var fslice []int

// Then we define a slice, and initialize its data

slice := []byte {'a', 'b', 'c', 'd'}

// Slice can redefine existing slice or array.slice uses array[i:j]
// where i is the start index and j is end index, but notice that array[j]
// will not be sliced since the length of the slice is j - i

// define an array with 10 elements whose types are bytes

var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

// define two slices with type []byte

var a, b []byte

// 'a' points to elements from 3rd to 5th in array ar

a = ar[2:5]

// 'a' now has elements with index [2], [3], [4]
// which are = 'c', 'd', 'e'

// 'b' is another slice of array ar

b = ar[3:5]
// b is = to index [3], [4] 
// returns 'd', 'e'



// define an array 

var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

// define two slices 

var aSlice, bSlice []byte

// convenient operations

aSlice = array[:3] // aSlice = array[0:3] => a, b, c
bSlice = array[5:] // bSlice = array[5:10] => f, g, h, i, j

aSlice = array[:] // aSlice = array[0:10] has all element




// slice from slice 

aSlice = array[3:7] // aSlice has elements d, e, f, g, h

// bSlices of aSlice //
bSlice = aSlice[1:3] // aSlice is still d, e, f ,g ,h
// bSlice will equal aSlice[1], aSlice[2]
// bSlice == e, f

bSlice = aSlice[:3] // aSlice is still d, e, f, g, h
// bSlice will equal aSlice[0], aSlice[1], aSlice[2]
// bSlice == d, e, f

bSlice = aSlice[0:5] // bSlice == d, e, f, g, h

bSlice = aSlice[:] // aSlice is still d, e, f, g, h
//bSlice == d, e, f, g, h



// slice is a reference type, any changes will affect other variables
// pointing to the same slice or array.

// In the case of aSlice and bSlice above, if you change the value of an element 
// in aSlice, bSlice will be changes as well


// Slice is like a struct by definition and it contains 3 parts:
// - Slice is a pointer that points to where slice starts
// - The length of lice
// - Capacity, the length from start to index to end index of slice

// EXample

Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

Slice_a := Array_a[2:5]

// Slice_a == c, d, e


// Built in functions for slice

// - len -- gets the length of slice
// - cap -- gets the maximum length of slice 
// - append -- appends one or more elements to slice

// - copy -- copies elements from one slice to the other, and returns the number
// of elements that were copied

// !!!!!! using 'append' will change the array that slice points to, and affect other
// slices that point to the same array 

// ALSO if there is not enough length for the slice -- capacity - length == 0
// appending will return a new array for this slice
// Other slices pointing to old array will not be affected



// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


// MAP

// Map is a key value pair, like a dictionary in Python.
// Use the form map[keyType]valueType to define it

// The 'set' and 'get' values in map are similar to slice,
// however the index in slice can only be of type 'int' 
// while map can use more types like: string, int etc. 
// With map you are able to make comparisons with == and !=, in slice 
// you cannot compare values

// Example: 

// Use string as the key type, int as a the value type. Make 'initialize' it

var numbers map[string] int

// OR 

numbers := make(map[string]int)

numbers["one"] = 1
numbers["ten"] = 10
numbers["three"] = 3

fmt.Printf("The third number is: ", numbers["three"])
// The third number is 3

// Some things to be careful about when using map:

// - map is disorderly -- Everytime you use map you will get different results 
		// It is impossible retrieve values by index. You have to use key.
// - map does not have a fixed length -- It is a reference type just like slice.
// - len works for map. It returns how many keys map has.
// - It is easy to change the value through map. Use numbers["one"]=11 to change
		// the value of key one to 11.

// You can use form key:val to intialize map's













