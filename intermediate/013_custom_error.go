package intermediate

import (
	"errors"
	"fmt"
)

func main() {

	// Step: 4)
	err := doSomething()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Operation completed successfully!")

}

// Step: 1)
// type customError struct {
// 	code    int
// 	message string
// }

// Step: 5)
type customError struct {
	code    int
	message string
	er      error
}

// Step: 2)
// // Error returns the error message. Implementing Error() method of error interface
//	func (e *customError) Error() string {
//		return fmt.Sprintf("Error %d: %s", e.code, e.message)
//	}

// Step: 8)
// Error returns the error message. Implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.er)
}

// Step: 3)
// Function that return a custom error
//
//	func doSomething() error {
//		return &customError{
//			code:    500,
//			message: "Something went wrong!",
//		}
//	}

// Step: 7)
func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er:      err,
		}
	}
	return nil
}

// Step: 6)
func doSomethingElse() error {
	return errors.New("internal error")
}
