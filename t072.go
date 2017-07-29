// Multiple structs implement an interface. How to:
//
// - Create any of them from a single function using an argument
//
// - Assign any of them to the same interface variable and call the function
//   declared in the interface
//
// - Custom error types
//
// Rule of thumb: There is no such thing as pointer to interface!
package main

import "fmt"

//
type PortScanner interface {
	Scan(int, int) (map[int]bool, *PortScannerError)
}

//
type VanillaTCPScanner struct {
	x string
}

//
func (s *VanillaTCPScanner) Scan(from, to int) (res map[int]bool, pse *PortScannerError) {
	return nil, &PortScannerError{
		descr: "VanillaTCPScanner.Scan() successful!"}
}

//
type Cat struct {
	x string
}

//
func (s *Cat) Scan(from, to int) (res map[int]bool, pse *PortScannerError) {
	return nil, &PortScannerError{
		descr: "Cat.Scan(): Dude, WTF? I'm a cat."}
}

// MOD(1): WORKS
func New(what string) PortScanner {
	switch what {
	case "vanilla":
		return &VanillaTCPScanner{x: "vanilla-scanner"}
	case "cat":
		return &Cat{x: "Meow"}
	}
	return nil
}

/* MOD(2): WORKS TOO
//
func NewVanilla() (s PortScanner) {
	s = &VanillaTCPScanner{x: "vanilla-scanner"}
	return s
}

//
func NewCat() PortScanner {
	return &Cat{x: "Meow"}
}
*/

//
type PortScannerError struct {
	descr string
}

//
func (pse *PortScannerError) Error() string {
	return fmt.Sprintf("PortScannerError: %v\n", pse.descr)
}

func main() {
	// MOD(1): WORKS
	var scanner PortScanner = New("vanilla")
	if _, err := scanner.Scan(0, 1000); err != nil {
		fmt.Println(err)
	}

	scanner = New("cat")
	if _, err := scanner.Scan(0, 1000); err != nil {
		fmt.Println(err)
	}
	/* MOD(2): WORKS TOO
	var scanner PortScanner = NewVanilla()
		if _, err := scanner.Scan(0, 1000); err != nil {
			fmt.Println(err)
		}

		scanner = NewCat()
		if _, err := scanner.Scan(0, 1000); err != nil {
			fmt.Println(err)
		}*/
}
