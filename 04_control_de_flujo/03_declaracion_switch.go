package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/*
	switch valor {
	case ... :
		...
		break
	default:
		...
		break
	}
	*/

/*
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Go run -> 	Windows")
		break
	case "linux":
		fmt.Println("Go run -> Linux")
		break
	case "darwin":
		fmt.Println("Go run -> MacOS")
		break
	default:
		fmt.Println("Go run -> Otro OS")
		break
	}
*/
	// los sintaxis son mismos como anterior los sintaxis
	
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> 	Windows")
		break
	case "linux":
		fmt.Println("Go run -> Linux")
		break
	case "darwin":
		fmt.Println("Go run -> MacOS")
		break
	default:
		fmt.Println("Go run -> Otro OS")
		break
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("¡Mañana!")
		break
	case t.Hour() < 17:
		fmt.Println("¡Tarde!")
		break
	default:
		fmt.Println("¡Noche!")
		break
	}

}