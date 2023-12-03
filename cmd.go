// @Name: cmd.go
// @Date: 2023-06-03
// @Author: ls

package main

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

//go:embed file.tpl
var tpl embed.FS

func main() {
	app := cli.NewApp()
	app.Name = "ginctl"
	app.Usage = "generate app module"
	app.UsageText = "ginctl new [-di -u username] app [package]"
	app.Version = "1.0.0"
	app.Commands = []*cli.Command{
		newCmd(),
		initCmd(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func newCmd() *cli.Command {
	return &cli.Command{
		Name:      "new",
		Usage:     "generate app module",
		UsageText: "ginctl new [-di -u username] app [package]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "wire",
				Aliases:  []string{"di"},
				Usage:    "google wire",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "author",
				Aliases:  []string{"u"},
				Usage:    "author who created files",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			args := c.Args()
			appName := args.First()
			if strings.TrimSpace(appName) == "" {
				return errors.New("generate app error: app name is empty")
			}
			packageName := args.Get(1)
			di := c.Bool("wire")
			authorName := c.String("author")
			err := createFile(di, appName, packageName, authorName)
			return err
		},
	}
}

func initCmd() *cli.Command {
	return &cli.Command{
		Name:      "init",
		Usage:     "go mod init project_name [mod_name]",
		UsageText: "ginctl init your_project [your_project.com]",
		Action: func(c *cli.Context) error {
			args := c.Args()
			proName := args.First()
			if strings.TrimSpace(proName) == "" {
				return errors.New("init project error: project name is empty")
			}
			modName := args.Get(1)
			if strings.TrimSpace(modName) == "" {
				modName = proName
			}
			err := os.Mkdir(proName, 0666)
			if err != nil {
				return err
			}

			platform := runtime.GOOS
			shellPre := "bash"
			shellArg := "-c"
			if "windows" == platform {
				shellPre = "cmd"
				shellArg = "/C"
			}

			shell := "cd " + proName + " && go mod init " + modName
			cmd := exec.Command(shellPre, shellArg, shell)
			err = cmd.Run()

			return err
		},
	}
}
