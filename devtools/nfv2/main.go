// Copyright 2023 Peridot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main implements nfv2 (nofussvendor 2) which is a tool to make sure the generated
// artifacts from Bazel can be used with native Go tooling. This is done by generating
// a go.mod file in bazel-bin to match the package name and then add a replace directive
// to the go.mod file in the root of the repo.
package main

import (
	"bytes"
	"fmt"
	base "go.resf.org/peridot/base/go"
	blaze_query "go.resf.org/peridot/third_party/bazel/src/main/protobuf"
	"golang.org/x/mod/modfile"
	"google.golang.org/protobuf/proto"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func callBazel(args ...string) []byte {
	cmd := exec.Command(
		"bazel",
		args...,
	)

	var out bytes.Buffer
	cmd.Stdout = &out

	var errOut bytes.Buffer
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		base.LogFatalf("error: %s", err)
	}

	return out.Bytes()
}

func main() {
	// get bazel workspace
	// exit if not invoked with bazel
	searchDirectory := os.Getenv("BUILD_WORKSPACE_DIRECTORY")
	if searchDirectory == "" {
		base.LogFatalf("error: BUILD_WORKSPACE_DIRECTORY not found")
	}

	// change directory to bazel workspace
	err := os.Chdir(searchDirectory)
	if err != nil {
		base.LogFatalf("could not change directory: %v", err)
	}

	queryProto := callBazel("query", "kind(go_proto_library, //... - //vendor/... - //third_party/... + //third_party/bazel/... + //third_party/googleapis/google/longrunning/...)", "--output", "proto")

	var query blaze_query.QueryResult
	err = proto.Unmarshal(queryProto, &query)
	if err != nil {
		base.LogFatalf("could not unmarshal query result: %v", err)
	}

	replaceList := map[string]string{}
	for _, rule := range query.Target {
		if *rule.Rule.RuleClass != "go_proto_library" {
			continue
		}

		fullTarget := *rule.Rule.Name
		ruleName := strings.Split(fullTarget, ":")[1]

		for _, nstring := range rule.Rule.Attribute {
			if *nstring.Name == "importpath" {
				origImportPath := *nstring.StringValue
				importpath := origImportPath

				buildLocation := strings.Replace(*rule.Rule.Location, searchDirectory+"/", "", 1)
				buildLocation = filepath.Dir(strings.Split(buildLocation, ":")[0])
				buildLocation = fmt.Sprintf("bazel-bin/%s/%s_", buildLocation, ruleName)

				modDir := filepath.Join(buildLocation, importpath)
				if err := os.MkdirAll(modDir, 0755); err != nil && !os.IsExist(err) {
					base.LogFatalf("could not generate directory for importpath %s", importpath)
				}

				if strings.HasSuffix(importpath, "/v1") {
					importpath = strings.TrimSuffix(importpath, "/v1")
					modDir = filepath.Join(modDir, "..")
				}

				modContent := []byte(fmt.Sprintf("module %s", importpath))
				if err := os.WriteFile(filepath.Join(modDir, "go.mod"), modContent, 0644); err != nil {
					base.LogFatalf("could not write go.mod file: %v", err)
				}

				replaceList[importpath] = modDir
			}
		}
	}

	// parse go.mod file
	modBts, err := os.ReadFile("go.mod")
	if err != nil {
		base.LogFatalf("could not read go.mod file: %v", err)
	}

	mod, err := modfile.Parse("go.mod", modBts, nil)
	if err != nil {
		base.LogFatalf("could not parse go.mod file: %v", err)
	}

	// add replace directives
	for from, to := range replaceList {
		err := mod.AddReplace(from, "", "./"+to, "")
		if err != nil {
			base.LogFatalf("could not add replace directive: %v", err)
		}
	}

	// write go.mod file
	bts, err := mod.Format()
	if err != nil {
		base.LogFatalf("could not format go.mod file: %v", err)
	}

	err = os.WriteFile("go.mod", bts, 0644)
	if err != nil {
		base.LogFatalf("could not write go.mod file: %v", err)
	}

	base.LogInfof("added %d replace directives pointing to generated modules", len(replaceList))
}
