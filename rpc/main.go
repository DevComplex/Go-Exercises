package main

import (
	"net"
	"net/http"
	"net/rpc"
	"log"
)

type Item struct {
	Title string
	Body string
}

var database []Item

type API int

func main() {
	api := new(API)

	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}
}

func (a *API) GetByName(Title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == Title {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}

func (a *API) CreateItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = Item{edit.Title, edit.Body}
			changed = edit
			break
		}
	}

	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var deleted Item

	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			deleted = item
			break
		}
	}

	*reply = deleted
	return nil
}

func (a *API) GetDB(Title string, reply *[]Item) error {
	*reply = database
	return nil
}