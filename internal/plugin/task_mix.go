package plugin

import (
	"github.com/paladin-devops/waypoint-plugin-sdk/component"
)

type mix_TaskLauncher_Authenticator struct {
	component.ConfigurableNotify
	component.TaskLauncher
	component.Documented
}
