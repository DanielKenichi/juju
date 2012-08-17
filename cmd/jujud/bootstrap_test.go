package main

import (
	"encoding/base64"
	. "launchpad.net/gocheck"
	"launchpad.net/juju-core/juju/testing"
	coretesting "launchpad.net/juju-core/testing"
)

type BootstrapSuite struct {
	coretesting.LoggingSuite
	testing.JujuConnSuite
	path string
}

var _ = Suite(&BootstrapSuite{})

func (s *BootstrapSuite) SetUpTest(c *C) {
	s.LoggingSuite.SetUpTest(c)
	s.path = "/watcher"
	s.JujuConnSuite.SetUpTest(c)
}

func (s *BootstrapSuite) TearDownTest(c *C) {
	s.JujuConnSuite.TearDownTest(c)
	s.LoggingSuite.TearDownTest(c)
}

func initBootstrapCommand(args []string) (*BootstrapCommand, error) {
	c := &BootstrapCommand{}
	return c, initCmd(c, args)
}

func (s *BootstrapSuite) TestParse(c *C) {
	args := []string{}
	_, err := initBootstrapCommand(args)
	c.Assert(err, ErrorMatches, "--instance-id option must be set")

	args = append(args, "--instance-id", "iWhatever")
	_, err = initBootstrapCommand(args)
	c.Assert(err, ErrorMatches, "--env-type option must be set")

	args = append(args, "--env-type", "dummy")
	cmd, err := initBootstrapCommand(args)
	c.Assert(err, IsNil)
	c.Assert(cmd.StateInfo.Addrs, DeepEquals, []string{"127.0.0.1:2181"})
	c.Assert(cmd.InstanceId, Equals, "iWhatever")
	c.Assert(cmd.EnvType, Equals, "dummy")

	args = append(args, "--zookeeper-servers", "zk1:2181,zk2:2181")
	cmd, err = initBootstrapCommand(args)
	c.Assert(err, IsNil)
	c.Assert(cmd.StateInfo.Addrs, DeepEquals, []string{"zk1:2181", "zk2:2181"})

	args = append(args, "haha disregard that")
	_, err = initBootstrapCommand(args)
	c.Assert(err, ErrorMatches, `unrecognized args: \["haha disregard that"\]`)
}

func (s *BootstrapSuite) TestSetMachineId(c *C) {
	args := []string{"--zookeeper-servers"}
	args = append(args, s.StateInfo(c).Addrs...)
	args = append(args, "--instance-id", "over9000", "--env-type", "dummy")
	cmd, err := initBootstrapCommand(args)
	c.Assert(err, IsNil)
	err = cmd.Run(nil)
	c.Assert(err, IsNil)

	machines, err := s.State.AllMachines()
	c.Assert(err, IsNil)
	c.Assert(len(machines), Equals, 1)

	instid, err := machines[0].InstanceId()
	c.Assert(err, IsNil)
	c.Assert(instid, Equals, "over9000")
}

var base64ConfigTests = []struct {
	input    []string
	err      string
	expected map[string]interface{}
}{
	{
		// no value supplied
		nil,
		"",
		nil,
	}, {
		// empty 
		[]string{"--env-config", ""},
		"",
		nil,
	}, {
		// wrong, should be base64
		[]string{"--env-config", "name: banana\n"},
		".*illegal base64 data at input byte.*",
		nil,
	}, {
		[]string{"--env-config", base64.StdEncoding.EncodeToString([]byte("name: banana\n"))},
		"",
		map[string]interface{}{"name": "banana"},
	},
}

func (s *BootstrapSuite) TestBase64Config(c *C) {
	for _, t := range base64ConfigTests {
		args := []string{"--zookeeper-servers"}
		args = append(args, s.StateInfo(c).Addrs...)
		args = append(args, "--instance-id", "over9000", "--env-type", "dummy")
		args = append(args, t.input...)
		cmd, err := initBootstrapCommand(args)
		if t.err == "" {
			c.Assert(cmd, NotNil)
			c.Assert(err, IsNil)
			c.Assert(cmd.EnvConfig, DeepEquals, t.expected)
		} else {
			c.Assert(err, ErrorMatches, t.err)
		}
	}
}
