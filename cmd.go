// @Name: cmd.go
// @Date: 2023-06-03
// @Author: ls

package main

import (
	"embed"
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/urfave/cli/v2"
)

//go:embed tpl
var tpl embed.FS

func main() {
	app := cli.NewApp()
	app.Name = "ginctl"
	app.Usage = "generate app module"
	app.UsageText = "ginctl new [-di -u username] app [package]"
	app.Version = "1.0.0"
	app.Commands = []*cli.Command{
		{
			Name: "new",
			//Aliases: []string{"new"},
			Usage:     "generate app module",
			UsageText: "ginctl new [-di -u username] app [package]",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:     "wire",
					Aliases:  []string{"di"},
					Usage:    "google wire di",
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
		},
		{
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

				shellPre := "bash"
				shellArg := "-c"
				if "windows" == runtime.GOOS {
					shellPre = "cmd"
					shellArg = "/C"
				}

				shell := "cd " + proName + " && go mod init " + modName
				cmd := exec.Command(shellPre, shellArg, shell)
				err = cmd.Run()

				return err
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func createFile(wire bool, app, packageName, author string) error {
	rootDir, err := os.Getwd()
	if err != nil {
		return err
	}
	newFileDir := filepath.Join(rootDir, app)
	err = mkDir(newFileDir)
	if err != nil {
		return err
	}

	if packageName == "" {
		packageName = filepath.Base(newFileDir)
	}
	packageUpperName := cases.Title(language.English).String(packageName)
	if author == "" {
		pcUser, _ := user.Current()
		author = concat(pcUser.Name, "-", pcUser.Username)
	}

	var genMap = map[string]any{
		"packageName": packageName,
		"PackageName": packageUpperName,
		"wire":        wire,
		"createAt":    time.Now().Format(time.DateOnly),
		"author":      author,
	}

	var filenames = []string{"admin", "app", "dao", "router", "serializer", "service"}
	if wire {
		filenames = append(filenames, "handler", "provider")
	}

	for _, filename := range filenames {
		tmplName := cases.Title(language.English).String(filename)
		genMap["fileName"] = filename
		genMap["FileName"] = tmplName
		switch filename {
		case "handler":
			genMap["param"] = "svc"
			genMap["di"] = concat("I", packageUpperName, "Service")
			genMap["file"] = true
		case "service":
			genMap["param"] = "dao"
			genMap["di"] = concat("I", packageUpperName, "Dao")
			genMap["file"] = true
		case "dao":
			genMap["file"] = true
			genMap["param"] = ""
			genMap["di"] = ""
		case "router":
			genMap["file"] = true
			genMap["param"] = "handler"
			genMap["di"] = concat("I", packageUpperName, "Handler")
		case "provider":
			genMap["file"] = true
		}

		if filename == "provider" {
			err = createGoFiles(newFileDir, filename, "tpl/provider.tpl", genMap)
			if err != nil {
				return err
			}
		} else if filename == "router" {
			err = createGoFiles(newFileDir, filename, "tpl/router.tpl", genMap)
			if err != nil {
				return err
			}
		} else {
			err = createGoFiles(newFileDir, filename, "tpl/file.tpl", genMap)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func createGoFiles(dir, filename, fsPatterns string, paddingMap map[string]any) error {
	t := template.Must(template.ParseFS(tpl, fsPatterns))
	f, err := os.Create(filepath.Join(dir, concat(filename, ".go")))
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, paddingMap)
	return err
}

func fileOrDirIsExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkDir(filePath string) error {
	isExist, err := fileOrDirIsExist(filePath)
	if err != nil {
		return err
	}
	if !isExist {
		err = os.MkdirAll(filePath, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

func currentPath() string {
	_, filename, _, _ := runtime.Caller(0)
	thisPath, _ := filepath.Abs(filename)
	return filepath.Dir(thisPath)
}

func concat(strs ...string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}
