# Gomiko
[![Build Status](https://travis-ci.org/Ali-aqrabawi/gomiko.svg?branch=master)](https://travis-ci.org/Ali-aqrabawi/gomiko)
[![GolangCI](https://golangci.com/badges/github.com/Ali-aqrabawi/gomiko.svg)](https://golangci.com)
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/Ali-aqrabawi/gomiko)

Gomiko is a `Go` implementation of [netmiko](https://github.com/ktbyers/netmiko). It serves as multi-vendor networking SDK that helps communicate and execute commands via an interactive `shell`
without needing to care about handling device prompts and terminal modes.
 
## Supports
* Cisco IOS
* Cisco IOS XR
* Cisco ASA
* Cisco NX-OS
* Mikrotik RouterOS
* Arista EOS
* Juniper JunOS
* Huawei OLT (New device in this fork)

## Installation
get gomiko pkg: `go get -u github.com/Ali-aqrabawi/gomiko/pkg`.

## Examples :
```go
import (
	"fmt"
	"log"
	"github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	
     device, err := gomiko.NewDevice("192.168.1.1", "admin", "password", "cisco_ios", 22)
     
     if err != nil {
     	log.Fatal(err)
     }
     
     //Connect to device
     if err := device.Connect(); err != nil {
     	log.Fatal(err)
     }
     
     // send command
     output1, _ := device.SendCommand("show vlan")
     
     // send a set of config commands
     commands := []string{"vlan 120", "name v120"}
     output2, _ := device.SendConfigSet(commands)
     
     device.Disconnect()
     
     fmt.Println(output1)
     fmt.Println(output2)
 
}
```
or use new device:
```go
import (
	"fmt"
	"log"
	"github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	
     device, err := gomiko.NewDevice("192.168.1.1", "admin", "password", "huawei_olt", 22)
     
     if err != nil {
     	log.Fatal(err)
     }
     
     //Connect to device
     if err := device.Connect(); err != nil {
     	log.Fatal(err)
     }
     
     // send command
     output1, _ := device.SendCommand("display board 0")
     
     // send a set of config commands
     commands := []string{
     	"interface gpon 0/0",
	"ont add 0 0 sn-auth 4857XXXXXXXXXX\n\n",
	"quit",
	"service-port 1 vlan 1 gpon 0/0/0 ont 0"}
     output2, _ := device.SendConfigSet(commands)
     
     device.Disconnect()
     
     fmt.Println(output1)
     fmt.Println(output2)
 
}
```


## create device with enable password:
```go
import (
	"fmt"
	"log"
	"github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	
     device, err := gomiko.NewDevice("192.168.1.1", "admin", "password", "cisco_ios", 22, gomiko.SecretOption("enablePass"))
     
     if err != nil {
     	log.Fatal(err)
     }     

}
```

## create device Huawei OLT with enable or config mode:
```go
import (
	"fmt"
	"log"
	"github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	// or use SetMode("enable")
     device, err := gomiko.NewDevice("192.168.1.1", "admin", "password", "cisco_ios", 22, gomiko.SetMode("config"))
     
     if err != nil {
     	log.Fatal(err)
     }     

}
```
