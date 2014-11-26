package iptables

import (
    "fmt"
    "os/exec"
)
var (
    xtableSupport = false
)
// Refers to https://github.com/docker/docker/blob/master/pkg/iptables/iptables.go


func init() {
    xtableSupport = exec.Command("iptables","--wait","-L","-n").Run() == nil
    fmt.Printf("Xtables is support: %t\n",xtableSupport)
}

func Opps() {
    fmt.Println("Opps")
}
