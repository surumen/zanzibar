// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/uber/zanzibar/codegen"
	"github.com/uber/zanzibar/runtime"
)

var configFile = flag.String("config", "", "the config file path")

const templateDir = "./codegen/templates/*.tmpl"

func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("%s:\n %s \n", message, err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	if *configFile == "" {
		flag.Usage()
		os.Exit(1)
		return
	}

	configDirName := filepath.Dir(*configFile)
	config := zanzibar.NewStaticConfigOrDie([]string{
		*configFile,
	}, nil)

	configDirName, err := filepath.Abs(configDirName)
	checkError(
		err, fmt.Sprintf("can not get abs path of config dir %s", configDirName),
	)

	copyright, err := ioutil.ReadFile(filepath.Join(
		configDirName,
		config.MustGetString("copyrightHeader"),
	))
	if err != nil {
		// Default to an empty copyright for now
		copyright = []byte("")
	}

	packageHelper, err := codegen.NewPackageHelper(
		filepath.Join(configDirName, config.MustGetString("thriftRootDir")),
		config.MustGetString("genCodePackage"),
		filepath.Join(configDirName, config.MustGetString("targetGenDir")),
		config.MustGetString("gatewayNamespace"),
		string(copyright),
	)
	checkError(
		err, fmt.Sprintf("Can't build package helper %s", configDirName),
	)

	moduleSystem, err := codegen.NewDefaultModuleSystem(packageHelper)
	checkError(
		err, fmt.Sprintf("Error creating module system %s", configDirName),
	)

	gatewaySpec, err := codegen.NewGatewaySpec(
		packageHelper,
		configDirName,
		config.MustGetString("clientConfig"),
		config.MustGetString("endpointConfig"),
		config.MustGetString("middlewareConfig"),
		config.MustGetString("gatewayName"),
	)
	checkError(
		err, fmt.Sprintf("Cannot create gateway spec %#v", gatewaySpec),
	)

	// TODO: All components should be generated by the module system
	fmt.Printf("Generating module system components \n")
	err = moduleSystem.GenerateBuild(
		configDirName,
		packageHelper.CodeGenTargetPath(),
	)
	checkError(err, "Failed to generate module system components.")

	// TODO: Module system should generate an index for clients
	fmt.Printf("Generating client index code for gateway \n")
	err = gatewaySpec.GenerateClientsInit()
	checkError(err, "Failed to generate client index.")

	fmt.Printf("Generating endpoint code for gateway \n")
	err = gatewaySpec.GenerateEndpoints()
	checkError(err, "Failed to generate endpoint files.")

	fmt.Printf("Generating main.go for gateway \n")
	err = gatewaySpec.GenerateMain()
	checkError(err, "Failed to generate main.go.")

}
