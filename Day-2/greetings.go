//  the function takes the name parameter with the type string

package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
