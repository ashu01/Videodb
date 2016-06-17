package main

import (
	"fmt"
	"os"
	// "bufio"
	// "bytes"
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


func writeoperation(buffer []byte, conn *websocket.Conn){
	conn.WriteMessage(2, buffer)	
}


func readoperation(conn *websocket.Conn){
	fs, err := os.Open("bike.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()
	buffer := make([]byte, 1024)
	for {
		buffer = buffer[:cap(buffer)]
		n,err := fs.Read(buffer)  //read data's bytes and error if exists
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
			return
		}
		if n == 0 {
			break
		}
		buffer = buffer[:n]
		go writeoperation(buffer,conn)	
	}
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
			

			go func(conn *websocket.Conn){	

				go readoperation(conn) 				
				// go writeoperation(conn)

		}(conn)		
	})

	fmt.Println("Live on :3000")
	http.ListenAndServe(":3000", nil)	
}