package main

import (
	"bufio"
	"fmt"
	"os"
)




func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your address: ")
    address, err := reader.ReadString('\n')
	fmt.Errorf("%v", err)
	
	weather.fetchWeather(address)

}