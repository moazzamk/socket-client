package socketclient

import (
	"net"
	"fmt"
	"io"
	"bytes"
	"log"
)

func GetResponse(host string, input string) (string, error) {
	log.Println(`host`, `'`  + host + `'`)
	log.Println(`input`, input)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Println(`errrrr`)
		return ``, err
	}

	defer conn.Close()

	fmt.Fprintf(conn, input + "\r\n\r\n")

	var buf bytes.Buffer
	var s []byte
	log.Println(conn.Read(s))

	_, err = io.Copy(&buf, conn)
	if err != nil {
		log.Println(`hiiiiii`)
		return ``, err
	}

	return buf.String(), nil
}
