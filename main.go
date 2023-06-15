package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	templateFilename := flag.String("template", "template.txt", "The template file")
	configFilename := flag.String("config", "config.txt", "The configuration file")

	flag.Parse()

	fmt.Fprintf(os.Stderr, "Template: %s\n", *templateFilename)
	fmt.Fprintf(os.Stderr, "Config: %s\n", *configFilename)

	if *templateFilename == "" {
		panic("-template must be specified")
	}
	if *configFilename == "" {
		panic("-config must be specified")
	}

	templateFile, err := os.Open(*templateFilename)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, templateFile)

	configFile, err := os.Open(*configFilename)
	if err != nil {
		panic(err)
	}
	values := configValues(configFile)

	output := buf.String()
	for k, v := range values {
		output = strings.ReplaceAll(output, k, v)
	}

	fmt.Println(output)

}

func configValues(r io.Reader) map[string]string {
	scanner := bufio.NewScanner(r)
	output := make(map[string]string)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		if len(parts) != 2 {
			panic("format must be KEY=VALUE")
		}
		output[parts[0]] = parts[1]
	}
	return output
}
