package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var hashes = [3]string{
	"1115dd800feaacefdf481f1f9070374a2a81e27880f187396db67958b207cbad",
	"3a7bd3e2360a3d29eea436fcfb7e44c735d117c42d1c1835420b6b9942dd4f1b",
	"74e1bb62f8dabb8125a58852b63bdf6eaef667cb56ac7f7cdba6d7305c50a22f",
}

func compareHash(hash []byte, password string) {
	for i := 0; i < len(hashes); i++ {
		if hashes[i] == hex.EncodeToString(hash[:]) {
			fmt.Println("password is: ", password)
		}
	}
}

func hashPassword(w *sync.WaitGroup, string string) {
	defer w.Done()
	data := []byte(string)
	hash := sha256.Sum256(data)
	compareHash(hash[:], string)
}

func file() {
	var waitGroup sync.WaitGroup

	file, err := os.Open("passwords.txt")
	if err != nil {
		fmt.Println("Smth goes wrong: ", err)
		os.Exit(1)
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	fmt.Println(time.Now())
	for scan.Scan() {
		if runtime.NumGoroutine() > 100 {
			waitGroup.Wait()
		}
		waitGroup.Add(1)
		go hashPassword(&waitGroup, scan.Text())
	}
	waitGroup.Wait()
	fmt.Println(time.Now())

	if err := scan.Err(); err != nil {
		fmt.Println("Smth goes wrong: ", err)
		os.Exit(1)
	}

}

func main() {
	file()
}
