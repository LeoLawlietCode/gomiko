package connections

import (
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh"
)

var Ch chan string = make(chan string)

func KeyScanCallback(hostname string, remote net.Addr, key ssh.PublicKey) error {
	keyHost := string(ssh.MarshalAuthorizedKey(key))
	host := strings.Split(hostname, ":")[0]
	port := strings.Split(hostname, ":")[1]

	if port == "22" {
		Ch <- host + " " + keyHost
	} else {
		Ch <- "[" + host + "]:" + port + " " + keyHost
	}
	return nil
}

func dial(server string, config *ssh.ClientConfig, wg *sync.WaitGroup) {
	_, err := ssh.Dial("tcp", server, config)
	if err != nil {
		log.Fatalln("Failed to dial:", err)
	}
	wg.Done()
}

func out(file *os.File, wg *sync.WaitGroup) {
	for s := range Ch {
		if _, err := file.WriteString(s); err != nil {
			log.Fatalf("Error writing file: %v", err)
		}
		wg.Done()
	}
}

func generateHandshake(session *SSHConn) {
	hostname := strings.Split(session.addr, ":")[0]
	KnownHostsFile := "./engine/connections/keys/known_hosts"
	// Reading the known_hosts file
	content, err := os.ReadFile(KnownHostsFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Listing content per line
	lines := strings.Split(string(content), "\n")
	isKnownHost := false
	// Find hostname in each line
	for _, line := range lines {
		if line != "" {
			knownHost := strings.Split(line, " ")[0]
			if strings.Contains(knownHost, "[") {
				knownHost = strings.Split(knownHost, "]")[0][1:]
			}
			if hostname == knownHost {
				isKnownHost = true
				break
			}
		}
	}

	if !isKnownHost {
		// Open the file in append mode with write permissions
		file, err := os.OpenFile(KnownHostsFile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer file.Close()
		// Edit a file
		auths := []ssh.AuthMethod{ssh.Password(session.password)}

		config := &ssh.ClientConfig{
			User:            session.username,
			Auth:            auths,
			HostKeyCallback: KeyScanCallback,
			Config: ssh.Config{
				KeyExchanges: []string{"diffie-hellman-group-exchange-sha1"},
				Ciphers:      []string{"aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
			},
		}

		var wg sync.WaitGroup
		go out(file, &wg)
		wg.Add(2) // dial and print
		dial(session.addr, config, &wg)

	}
}
