package docker

import(
    "os/exec"
    "os"
    "fmt"
    "env"
)
var Path string
var err error

func init() {
   Path, err = exec.LookPath("docker")
   if err != nil {
        fmt.Printf("Error pkg/docker : %s",err);
        os.Exit(-1);
   }
}

func Exec() {
    if env.DEBUG == true {
        fmt.Printf("Docker path is %s\n",Path) 
    }
}
