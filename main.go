package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	fmt.Println(conn)
	// conn -> a connection object

	buf := make([]byte, 1024) // 1kb buffer ( byte array )

	noOfBytesRead, err := conn.Read(buf) // read from the connection
	if err != nil {
		log.Fatal(err)
	}
	// reading is a blocking system call , it will wait until the data is available

	time.Sleep(8 * time.Second) // sleep for x second

	fmt.Println("No of bytes read : ", noOfBytesRead)

	// message := "Hello from serv	er"
	// conn.Write([]byte("Hello from server"))
	// write to the connection( But it will not send the data to the client until the buffer is full or the connection is closed)
	// and for curl request , it is not good , as curl is waiting for an HTTP response

	// Making the same message as a HTTP response
	// conn.Write([]byte("HTTP/1.1 200 OK\r\n\rHello, World!\r\n"))
	log.Println("started Processing the request") ;
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	log.Println("Finished Processing the request") ;

	// HTTP/1.1 200 OK\r\n\r\n --> HTTP response header
	// Hello, World!\r\n --> HTTP response body
	// HTTP response header is terminated by \r\n\r\n
	// HTTP response body is terminated by \r\n

	conn.Close() // close the connection

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Initial Problem : It is not able to handle multiple requests at the same time 
	// Solution : Goroutines
	// Goroutines : It is a lightweight thread managed by the Go runtime
	// Goroutines are not OS threads , they are managed by the Go runtime
	// Goroutines are multiplexed to fewer number of OS threads


	// acting as a sever due to continuous listening on port 8080 for incoming connections
	for {
		// accept connection on port --> A blocking system call
		conn, err := listener.Accept()
		// listener.Accepts what does it do ?
		// It will wait until a client connects to the server
		// It will return a connection object

		log.Println("Connection accepted") ;

		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(conn)
		go do(conn)
	}

}
