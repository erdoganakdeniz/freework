package utils

import (
	"fmt"
	"freework/cache"
	"os"
	"time"
)

var file = "tmp/TIMESTAMP-data.gob"

func SavetoFile(c *cache.Cache) {
	_, err := os.Create(file)
	if err != nil {
		fmt.Println("Error : Save to File")
	}
	c.SaveFile(file)
}
func SaveInterval(c *cache.Cache) {
	for range time.Tick(10 * time.Minute) {
		c.SaveFile(file)
		fmt.Println("Automatic Save to File")
	}
	time.Sleep(10 * time.Second)

}
func LoadfromFile(c *cache.Cache) {
	_, err := os.Open(file)
	if err == nil {
		err := c.LoadFile(file)
		if err != nil {
			fmt.Println("Error : Load From File")
		}

	}

}
