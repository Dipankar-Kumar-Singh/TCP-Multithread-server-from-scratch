package main

import (
	"fmt"
	"log"
	"net"
	"time"
);

func do(conn net.Conn) {
	fmt.Println(conn)
	// conn -> a connection object

	buf := make([]byte, 1024) // 1kb buffer ( byte array )

	noOfBytesRead , err := conn.Read(buf) // read from the connection
	if err != nil {
		log.Fatal(err)
	}
	// reading is a blocking system call , it will wait until the data is available

	time.Sleep(1 * time.Second) // sleep for 1 second

	fmt.Println("No of bytes read : " , noOfBytesRead);


	string message = "Hello from server" ;
	// conn.Write([]byte("Hello from server"))
	 // write to the connection( But it will not send the data to the client until the buffer is full or the connection is closed)
	 // and for curl request , it is not good , as curl is waiting for an HTTP response

	 // Making the same message as a HTTP response
	 conn.Write([]byte("HTTP/1.1 200 OK\r\n\rHello, World!\r\n"))



}

func main() {
	listener , err := net.Listen("tcp", ":8080") ;
	if(err != nil){
		log.Fatal(err) 
	}

	// accept connection on port --> A blocking system call 
	conn , err := listener.Accept() ;
	if(err != nil){
		log.Fatal(err) 
	}

	fmt.Println(conn)
};