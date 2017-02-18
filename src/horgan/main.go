package main

import (
	"os/exec"
	"fmt"
	"sync"
)

func flywayMigration(wg *sync.WaitGroup) {
	out, err := exec.Command("flyway", "migrate").Output()
	
	if err != nil {
		fmt.Println("Error occured running flyway migration")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	
	wg.Done()
}

func initialize() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go flywayMigration(wg)
	
	wg.Wait()
}

func main() {
	initialize()
}
