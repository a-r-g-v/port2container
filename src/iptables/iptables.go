package iptables

import (
    "fmt"
    "os/exec"
    "env"
    "errors"
    "strings"
)
var (
    xtableSupport = false
    nat = []string{"-t", "nat"}
)
// Refers to https://github.com/docker/docker/blob/master/pkg/iptables/iptables.go


func init() {
    xtableSupport = exec.Command("iptables","--wait","-L","-n").Run() == nil

    if env.DEBUG {
    fmt.Printf("Xtables is support: %t\n",xtableSupport)
    }
}


func Exists(args ...string) bool {
        if _, err := Exec(append([]string{"-C"}, args...)...); err == nil {
            return true
        }


        rule := strings.Replace(strings.Join(args, " "), "-t nat", "", -1)
        fmt.Printf("%s",rule)
        return false
}

// Exec Add iptables-entry
func Exec(args ...string)([]byte, error) {

    path, err := exec.LookPath("iptables")
    if err != nil{
        return nil, errors.New("Iptables Not Found")
    }
    if xtableSupport {
        args = append([]string{"--wait"},args...)
    }
    
    if env.DEBUG{
      fmt.Printf("%s",args)
    }

    output, err := exec.Command(path, args...).CombinedOutput()
    if err!= nil{
        return nil, fmt.Errorf("iptables failed: iptables %v : %s(%s) ", strings.Join(args, " "), output, err)
    }

    
    if env.DEBUG{
      fmt.Printf("%s",output)
    }

    return output,nil
}


