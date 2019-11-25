package main

import "testing"
import "os"
import "github.com/urfave/cli"
import "flag"
import "github.com/alexyans/scooba/handlers"
import "fmt"
import "strconv"

func contextWithFlag(t *testing.T, name string, value interface{}) *cli.Context {
	flag := &flag.FlagSet{}
	switch value.(type) {
	case string:
		ret := flag.String(name, "", "")
		if ret == nil {
			t.Error("Failed to create string flag.")
		}
	case bool:
		ret := flag.Bool(name, false, "")
		if ret == nil {
			t.Error("Failed to create bool flag.")
		}
	}
	
	context := cli.NewContext(nil, flag, nil)
	var err error
	switch value.(type) {
	case string:
		err = context.Set(name, value.(string))
	case bool:
		err = context.Set(name, strconv.FormatBool(value.(bool)))
	}
	
	if err != nil {
		t.Error("Failed to set flag")
	}

	return context
}

func setup() string {
	path := "fixtures/testrepo"
	os.Chdir(path)

	return path
}

func TestGetRepo(t *testing.T) {
	path := setup()

	_, err := handlers.GetRepoFromPwd()
	if err != nil {
		t.Error(fmt.Sprintf("Failed to read repository in path %s.", path))
	}
}

func TestDiveDefault(t *testing.T) {
	path := setup()
	context := cli.NewContext(nil, &flag.FlagSet{}, nil)

	err := handlers.DiveHandler(context)
	if err != nil {
		t.Error(fmt.Sprintf("Dive without argument failed in path %s", path))
	}
}

func TestDiveWithCommit(t *testing.T) {
	path := setup()
	commitHash := "dc67085c71222232a9fa6feb11821192f260fe34"
	context := contextWithFlag(t, "commit", commitHash)

	defer func() {
		if r := recover(); r != nil {
			t.Error(fmt.Sprintf("Dive for commit %s in path %s panicked.", commitHash, path))
		}
	}()

	err := handlers.DiveHandler(context)
	if err != nil {
		t.Error(fmt.Sprintf("Dive failed for commit %s in path %s.", commitHash, path))
	}
}

func TestDiveNonExistentCommit(t *testing.T) {
	path := setup()
	commitHash := "dcdcdc"
	context := contextWithFlag(t, "commit", commitHash)

	defer func() {
		if r := recover(); r == nil {
			t.Error(fmt.Sprintf("Dive for fake commit %s in path %s should have panicked.", commitHash, path))
		}
	}()

	_ = handlers.DiveHandler(context)
}

func TestForward(t *testing.T) {
	path := setup()
	context := cli.NewContext(nil, &flag.FlagSet{}, nil)
	_ = handlers.DiveHandler(context)

	context = contextWithFlag(t, "forward", true)
	err := handlers.ForwardHandler(context)
	if err != nil {
		t.Error(fmt.Sprintf("Forward in path %s failed.", path))
	}

	err = handlers.ForwardHandler(context)
	if err != nil {
		t.Error(fmt.Sprintf("Forward in path %s failed.", path))
	}
}

func TestBackward(t *testing.T) {
	path := setup()
	context := cli.NewContext(nil, &flag.FlagSet{}, nil)
	_ = handlers.DiveHandler(context)

	context = contextWithFlag(t, "backward", true)
	err := handlers.BackwardHandler(context)
	if err != nil {
		t.Error(fmt.Sprintf("Backward in path %s failed.", path))
	}

	err = handlers.BackwardHandler(context)
	if err != nil {
		t.Error(fmt.Sprintf("Forward in path %s failed.", path))
	}
}