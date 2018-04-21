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
// prints out => The third number is 3

// Some things to be careful about when using map:

// - map is disorderly -- Everytime you use map you will get different results 
		// It is impossible retrieve values by index. You have to use key.
// - map does not have a fixed length -- It is a reference type just like slice.
// - len works for map. It returns how many keys map has.
// - It is easy to change the value through map. Use numbers["one"]=11 to change
		// the value of key one to 11.

// You can use form key:val to intialize map's values
// Built-in methods to check if keys exist: 



// Use 'delete' to delete an element in map.
// First
// initialze a map

rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2}

// map has two return values, the second return value will return false if the 
// key does not exist. Otherwise it returns true.

csharpRating, ok := rating["C#"]

if ok {
	fmt.Println("C# is in the map and its rating is ", csharpRating)
} else {
	fmt.Println("We have no rating associated with C# in the map")
}

delete(rating, "C") // deletes the element with key 'C'

// REMEMBER - Map is a reference type. Any changes to the data affects any pointers
// i.e slices, maps pointing to the same data.

// One more example:

m := make(map[string]string)

m["Hello"] = "Bonjour"
m1 := m
m1["Hello"] = "Salut" // m1 changes the value of "Hello" to "Salut"

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// make, new

// make performs memory allocation for built-in models, such as map, slicem, 
// and channel.

//"while new" is for types' memory allocation.

// "new(T)" allocates zero-value to type T's  memory, returns its memory address,
// which is the value of type *T. 

// It then returns a pointer which points to type T's zero-value.

// "new" returns pointers


// the built it function 
make(T, args) 
// has different purposes than new(T)

// "make" can be used for slice, map, and channel, and returns a type T with an 
// initial value. The reason for doing this is because the underlying data of these three
// types must be initialized before they point to them. 

// Example: a "slice" contains a pointer that points to the underlying array 
//, length and capacity
// Before these are initialized, "slice" is nil, so for "slice", "map", and "channel"
make // initializes their underlying data and assigns some suitable values.

make // returns non-zero values.


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// RECAP

int  	0 
int8 	0
int32  	0
int64	0
uint 	0x0
rune  	0 // actual type is int32
byte 	0x0 // actual type is uint8
float32 0 // length is 4 byte
float64 0 // length is 8 byte
bool    false
string  ""	


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// CONTROL STATEMENTS AND FUNCTIONS 

// CONTROL STATEMENT

if 

// "if" does not need parenthesis in Go. 
if x > 10 {
	// When x is greater than 10 the program enters this block
	fmt.Println("x is greater than 10")
} else {
	// when x is less than or equal to 10
	fmt.Println("x is less than or equal to 10")
}


// else ifs
if integer == 3 {
	fmt.Println("The int is equal to 3")
} else if integer < 3 {
	fmt.Println("The int is less than 3")
} else {
	fmt.Println("The int is greater than 3")
}


// GOTO
// Go has a goto keyword, but take caution when using it. 
goto // reroutes the control flow to a previously defined label within the body of the same code block

func myFunc(){
	i := 0 
Here: // label ends with ":"
	fmt.Println(i)
	i++
	goto Here // jumps to label 'Here'
}


// label name is case sensitive 



// Go does not have a while, nor a do while, only a 'for', which happens to be 
// the most powerful control logic.

// It is possible to read data in loops and iterative operations, just like while.
// Like 'if', for does not need parenthesis.

for expression1; expression2; expression3 {
	// stuff here
}

// EXAMPLES

package main 
import "fmt"

func main(){
	sum := 0
	for index:=0; index < 10; index++{
		sum += index
	}

	fmt.Println("Sum is equal to ", sum)
}

// prints out: Sum is equal to 45

// ANOTHER ONE
sum := 1
for ; sum < 1000; {
	sum += sum
}

for {
	// this is an infinite loop
}

// Using 'for' like a 'while' loop

sum := 1 

for sum < 1000 {
	sum += sum
}



// BREAKS 

// Break jumps out of the loop. If you have nested loops, use break along with labels.
// 'continue' skips the current loop and starts the next one.

