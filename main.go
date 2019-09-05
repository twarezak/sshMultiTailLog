package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
)

var tailLogCmd = "tail -f /var/log/php/all.log"

func main() {
	config := getConfig()

	MultiTail(config.Hosts)
}

func getConfig() Hosts {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var hosts Hosts

	_ = json.Unmarshal([]byte(byteValue), &hosts)

	return hosts
}

func TailLog(name string, client *ssh.Client, lines chan<- string) {
	sess, _ := client.NewSession()
	defer sess.Close()

	out, _ := sess.StdoutPipe()

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)

	_ = sess.Start(tailLogCmd)

	for scanner.Scan() {
		lines <- fmt.Sprintf("[%s] %s", name, scanner.Text())
	}

	_ = sess.Wait()
}

func MultiTail(hosts []Host) {
	lines := make(chan string)

	for _, host := range hosts {
		var sshConfig = &ssh.ClientConfig{
			User: host.User,
			Auth: []ssh.AuthMethod{
				ssh.Password(host.Password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		client, err := ssh.Dial("tcp", host.Address + ":22", sshConfig)
		if err != nil {
			log.Fatal("Failed to dial: ", err)
		}

		go TailLog(host.Address, client, lines)
	}

	for l := range lines {
		fmt.Println(l)
	}
}
