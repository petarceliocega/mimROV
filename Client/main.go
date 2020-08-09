package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer keyboard.Close()

	addrServ := "169.254.106.87" // IP of Raspberry Pi Server
	portIP := "8085"
	remIPAddrPort := addrServ + ":" + portIP
	fmt.Println(remIPAddrPort)
	// Spajanje na socket
	conn, err := net.Dial("tcp4", remIPAddrPort) // Connection to RPi Server
	if err != nil {
		fmt.Println(err)
	} else {
		// Main loop until Esc key is pressed
		for {
			// Reads stdin from console
			fmt.Println("------------------------------------------------")
			fmt.Print("Send to server: ")
			char, key, _ := keyboard.GetSingleKey()
			sc := string(char)
			fmt.Printf("Sending server: %q\r\n", char)
			if err != nil {
				panic(err)
			} else if key == keyboard.KeyEsc {
				break
			}

			timeNow := time.Now().Format("2006-01-02 15:04:05.000")
			fmt.Println(timeNow + " : Sent to server: " + sc)
			fmt.Fprintf(conn, sc+"\n") // Sends sc string on socket

			// Listening on socket for reply
			serverMessage, _ := bufio.NewReader(conn).ReadString('\n') // Saves returned message
			serverMessage = strings.TrimSuffix(serverMessage, "\n")    // Removes newline from returned message
			fmt.Println("Server returned: ", serverMessage)

			fmt.Println("------------------------------------------------")
			fmt.Println("")
		}
	}
}
