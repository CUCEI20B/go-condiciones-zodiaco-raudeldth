package main

import "fmt"

func main()  {
	var dia int32
	var mes int32
	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch mes {
		case 1 :
			if dia >= 20 {
			fmt.Println("acuario")
			} else {
				fmt.Println("capricornio")
			}
		case 2 :
			if dia >= 19 {
			fmt.Println("piscis")
			} else {
				fmt.Println("acuario")
			}
		case 3 :
			if dia >= 21 {
			fmt.Println("aries")
			} else {
				fmt.Println("piscis")
			}
		case 4 :
			if dia >= 20 {
			fmt.Println("tauro")
			} else {
				fmt.Println("aries")
			}
		case 5 :
			if dia >= 21 {
			fmt.Println("geminis")
			} else {
				fmt.Println("tauro")
			}
		case 6 :
			if dia >= 21 {
			fmt.Println("cancer")
			} else {
				fmt.Println("geminis")
			}
		case 7 :
			if dia >= 23 {
			fmt.Println("leo")
			} else {
				fmt.Println("cancer")
			}
		case 8 :
			if dia >= 23 {
			fmt.Println("virgo")
			} else {
				fmt.Println("leo")
			}
		case 9 :
			if dia >= 23 {
			fmt.Println("libra")
			} else {
				fmt.Println("virgo")
			}
		case 10:
			if  dia >= 23 {
			fmt.Println("escorpio")
			} else {
				fmt.Println("libra")
			}
		case 11:
			if  dia >= 22 {
			fmt.Println("sagitario")
			} else {
				fmt.Println("escorpio")
			}
		case 12:
			if  dia >= 22 {
			fmt.Println("capricornio")
			} else {
				fmt.Println("sagitario")
			}
	}
}
