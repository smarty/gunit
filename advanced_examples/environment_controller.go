package examples

type EnvironmentController struct {
	hardware HVAC
}

func NewController(hardware HVAC) *EnvironmentController {
	hardware.DeactivateBlower()
	hardware.DeactivateHeater()
	hardware.DeactivateCooler()
	hardware.DeactivateHighTemperatureAlarm()
	hardware.DeactivateLowTemperatureAlarm()
	return &EnvironmentController{hardware: hardware}
}

func (self *EnvironmentController) Regulate() {
	temperature := self.hardware.CurrentTemperature()

	if temperature >= WAY_TOO_HOT {
		self.hardware.ActivateHighTemperatureAlarm()
	} else if temperature <= WAY_TOO_COLD {
		self.hardware.ActivateLowTemperatureAlarm()
	}

	if temperature >= TOO_HOT {
		self.hardware.DeactivateHeater()
		self.hardware.ActivateBlower()
		self.hardware.ActivateCooler()
	} else if temperature <= TOO_COLD {
		self.hardware.DeactivateCooler()
		self.hardware.ActivateBlower()
		self.hardware.ActivateHeater()
	}
}

const (
	WAY_TOO_HOT  = 80
	TOO_HOT      = 70
	TOO_COLD     = 60
	WAY_TOO_COLD = 50
	COMFORTABLE  = (TOO_HOT + TOO_COLD) / 2
)
