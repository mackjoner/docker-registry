package main

import (
  "github.com/astaxie/beego"
  "github.com/codegangsta/cli"
  "github.com/dockboard/docker-registry/cmd"
  "os"
)

func main() {
  app := cli.NewApp()
  app.Name = beego.AppConfig.String("appname")
  app.Usage = beego.AppConfig.String("usage")
  app.Version = beego.AppConfig.String("version")

  app.Commands = []cli.Command{
    cmd.CmdWeb,
  }

  app.Flags = append(app.Flags, []cli.Flag{}...)
  app.Run(os.Args)
}
