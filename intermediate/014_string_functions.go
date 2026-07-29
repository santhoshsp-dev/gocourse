package main

import (
	"fmt"
	"strings"
)

func main() {

	// Step: 1)
	// str := "Hello Go!"
	// fmt.Println(len(str))

	// Step: 2)
	// str1 := "Hello"
	// str2 := "World"
	// result := str1 + " " + str2
	// fmt.Println(result)

	// Step: 3)
	// fmt.Println(str[0])

	// Step: 4)
	// fmt.Println(str[1:7])

	// Step: 5)
	// // String Conversion
	// num := 18
	// str3 := strconv.Itoa(num) // convert integer to string
	// fmt.Println(len(str3))

	// Step: 6)
	// // strings splitting
	// fruits := "apple, orange, banana"
	// parts := strings.Split(fruits, ",")
	// fmt.Println(parts)

	// Step: 7)
	// fruits1 := "apple-orange-banana"
	// parts1 := strings.Split(fruits1, "-")
	// fmt.Println(parts1)

	// Step: 8)
	// countries := []string{"Germany", "France", "Italy"} // string slice
	// joined := strings.Join(countries, ", ")
	// fmt.Println(joined)

	// Step: 9)
	// fmt.Println(strings.Contains(str, "Go?"))

	// Step: 10)
	// replaced := strings.Replace(str, "Go", "Universe", 1)
	// fmt.Println(replaced)

	// Step: 11)
	// strwspace := " Hello Everyone! "
	// fmt.Println(strwspace)
	// fmt.Println(strings.TrimSpace(strwspace))

	// Step: 12)
	// fmt.Println(strings.ToLower(strwspace))
	// fmt.Println(strings.ToUpper(strwspace))

	// Step: 13)
	// fmt.Println(strings.Repeat("foo ", 3))

	// Step: 14)
	// fmt.Println(strings.Count("Hello", "l"))
	// fmt.Println(strings.HasPrefix("Hello", "he"))
	// fmt.Println(strings.HasSuffix("Hello", "lo"))
	// fmt.Println(strings.HasSuffix("Hello", "la"))

	// Step: 15)
	// str5 := "Hel1lo, 123 Go 11!"
	// re := regexp.MustCompile(`\d+`) // d means digit; + means one or more
	// matches := re.FindAllString(str5, -1) // -1 means all the matching expressions
	// fmt.Println(matches)

	// Step: 16)
	// str6 := "Hello, 世界"
	// fmt.Println(utf8.RuneCountInString(str6))

	// Step: 17)
	// STRING BUILDER
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result := builder.String()
	fmt.Println(result)

	// Step: 18)
	// Using Writerune to add a character
	builder.WriteRune(' ') // when we are dealing with strings it has to be single quotes.
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	// Step: 19)
	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)
}
