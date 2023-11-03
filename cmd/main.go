/*
Copyright 2021 The Wick Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"

	"github.com/WickCloud/wick/pkg/skel"
	"github.com/WickCloud/wick/pkg/util"
)

func cmdAdd(args *skel.CmdArgs) error {
	fmt.Println("进入到cmdAdd")
	data := map[string]interface{}{
		"containerId": args.ContainerID,
		"netns":       args.Netns,
		"ifName":      args.IfName,
		"args":        args.Args,
		"path":        args.Path,
		"stdinData":   string(args.StdinData),
	}
	util.Print(data)
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	fmt.Println("进入到cmdDel")
	data := map[string]interface{}{
		"containerId": args.ContainerID,
		"netns":       args.Netns,
		"ifName":      args.IfName,
		"args":        args.Args,
		"path":        args.Path,
		"stdinData":   string(args.StdinData),
	}
	util.Print(data)
	return nil
}

func cmdCheck(args *skel.CmdArgs) error {
	fmt.Println("进入到cmdCheck")
	data := map[string]interface{}{
		"containerId": args.ContainerID,
		"netns":       args.Netns,
		"ifName":      args.IfName,
		"args":        args.Args,
		"path":        args.Path,
		"stdinData":   string(args.StdinData),
	}
	util.Print(data)
	return nil
}

func main() {
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, bv.BuildString("wick"))
}
