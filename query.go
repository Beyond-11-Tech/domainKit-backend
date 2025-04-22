package main

import (
	"bytes"
	"log"
	"os/exec"
)

func executeDigQuery(registrar, domain string) (res bytes.Buffer, err error) {
	var out bytes.Buffer

	cmd := exec.Command("dig", "+short", "@"+registrar, domain)
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		log.Printf("Error executing dig for %s: %v", domain, err)
		return out, err
	}

	return out, nil

}
