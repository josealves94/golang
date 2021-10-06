package main


import (

        "fmt"
)

func add(nb int, nb2 int) int {
        result := nb + nb2
        return result
}


func sub(nb3 int, nb4 int ) int {
        result := nb3 - nb4
        return result
}

func mul(nb5 int, nb6 int) int {
        result := nb5 * nb6
        return result
}

func div(nb7 int, nb8 int) int {
        result := nb7 / nb8
        return result
}

func main() {
        var op1 string
        var nb1, nb2, nb3, nb4, nb5, nb6, nb7, nb8 int

        fmt.Println("Please choose an operation for calculation , possibles choices are : + or - or * or  /")
        fmt.Scanln(&op1)
                switch op1 {
                case "+":
                        fmt.Println("You choose Please put your first  number below")
                        fmt.Scanln(&nb1)
                        fmt.Println("Please put your second number below")
                        fmt.Scanln(&nb2)
                        res := add(nb1, nb2)
                        fmt.Printf("Result is %d ", res)

                case "-":
                        fmt.Println("You choose substraction, Please put the first number below")
                        fmt.Scanln(&nb3)
                        fmt.Println("Please put your second number below")
                        fmt.Scanln(&nb4)
                        res := sub(nb3, nb4)
                        fmt.Printf("Result is %d", res)
                case "*":
                        fmt.Println("You choose multiplication, Please put the first number below")
                        fmt.Scanln(&nb5)
                        fmt.Println("Please put your second number below")
                        fmt.Scanln(&nb6)
                        res := mul(nb5, nb6)
                        fmt.Printf("Result is %d", res)
                case "/":
                        fmt.Println("You choose division, Please put the first number below")
                        fmt.Scanln(&nb7)
                        fmt.Println("Please put your second number below")
                        fmt.Scanln(&nb8)
                        res := div(nb7, nb8)
                        fmt.Printf("Result is %d", res)
                default:
                        fmt.Println("Please choose + or - or * or /")
                }

}
