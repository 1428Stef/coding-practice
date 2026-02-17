package main

func ReverseSeq(n int) []int {
    result := []int{} 
    for i := n; i >= 1; i-- { 
        result = append(result, i)  
    }
    return result  
}

func main() {
    ReverseSeq(5)  
}
