package state_test

import (
	. "launchpad.net/gocheck"
	"launchpad.net/juju-core/charm"
	"launchpad.net/juju-core/state"
	"time"
)

type ServiceSuite struct {
	ConnSuite
	charm   *state.Charm
	service *state.Service
}

var _ = Suite(&ServiceSuite{})

func (s *ServiceSuite) SetUpTest(c *C) {
	s.ConnSuite.SetUpTest(c)
	s.charm = s.AddTestingCharm(c, "dummy")
	var err error
	s.service, err = s.St.AddService("wordpress", s.charm)
	c.Assert(err, IsNil)
}

func (s *ServiceSuite) TestServiceCharm(c *C) {
	// Check that getting and setting the service charm URL works correctly.
	testcurl, err := s.service.CharmURL()
	c.Assert(err, IsNil)
	c.Assert(testcurl.String(), Equals, s.charm.URL().String())

	// TODO BUG shouldn't it be an error to set a charm URL that doesn't correspond
	// to a known charm??
	testcurl = charm.MustParseURL("local:myseries/mydummy-1")
	err = s.service.SetCharmURL(testcurl)
	c.Assert(err, IsNil)
	testcurl, err = s.service.CharmURL()
	c.Assert(err, IsNil)
	c.Assert(testcurl.String(), Equals, "local:myseries/mydummy-1")
}

func (s *ServiceSuite) TestServiceExposed(c *C) {
	// Check that querying for the exposed flag works correctly.
	exposed, err := s.service.IsExposed()
	c.Assert(err, IsNil)
	c.Assert(exposed, Equals, false)

	// Check that setting and clearing the exposed flag works correctly.
	err = s.service.SetExposed()
	c.Assert(err, IsNil)
	exposed, err = s.service.IsExposed()
	c.Assert(err, IsNil)
	c.Assert(exposed, Equals, true)
	err = s.service.ClearExposed()
	c.Assert(err, IsNil)
	exposed, err = s.service.IsExposed()
	c.Assert(err, IsNil)
	c.Assert(exposed, Equals, false)

	// Check that setting and clearing the exposed flag multiple doesn't fail.
	err = s.service.SetExposed()
	c.Assert(err, IsNil)
	err = s.service.SetExposed()
	c.Assert(err, IsNil)
	err = s.service.ClearExposed()
	c.Assert(err, IsNil)
	err = s.service.ClearExposed()
	c.Assert(err, IsNil)

	// Check that setting and clearing the exposed flag on removed services also doesn't fail.
	// TODO BUG this doesn't appear to be sane.
	err = s.St.RemoveService(s.service)
	c.Assert(err, IsNil)
	err = s.service.ClearExposed()
	c.Assert(err, IsNil)
}

func (s *ServiceSuite) TestAddUnit(c *C) {
	// Check that principal units can be added on their own.
	unitZero, err := s.service.AddUnit()
	c.Assert(err, IsNil)
	c.Assert(unitZero.Name(), Equals, "wordpress/0")
	principal := unitZero.IsPrincipal()
	c.Assert(principal, Equals, true)
	unitOne, err := s.service.AddUnit()
	c.Assert(err, IsNil)
	c.Assert(unitOne.Name(), Equals, "wordpress/1")
	principal = unitOne.IsPrincipal()
	c.Assert(principal, Equals, true)

	// Check that principal units cannot be added to principal units.
	_, err = s.service.AddUnitSubordinateTo(unitZero)
	c.Assert(err, ErrorMatches, `can't add unit of principal service "wordpress" as a subordinate of "wordpress/0"`)

	// Assign the principal unit to a machine.
	m, err := s.St.AddMachine()
	c.Assert(err, IsNil)
	err = unitZero.AssignToMachine(m)
	c.Assert(err, IsNil)

	// Add a subordinate service.
	subCharm := s.AddTestingCharm(c, "logging")
	logging, err := s.St.AddService("logging", subCharm)
	c.Assert(err, IsNil)

	// Check that subordinate units can be added to principal units
	subZero, err := logging.AddUnitSubordinateTo(unitZero)
	c.Assert(err, IsNil)
	c.Assert(subZero.Name(), Equals, "logging/0")
	principal = subZero.IsPrincipal()
	c.Assert(principal, Equals, false)

	// Check the subordinate unit has been assigned its principal's machine.
	id, err := subZero.AssignedMachineId()
	c.Assert(err, IsNil)
	c.Assert(id, Equals, m.Id())

	// Check that subordinate units must be added to other units.
	_, err = logging.AddUnit()
	c.Assert(err, ErrorMatches, `cannot directly add units to subordinate service "logging"`)

	// Check that subordinate units cannnot be added to subordinate units.
	_, err = logging.AddUnitSubordinateTo(subZero)
	c.Assert(err, ErrorMatches, "a subordinate unit must be added to a principal unit")
}

