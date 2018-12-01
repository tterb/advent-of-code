package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString(' ')
	str_data := strings.Fields(input)
	var data = []int{}
	for _, j := range str_data {
		k, err := strconv.Atoi(j)
    if err != nil {
	    panic(err)
    }
    data = append(data, k)
	}
	i, steps := 0, 0
	for(i < len(data)) {
		temp := i
		i += data[i]
		data[temp] += 1
		steps += 1
	}
	fmt.Println(steps)
}
