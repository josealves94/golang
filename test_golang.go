package main


import (
        "fmt"
        "os/exec"
        "runtime"
)

func command_execute() {
        out, err := exec.Command("ls", "-ltr").Output()
        if err != nil {
                fmt.Printf("%s", err)
        }
       fmt.Println("Command Successfully Executed")
        output := string(out[:])
        fmt.Println(output)

}



func main() {
        if runtime.GOOS == "linux" {
                command_execute()
        }
}
