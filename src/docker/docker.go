package docker

import(
    "os/exec"
    "os"
    "fmt"
    "env"
    "strings"
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

func Exec(args ...string)([]byte, error) {
    if env.DEBUG == true {
        fmt.Printf("Docker path is %s\n",Path) 
    }
    output, err := exec.Command(Path, args...).CombinedOutput();
    if err != nil {
        return nil, fmt.Errorf("docker failed %v : %s(%s)",strings.Join(args, " "), output,err)
    }

    return output, nil 
}
