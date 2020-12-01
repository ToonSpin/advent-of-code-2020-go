package main

import (
    "fmt"
    "errors"
)

func part1(input []int) (int, error) {
    for j, a := range input {
        for i, b := range input {
            if i == j {
                break
            }
            if a + b == 2020 {
                return a * b, nil
            }
        }
   }
   return 0, errors.New("No two entries sum to 2020")
}

func part2(input []int) (int, error) {
    for k, a := range input {
        for j, b := range input {
            if j == k {
                break
            }
            for i, c := range input {
                if i == j {
                    break
                }
                if a + b + c == 2020 {
                    return a * b * c, nil
                }
            }
        }
   }
   return 0, errors.New("No three entries sum to 2020")
}

func main() {
    var input []int
    for {
        var entry int
        if _, err := fmt.Scanf("%d\n", &entry); err == nil {
            input = append(input, entry)
        } else {
            break
        }
    }

    part1, _ := part1(input)
    part2, _ := part2(input)

    fmt.Println("The product of the two values:", part1)
    fmt.Println("The product of the three values:", part2)
}