/*****************************************************************************
 * Name: Ted Kang
 * NetId: tikang
 *
 * Description: A simple client using Go Net API.  Sends up to
 *              SEND_BUFFER_SIZE bytes or EOF.   Limited to IP address
 *              of server (e.g., 10.0.2.15  or 127.0.0.1) instead of domain
 *              name (e.g., mycomputer.cs.princeton.edu or localhost)
 *
 * Usage:       Usage: ./client-g [server IP] [server port] < [message]
 *              (Assumes server is running.)
 *
 * Example:     ./client-g 127.0.0.1 8999 < somefile.txt
 *
 *****************************************************************************/

package main

import (
	"os"
	"bufio"
)
const SEND_BUFFER_SIZE = 2048


func main() {
	// get the server's IP and port from the command line
		serverIP
        for i, arg := range os.Args[1:] {
        	if i == 0 {

        	}
        }

	// use net.Dial to connect ip and port, with the specified protocol.
        connection, err := net.Dial("tcp4", "127.0.0.1:8080")

	// Wrap the connection in bufio.Reader and get input from stdin
	reader := bufio.NewReader(os.Stdin)
	message := make([]byte, SEND_BUFFER_SIZE)
	bytes_read, err := reader.Read(message)

	// net.Conn.Write writes to the connection
        _______________

	// net.Conn.Close closes to the connection
        _______________

}
