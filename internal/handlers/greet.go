package handlers

import (
	"fmt"
	"net/http"
	"os"
)

type GreetHandler struct {
	name string
}

func NewGreetHandler(name string) *GreetHandler {
	return &GreetHandler{name: name}
}

func (g *GreetHandler) BasicHandler(w http.ResponseWriter, r *http.Request) {
	greet := []byte("Hello " + g.name)
	_, err := w.Write(greet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problems with writing: %v\n", err)
	}
}
