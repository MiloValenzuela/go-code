package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	
	// if t := time.Now(); t.Hour() < 12 {
	// 	fmt.Println("Ma;ana!")
	// } else if t.Hour() < 17 {
	// 	fmt.Println("Tarde!")
	// } else {
	// 	fmt.Println("Noche!")
	// }

	switch  t := time.Now(); {
	case t.Hour() < 12:
			fmt.Println("Ma;ana!")
	case t.Hour() < 17:
			fmt.Println("Tarde!")
	default:
		fmt.Println("Noche!")
	}

	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}
}