package visca_cli

import (
	"fmt"
	"flag"
	"log"
	"github.com/jacobsa/go-serial/serial"
)


func visca_cli() {


	// Find Visca cam.
	// Create some kind of communication bridge to something else



	// examples of using the Flag lib
	textPtr := flag.String("text", "", "Text to parse.")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")

	flag.Parse()
	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)


	// Set up options.
		options := serial.OpenOptions{
		PortName: "/dev/tty.usbserial-A8008HlV",
		BaudRate: 19200,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,
		}
	
		// Open the port.
		port, err := serial.Open(options)
		if err != nil {
		log.Fatalf("serial.Open: %v", err)
		}
	
		// Make sure to close it later.
		defer port.Close()
	
		// Write 4 bytes to the port.
		b := []byte{0x00, 0x01, 0x02, 0x03}
		n, err := port.Write(b)
		if err != nil {
		log.Fatalf("port.Write: %v", err)
		}
	
		fmt.Println("Wrote", n, "bytes.")
}