package kata

import (
    "strconv"
    "strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
    var builder strings.Builder
    builder.Grow(14) 
    
    builder.WriteByte('(')
    for i := 0; i < 3; i++ {
        builder.WriteString(strconv.Itoa(int(numbers[i])))
    }
    builder.WriteString(") ")
	
    for i := 3; i < 6; i++ {
        builder.WriteString(strconv.Itoa(int(numbers[i])))
    }
    builder.WriteByte('-')

    for i := 6; i < 10; i++ {
        builder.WriteString(strconv.Itoa(int(numbers[i])))
    }

    return builder.String()
}
