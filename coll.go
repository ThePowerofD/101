package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	binary_global()

}

func prueba() {

}
// NAamas para checar Git vamos a seguir modificando a ver que onda/ 
func binary_global() {
	var counter int
	num := 0
	num2 := 0
	fmt.Println("Dame un numero para convertir en Binario")
	fmt.Scan(&num)
	fmt.Println("Dame un segndo numero para convertir en Binario")
	fmt.Scan(&num2)

	fmt.Println("")
	var bin []int
	var binary []int
	numb := num
	for i := num; num != 1; {
		if num == 0 {
			bin = append(bin, 0)
			break
		}
		bin = append(bin, num%2)
		num = i / 2
		i = num

	}
	bin = append(bin, 1)
	s := 0
	for c := len(bin); c > 0; c-- {
		s = c - 1
		binary = append(binary, bin[s])

	}

	fmt.Println("Numero", numb, "= Binario : ", binary)
	for s := 0; s < len(binary); s++ {
		if 1 == binary[s] {
			counter += 1
		}
	}
	fmt.Println("La cantidad de 1's en la wea esta binaria son: ", counter)
	// Segundo Binario

	fmt.Println("")
	var bin2 []int
	var binary2 []int
	numb2 := num2
	for i := num2; num2 != 1; {
		if num2 == 0 {
			bin2 = append(bin2, 0)
			break
		}
		bin2 = append(bin2, num2%2)
		num2 = i / 2
		i = num2

	}
	bin2 = append(bin2, 1)
	s = 0
	for c := len(bin2); c > 0; c-- {
		s = c - 1
		binary2 = append(binary2, bin2[s])

	}

	fmt.Println("Numero", numb2, "= Binario : ", binary2)
	counter = 0
	for s = 0; s < len(binary2); s++ {
		if 1 == binary2[s] {
			counter += 1
		}
	}
	fmt.Println("La cantidad de 1's en la wea esta binaria son: ", counter)

	//Aqui comparamos ambos Binairos
	fmt.Println("")
	var comp []int
	var y int
	if len(binary2) > len(binary) {
		for z, item := range binary {
			y = binary2[z]
			if y == item && y != 0 {
				comp = append(comp, z)
			}
		}
	} else if len(binary) > len(binary2) {
		for z, item := range binary2 {
			y = binary2[z]
			if y == item && y != 0 {
				comp = append(comp, z)
			}
		}
	} else if len(binary) == len(binary2) {
		for z, item := range binary {
			y = binary2[z]
			if y == item && y != 0 {
				comp = append(comp, z)
			}
		}
	}
	fmt.Println("Las posiciones iguales con 1 en ambos numero binarios son las siguientes\n", comp)

}

func binary1count(num int) {
	/*var j int
	fmt.Scan(&j)*/

	counter := 0
	fmt.Println("")
	var bin []int
	var binary []int
	numb := num
	for i := num; num != 1; {
		if num == 0 {
			bin = append(bin, 0)
			break
		}
		bin = append(bin, num%2)
		num = i / 2
		i = num

	}
	bin = append(bin, 1)
	s := 0
	for c := len(bin); c > 0; c-- {
		s = c - 1
		binary = append(binary, bin[s])

	}

	fmt.Println("Numero", numb, "= Binario : ", binary)
	for s := 0; s < len(binary); s++ {
		if 1 == binary[s] {
			counter += 1
		}
	}
	fmt.Println("La cantidad de 1's en la wea esta binaria son: ", counter)

	/*fmt.Println("Escribe el valor de el numero que quieras traducir a Binairo")
	j := 0
	counter := 0
	fmt.Scan(&j)
	binary_calc(j)
	for s := 0; s < len(binary); s++ {
		if 1 == binary[s] {
			counter += 1
		}
	}
	fmt.Println("La cantidad de 1's en la wea esta binaria son: ", counter)*/
}

func binary_calc(num int) {
	fmt.Println("")
	var bin []int
	var binary []int
	numb := num
	for i := num; num != 1; {
		if num == 0 {
			bin = append(bin, 0)
			break
		}
		bin = append(bin, num%2)
		num = i / 2
		i = num

	}
	bin = append(bin, 1)
	s := 0
	for c := len(bin); c > 0; c-- {
		s = c - 1
		binary = append(binary, bin[s])

	}

	fmt.Println("Numero", numb, "= Binario : ", binary)

}