for index := 10; index > 0; index-- {
	if index == 5 {
		break // or continue 
	}

	fmt.Println(index)
}

// break prints out :  109876
// continue prints out: 1098764321 // skips over the number 5

for // can read data from slice and map when it is used together with range.

for k, v:=range map {
	fmt.Println("map's key:", k)
	fmt.Println("map's value:", v)
}

// Because go supports multi-value returns and gives compile errors when you don't use 
// values taht were defined, as discussed earlier, _ is used to discard certain return values


for _. v := range map {
	fmt.Println("map's value:", v)
}




// Switch

// Use switch to avoid long if-else statements uwu 

switch sExpression {

case expression1:
	some stuff
case expression2:
	some other stuff
case expression3:
	some other other stuff
default: 
	other code
}

// EXAMPLE

i := 10 
switch i {
case 1: 
	fmt.Println("i is equal to 1")	
case 2, 3, 4: 
	fmt.Println("i is equal to 2, 3, or 4")
case 10: 
	fmt.Println("i is equal to 10")
default: 
	fmt.Println("D A W G all I know is that i is an integer")
}

// Cases can have more than one value separated by a comma. By default switch executes 
// only the matching case, however, we can make switch execute th matching case and all cases 
// below it using the 'fallthrough' statement.

// EXAMPLE fallthrough statement

integer := 6
switch integer {
case 4: 
	fmt.Println("Integer <= 4")
case 5: 
	fmt.Println("Integer <= 5")
case 6: 
	fmt.Println("Integer <= 6")
case 7: 
	fmt.Println("Integer <= 7")
case 8: 
	fmt.Println("Integer <= 8")
default: 
	fmt.Println("default case")
}

// The program prints the following: 
// integer <= 6 integer <= 7 integer <= 8 default case

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


// Functions +++ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func // keyword is used to define a function

func funcName(input1 type1, input2 type2)(ouput1 type1, output2 type2){
	// function body
	// multi-value return
	return value1, return value2
}

// - Functions can have zero, one or more arguments. The argument type comes after 
//   the argument name and arguments are separated by ','

// - Functions can return multiple values

// - There are two return values named output1 and output2, you can omit their names and 
//   use their type only

// - If there is only noe return value and you omitted the name, you do not 
//   need brackets for the return values

// - If the function does not have return valuesm you can omit the return parameters
//   altogether

// - If the function has return values, you have to use the 'return' statement soewhere
//   in the body of the function


// -- A simple program that calculates maximum value --

package main
import 'fmt'

// returns the greater value between a and b

func max(a, b int) int {
	if a > b {
		return a
	}
	// else 
	return b
}


func main(){
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) // calls function max(x, y)
	max_xz := max(x, z) // calls function max(x, z)

	// %d interpolates in the variables x, y, max_xy etc.
	fmt.Println("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Println("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Println("max(%d, %d) = %d\n", y, z, max(y,z)) // calls the function
}



// In a function call, if two or more arguments have the same data type, then we
// can put the data type only after the last argument.

func max(a, b int, c, d string): // there are four arguements in this line
// a,b:integer, c,d:string





// MULTI-VALUE RETURN

package main 
import "fmt"

// return results of A + B and A * B

func SumandProduct(A, B int) (int, int){
	return A + B, A * B
}

