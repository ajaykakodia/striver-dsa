package main

import "fmt"

/*
* * * * *
* * * * *
* * * * *
* * * * *
* * * * *
 */
func pattern1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
*
* *
* * *
* * * *
* * * * *
 */
func pattern2(n int) {
	for row := 0; row < n; row++ {
		for col := 0; col <= row; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
1
1 2
1 2 3
1 2 3 4
1 2 3 4 5
*/
func pattern3(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(c+1, " ")
		}
		fmt.Println()
	}
}

/*
1
2 2
3 3 3
4 4 4 4
5 5 5 5 5
*/
func pattern4(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(r+1, " ")
		}
		fmt.Println()
	}
}

/*
* * * * *
* * * *
* * *
* *
*
 */
func pattern5(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n-r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
1 2 3 4 5
1 2 3 4
1 2 3
1 2
1
*/
func pattern6(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n-r; c++ {
			fmt.Print(c+1, " ")
		}
		fmt.Println()
	}
}

/*
    *
   ***
  *****
 *******
*********
*/
func pattern7(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n-r-1; c++ {
			fmt.Print(" ")
		}
		for c := 0; c <= r; c++ {
			fmt.Print("*")
		}
		for c := 0; c < r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
*********
 *******
  *****
   ***
    *
*/
func pattern8(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < r; c++ {
			fmt.Print(" ")
		}
		for c := 0; c < n-r; c++ {
			fmt.Print("*")
		}
		for c := 0; c < n-r-1; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
    *
   ***
  *****
 *******
*********
 *******
  *****
   ***
    *
*/
func pattern9(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n-r-1; c++ {
			fmt.Print(" ")
		}
		for c := 0; c <= r; c++ {
			fmt.Print("*")
		}
		for c := 0; c < r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for r := 0; r < n-1; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(" ")
		}
		for c := 0; c < n-r-1; c++ {
			fmt.Print("*")
		}
		for c := 0; c < n-r-2; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
*
* *
* * *
* * * *
* * * * *
* * * *
* * *
* *
*
 */
func pattern10(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for r := 0; r < n-1; r++ {
		for c := 0; c < n-r-1; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
1
0 1
1 0 1
0 1 0 1
1 0 1 0 1
*/
func pattern11(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print((r + c + 1) & 1)
		}
		fmt.Println()
	}
}

/*
1                 1
1 2             2 1
1 2 3         3 2 1
1 2 3 4     4 3 2 1
1 2 3 4 5 5 4 3 2 1
*/
func pattern12(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(c + 1)
		}
		for c := 0; c < n-r-1; c++ {
			fmt.Print("  ") // two prints for cover 2nd area and save one loop
		}
		for c := 0; c <= r; c++ {
			fmt.Print(r - c + 1)
		}
		fmt.Println()
	}
}

/*
1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
*/
func pattern13(n int) {
	count := 1
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(count, " ")
			count++
		}
		fmt.Println()
	}
}

/*
A
AB
ABC
ABCD
ABCDE
*/
func pattern14(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(string('A'+c), " ")
		}
		fmt.Println()
	}
}

/*
ABCDE
ABCD
ABC
AB
A
*/
func pattern15(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n-r; c++ {
			fmt.Print(string('A'+c), " ")
		}
		fmt.Println()
	}
}

/*
A
BB
CCC
DDDD
EEEEE
*/
func pattern16(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(string('A'+r), " ")
		}
		fmt.Println()
	}
}

/*
   A
  ABA
 ABCBA
ABCDCBA
*/
func pattern17(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n-r-1; c++ {
			fmt.Print(" ")
		}
		for c := 0; c <= r; c++ {
			fmt.Print(string('A' + c))
		}
		for c := 0; c < r; c++ {
			fmt.Print(string('A' + r - 1 - c))
		}
		fmt.Println()
	}
}

/*
E
E D
E D C
E D C B
E D C B A
*/
func pattern18(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(string('A'+n-1-c), " ")
		}
		fmt.Println()
	}
}

/*
****
*  *
*  *
****
 */
func pattern21(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if r == 0 || r == n-1 || c == 0 || c == n-1 {
				fmt.Print("*")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

/*
Bheja Fry Pattern -
for input 4 :

4 4 4 4 4 4 4
4 3 3 3 3 3 4
4 3 2 2 2 3 4
4 3 2 1 2 3 4
4 3 2 2 2 3 4
4 3 3 3 3 3 4
4 4 4 4 4 4 4
*/
func pattern22(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(n-c, " ")
		}
		for c := 0; c < n-r-1; c++ {
			fmt.Print(n-r, " ", n-r, " ")
		}
		for c := 0; c < r; c++ {
			fmt.Print(n-r+c+1, " ")
		}
		fmt.Println()
	}
	for r := 0; r < n-1; r++ {
		for c := 0; c < n-r-1; c++ {
			fmt.Print(n-c, " ")
		}
		for c := 0; c <= r; c++ {
			fmt.Print(2+r, " ", 2+r, " ")
		}
		for c := 0; c < n-r-2; c++ {
			fmt.Print(n-1+c+r, " ")
		}

		fmt.Println()
	}
}

/*
**********
****  ****
***    ***
**      **
*        *
*        *
**      **
***    ***
****  ****
**********
 */
func pattern19(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n-r; c++ {
			fmt.Print("*")
		}
		for c := 0; c < r; c++ {
			fmt.Print("  ")
		}
		for c := 0; c < n-r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print("*")
		}
		for c := 0; c < n-1-r; c++ {
			fmt.Print("  ")
		}
		for c := 0; c <= r; c++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

/*
*        *
**      **
***    ***
****  ****
**********
****  ****
***    ***
**      **
*        *
 */
func pattern(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print("*")
		}
		for c := 0; c < n-r-1; c++ {
			fmt.Print("  ")
		}
		for c := 0; c <= r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for r := 0; r < n-1; r++ {
		for c := 0; c < n-r-1; c++ {
			fmt.Print("*")
		}
		for c := 0; c <= r; c++ {
			fmt.Print("  ")
		}
		for c := 0; c < n-r-1; c++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
