package main

import (
	// "bufio"
	"fmt"
	"os"
	"bytes"
	"io"
	// "strings"
	// "encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {																																																										
	if e != nil {																																																												
		panic(e)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "indexserver.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, errors = upgrader.Upgrade(w, r, nil)
		if errors != nil {
			fmt.Println(errors)
		}

		go func(conn *websocket.Conn) {

			fs, err := os.Open("bike.mp4")
			if err != nil {
				log.Fatalln(err)
			}
			defer fs.Close()

			var buf bytes.Buffer

			data := make([]byte, 1024)

			flag := false

			go func() { // Read from video file in buffer
				for {
					n, err := bufio.NewReader(fs).Read(data)
					n,err = buf.Write(data)
					// conn.WriteMessage(2, data)
					if err == io.EOF {
						flag = true
						break
					} else if err != nil {
						log.Fatalln(err)
						return
					}
					if n == 0 {
						break
					}
				}
			}()
			go func() {
				for {
					readdata := make([]byte, 1024)
					n1, err1 := buf.Read(readdata)
					conn.WriteMessage(2, readdata)

					if err1 == io.EOF && flag {
						fmt.Println("Err1 EOF : ",err1)

						break
					} else if err1 != nil && flag {
						fmt.Println("Err1 err1 != nil : ",err1)
						log.Fatalln(err1)
						return
					}
					if n1 == 0 && flag {
						fmt.Println("Err1 : all done  ",err1)
						break
					}
				}
			}()
			// for {
			// 	 n1,err1 := bufio.NewWriter(conn).Write(data)
			// }

		}(conn)
	})

	fmt.Println("Live on :3000")
	http.ListenAndServe(":3000", nil)
}

//gofmt -w readvideofiles.go
