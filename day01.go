package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

func get_fuel(mass int) int {
    fuel := (mass / 3) - 2
    if fuel < 0 {
        fuel = 0
    }

    return fuel
}

func solve1(contents string) int {
    lines := strings.Split(contents, "\n")

    sum := 0
    for _, line := range lines {
        num := 0

        _, err := fmt.Sscanf(line, "%d", &num)
        if err == nil {
            sum += get_fuel(num)
        }
    }

    return sum
}

func solve2(contents string) int {
    lines := strings.Split(contents, "\n")

    sum := 0
    for _, line := range lines {
        num := 0

        _, err := fmt.Sscanf(line, "%d", &num)
        if err == nil {
            module_sum := 0
            for fuel := get_fuel(num); fuel > 0; fuel = get_fuel(fuel) {
                module_sum += fuel
            }

            sum += module_sum
        }
    }

    return sum
}

func main() {
    for _, arg := range os.Args[1:] {
        fmt.Printf("%s:\n", arg)

        file, err := os.Open(arg)
        if err != nil {
            log.Fatal(err)
        }

        defer func() {
            err := file.Close()
            if err != nil {
                log.Fatal(err)
            }
        }()

        contents, err := io.ReadAll(file)
        if err != nil {
            log.Fatal(err)
        }

        result1 := solve1(string(contents[:]))
        fmt.Printf("  1: %d\n", result1)

        result2 := solve2(string(contents[:]))
        fmt.Printf("  2: %d\n", result2)
    }
}
