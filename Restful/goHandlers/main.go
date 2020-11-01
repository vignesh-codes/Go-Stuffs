
//go handlers with timeouts and graceful shutdowns

package main

import (
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"os"
	"os/signal"
	"context"
)

type Hello struct {
	l *log.Logger
}

type GoodBye struct {
	l *log.Logger
}


func NewHello(l *log.Logger) *Hello{
	return &Hello{l}

}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}



func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(rw, "Opps", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s \n", d)
}


func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	rw.Write([]byte("Bye"))
	g.l.Println("Good Bye World")
}


//Our Main Function
func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := NewHello(l)
	gh := NewGoodBye(l)

	sm :=http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)


	//Defining Timeouts
	s := &http.Server{
		Addr:":9000",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()

	sigChan := make (chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved Terminate, Gracefully ShutDown", sig)
	

	//graceful shutdown
	tc,_:= context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}