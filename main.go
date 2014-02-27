package main

import (
	"code.google.com/p/go.net/websocket"
	//"encoding/json"
	//	"github.com/howeyc/fsnotify"
	//	"io/ioutil"
	"log"
	"net/http"
	//	"os"
	//	"os/exec"
	//	"path/filepath"
	//	"sync"
)

const dart2js = "c:\\dart\\dart-sdk\\bin\\dart2js.bat"

var filesForWatch = [...]string{
	`d:\Silent\Golang\Projects\GoGame_v1\index.css`,
	`d:\Silent\Golang\Projects\GoGame_v1\index.dart`,
	`d:\Silent\Golang\Projects\GoGame_v1\index.html`,
	`d:\Silent\Golang\Projects\GoGame_v1\main.go`}

var needReload chan bool = make(chan bool)

/*
func init() {
	//сгенерим все *.dart.js заново
	//установим нотификатор
	var wg sync.WaitGroup
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Println("Init error")
		os.Exit(0)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".dart" {
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()
				recompile(filename)
			}(file.Name())
		}
	}
	wg.Wait()
	log.Println("Server start success")
}
*/

func main() {
	/*
		watcher, _ := fsnotify.NewWatcher()
		for _, file := range filesForWatch {
			watcher.Watch(file)
		}

		//watcher.Watch("./") //следим за текущим каталогом
		go func() {
			for {
				select {
				case ev := <-watcher.Event:
					log.Println("event:", ev)
					if filepath.Ext(ev.Name) == ".dart" {
						recompile(ev.Name)
						needReload <- true
					}
					needReload <- true
				case err := <-watcher.Error:
					log.Println("error:", err)
				}
			}
		}()
	*/
	//основная часть
	http.Handle("/ws", websocket.Handler(WebSocketServer))
	//	http.Handle("/needReload", websocket.Handler(NeedReload))
	http.HandleFunc("/", mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:]
	//log.Println(url)
	http.ServeFile(w, r, url)
}

/*
func recompile(filename string) {
	cmd := exec.Command(dart2js, "-o", filename+".js ", filename)
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Println(err)
		res, _ := cmd.Output()
		log.Println(string(res))
	} else {
		log.Println(filename, "recompile success")
	}

}

func NeedReload(ws *websocket.Conn) {
	for {
		select {
		case <-needReload:
			ws.Write([]byte("needReload")) //можно отправить любые данные, важен лишь факт отправки
		}
	}
}
*/

func WebSocketServer(ws *websocket.Conn) {
	for {
		var in []byte
		err := websocket.Message.Receive(ws, &in)
		log.Printf("Receive: %v\n", string(in))
		if err != nil {
			log.Println(err)
			ws.Close()
			return
		}
		//msg, _ := json.Marshal(struct{ id, value string }{"0", "pong"})
		msg := "pong"
		if string(in) == "ping" {
			websocket.JSON.Send(ws, msg)
		} else {
			websocket.JSON.Send(ws, in)
		}
	}
}