func semana2() {
	var s, c int
	var list2 []int
	var liststr [4]string
	var arrstr []string
	var frase string
	var frase1 string

	//Problema 1 Buscar el elemento mayor en una lista
	list := []int{1, 2, 7, 4, 5, 6} // For the real thing just add more things to the array

	fmt.Println("Check", list)

	s = 0
	for _, element := range list {
		if element > s {
			s = element

		}

	}
	fmt.Println("El numero mayor es:", s)

	//Problema 1.1
	frase = "c"
	liststr[0] = "Jorge"
	liststr[1] = "Fer"
	liststr[2] = "Joaquin"
	liststr[3] = "Jolgais"
	fmt.Println("Esta es la Lista\n", liststr)

	for i := 0; i < len(liststr); i++ {
		if len(frase) < len(liststr[i]) {
			frase = liststr[i]
		}
	}

	for i := 0; i < len(liststr); i++ {
		if len(frase) == len(liststr[i]) {
			arrstr = append(arrstr, liststr[i])
		}
	}
	fmt.Println("Los elementos mas grandes son:", arrstr, "Y tienen un tamaño de:", len(frase))
	// Problema 2 Write function that reverses a list, preferably in place.
	c = len(list)
	fmt.Println(" ")

	fmt.Println("Revertiendo numeros con array")
	for c = len(list); c > 0; c-- {
		s = c - 1
		list2 = append(list2, list[s])

	}
	fmt.Println(list2)

	// 2.1 Como hacerlo con commandos de Sort
	sort.Ints(list2)
	fmt.Println("Sorted again", list2)
	sort.Sort(sort.Reverse(sort.IntSlice(list2)))
	fmt.Println("Reversed without using arrays", list2)

	//Write a function that checks whether an element occurs in a list
	fmt.Println("Check if the number is in the array")
	fmt.Scan(&s)
	c = 0
	for _, element := range list {
		if element == s {
			c += 1
			fmt.Println("The Element", s, "is in the Array", c, "times")
			break
		}
	}
	if c == 0 {
		fmt.Println("Your number is not in the array")
	}
	// 4.- Write a function that returns the elements on odd positions in a list
	fmt.Println("Imprimir espacios nones de la lista")
	for i := 0; i < len(liststr); i = i + 2 {
		fmt.Println(liststr[i])
	}
	fmt.Println(" ")

	// 5.- Writa a Function that computes the runing total of a list

	//6.- Write a function that tests whether a string is a palindrome
	frase = "ala"

	frase1 = strings.ReplaceAll(frase, " ", "")
	fmt.Println(frase, "\n", frase1)
	c = 1
	for i, j := 0, (len(frase1) - 1); i < j; i, j = i+1, j-1 {
		fmt.Println("i ", i, "\n j ", j)
		if frase1[i] != frase1[j] {
			fmt.Println("It is not a Palindrome")
			c = 0
			break
		}
	}
	if c == 1 {
		fmt.Println("Yes it is a Palindrome")
	}
	// Aqui hacemos la chulada esa que queria el Juan
	fmt.Println("Mira Esto es el valor de la variable s=", s)
	s = 7
	x := &s
	fmt.Println("Mira este es el lugar donde esta guardado", x)
	*x = 10
	fmt.Println("Cambiamos s usando su direccion de veras creeme ;) s= ", s)
	fmt.Println(" ")

	// 7 Compute sum of nummers in a list using for while and recursion
	// 8  Write afucntion that aplies funciton to each element of a list.

	// 9 Write a Function that concatanatees two lists
	abc := []string{"a", "b", "c"}
	cde := []string{"1", "2", "3"}

	var slice []string

	slice = append(abc, cde...)
	fmt.Println("First string arr", abc, "\nSecond string array", cde, "\nConcatenated Array", slice)
	fmt.Println(" ")

	// 10 Concatenate two lists making alternating elements
	fmt.Println("Concatenated alternating lists")

	slice = nil

	for i := 0; i < len(abc); i++ {

		slice = append(slice, abc[i])
		slice = append(slice, cde[i])

	}
	fmt.Println(slice)

	// 11 Merge two lists and sort them at the same time
	fmt.Println("Merge and Sort using [a,b,c,] & [1,2,3]")
	fmt.Println("")
	slice = nil

	slice = append(cde, abc...)
	fmt.Println("Merged", slice)
	sort.Strings(slice)
	fmt.Println("Sorted out string:", slice)

}

