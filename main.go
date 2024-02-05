package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// frontend http-in
//         bind *:80

//         acl host_d1 hdr(host) -i d1.com
//         acl host_d2 hdr(host) -i d2.com

//         use_backend be_d1 if host_d1
//         use_backend be_d2 if host_d2

// backend be_d1
//         server D1M1 10.0.0.1:8080
//         server D1M2 10.0.0.2:8080

// backend be_d2
//         server D2M1 10.0.0.1:8080
//         server D2M2 10.0.0.2:8080

func main() {
	godotenv.Load()

	acls := "defaults\n"
	acls += "    timeout connect 5s\n"
	acls += "    timeout client 50s\n"
	acls += "    timeout server 50s\n\n"

	acls += "frontend http-in\n    bind *:80\n"
	uses := ""
	backends := ""

	idx_rule := 0
	for _, e := range os.Environ() {
		// Check start with "BIND_"
		if strings.HasPrefix(e, "BIND_") {
			rules := strings.SplitAfter(e, "=")[1:]
			for _, rule := range rules {
				ips := strings.Split(rule, " ")
				for idx, ip := range ips {
					if idx == 0 {
						acls += fmt.Sprintf("    acl host_d%d hdr(host) -i %s\n", idx_rule, ip)
						uses += fmt.Sprintf("    use_backend be_d%d if host_d%d\n", idx_rule, idx_rule)
						backends += fmt.Sprintf("\nbackend be_d%d\n", idx_rule)
					} else {
						// fmt.Println("to", ip)
						backends += fmt.Sprintf("    server d%dm%d %s\n", idx_rule, idx, ip)
					}
				}
				fmt.Println()
			}
			idx_rule++
		}

	}
	// read file from .env

	fmt.Println(acls)
	fmt.Println(uses)
	fmt.Println(backends)

	os.WriteFile("haproxy.cfg", []byte(acls+uses+backends), 0644)
	defer fmt.Println("haproxy.cfg created")

}
