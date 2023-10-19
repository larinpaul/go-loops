//// 2023/10/19 // 21:44 //

//// Loops in Go

// // The basic loop in Go is written in standard C-like syntax:
// for INITIAL; CONDITION; AFTER {
// 	// do something
// }

// // INITIAL is run once at the beginning of the loop and can create variables within the scope of the loop.
// // CONDITION is checked before each iteration. If the condition doesn't pass then the loop breaks.
// // AFTER is run after each iteration.

// // For example:
// for i := 0; i < 10; i++ {
// 	fmt.Println(i)
// }
// // Prints 0 throught 9

// //// Assignment

// // At Textio we have a dynamic formula for determening how much a batch of bulk messages costs to send.

// // Complete the bulkSend() function

// // This function should return the total cost (as a float64) to send a batch of numMessages messages. Each
// // messages costs 1.0, plus an additional fee. The fee structure is:
// // * 1st message: 1.0 + 0.00
// // * 2nd message: 1.0 + 0.01
// // * 3rd message: 1.0 + 0.02
// // * 4th message: 1.0 + 0.03
// // Use a loop to calculate the total cost and return it.

// package main

// import (
// 	"fmt"
// )

// func bulkSend(numMessages int) float64 {
// 	totalCost := 0.0
// 	for i := 0; i < numMessages; i++ {
// 		totalCost += 1 + (float64(i) * 0.01)
// 	}
// 	return totalCost
// }

// func test(numMessages int) {
// 	fmt.Printf("Sending %v messages\n", numMessages)
// 	cost := bulkSend(numMessages)
// 	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
// 	fmt.Println("===========================================================")
// }

// func main() {
// 	test(2)
// 	test(10)
// 	test(20)
// 	test(30)
// 	test(40)
// 	test(50)
// }

//// 21:54
//// Omitting conditions from a for loop in Go

// // Loops in GO can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted
// // which causes the loop to run forever.
// for INITIAL; ; AFTER {
// 	// do something forever
// }

//// ASSIGNMENT

// Complete teh maxMessages function. Given a cost threshold, it should calculate the maximum number of
// messages that can be sent.

// Each message costs 1.0, plus an additional fee. The fee structure is:
// * 1st message: 1.0 + 0.00
// * 2nd message: 1.0 + 0.01
// * 3rd message: 1.0 + 0.02
// * 4th message: 1.0 + 0.03

// Browser freeze
// If you lock up your browser by creating an infinite loop that isn't breaking, just click the cancel button.

// package main

// import (
// 	"fmt"
// )

// func maxMessages(thresh float64) int {
// 	totalCost := 0.0
// 	for i := 0; ; i++ {
// 		totalCost += 1 + (float64(i) * 0.01)
// 		if totalCost > thresh {
// 			return i
// 		}
// 	}
// }

// func test(thresh float64) {
// 	fmt.Printf("Threshold: %.2f\n", thresh)
// 	max := maxMessages(thresh)
// 	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
// 	fmt.Println("================================================================")
// }

// func main() {
// 	test(10.00)
// 	test(20.00)
// 	test(30.00)
// 	test(40.00)
// 	test(50.00)
// }

//// 22:00
//// There is no while loop in Go

// // Most programming languages have a concept of a while loop. Because Go allows for the omission of
// // sections of a for loop, a while loop is jsut a for loop that only has a CONDITION.

// for CONDITION {
// 	// do some stuff while CONDITION is true
// }

// // For example:
// plantHeight := 1
// for plantHeight < 5 {
// 	fmt.Println("still growing! current height:", plantHeight)
// 	plantHeight++
// }
// fmt.Println("plannt has grown to ", plantHeight, "inches")

// // Which prints:
// // still growing! current height: 1
// // still growing1 current height: 2
// // still growing! current height: 3
// // still growing! current height: 4
// // plant has grown to 5 inches

//// Assignment

