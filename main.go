package main

import (
  "os"
  "os/exec"
)

var gitLocation = ""
var gitPullDirectory = ""

var javaWLocation = "C:\\Program Files (x86)\\Minecraft Launcher\\runtime\\jre-x64\\bin\\javaw.exe"

func main() {
    
    cmd1 := exec.Command(gitLocation, "-C", gitPullDirectory, "pull")
    cmd1.Stdout = os.Stdout
    cmd1.Run()
    
    cmd2 := exec.Command(javaWLocation, os.Args[1:]...)
    cmd2.Stdout = os.Stdout
    cmd2.Run()
    
}