func main(){
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumandProduct(x, y)

	fmt.Println("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Println("%d * %d = %d\n", x, y, xTIMESy)
}

// SumandProduct will return two values wothout names
// Go allows us to have named return arguments. By using name arguments, the respective 
// variables are returned automatically, we would only need to use 'return' to do so.

// NOTE: If functions are going to be used outside of current programs, it is better 
// to explicitly write return statements.

func SumandProduct(A, B int) (add int, multiplied int) {
	add = A + B
	multiplied = A * B
	return
}

// Since the return arguments are named, the function automatically returns them

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// VARIADIC ARGUMENTS TO FUNCTIONS

// When we do not know how many arguments can be passed, we use 'variadic arguments' 
// instead

func myFunc(arg ...int){

}

// arg ...int tells Go that this is a function that has variable arguments. Notice that 
// these arguments are type int 
// In the body of a function, the arg becomes a slice of int.

for _, n := range arg {
	fmt.Println("Adn the number is: %d", n)
}



// PASS BY VALUE AND POINTERS

// Arguments are passed by value to the functions, the argument change inside the function
// does not affect the arguments used to call the function.

package main 
import "fmt"

// function that adds 1 to a

func add1(a int) int{
	a = a + 1 // changes value of a 
	return a // returns new value of a
}

func main(){
	x := 3

	fmt.Println("x = ", x) // prints out x = 3

	x1 := add1(x) // calls the function add1(x)

	fmt.Println("x+1 = ", x1) // prints out "x+1 = 4"
	fmt.Println("x = ", x) // prints out "x = 3"
}


// Whats goin on?
// Original value of x does not changes as we are passing x as a value
// The function add1 creates a copy of x. 
// When we want to modify the arguments value, we use a pass by reference using pointers

// Remember that a variable is nothing more than a pointer to a location is memory. 
// Each variable has a unique memory address. If we want to change the value of a variable, we must change its memory address
// Therefore, the function add1 has to know the memory address of x in order to change its value. 

// Below we pass &x to the function and change the argument's type to the pointer type*int. Be aware that we pass a copy of the pointer, not the copy of value.


package main 
import "fmt"
// Same function as before, adds 1 to a

func add1(a *int) int{
	*a = *a+1 // changed the value of a
	return *a // returns new value of a
}

func main(){
	x := 3 
	fmt.Println("x = ", x) // should print "x = 3"
	x1 := add1(&x) // calls function add1(&x) - passes memory address of x

	fmt.Println("x+1 = ", x1) // should print out "x+1 = 4"
	fmt.Println("x = ", x) // prints out "x = 4"
}


// -- Advantage of pointers -- 

// - Allows us to use more functions to operate on one varible
// - Low cost by passing memory address(8 bytes), copy is not an efficient way, both in terms of time and space, to pass variables.
// - string, slice and map are reference types, so they use pointers when passsing to functions by default. 
// **BIG NOTE**: IF the length of a slice needs to be changed, pointers have to be passed explicitly 

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// defer

// 'Defer' postpones the execution of a function till the calling function has finished executing. You can have many 
// 'defer' statements in one function; HOWEVER ; When the program reaches the end they will be executed in reverse order.
// As an example, in the case where the program opens some resource files, these files would have to be closed before the function can return with errors. 
// Some exaples below:









func ReadWrite() bool {
	file.Open("file")
	// perform some work

	if failureX {
		file.Close()
		return false
	}

	if failureY {
		file.Close()
		return false
	}

	file.Close()
	return true
}

// Several lines of repeated code as show above can be refactored with 'defer', which not only makes it cleaner, but easier to understand

// Same code, refactored

func ReadWrite() bool {
	file.Open("file")

	defer file.Close()
	if failureX {
		return false
	}

	if failureY {
		return false
	}
	
	return true
}

// If there are more than one defers, they will execute by reverse order. The following example will print 4 3 2 1 0

for i := 0; i < 5; i++ {
	defer fmt.Printf("%d", i)
}

// FUNCTIONS AS VALUES AND TYPES

// Functions are also variables in Go, we can use 'type' to define them. Functions that have the same signature can be seen as the same type.

type typeName func(input1 inputType1, input2 inputType2, ...) (result1 resultType1...)

// This makes Go a functional language as functions are a first class citizen.

package main
import "fmt"

type testInt func(int) bool // define a function type of varible = boolean

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// passing the function `f` as an arument to another function 

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value){
			result = append(result, value)
		}
	}
	return result
}

func main(){
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // uses functions as values
	fmt.Println("Odd elements of slices are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even the elements of slices are: ", even)
}


// It is useful to have interfaces. testInt serves as a variable that has a function as type and the returned 
// values and arguments of 'filter' are the same as those of testInt.
// This means we can have complex logic while maintaining flexibility 

// Panic and Recover

// Go does not have try-catch structure such as java.. Instead of throwing exceptions, Go uses 'panic' and 'recover' to deal with errors.
// However, you should not use 'panic' as much.

