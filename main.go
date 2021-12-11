package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mvogttech/go-weather-cli/internal/weather"
)




func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your address: ")
    address, err := reader.ReadString('\n')
	fmt.Errorf("%v", err)
	
	weather.FetchWeather(address)

}