//go handlers with timeouts and graceful shutdowns
//using standard Library
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
	"./data"
	"strconv"
	"regexp"
	
	
)
//Defining Structs
type Hello struct {
	l *log.Logger
}

type GoodBye struct {
	l *log.Logger
}

type Products struct {
	l *log.Logger
}


//defining Functions to return its func
func NewHello(l *log.Logger) *Hello{
	return &Hello{l}

}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}



//Functions
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

//returning product lists
func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request){
	//with standard library
	if h.Method == http.MethodGet{
		p.getProducts(rw, h)
		return
	}

	if h.Method == http.MethodPost{
		p.postProducts(rw,h)
		return
	}
	if h.Method == http.MethodPut{
		//get id using Regex
		p.l.Println("PUT", h.URL.Path)
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(h.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one ID")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one cap grp")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)

		}
		idString := g[0][1]
		id, err:= strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to number", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, h)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
	
}

//GET
func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request){
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil{
	http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
	}
}
//POST
func (p *Products) postProducts(rw http.ResponseWriter, h *http.Request){
	//p.l.Println("Handle Post Products")
	
	prod := &data.Product{}
	err := prod.FromJSON(h.Body)
	if err != nil{
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	
	data.PostProducts(prod)

}

//PUT
func (p *Products) updateProducts(id int, rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJSON(h.Body)
	if err != nil{
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProdNotFound{
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil{
		http.Error(rw, "Product not found", http.StatusBadRequest)
		return
	}
	data.UpdateProduct(id, prod)
}

//Our Main Function
func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//hh := NewHello(l)
	gh := NewGoodBye(l)

	ph := NewProducts(l)

	sm :=http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/goodbye", gh)
	sm.Handle("/products", ph)

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