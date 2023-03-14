package main //everything with this name at the beginning belongs to the same package
// "main" is specifically used for executables
// Literally any other word makes it non-executable (reusable)
import "fmt" //import fmt functions into main package

func main() { //must use main func to automatically run
	fmt.Println("Hi there!")
}
