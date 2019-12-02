package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)
/**
	https://blog.csdn.net/yjp19871013/article/details/82316299
 */
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	serverAddr := "127.0.0.1:10006"
	conn, err := net.Dial("udp", serverAddr)
	checkError(err)

	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()

		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen-written > 512 {
				toWrite = line[written : written+512]
			} else {
				toWrite = line[written:]
			}
			toWrite = toWrite
			n, err = conn.Write([]byte(toWrite))
			checkError(err)

			fmt.Println("Write:", toWrite)

			msg := make([]byte, 512)
			n, err = conn.Read(msg)
			checkError(err)

			fmt.Println("Response:", string(msg))
		}
	}
}

