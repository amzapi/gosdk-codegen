package main

import (
	"flag"
	"fmt"
	"gosdk-codegen/pkg/codegen"
	"gosdk-codegen/pkg/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func errExit(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

var (
	flagPackageName string
	flagGenerate    string
	flagOutputFile  string
)

type configuration struct {
	PackageName     string   `yaml:"package"`
	GenerateTargets []string `yaml:"generate"`
	OutputFile      string   `yaml:"output"`
}

func main() {

	flag.StringVar(&flagPackageName, "package", "", "The package name for generated code")
	flag.StringVar(&flagGenerate, "generate", "types,client", `Comma-separated list of code to generate; valid options: "types", "client"`)
	flag.StringVar(&flagOutputFile, "o", "", "Where to output generated code, stdout is default")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Please specify a path to a OpenAPI 3.0 spec file")
		os.Exit(1)
	}

	cfg := configFromFlags()

	// If the package name has not been specified, we will use the name of the
	// swagger file.
	if cfg.PackageName == "" {
		path := flag.Arg(0)
		baseName := filepath.Base(path)
		// Split the base name on '.' to get the first part of the file.
		nameParts := strings.Split(baseName, ".")
		cfg.PackageName = codegen.ToCamelCase(nameParts[0])
	}

	opts := codegen.Options{}
	for _, g := range cfg.GenerateTargets {
		switch g {
		case "client":
			opts.GenerateClient = true
		case "types":
			opts.GenerateTypes = true
		default:
			fmt.Printf("unknown generate option %s\n", g)
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	swagger, err := util.LoadSwagger(flag.Arg(0))
	if err != nil {
		errExit("error loading swagger spec\n: %s", err)
	}

	code, err := codegen.Generate(swagger, cfg.PackageName, opts)
	if err != nil {
		errExit("error generating code: %s\n", err)
	}

	if cfg.OutputFile != "" {
		err = ioutil.WriteFile(cfg.OutputFile, []byte(code), 0644)
		if err != nil {
			errExit("error writing generated code to file: %s", err)
		}
	} else {
		fmt.Println(code)
	}
}

// configFromFlags parses the flags and the config file. Anything which is
// a zerovalue in the configuration file will be replaced with the flag
// value, this means that the config file overrides flag values.
func configFromFlags() *configuration {
	var cfg configuration
	if cfg.PackageName == "" {
		cfg.PackageName = flagPackageName
	}
	if cfg.GenerateTargets == nil {
		cfg.GenerateTargets = util.ParseCommandLineList(flagGenerate)
	}
	if cfg.OutputFile == "" {
		cfg.OutputFile = flagOutputFile
	}
	return &cfg
}
