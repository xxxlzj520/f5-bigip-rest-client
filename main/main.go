package main

import (
	"f5-bigip-rest-client/f5"
	"f5-bigip-rest-client/f5/ltm"
	"fmt"
	"log"
)

func main() {
	f5Client, err := f5.NewBasicClient("https://192.168.1.11", "admin", "admin")
	if err != nil {
		log.Fatal(err)
	}
	// Disable the cert chech.
	f5Client.DisableCertCheck()
	if f5Client.IsActive("bigip1") {
		fmt.Println("ssss")
	} else {
		fmt.Println("error")
	}
	// setup client for the LTM API
	ltmClient := ltm.New(f5Client)
	// query the /ltm/virtual API
	vsConfigList, err := ltmClient.Virtual().ListAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(vsConfigList)
}
