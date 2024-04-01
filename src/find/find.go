package main

import (
	"args"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func exitWithErrorAndCode(err error, code int) {
	log.Println(err)
	os.Exit(code)
}

func getDefaultArguments() []args.Arg {
	return []args.Arg{
		{
			Name:         "sl",
			Description:  "print finded symlinks",
			DefaultValue: false,
			Required:     false,
		},
		{
			Name:         "d",
			Description:  "print finded directories",
			DefaultValue: false,
			Required:     false,
		},
		{
			Name:         "f",
			Description:  "print finded files",
			DefaultValue: false,
			Required:     false,
		},
		{
			Name:         "ext",
			Description:  "print files with specified extension. Can be enabled only if -f is enabled",
			DefaultValue: "\\",
			Required:     false,
		},
	}
}

func checkArguments() error {
	fPresence, extPresence := false, false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "f" {
			fPresence = true
		} else if f.Name == "ext" {
			extPresence = true
		}
	})

	if extPresence && !fPresence {
		return fmt.Errorf("-ext specified without -f")
	}

	return nil
}

func createFindParameters(args map[string]any) FindParams {
	result := FindParams{
		Symlinks:    args["sl"].(bool),
		Directories: args["d"].(bool),
		Files:       args["f"].(bool),
		Extension:   args["ext"].(string),
	}

	if !result.Directories && !result.Symlinks && !result.Files {
		result.Directories = true
		result.Symlinks = true
		result.Files = true
	}

	if result.Extension != "" && result.Extension != "\\" {
		result.Extension = "." + result.Extension
	}

	return result
}

type FindParams struct {
	Symlinks    bool
	Directories bool
	Files       bool
	Extension   string
}

func isSymlink(info os.FileInfo) bool {
	return info.Mode()&os.ModeSymlink == os.ModeSymlink
}

func recursiveTraverse(files []string, params FindParams) {
	for _, path := range files {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return nil
			}

			if params.Symlinks && isSymlink(info) {
				destination, err := os.Readlink(path)
				if err != nil {
					destination = "[broken]"
				}

				_, err = os.Stat(path)
				if err != nil {
					destination = "[broken]"
				}

				fmt.Printf("%s -> %s\n", path, destination)

			} else if params.Directories && info.IsDir() {
				fmt.Println(path)

			} else if params.Files && info.Mode().IsRegular() {
				if params.Extension == "\\" || params.Extension == filepath.Ext(path) {
					fmt.Println(path)
				}
			}

			return nil
		})

		if err != nil {
			log.Printf("can't open file %s\n", path)
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile)

	params, rest, err := args.ParseArgs(getDefaultArguments()...)
	if err != nil {
		exitWithErrorAndCode(err, 1)
	}

	if err = checkArguments(); err != nil {
		exitWithErrorAndCode(err, 2)
	}

	if len(rest) == 0 {
		exitWithErrorAndCode(fmt.Errorf("nothing to search"), 3)
	}

	recursiveTraverse(rest, createFindParameters(params))
}
