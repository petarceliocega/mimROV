package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os/exec"
	"strings"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	l, err := net.Listen("tcp", ":8085") // Listening on port for incoming connections

	fmt.Println("█▀▀▄░░░░░░░░░░░▄▀▀█")
	fmt.Println("░█░░░▀▄░▄▄▄▄▄░▄▀░░░█")
	fmt.Println("░░▀▄░░░▀░░░░░▀░░░▄▀")
	fmt.Println("░░░░▌░▄▄░░░▄▄░▐▀▀")
	fmt.Println("░░░▐░░█▄░░░▄█░░▌▄▄▀▀▀▀█")
	fmt.Println("░░░▌▄▄▀▀░▄░▀▀▄▄▐░░░░░░█")
	fmt.Println("▄▀▀▐▀▀░▄▄▄▄▄░▀▀▌▄▄▄░░░█")
	fmt.Println("█░░░▀▄░█░░░█░▄▀░░░░█▀▀▀")
	fmt.Println("░▀▄░░▀░░▀▀▀░░▀░░░▄█▀")
	fmt.Println("░░░█░░░░░░░░░░░▄▀▄░▀▄")
	fmt.Println("░░░█░░░░░░░░░▄▀█░░█░░█")
	fmt.Println("░░░█░░░░░░░░░░░█▄█░░▄▀")
	fmt.Println("░░░█░░░░░░░░░░░████▀")
	fmt.Println("░░░▀▄▄▀▀▄▄▀▀▄▄▄█▀")
	fmt.Println("-----------------------")
	fmt.Println("-| Listening@: 8085  |-")
	fmt.Println("-----------------------")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Handling connections (-(-.(-.(-.-).-).-)-)")
		go handleConnection(c)
	}
}

// Handling client connections
func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		commandGPIO(temp)

		// Sends command we got from client back at them so we can confirm correct transmission client->server
		rezultat := time.Now().Format("2006-01-02 15:04:05.000") + " : Server obradio : " + temp + "\n"
		c.Write([]byte(string(rezultat)))
	}
	c.Close()
}

func commandGPIO(commandr string) {
	errgpio := rpio.Open()
	defer rpio.Close()
	if errgpio != nil {
		fmt.Println("Problem pri otvaranju GPIO: ", errgpio.Error())
	}

	pin18 := rpio.Pin(18)
	pin19 := rpio.Pin(19)
	pin26 := rpio.Pin(26)
	pin18.Output()
	pin19.Output()
	pin26.Output()

	switch commandr {
	case "r": // To use LED PWM dimming, sudo should be used
		fmt.Println(" : Received r: pin26=Dimming")
		pin26.Mode(rpio.Pwm)
		pin26.Freq(64000)
		pin26.DutyCycle(0, 32)

		// Blinking a LED using 2000Hz frequency with PWM
		// (source frequency divided by working cycle => 64000/32 = 2000)
		for i := 0; i < 5; i++ {
			for i := uint32(0); i < 32; i++ { // Increasing brightness
				pin26.DutyCycle(i, 32)
				time.Sleep(time.Second / 32)
			}
			for i := uint32(32); i > 0; i-- { // Lowering brightness
				pin26.DutyCycle(i, 32)
				time.Sleep(time.Second / 32)
			}
			fmt.Println("Cycle completed: ", i)
		}
	case "w": // Turn on LED at pin18 and turning off at pin19
		fmt.Print(time.Now().Format("2006-01-02 15:04:05.000000"))
		fmt.Println(" : Received w: pin18=HIGH | pin19=LOW")
		pin19.Low()
		pin18.High()
	case "s": // Turn on LED at pin19 and turn off at pin18
		fmt.Print(time.Now().Format("2006-01-02 15:04:05.000000"))
		fmt.Println(" : Received s: pin18=LOW | pin19=HIGH")
		pin18.Low()
		pin19.High()
	case "a": // Turn off LED at pin18 and pin19
		fmt.Print(time.Now().Format("2006-01-02 15:04:05.000000"))
		fmt.Println(" : Received a: svi upravljacki pinovi na LOW")
		pin18.Low()
		pin19.Low()
	case "p": // Starting video stream using bash script in console
		cmdStartStream := exec.Command("bash", "start_stream.sh")
		if err := cmdStartStream.Start(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Stream start")
	case "o":
		cmdStopStream := exec.Command("bash", "stop_stream.sh")
		// Turning off video stream using bash script in console
		if err := cmdStopStream.Start(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Stream end")
	}
}
