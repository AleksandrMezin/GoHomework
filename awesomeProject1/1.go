package main

import "fmt"

func main() {
	/*Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
	Программа в первой строке получает на вход число n - количество чисел в последовательности,
	во второй строке -- n чисел, входящих в данную последовательность.
	*/
	var a int
	var b int
	var sum int
	fmt.Scan(&a)
	for; a < 5; a++ {
		fmt.Println(a)
	}

	fmt.Scan(&b)
	if 10 <= b && b%8 == 0 && b <= 99 {
		sum += b
	}
	fmt.Println(sum)
}
	/*	Вообще-м ребят мучался с этим примером 2 дня, т.к. условие написано не очень понятно. Если б не коменты не решил бы!!!
			1. Объявляем 3 переменные просто интовые (a,b,sum )
		2. Cканируем через  fmt.Scan(&a) переменную а
		3.Создаем цикл for и подставляем вместо числа (количество итераций )  переменную а (т.к. а=5 цикл будет крутится 5 раз )
		4.Внутри цикла сканируем  переменную b  fmt.Scan(&b)
		5.Создаем условие (подходит нам число или нет)  if 10<=b && b%8==0 && b<=99
			6. Если число подходит , то прибавляем  к пиременной sum+=b
		7.Выводим переменную sum
	*/
}
