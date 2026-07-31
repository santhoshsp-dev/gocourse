package main

import (
	"errors"
	"fmt"
)

// Step: 1)
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: square root of negative number")
	}
	// compute the square root
	return 1, nil
}

// Step: 4)
func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data")
	}
	// Process data
	return nil
}

func main() {

	// Step: 2)
	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// Step: 3)
	// result1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println(result1)

	// Step: 5)
	// data := []byte{}
	// // if err := process(data); err != nil {
	// err := process(data)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Data Processed Successfully")

	// Step: 9)
	// --- error interface of builtin package
	// err1 := eprocess()
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// println("")

	// Step: 12)
	err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully.")

}

// Step: 6)
type myError struct {
	message string
}

// Step: 7)
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

// Step: 8)
func eprocess() error {
	return &myError{"Custom error message"}
}

// Step: 11)
func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

// Step: 10)
func readConfig() error {
	return errors.New("config error")
}
