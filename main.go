package main

import (
	"time"
)

func init() {
	OpenConnection()
}

func main() {
	defer DB.Close()
	go func() {
		for {
			RemoveExpiredURLs()
			time.Sleep(1 * time.Hour)
		}
	}()
	Routes()

}