func (s *ServiceSuite) TestReadUnit(c *C) {
	_, err := s.service.AddUnit()
	c.Assert(err, IsNil)
	_, err = s.service.AddUnit()
	c.Assert(err, IsNil)
	// Check that retrieving a unit works correctly.
	unit, err := s.service.Unit("wordpress/0")
	c.Assert(err, IsNil)
	c.Assert(unit.Name(), Equals, "wordpress/0")

	// Check that retrieving a non-existent or an invalidly
	// named unit fail nicely.
	unit, err = s.service.Unit("wordpress")
	c.Assert(err, ErrorMatches, `can't get unit "wordpress" from service "wordpress": "wordpress" is not a valid unit name`)
	unit, err = s.service.Unit("wordpress/0/0")
	c.Assert(err, ErrorMatches, `can't get unit "wordpress/0/0" from service "wordpress": "wordpress/0/0" is not a valid unit name`)
	unit, err = s.service.Unit("pressword/0")
	c.Assert(err, ErrorMatches, `can't get unit "pressword/0" from service "wordpress": unit not found`)

	// Add another service to check units are not misattributed.
	mysql, err := s.St.AddService("mysql", s.charm)
	c.Assert(err, IsNil)
	_, err = mysql.AddUnit()
	c.Assert(err, IsNil)

	unit, err = s.service.Unit("mysql/0")
	c.Assert(err, ErrorMatches, `can't get unit "mysql/0" from service "wordpress": unit not found`)

	// Check that retrieving all units works.
	units, err := s.service.AllUnits()
	c.Assert(err, IsNil)
	c.Assert(len(units), Equals, 2)
	c.Assert(units[0].Name(), Equals, "wordpress/0")
	c.Assert(units[1].Name(), Equals, "wordpress/1")
}

func (s *ServiceSuite) TestRemoveUnit(c *C) {
	_, err := s.service.AddUnit()
	c.Assert(err, IsNil)
	_, err = s.service.AddUnit()
	c.Assert(err, IsNil)

	// Check that removing a unit works.
	unit, err := s.service.Unit("wordpress/0")
	c.Assert(err, IsNil)
	err = s.service.RemoveUnit(unit)
	c.Assert(err, IsNil)

	units, err := s.service.AllUnits()
	c.Assert(err, IsNil)
	c.Assert(units, HasLen, 1)
	c.Assert(units[0].Name(), Equals, "wordpress/1")

	// Check that removing a non-existent unit fails nicely.
	// TODO BUG is this sane?
	err = s.service.RemoveUnit(unit)
	c.Assert(err, ErrorMatches, `can't unassign unit "wordpress/0" from machine: environment state has changed`)
}

func (s *ServiceSuite) TestReadUnitWithChangingState(c *C) {
	// Check that reading a unit after removing the service
	// fails nicely.
	err := s.St.RemoveService(s.service)
	c.Assert(err, IsNil)
	_, err = s.St.Unit("wordpress/0")
	c.Assert(err, ErrorMatches, `can't get unit "wordpress/0": can't get service "wordpress": service with name "wordpress" not found`)
}

var serviceWatchConfigData = []map[string]interface{}{
	{},
	{"foo": "bar", "baz": "yadda"},
	{"baz": "yadda"},
}

func (s *ServiceSuite) TestWatchConfig(c *C) {
	config, err := s.service.Config()
	c.Assert(err, IsNil)
	c.Assert(config.Keys(), HasLen, 0)

	configWatcher := s.service.WatchConfig()
	defer func() {
		c.Assert(configWatcher.Stop(), IsNil)
	}()

	// Two change events.
	config.Set("foo", "bar")
	config.Set("baz", "yadda")
	changes, err := config.Write()
	c.Assert(err, IsNil)
	c.Assert(changes, DeepEquals, []state.ItemChange{{
		Key:      "baz",
		Type:     state.ItemAdded,
		NewValue: "yadda",
	}, {
		Key:      "foo",
		Type:     state.ItemAdded,
		NewValue: "bar",
	}})
	time.Sleep(100 * time.Millisecond)
	config.Delete("foo")
	changes, err = config.Write()
	c.Assert(err, IsNil)
	c.Assert(changes, DeepEquals, []state.ItemChange{{
		Key:      "foo",
		Type:     state.ItemDeleted,
		OldValue: "bar",
	}})

	for _, want := range serviceWatchConfigData {
		select {
		case got, ok := <-configWatcher.Changes():
			c.Assert(ok, Equals, true)
			c.Assert(got.Map(), DeepEquals, want)
		case <-time.After(200 * time.Millisecond):
			c.Fatalf("didn't get change: %#v", want)
		}
	}

	select {
	case got := <-configWatcher.Changes():
		c.Fatalf("got unexpected change: %#v", got)
	case <-time.After(100 * time.Millisecond):
	}
}