// We have an interesting new cost structure from out SMS vendor. They charge exponentially more money
// for each consecurtive text we send! Let's write a function that can calculate how many messages we can
// send in a given batch given a costMultiplier and a maxCostInPennies.

// In a nutshell, the first message costs a penny, and each message after that costs the same as the previous
// message multiplied by the costMultiplier. That gets expensive!

// There is an infinite loop in the code! Let's add a condition to fix the bug. The loop should exit before
// incrementing maxMessagesToSend if the cost of the next message would go over the max cost.

// package main

// import (
// 	"fmt"
// )

// func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
// 	actualCostInPennies := 1.0
// 	maxMessagesToSend := 0
// 	for actualCostInPennies <= float64(maxCostInPennies) {
// 		maxMessagesToSend++
// 		actualCostInPennies *= costMultiplier
// 	}
// 	return maxMessagesToSend
// }

// func test(costMultiplier float64, maxCostInPennies int) {
// 	maxMessagesToSend := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
// 	fmt.Printf("Multiplier: %v\n",
// 		costMultiplier,
// 	)
// 	fmt.Printf("Max cost: %v\n",
// 		maxCostInPennies,
// 	)
// 	fmt.Printf("Max messages you can send: %v\n",
// 		maxMessagesToSend,
// 	)
// 	fmt.Println("======================================")
// }

// func main() {
// 	test(1.1, 5)
// 	test(1.3, 10)
// 	test(1.35, 25)
// }

// //// 22:54
// //// Fizzbuzz

// // Go supports the standard modulo operator:
// 7 % 3 // 1

// // Logical AND operator:
// true && false // false
// true && true // true

// // Logical OR operator:
// true || false // true
// false || false // false

//// Assignment

// We're hiring engineers at Textioo, so time to brush up on the classic "Fizzbuz" game, a coding exercise that
// has been dramatically overused in coding interviews across the world.

// Complete the fizzbuzz function that prints the numbers 1 to 100 inclusive each on their own line, but
// substitutes multiples of 3 for the text fizz and multiples of 5 for buzz. For multiples of 3 AND 5 print
// instead fizzbuzz.

// package main

// import "fmt"

// func fizzbuzz() {
// 	for i := 1; i < 101; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("fizzbuzz")
// 		} else if i%3 == 0 {
// 			fmt.Println("fizz")
// 		} else if i%5 == 0 {
// 			fmt.Println("buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}

// }

// func main() {
// 	fizzbuzz()
// }

//// 23:00

// //// Continue

// continue

// // The continue keyword stops the current iteration of a loop and continues to the next iteration. continue
// // is a powerful way to use the "guard clause" pattern within loops.
// for i := 0; i < 10; i++ {
// 	if i%2 == 0 {
// 		continue
// 	}
// 	fmt.Println(i)
// }
// // 1
// // 3
// // 5
// // 7
// // 9

// break

// // The break keyword stops the current iteration of a loop and exits the loop.
// for i := 0; i < 10; i++ {
// 	if i ==5 {
// 		break
// 	}
// 	fmt.Println(i)
// }
// // 0
// // 1
// // 2
// // 3
// // 4

// // Assignment

// // As an easter egg, we decide to reward our users with a free text message if they send a prime number of
// // text messages this year.

// // Complete the printPrimes function. It should print all of the prime numbers up to and including max. It
// // should skip any numbers that are not prime.

// // Here's the pseudocode:
// printPrimes(max):
// 	for n in range(2, max+1):
// 		if n is 2:
// 			n is prime, print it
// 		if n is even:
// 			n is not prime, skip to next n
// 		for i in range (3, sqrt(n) + 1):
// 			if i can be multipled into n:
// 				n is not prime, skip to next n
// 		n is prime, print i

// // Breakdown
// // * We skip even numbers because they can't be prime
// // * We only check up to the square root because anything higher than the square root has no chance of
// // multiplying evenly into n
// // * We start checking at 2 because 1 is not prime

package main

import (
	"fmt"
)

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}

		if n%2 == 0 {
			continue
		}

		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("================================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
