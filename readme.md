# TCP: Transmission Control Protocol

It is a widely used protocol for sending and receiving data over the internet.
TCP uses **sockets** to establish and maintain connections between devices.
When a TCP connection is established, each device creates a socket that is associated with that connection. The devices use their sockets to send and receive data over the connection.

#### What I am trying to do  : 
Playing out with Low-Level Handling of TCP Request and Processing it
Figuring out how does the sever work ? 🤔
How does a web server keep on running and accepting multiple request  ? 🤔

Language used : **Go**
and using **CURL** for sending testing and sending **TCP packets** ( as CURL is used for HTTP) and HTTP uses TCP , so CURL will establish TCP connection with our Listener ( on port 8080 )

In side main function  : 
```GO
listener, err := net.Listen("tcp", ":8080")
// Listening on port 8080 for any TCP Connection 

conn, err := listener.Accept()  // conn : connection object 
// listener.Accepts what does it do ?
// accept connection on port --> A blocking system call
// It will wait until a client connects to the server
// It will return a connection object
// without accpeting , program will not connect with the request sender and exit.

processIt(conn) ; // for doing some oprations // user defined Function 
```

what is net ?  : In Go programming language, the "net" package provides support for various network-related functionality, such as creating TCP/UDP connections, implementing servers, performing DNS lookups, and more.

```Go
// Implementation of processIt() function
func processIt(conn net.Conn) {
	buf := make([]byte, 1024) // 1kb buffer ( byte array )
	noOfBytesRead, err := conn.Read(buf) // read from the connection
	// Read is a blocking system call , it will wait until the data is available
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	// HTTP/1.1 200 OK\r\n\r\n --> HTTP response header
    // Hello, World!\r\n --> HTTP response body
    // HTTP response header is terminated by \r\n\r\n
    // HTTP response body is terminated by \r\n
    conn.Close() // close the connection
}
```

```Bash
curl http://localhost:8080
// A Connection will be made when Connection get Accepted 
// data can be send and recived via this connection 
// and it will be Listened by our TCP Listener on port 8080
```

# What happens when multiple curl request are made at once ?
##### Yes I am talking about Multithreading 🧵🤹‍♂️  🙂 . 

![Behavior on Multiple request -processing one another_Synchronous](https://user-images.githubusercontent.com/66475186/233722548-bd52a92c-7fc2-45cf-aea8-5824d4ca7f3d.png)


