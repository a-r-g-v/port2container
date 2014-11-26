package iptables

import (
    "fmt"
    "os/exec"
    "env"
)
var (
    xtableSupport = false
)
// Refers to https://github.com/docker/docker/blob/master/pkg/iptables/iptables.go


func init() {
    xtableSupport = exec.Command("iptables","--wait","-L","-n").Run() == nil

    if env.DEBUG {
    fmt.Printf("Xtables is support: %t\n",xtableSupport)
    }
}

func Opps() {
    if env.DEBUG {
    fmt.Println("Opps")
    }
}
