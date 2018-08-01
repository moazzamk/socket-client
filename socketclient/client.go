package socketclient

import (
	"net"
	"fmt"
	"io"
	"time"
	"bytes"
)

// GetResponse connects to a server through sockets and returns the response
func GetResponse(host string, input string) (string, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return ``, err
	}

	defer conn.Close()

	t := time.Time{}
	t.Add(30 * time.Second)

	conn.SetDeadline(t)
	conn.Write([]byte(input ))
	//fmt.Fprintf(conn, input)

	var buf bytes.Buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		fmt.Println(string(tmp[:n]), n)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}

			break
		}
		buf.Write(tmp)
	}

	fmt.Println("total size:", buf.Len())

	return buf.String(), nil
}
