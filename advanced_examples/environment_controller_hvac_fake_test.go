package examples

import "strings"

type FakeHVAC struct {
	state       map[string]bool
	temperature int
}

func NewFakeHardware() *FakeHVAC {
	return &FakeHVAC{
		state: map[string]bool{
			"heater":          false,
			"blower":          false,
			"cooler":          false,
			"high-temp-alarm": false,
			"low-temp-alarm":  false,
		},
	}
}

func (self *FakeHVAC) ActivateHeater()               { self.state["heater"] = true }
func (self *FakeHVAC) ActivateBlower()               { self.state["blower"] = true }
func (self *FakeHVAC) ActivateCooler()               { self.state["cooler"] = true }
func (self *FakeHVAC) ActivateHighTemperatureAlarm() { self.state["high"] = true }
func (self *FakeHVAC) ActivateLowTemperatureAlarm()  { self.state["low"] = true }

func (self *FakeHVAC) DeactivateHeater()               { self.state["heater"] = false }
func (self *FakeHVAC) DeactivateBlower()               { self.state["blower"] = false }
func (self *FakeHVAC) DeactivateCooler()               { self.state["cooler"] = false }
func (self *FakeHVAC) DeactivateHighTemperatureAlarm() { self.state["high"] = false }
func (self *FakeHVAC) DeactivateLowTemperatureAlarm()  { self.state["low"] = false }

func (self *FakeHVAC) IsHeating() bool            { return self.state["heater"] }
func (self *FakeHVAC) IsBlowing() bool            { return self.state["blower"] }
func (self *FakeHVAC) IsCooling() bool            { return self.state["cooler"] }
func (self *FakeHVAC) HighTemperatureAlarm() bool { return self.state["high"] }
func (self *FakeHVAC) LowTemperatureAlarm() bool  { return self.state["low"] }

func (self *FakeHVAC) SetCurrentTemperature(value int) { self.temperature = value }
func (self *FakeHVAC) CurrentTemperature() int         { return self.temperature }

// String returns the status of each hardware component encoded in a single space-delimited string.
// UPPERCASE components are activated.
// lowercase components are deactivated.
func (self *FakeHVAC) String() string {
	current := []string{"heater", "blower", "cooler", "low", "high"}
	for i, component := range current {
		if self.state[component] {
			current[i] = strings.ToUpper(current[i])
		}
	}
	return strings.Join(current, " ")
}
