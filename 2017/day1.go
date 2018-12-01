package main

import (
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString(' ')
  input = strings.TrimSuffix(input, "\n")
  arr := strings.Split(input, "")
  data := []int{}
  // convert to int array
  for _, i := range arr {
    k, err := strconv.Atoi(i)
    if err != nil {
      panic(err)
    }
    data = append(data, k)
  }
  var sum int
  temp := data[len(data)-1]
  for _, i := range data {
    if temp == i {
      sum += i
    }
    temp = i
  }
  print(sum)
}
