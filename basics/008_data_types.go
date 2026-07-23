// 008 Data Types

package main

import "fmt"

func main() {
	fmt.Println("Hello" + "World")
	fmt.Println("9 X 10 =", 9*10)
	fmt.Println("180.18/2.0 = ", 180.18/2.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

/*
In go programming data types define the type of data that variables can store. Each variable in go must have a specific data type, which determines the size, range of values, and operations that can be performed on the variable. Some of the data types that go language supports are as follows. In the numeric types, we have integers which represent whole numbers without fractional components. Examples include int which is platform dependent size, int8, int16, int32 or int 64, and their unsigned counterparts like Uint8, Uint16, and so on. And next is floating point numbers. Floating point numbers represent numbers with a fractional component. Examples include float 32, which is single precision, and float 64, which is double precision. Next, we come to complex numbers. Complex numbers are rarely used in day to day programming, but yes, it has a very specific use, and for that go has included complex numbers as well. In its standard library. So the standard library contains complex 64 data type and complex 128 to represent complex numbers and to perform operations on complex numbers. We have a package called Math Forward slash C. And this package includes functions such as real or imag. Imag, which returns the imaginary part of a complex number and conjugate function. Absolute value function and many more. So this was the part about complex numbers. The next data type is Boolean type, which represents logical values like true and false. Next is string type which represents a sequence of characters. And next we come to composite data types. In composite data types we have arrays which are fixed size collection of elements of the same type. Example an array of five integers. And then we come to slices which are dynamic and flexible sequence built on top of arrays. And then we come to maps. Maps are key value pairs where all keys are of the same type, and all values are of the same type. In other languages we can call them hashes, hash, map or dictionary, or they can go by any other name. So in go we call them maps. Next composite data type is structs. Structs are user defined composite types that group together variables of different types under one name. And we will get more clarity on all these data types when we will take them up in their separate lectures. And apart from these we also have pointers, function, data type channels, JSON text and HTML templates, and some more as well. And we will explore all the data types as we move ahead in this course. Now let's also come to zero values. So variables declared without an explicit initialization are assigned a default zero value based on their type. Numeric types are given a value zero by default, and boolean types are defaulted to False string type is an empty string by default, and pointers, slices, maps, functions, and structs are initialized with nil value. Nil, nil. So if we have not initialized any of these with a value, they are initialized with nil values. Understanding data types in go is fundamental to writing effective and efficient programs. Each data type has its characteristics and usage scenarios, influencing how you structure and manipulate data within your go applications. Mastery of data types empowers developers to leverage Go's strengths in performance, concurrency, and simplicity in building scalable and robust software solutions. By comprehending these concepts and practicing with various data types, you will be well equipped to tackle a wide range of programming challenges using go.
*/
