package main

func BoolToWord(word bool) string {
  var answer string
  if word == true {
    answer = "Yes"
  } else if word == false {
    answer = "No"
  }
  
  return answer
}