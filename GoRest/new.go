package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Types
type User struct {
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

//to show all available users
type allUsers []User

var Users = allUsers{
	{
		ID:          1,
		Name:        "User One",
		Description: "Some Description",
	},
}
var addr = flag.String("addr", "localhost:5000", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	newUser.ID = len(Users) + 1
	Users = append(Users, newUser)

	w.Header().Set("Description-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Description-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	for _, User := range Users {
		if User.ID == UserID {
			w.Header().Set("Description-Type", "application/json")
			json.NewEncoder(w).Encode(User)
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserID, err := strconv.Atoi(vars["id"])
	var updatedUser User

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedUser)

	for i, t := range Users {
		if t.ID == UserID {
			Users = append(Users[:i], Users[i+1:]...)

			updatedUser.ID = t.ID
			Users = append(Users, updatedUser)

			// w.Header().Set("Description-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedUser)
			fmt.Fprintf(w, "The User with ID %v has been updated successfully", UserID)
		}
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	for i, t := range Users {
		if t.ID == UserID {
			Users = append(Users[:i], Users[i+1:]...)
			fmt.Fprintf(w, "The User with ID %v has been remove successfully", UserID)
		}
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/Users", createUser).Methods("POST")
	router.HandleFunc("/Users", getUsers).Methods("GET")
	router.HandleFunc("/Users/{id}", getOneUser).Methods("GET")
	router.HandleFunc("/Users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/Users/{id}", updateUser).Methods("PUT")

	flag.Parse()
	log.SetFlags(0)
	router.HandleFunc("/echo", echo)
	router.HandleFunc("/chat", home)
	log.Fatal(http.ListenAndServe(*addr, router))
	//log.Fatal(http.ListenAndServe(*addr, nil))

	fmt.Println("Server running in Port 5000")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
