package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	bin := NewBin("first-bin", false)
	fmt.Printf("Bin has been created: ID=%s, Name=%s, Private=%t, CreatedAt=%v", bin.ID, bin.Name, bin.Private, bin.CreatedAt)

	bins := []Bin{
		{ID: "1", Name: "Alice", Private: false, CreatedAt: time.Now()},
		{ID: "2", Name: "Bob", Private: true, CreatedAt: time.Now()},
	}
	binList := NewBinList(bins)
	fmt.Printf("Binlist: %d, elements\n", len(binList))
}

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}
type BinList []Bin

func NewBin(name string, private bool) *Bin {
	return &Bin{
		ID:        generateID(),
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
func generateID() string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}
func NewBinList(bins []Bin) BinList {
	return BinList(bins)
}