// 'panic' is a built-in function to break the normal flow of programs and get into panic status. When a function 'F' calls 'panic', 'F' will not conintue 
// to execute but its 'defer' functions will continue to execute. 
// Afterwards, F goes back to break point which caused the panic status. 
// The program will not terminate until all of these functions reurn with panic to the first level of that goroutine. 'panic' can 
// be produced by calling 'panic' in the program, and some errors also cause 'panic' like array access out of bounds error


// 'recover' is a built-in function to recover 'goroutines' from the panic status. Calling recover in defer functions is useful because normal functions
// will not be executed when the program is in the 'panic' status.

// It catches 'panic' values if the program is in the panic status, and it gets nil if the program is not in panic status.



// The following example shows how to utilize 'panic'
var user = os.Getevn("USER")

func init(){
	if user == "" {
		panic("no value for $USER")
	}
}


// The following example shows how to check 'panic'

func throwsPanic(f func())(b bool) {
	defer func(){
		if x := recover(); x != nil {
			b = true
		}
	}
	f() // if f causes panic it will recover return
}




// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// MAIN AND INIT FUNCTIONS

// Go has two retention which are called main and init.
// - Init can be used in all packages
// - Main can only be used in the main package
// Both functions cannot have arguments or return values
// Although many init functions can be written in one package, only write one init function per package.

// Go programs automatically call init() and main() functions, so there is not need to call them.
// For every package, the init() function is optional, but package main() has one and ONLY one main() function

/*
	Programs utilize and begin execution from the main package.
	IF the main package imports other packages, they will be imported in the compile time. 

	if one package is imported many times, it will be only compiled once.

	After importing packages, programs will initialze the constants and variables within the imported packages, then execute the init() function it it 
	exists.

*/


// IMPORT

// import is very often used in Go programs

import(
	"fmt"
)

// Methods of fmt are called as follows: 

fmt.Println("hello world")

// fmt is from Go standard library, is is located within $GOROOT/pkg.
// Go supports third-party packages in two ways

/*
	1. Relative path import "./model" - load package in the same directory, NOT RECOMMENDED

	2. Absolute path import "shorturl/model" - load package in path "$GOPATH/pkg/shorturl/model"
*/ 

/*
	There are some special operators when we import packages, and beginners are always confused by these operators(LOL not me 😹👌)

	- Dot operator 
	  -  A common way to import packages
*/ 


import(
	. "fmt"
)

 
// The dot operator mmeans you can omit the package name when you call functions inside of that package. 
// Now => 
	fmt.Printf("Hello world") 
	// becomes 
	Printf("Hello world")

// Alias Operation
	// changes the name of the package that was imported when we call functions that belong to that package
import(
	f "fmt"
)
 
// Now 
fmt.Printf("Hello world")
// becomes
f.Printf("Hello world")

// _operator

// Difficult operator to understand but here we got !

import(
	"database/sql"
	_ "github.com/ziutek/mysql/godrv"
)
	
/*
	The _ operator actually means we just want to import that package and execute its init function, not sure if we want to use the functions
	belonging to that package
*/ 


// STRUCT

// Basics of Struct 

/*
	Struct can be used to define custom data types in Go. Often times, we cannot handle the real world information using the standard 
	data types which come with some languages.

	While it is not impossible. it is highly inefficient.

	Example: 

		- in an eCommerce application, we have the ShoppingCart in which we put products for checkout.
*/ 

type Product struct {
	name 			string
	itemID 			int
	cost 			float32
	isAvailable 	bool
	inventoryLeft 	int
}

// There are a lot of attrubutes of Product, its name, the ID used internally, the cost, number of products in stock, etc.
/* 
	'name' -- is a string used to store a product's name
	'itemID' -- is an int used to store for reference. 
	'cost' -- is a float 32 of the item
	'isAvailable' -- is a 'bool' which is true if the item is in stock, false, otherwise.
	'inventoryLeft' -- is an int of the number of products left in stock 
*/ 

// Initializing

// Struct example

// 1. Define goBook as a Product Type
	var goBook Product

