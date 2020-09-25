package main

import "fmt"

func main() {
	var dia int64
	var mes int64
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	var signo string
	switch mes {
	case 1:

		if dia < 21 {
			signo = "Capricornio"
		} else if dia <= 31 {
			signo = "Acuario"
		} else {
			signo = "Fecha no valida"

		}

	case 2:
		if dia < 17 {
			signo = "Acuario"
		} else if dia <= 29 {
			signo = "Piscis"
		} else {
			signo = "Fecha no valida"

		}
	case 3:
		if dia < 21 {
			signo = "Piscis"
		} else if dia <= 31 {
			signo = "Aries"
		} else {
			signo = "Fecha no valida"

		}
	case 4:
		if dia < 21 {
			signo = "Aries"
		} else if dia <= 30 {
			signo = "Tauro"
		} else {
			signo = "Fecha no valida"

		}
	case 5:
		if dia < 21 {
			signo = "Tauro"
		} else if dia <= 31 {
			signo = "Geminis"
		} else {
			signo = "Fecha no valida"

		}
	case 6:
		if dia < 22 {
			signo = "Geminis"
		} else if dia <= 30 {
			signo = "Cancer"
		} else {
			signo = "Fecha no valida"

		}
	case 7:
		if dia < 23 {
			signo = "Cancer"
		} else if dia <= 31 {
			signo = "Leo"
		} else {
			signo = "Fecha no valida"

		}
	case 8:
		if dia < 23 {
			signo = "Leo"
		} else if dia <= 31 {
			signo = "Virgo"
		} else {
			signo = "Fecha no valida"

		}
	case 9:
		if dia < 23 {
			signo = "Virgo"
		} else if dia <= 30 {
			signo = "Libra"
		} else {
			signo = "Fecha no valida"

		}
	case 10:
		if dia < 23 {
			signo = "Libra"
		} else if dia <= 31 {
			signo = "Escorpio"
		} else {
			signo = "Fecha no valida"

		}
	case 11:
		if dia < 23 {
			signo = "Escorpio"
		} else if dia <= 30 {
			signo = "Sagitario"
		} else {
			signo = "Fecha no valida"

		}
	case 12:
		if dia < 22 {
			signo = "Sagitario"
		} else if dia <= 31 {
			signo = "Capricornio"
		} else {
			signo = "Fecha no valida"

		}
	default:
		signo = "Fecha no valida"

	}
	fmt.Println(signo)
}
