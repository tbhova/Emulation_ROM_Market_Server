package main

import (
	"os/exec"
	"fmt"
	"sync"
)

func flywayMigration(wg *sync.WaitGroup) {
	fmt.Println("Migration migrate")
	out, err := exec.Command("flyway", "migrate").Output()
	
	if err != nil {
		fmt.Println("Error occured running flyway migration")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	
	wg.Done()
}

func initialize() {
	fmt.Println("Migration init")
	wg := new(sync.WaitGroup)
	wg.Add(1)
	flywayMigration(wg)
	
	wg.Wait()
}

func main() {
	initialize()
	fmt.Println("Migration ran")
}
