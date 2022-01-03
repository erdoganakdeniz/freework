package utils

import (
	"fmt"
	"freework/store"
	"os"
	"time"
)

var file = "tmp/TIMESTAMP-data.gob"

func SavetoFile(s *store.Store) {
	_, err := os.Create(file)
	if err != nil {
	}
	s.SaveFile(file)

}
func SaveInterval(s *store.Store) {
	for range time.Tick(20 * time.Minute) {
		s.SaveFile(file)
		fmt.Println("Automatic Save to File")
	}

}
func LoadfromFile(s *store.Store) {
	_, err := os.Open(file)
	if err == nil {
		err := s.LoadFile(file)
		if err != nil {
			fmt.Println("Error : Load From File")
		}

	}

}