// 2. Assign "Webapps in Go" to the field 'name' of goBook
	goBook.name = "Webapps in Go"

// 3. Assign 10025 to field 'itemID' of goBook
	goBook.itemID = 10025

// 4. Access field 'name' of goBook
	fmt.Printf("The product's name is %s\n", goBook.name)




// There are three more ways to define a struct 
	// - Assign initial values by order 

goBook := Product("Webapps in Go", 10025)

	// - Use the format 'field:value' to initialize the struct without order

goBook := Product{name:"Webapps in Go", item:10025, cost:100} 

	// - name an annonymous struct, then initialize it

p := struct{name string; age init}{"Amy", 18}

// Full example below

// file code/Struct/Book/struct.go

package main

import "fmt"

// Product will denote a physical *object*

// which we will sell online to make some currency boi 😹👌👌👌

type Product struct {
	name 			string
	itemID			int 
	cost 			float32
	isAvailable		bool
	inventoryLeft	int
}

func main(){
	var goBook Product

	// we initialize 

	goBook.name, goBook.itemID, goBook.isAvailable, goBook.inventoryLeft = "Webapps in Go", 10025, true, 25

	// initialize four values by format "field:value"
	pythonBook := Product{itemID: 10026, name: "Learn Python", inventoryLeft: 0, isAvailable: false}

	// initialize all five values in order 
	rubyBook := Product{"Learn Ruby", 10043, 100, true, 12}

	if goBook.isAvailable{
		fmt.Printf("%d copies of %s are available\n", goBook.inventoryLeft, goBook.name)
	}

	if pythonBook.isAvailable{
		fmt.Printf("%d copies of %s are available\n", pythonBook.inventoryLeft, pythonBook.name)
	}

	if rubyBook.isAvailable{
		fmt.Printf("%d copies of %s are available\n", rubyBook.inventoryLeft, rubyBook.name)
	}
}


// Embedded fields in struct 

// Earlier we saw how to define a Struct with field names and type.
// Embedded fields can be thoght of as 'subclass' and 'superclass' in object oriented programming

// When the embedded field is a struct, all the fields in that struct will implicitly be in the fields in the struct in which it has been embedded 

// Example: 

// file: code/Struct/Human/human.go

package main
import  "fmt"

type Human struct{
	name 	string
	age 	int
	weight 	int
}


type Student struct {
	Human // embedded field, it means Student struct - kind of like inheritance from a class
	specialty string // includes all fields that human has
}

func main(){
 	// intializes a student
 	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
 	// specialty string is == "Computer Science"
// in ruby => mark = Student.new("Mark", 25, 120, "Computer Science")



 	// Accessing fields from mark

 	fmt.Println("His name is ", mark.name)
 	fmt.Println("His age is ", mark.age)
 	fmt.Println("His weight is ", mark.weight)
 	fmt.Println("His specialty is ", mark.specialty)

 	// Modifying fields values

 	mark.specialty = "AI" // changes specialty

 	mark.age = 46 // marks new age is 46

 	mark.weight += 60 // mark gained 60 lbs
}




// All types in Go can be used as embedded fields 

// file:code/Struct/Skills/skills.go

import "fmt"

type Skills []string

type Human struct {
	name 		string
	age 		int
	weight 		int
}

type Student struct {
	Human // struct as embedded field 
	Skills // string slice as embedded field 
	int // built-in type as embedded field
	specialty string
}

func main(){
	// initialize student Jane

	jane := Student{Human: Human{"Jane",  35, 100}, specialty: "Biology"}

	// accessing fields 

	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her specialty is ", jane.specialty)

	// modifying skills

	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She has acquired two new skills ") 
	jane.Skills = append(jane.Skills, "physics", "golang")

}



// ANOTHER ONE

package main

import "fmt"

type Human struct {
	name 		string
	age 		int
	phone 		string
}

type Employee struct {
	Human // embedded field human -- inheritance from human, Employe is also a human( ie has work phone, age , name)
	specialty 	string 	
	phone 		string // phone is in employee
}

func main(){
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}

	fmt.Println("Bob's work phone is: ", Bob.phone)

	// access phone in human
	fmt.Println("Bob's personal phone is: ", Bob.human.phone)
}



// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// OBJECT-ORIENTED WITH GO

/*
	OBject oriented languages allow programmers to declare a function inside the class definition
	Go does not allow us to do that, instead we have to declare a method on a struct via a special syntax.
*/ 


// methods
	// We defined a "rectangle" struct and we want to declare its area. Normally we would create a function, pass the struct's 
	// instance and calculate the area.

package main

import "fmt"

type Rectangle struct {
	width, height float32
}


func area(r Rectangle) float32 {
	return r.width*r.height
}

func main(){
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))
}

/* 
	A method is affiliated with the type. It has the same syntax as functions do except for an additional parameter after the func keyword
	 called the receiver, which is the main body of that method.

	 Using the same example, Rectangle.area() belongs directly to rectangle, instead of as a peripheral function. More specifically,
	 length, width, and area() all belong to rectangle.


	 Quote from Rob Pike: 
	 	"A method is a function with an implicit first argument, called a receiver."
*/


// Syntax of a method 

func (r ReceiverType) funcName(parameters)(results)
// r is the argument block |r|

// Above example using a method instead

package main

import(
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float32
}

type Circle struct {
	radius float32
}

func (r Rectangle) area() float64 {
	return r.width*r.height
} 

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main(){
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

// Note for using methods: 

/* 
	- If the name of methods are the same but they don't share the same receivers, they are not the same
	- Methods are able to access fields within receivers
	- Use . to call a method in the struct, the same way fields are called
*/ 

// CUSTOM DATA TYPES

// Use the following format to define a custom data type

	type typeName typeLiteral

// Examples

type ages int

type money float32

type months map[string]int

m := months {
	"January":31,
	"February":28,
	"December":31
}

// full example

package main
import(
	"fmt"
)

const(
	WHITE = iota
	BLACK 
	BLUE 
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box // slice of boxes

func (b Box) Volume() float64 {
	return b.widht * b.height 
}

func (b *Box) SetColor(c Color){
	b.color = c
}

func (b1 BoxList) BiggestColor() Color{
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
 	}

 	return k
}

func (b1 BoxList) PaintItBlack(){
	for i, _ := range b1 {
		b1[i].SetColor(BLACK)
	}
}


func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main(){
	boxes := BoxList {
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK}, 
		Box{10, 10, 1, BLUE},
        Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is ", boxes[0].Volume(), "cm")
	fmt.Println("The color of the last one is ", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is ", boxes.BiggestColor().String())
	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is ", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is ", boxes.BiggestColor().String())
}

// We define some constants and customized types

/*
	- Use 'Color' as alias of byte
	- Define a struct 'Box' which has fields height, width, length and color.
	- Define a struct BoxList which has box as its field
*/ 

// Then we defined some methods for our customized types.

/*
	- Volume() uses Box as its receiver and returns the volume of Box.
	- SetColor(c Color) changes the box's color.
	- BiggestColor() returns the color which has the biggest volume.
	- PaintItBlack sets color for all Box in BoxList to black.
	- String() use Color as its receiver, returns the string format of color name.
 */ 




// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// INTERFACE

/*
	One of the subltest design features in Go are interfaces. 

	What is an interface?

	An interface is a set of methods that we use to define a set of actions.
*/ 

// An interface defines a set of methods, so if a type implements all the methods we say that it implements the interface.

type Human struct {
	// fields for Human
	name string
	age int
	phone string
}

type Student struct {
	// student inherits fields from Human
	Human

	school string
	loan float32
}

type Employee struct {
	// inherits from Human
	Human
	company string
	money float32
}


func (h *Human) SayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human) Sing(lyrics string){
	fmt.Println("La la, la, la, la", lyrics)
}

func (h *Human) Guzzle(beerStein string){
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee overloads SayHi

func (e *Employee) SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func (s *Student) BorrowMoney(amount float32){
	s.loan += amount // again and again
}

func (e *Employee) SpendSalary(amount float32){
	e.money -= amount 
}

// defining an interface

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

// An interface can be implemented by any type, and one type can implement many interfaces simultaneously 






















