package main

import (
  "os"
  "os/exec"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "flag"
)

// Struct for unmarshalling config file
type command struct {
  Name  string
  Path  string
  Args  []string
}

func main() {
  // If parameter config is missing, use config.json as default config file
  var configFile = flag.String("config", "config.json", "Path to config file")
  flag.Parse()

  // Read config file content to config string
  config, err := ioutil.ReadFile(*configFile)
  if err != nil {
    fmt.Println("Error reading from", *configFile)
    os.Exit(1)
  }

  // Parse the json content to an array of command structs
  var commandsToRun []command
  if err:= json.Unmarshal(config, &commandsToRun); err != nil {
    fmt.Printf("Error while unmarshalling %s", err)
  }

  // Loop through the array of command structs, execute each command and present output
  for _, value := range commandsToRun {
    cmd := exec.Command(value.Path, value.Args...)
    cmdOutput, err := cmd.Output()
    if err != nil {
      fmt.Printf("Error: %s", err)
    }
    fmt.Println(string(cmdOutput))
  }
}
