package main

import (
    "os"
    "net/http"
    "io/ioutil"

    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "git-ignore"
    app.HideHelp = true
    app.Flags = []cli.Flag{
        cli.HelpFlag,
        cli.BoolTFlag{
            Name: "add, a",
            Usage: "Add ignore settings into the .gitignore file",
        },
        cli.BoolTFlag{
            Name: "create, c",
            Usage: "Create the .gitignore file",
        },
    }
    app.Version = Version
    app.Usage = ""
    app.Author = "Yoshiki Aoki"
    app.Email = "yoshiki_aoki@dwango.co.jp"
    app.Action = doMain
    cli.AppHelpTemplate = `
  NAME:
     {{.Name}} - {{.Usage}}
  
  USAGE:
     {{.Name}} [options] [arguments...]
  
  VERSION:
     {{.Version}}{{if or .Author .Email}}
  
  AUTHOR:{{if .Author}}
    {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
    {{.Email}}{{end}}{{end}}
  
  OPTIONS:
     {{range .Flags}}{{.}}
     {{end}}

`
    app.Run(os.Args)
}


func doMain(c *cli.Context) {
    if len(c.Args()) == 0 {
        cli.ShowAppHelp(c)
        os.Exit(1)
    }

    resp, err := http.Get("https://www.gitignore.io/api/" + c.Args()[0])
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    byteArray, _ := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    f.Write(byteArray)
    defer f.Close()
}

