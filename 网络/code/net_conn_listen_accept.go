//https://gist.github.com/hyper0x/8f724925c344f896b63c
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const (
	DELIMITER byte = '\n'
	QUIT_SIGN      = "quit!"
)

func Read(conn net.Conn, delim byte) (string, error) {
	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer
	for {
		ba, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		buffer.Write(ba)
		if !isPrefix {
			break
		}
	}
	return buffer.String(), nil
}

func Write(conn net.Conn, content string) (int, error) {
	writer := bufio.NewWriter(conn)
	number, err := writer.WriteString(content)
	if err == nil {
		err = writer.Flush()
	}
	return number, err
}

func main() {
	go func() {
		listener, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Printf("Listener: Listen Error: %s\n", err)
			os.Exit(1)
		}
		log.Println("Listener: Listening...")
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Listener: Accept Error: %s\n", err)
				continue
			}
			go func(conn net.Conn) {
				defer conn.Close()
				for {
					log.Println("Listener: Accepted a request.")
					log.Println("Listener: Read the request content...")
					content, err := Read(conn, DELIMITER)
					if err != nil {
						log.Printf("Listener: Read error: %s", err)
					}
					if content == QUIT_SIGN {
						log.Println("Listener: Quit!")
						break
					}
					log.Printf("Listener: Received content: %s\n", content)
					respContent := fmt.Sprintf("listener response: %s%c", content, DELIMITER)
					log.Printf("Listener: the response content: %s\n", respContent)
					num, err := Write(conn, respContent)
					if err != nil {
						log.Printf("Listener: Write Error: %s\n", err)
					}
					log.Printf("Listener: Wrote %d byte(s)\n", num)
				}
			}(conn)
		}
	}()
	time.Sleep(time.Millisecond * 500)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		conn, err := net.DialTimeout("tcp", "127.0.0.1:9090", time.Millisecond*200)
		if err != nil {
			log.Printf("Sender: DialTimeout Error: %s\n", err)
			os.Exit(1)
		}
		log.Println("Sender: Dial OK.")
		for i := 0; i < 10; i++ {
			reqContent := fmt.Sprintf("sender request %d.%c", i, DELIMITER)
			log.Printf("Sender: the request content: %s\n", reqContent)
			num, err := Write(conn, reqContent)
			if err != nil {
				log.Printf("Sender: Write Error: %s\n", err)
				break
			}
			log.Printf("Sender: Wrote %d byte(s)\n", num)
			respContent, err := Read(conn, DELIMITER)
			if err != nil {
				log.Printf("Sender: Read error: %s", err)
				break
			}
			log.Printf("Sender: Received content: %s\n", respContent)
		}
		reqContent := fmt.Sprintf("%s%c", QUIT_SIGN, DELIMITER)
		log.Printf("Sender: the request content: %s\n", reqContent)
		num, err := Write(conn, reqContent)
		if err != nil {
			log.Printf("Sender: Write Error: %s\n", err)
		}
		log.Printf("Sender: Wrote %d byte(s)\n", num)
	}()
	wg.Wait()
	time.Sleep(time.Millisecond * 500)
}