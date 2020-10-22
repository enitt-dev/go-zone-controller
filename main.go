package main

import (
	"log"

	"github.com/goburrow/modbus"
)

func bytesToIntArr(b []byte) (result []int, err error) {

	var tmp int
	for i, value := range b {
		if i%2 == 0 { // even
			tmp = int(value) * 256
		} else { // odd
			tmp += int(value)

			result = append(result, tmp)
		}
	}
	return
}

func main() {
	client := modbus.TCPClient("localhost:502")

	results, err := client.ReadHoldingRegisters(0, 4)
	if err != nil {
		log.Println(err)
	}

	result, err := bytesToIntArr(results)
	if err != nil {
		log.Println(err)
	}

	log.Println(result)
}
