// Package main implements taskrunner2.
// Taskrunner2 is used to specify a list of Bazel targets to run.
// Useful for development. Also supports running certain targets with ibazel.
// Does NOT require you to have ibazel installed.
package main

import (
	"bufio"
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sync"
)

var ibazelPath = "vendor/github.com/bazelbuild/bazel-watcher/cmd/ibazel/ibazel_/ibazel"
var interrupt = make(chan os.Signal, 1)

func getBazelArgs(target string) []string {
	args := []string{"run"}

	// check if we're on macOS
	if runtime.GOOS == "darwin" {
		args = append(args, "--config=macdev")
	}

	args = append(args, target)

	return args
}

func spawnCmd(bin string, target string) error {
	cmd := exec.Command(bin, getBazelArgs(target)...)

	// Capture stdout and stderr and print it with [target]: as prefix
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	// print stdout
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			base.LogInfof("[%s]: %s", target, scanner.Text())
		}
	}()

	// print stderr
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			base.LogInfof("[%s]: %s", target, scanner.Text())
		}
	}()

	// if user terminates the main process, kill the ibazel process
	go func() {
		<-interrupt
		cmd.Process.Kill()
	}()

	return cmd.Wait()
}

// spawnBazel spawns a simple bazel run call.
// on macOS adds `--config=macdev`
func spawnBazel(target string) error {
	return spawnCmd("bazel", target)
}

// spawnIbazel spawns an ibazel run call.
// similar to spawnBazel, but uses ibazel instead.
func spawnIbazel(target string) error {
	return spawnCmd(ibazelPath, target)
}

func run(ctx *cli.Context) error {
	// add interrupt handler
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	// get current wd and make the ibazel path relative to it
	wd, err := os.Getwd()
	if err != nil {
		return cli.Exit(err, 1)
	}

	ibazelPath = wd + "/" + ibazelPath

	// if ran with bazel run, move to the workspace
	if wrksp := os.Getenv("BUILD_WORKSPACE_DIRECTORY"); wrksp != "" {
		if err := os.Chdir(wrksp); err != nil {
			return cli.Exit(err, 1)
		}
	}

	var wg sync.WaitGroup

	// for each target, spawn a bazel run call
	for _, target := range ctx.StringSlice("targets") {
		wg.Add(1)
		go func(target string) {
			defer wg.Done()
			if err := spawnBazel(target); err != nil {
				base.LogFatalf("failed to spawn bazel run: %v", err)
			}
		}(target)
	}

	// for each target, spawn an ibazel run call
	for _, target := range ctx.StringSlice("watch-targets") {
		wg.Add(1)
		go func(target string) {
			defer wg.Done()
			if err := spawnIbazel(target); err != nil {
				base.LogFatalf("failed to spawn ibazel run: %v", err)
			}
		}(target)
	}

	wg.Wait()
	return nil
}

func main() {
	app := &cli.App{
		Name:   "taskrunner2",
		Action: run,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "targets",
				Aliases: []string{"t"},
				Usage:   "bazel targets to run",
			},
			&cli.StringSliceFlag{
				Name:    "watch-targets",
				Aliases: []string{"w"},
				Usage:   "targets to watch for changes",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start taskrunner2: %v", err)
	}
}
