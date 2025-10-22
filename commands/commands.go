package commands

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func executeDigQuery(registrar, domain, recordType string) (res bytes.Buffer, err error) {
	var out bytes.Buffer

	cmd := exec.Command("dig", "+short", "@"+registrar, domain, recordType)
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		log.Printf("Error executing dig for %s: %v", domain, err)
		return out, err
	}

	return out, nil

}

func ExecuteARecordQuery(registrar, domain string) []string {
	var results []string
	var buf bytes.Buffer

	buf, _ = executeDigQuery(registrar, domain, "A")

	results = strings.Fields(buf.String())

	return results
}

func ExecuteAAAARecordQuery(registrar, domain string) []string {
	var results []string
	var buf bytes.Buffer

	buf, _ = executeDigQuery(registrar, domain, "AAAA")

	results = strings.Fields(buf.String())

	return results
}

func ExecuteNSRecordQuery(registrar, domain string) []string {
	var results []string
	var buf bytes.Buffer

	buf, _ = executeDigQuery(registrar, domain, "NS")

	results = strings.Fields(buf.String())

	return results
}

func ExecuteTXTRecordQuery(registrar, domain string) []string {
	var results []string
	var buf bytes.Buffer

	buf, _ = executeDigQuery(registrar, domain, "TXT")

	results = strings.Split(strings.TrimSpace(strings.ReplaceAll(buf.String(), "\"", "")), "\n")
	return results
}