func (s *ServiceSuite) TestWatchConfigIllegalData(c *C) {
	configWatcher := s.service.WatchConfig()
	defer func() {
		c.Assert(configWatcher.Stop(), ErrorMatches, "unmarshall error: YAML error: .*")
	}()

	// Receive empty change after service adding.
	select {
	case got, ok := <-configWatcher.Changes():
		c.Assert(ok, Equals, true)
		c.Assert(got.Map(), DeepEquals, map[string]interface{}{})
	case <-time.After(100 * time.Millisecond):
		c.Fatalf("unexpected timeout")
	}

	// Set config to illegal data.
	_, err := s.zkConn.Set("/services/service-0000000000/config", "---", -1)
	c.Assert(err, IsNil)

	select {
	case _, ok := <-configWatcher.Changes():
		c.Assert(ok, Equals, false)
	case <-time.After(100 * time.Millisecond):
	}
}

var serviceExposedTests = []struct {
	test func(s *state.Service) error
	want bool
}{
	{func(s *state.Service) error { return nil }, false},
	{func(s *state.Service) error { return s.SetExposed() }, true},
	{func(s *state.Service) error { return s.ClearExposed() }, false},
	{func(s *state.Service) error { return s.SetExposed() }, true},
}

func (s *ServiceSuite) TestWatchExposed(c *C) {
	exposedWatcher := s.service.WatchExposed()
	defer func() {
		c.Assert(exposedWatcher.Stop(), IsNil)
	}()

	for i, test := range serviceExposedTests {
		c.Logf("test %d", i)
		err := test.test(s.service)
		c.Assert(err, IsNil)
		select {
		case got, ok := <-exposedWatcher.Changes():
			c.Assert(ok, Equals, true)
			c.Assert(got, Equals, test.want)
		case <-time.After(200 * time.Millisecond):
			c.Fatalf("didn't get change: %#v", test.want)
		}
	}

	select {
	case got := <-exposedWatcher.Changes():
		c.Fatalf("got unexpected change: %#v", got)
	case <-time.After(100 * time.Millisecond):
	}
}

func (s *ServiceSuite) TestWatchExposedContent(c *C) {
	exposedWatcher := s.service.WatchExposed()
	defer func() {
		c.Assert(exposedWatcher.Stop(), IsNil)
	}()

	s.service.SetExposed()
	select {
	case got, ok := <-exposedWatcher.Changes():
		c.Assert(ok, Equals, true)
		c.Assert(got, Equals, true)
	case <-time.After(200 * time.Millisecond):
		c.Fatalf("didn't get change: %#v", true)
	}

	// Re-set exposed with some data.
	_, err := s.zkConn.Set("/services/service-0000000000/exposed", "some: data", -1)
	c.Assert(err, IsNil)

	select {
	case got := <-exposedWatcher.Changes():
		c.Fatalf("got unexpected change: %#v", got)
	case <-time.After(200 * time.Millisecond):
	}
}

var serviceUnitTests = []struct {
	testOp string
	idx    int
}{
	{"none", 0},
	{"add", 0},
	{"add", 1},
	{"remove", 0},
}

func (s *ServiceSuite) TestWatchUnits(c *C) {
	unitsWatcher := s.service.WatchUnits()
	defer func() {
		c.Assert(unitsWatcher.Stop(), IsNil)
	}()
	units := make([]*state.Unit, 2)

	for i, test := range serviceUnitTests {
		c.Logf("test %d", i)
		var want *state.ServiceUnitsChange
		switch test.testOp {
		case "none":
			want = &state.ServiceUnitsChange{}
		case "add":
			var err error
			units[test.idx], err = s.service.AddUnit()
			c.Assert(err, IsNil)
			want = &state.ServiceUnitsChange{[]*state.Unit{units[test.idx]}, nil}
		case "remove":
			err := s.service.RemoveUnit(units[test.idx])
			c.Assert(err, IsNil)
			want = &state.ServiceUnitsChange{nil, []*state.Unit{units[test.idx]}}
			units[test.idx] = nil
		}
		select {
		case got, ok := <-unitsWatcher.Changes():
			c.Assert(ok, Equals, true)
			c.Assert(got, DeepEquals, want)
		case <-time.After(200 * time.Millisecond):
			c.Fatalf("didn't get change: %#v", want)
		}
	}

	select {
	case got := <-unitsWatcher.Changes():
		c.Fatalf("got unexpected change: %#v", got)
	case <-time.After(100 * time.Millisecond):
	}
}
