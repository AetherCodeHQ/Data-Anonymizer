package main

// Data-Anonymizer: PII (email, phone, SSN) maskeler
import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: data-anonymizer <file>")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	s := string(data)

	emailRe := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	phoneRe := regexp.MustCompile(`\b\d{3}[-.]?\d{3}[-.]?\d{4}\b`)
	ssnRe := regexp.MustCompile(`\b\d{3}-\d{2}-\d{4}\b`)
	ipRe := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)

	emailCount := len(emailRe.FindAllString(s, -1))
	phoneCount := len(phoneRe.FindAllString(s, -1))
	ssnCount := len(ssnRe.FindAllString(s, -1))
	ipCount := len(ipRe.FindAllString(s, -1))

	s = emailRe.ReplaceAllString(s, "[EMAIL]")
	s = phoneRe.ReplaceAllString(s, "[PHONE]")
	s = ssnRe.ReplaceAllString(s, "[SSN]")
	s = ipRe.ReplaceAllString(s, "[IP]")

	fmt.Print(s)
	fmt.Fprintf(os.Stderr, "masked: %d emails, %d phones, %d SSNs, %d IPs\n",
		emailCount, phoneCount, ssnCount, ipCount)
}
