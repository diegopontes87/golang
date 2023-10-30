package assistant

import "fmt"

// first letter: uppercase -> the method is visible for other packages
// first letter: lowercase -> the method is visible only for that package
func Write() {
	fmt.Println("Writing from assistant")
	write2()
}
