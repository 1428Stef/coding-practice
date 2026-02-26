package main

func splitInt(n int) []int {
    if n == 0 {
        return []int{0}
    }
    var digits []int
    for n > 0 {
        digits = append([]int{n % 10}, digits...)
        n /= 10
    }
    return digits
}

func DigitalRoot(n int) int {
    if n == 0 {
        return 0
    }
    
    value := n
    for {
        sli := splitInt(value)
        if len(sli) == 1 {
            return sli[0]
        }
        value = 0
        for _, i := range sli {
            value += i
        }
    }
}