func semana1() {
	//Problema 1 Hello World

	fmt.Println("Hello, World!")
	var name, names, option string
	var number, sum int
	var ans, prompt, boo, errcount int

	// Problema 2 Ask Input
	fmt.Println("Please insert your name:")
	fmt.Scan(&name)
	fmt.Println("Welcome ", name)
	// Problema 3 Analizar Input y responder acorde
	fmt.Println("Are you on the VIP List?")
	fmt.Println("Type your name again: ")
	fmt.Scan(&names)

	if names == "Alice" {
		fmt.Println("Hello ", names, " Welcome to the VIP section")

	} else if names == "Bob" {
		fmt.Println("Hello ", names, " Welcome to the VIP section")

	} else {
		fmt.Println("Sorry you are not on the VIP go sit with the commoners")
	}
	// Problema 4 n!
	fmt.Println("Give me a number: ")
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		sum += i
	}
	fmt.Print(sum, " This is the result of ", number, "!\n")
	// Problema 5 n! solo con mulitplos de 5 & 3
	sum = 0
	fmt.Println("This is the Test for n! for 3 & 5 Multiples:")
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("The Result is", sum)
	// Problema 6 += o *=
	fmt.Println("Give me a number for n: ")
	fmt.Scan(&number)
	fmt.Println("Choose do you want the sum or the product from 1 to n\n Type '+' to Sum otherwise Type '*' for product")
	fmt.Scan(&option)
	if option == "+" {
		sum = 0
		for i := 1; i <= number; i++ {
			sum += i
		}
		fmt.Println("Sum =", sum)
	} else if option == "*" {
		sum = 1
		for i := 1; i <= number; i++ {
			sum *= i
		}
		fmt.Println("Product =", sum)
	} else {
		fmt.Println("You did not choose a correct input BITCH")
	}

	//Problem 7
	var multable [12][12]int
	for i := 0; i <= 11; i++ {
		for c := 0; c <= 11; c++ {
			multable[i][c] = (c + 1) * (i + 1)
		}
	}
	fmt.Println(multable)

	// Problem 8 producing prime numbers
	fmt.Println("Up to What number do you wish to know the prime numbers?")
	fmt.Scan(&number)
	if number == 1 {
		fmt.Println("1")
	} else if number == 2 {
		fmt.Print("1")
		fmt.Print("2")
	} else if number >= 3 {

		for n := 1; n <= number; n++ {

			if n == 3 || n == 5 || n == 7 || n == 1 || n == 2 {
				fmt.Println(n)
			} else {
				if n%2 != 0 && n%3 != 0 && n%4 != 0 && n%5 != 0 && n%7 != 0 {
					fmt.Println(n)
				}
			}

		}
	} else {
		fmt.Println("Don't be Funny use a Natural Number.")
	}

	// Problem 9

	boo = 1
	ans = rand.Intn(1000 - 0)
	fmt.Println("Try to guess the number")

	fmt.Println("ans", ans)
	fmt.Scan(&prompt)
	var gametable []int
	gametable = append(gametable, prompt)
	for ans != prompt {
		fmt.Println("gametable check", gametable)
		if ans < prompt {
			fmt.Println("The number is Lower")
			fmt.Println("Try Again")
			//Agregando al contador
			for _, element := range gametable {
				if element != prompt {
					boo *= 1
				} else if element == prompt {
					boo *= 2
				}
			}
			if boo == 1 {
				errcount += 1
				gametable = append(gametable, prompt)
			}

		} else if ans > prompt {
			fmt.Println("The number is Higher")
			fmt.Println("Try Again")
			//Contador de Errores
			for _, element := range gametable {
				if element != prompt {
					boo *= 1
				} else if element == prompt {
					boo *= 2
				}
			}
			if boo == 1 {
				errcount += 1
				gametable = append(gametable, prompt)
			}
		}
		fmt.Scan(&prompt)
		boo = 1

	}
	fmt.Println("GAME OVER", "YOU WON")
	fmt.Println("# of ERRORs", errcount+1)
	fmt.Println("Array CHeck", gametable)
}
