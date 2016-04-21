package main

import (
	"fmt"
	"os"

	"github.com/juju/cmd"
	"github.com/juju/errors"
	"github.com/juju/loggo"
	"github.com/juju/names"
	"launchpad.net/gnuflag"

	"github.com/juju/juju/agent"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/mongo"
	"github.com/juju/juju/state"
)

var logger = loggo.GetLogger("juju")

func main() {
	ctx, err := cmd.DefaultContext()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
	os.Exit(cmd.Main(&FixitCommand{}, ctx, os.Args[1:]))
}

type FixitCommand struct {
	cmd.CommandBase

	dataDir    string
	machineTag names.MachineTag
}

func (c *FixitCommand) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "fix-wrong-tools",
		Args:    "<machine-id>",
		Purpose: "Fix wrong tools",
	}
}

func (c *FixitCommand) SetFlags(f *gnuflag.FlagSet) {
	f.StringVar(&c.dataDir, "data-dir", "/var/lib/juju", "directory for juju data")
}

func (c *FixitCommand) Init(args []string) error {
	if len(args) == 0 {
		return errors.New("missing machine-id")
	}
	var machineId string
	machineId, args = args[0], args[1:]

	if !names.IsValidMachine(machineId) {
		return errors.Errorf("%q is not a valid machine id", machineId)
	}
	c.machineTag = names.NewMachineTag(machineId)
	return cmd.CheckEmpty(args)
}

func (c *FixitCommand) Run(ctx *cmd.Context) error {

	loggo.GetLogger("juju").SetLogLevel(loggo.DEBUG)
	conf, err := agent.ReadConfig(agent.ConfigPath(c.dataDir, c.machineTag))
	if err != nil {
		return err
	}

	info, ok := conf.MongoInfo()
	if !ok {
		return errors.Errorf("no state info available")
	}
	st, err := state.Open(conf.Environment(), info, mongo.DefaultDialOpts(), environs.NewStatePolicy())
	if err != nil {
		return err
	}
	defer st.Close()

	storage, err := st.ToolsStorage()
	defer storage.Close()

	if err := storage.RemoveInvalid(); err != nil {
		return err
	}

	return nil
}
