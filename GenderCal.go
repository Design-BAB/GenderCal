package main

import (
  "fmt"
  "math"
)
func percent(total int, gender int) float64 {
  genderP := float64(gender) / float64(total)
  genderP = genderP * 100
  return math.Round(genderP*10) / 10
}

func main()  {
  var boys int
  var girls int
  
  fmt.Println("部屋に男がどのぐらいですか？")
  fmt.Scanln(&boys)
  fmt.Println("そして、女がどのぐらいですか？")
  fmt.Scanln(&girls)
  fmt.Println("ありがとう！\n")
  total := boys + girls
  fmt.Println("🧔　: ", percent(total, boys), "%\n")
  fmt.Println("👧　: ", percent(total, girls), "%\n")
  fmt.Println("今、部屋に", total, "人います。")
}
