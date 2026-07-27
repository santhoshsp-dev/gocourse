package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Step: 1)
	message := "Hello, \nGo!"

	// Step: 3)
	message1 := "Hello, \tGo!"

	// Step: 5)
	// message2 := "Hello, \rGo!" // Go!lo
	message2 := "Hello, Go!"

	// Step: 7)
	rawMessage := `Hello\nGo`

	// Step: 2)
	fmt.Println(message)

	// Step: 4)
	fmt.Println(message1)

	// Step: 6)
	fmt.Println(message2)

	// Step: 8)
	fmt.Println(rawMessage)

	// Step: 9)
	fmt.Println("Length of message variable is", len(message))

	// Step: 10)
	fmt.Println("Length of message2 variable is", len(message2))

	// Step: 11)
	fmt.Println("Length of rawmessage variable is", len(rawMessage))

	// Step: 12)
	fmt.Println("The first character in message var is", message[0]) // ASCII value of H

	// Step: 13)
	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	// Step: 14)
	// lexicographical comparison
	str1 := "Apple"  // A has an ASCII value of 65
	str := "apple"   // a has an ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	// Step: 15)
	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char)
	}

	// Step: 16)
	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	// Step: 17)
	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	// Step: 18)
	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	// Step: 19)
	var ch rune = 'a'
	jch := '日'

	fmt.Println(ch)
	fmt.Println(jch)

	// Step: 20)
	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	// Step: 21)
	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	// Step: 22)
	const NIHONGO = "日本語" // Japanese text
	fmt.Println(NIHONGO)

	// Step: 23)
	jhello := "こんにちは" // Japanese "Hello"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	// Step: 24)
	r := '😊'
	fmt.Printf("%v\n", r)
	// -- OR --
	fmt.Printf("%d\n", r)
	fmt.Printf("%c\n", r)
}
