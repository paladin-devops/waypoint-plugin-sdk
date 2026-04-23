package plugin

import (
	"github.com/paladin-devops/waypoint-plugin-sdk/component"
)

type mix_Builder_Authenticator struct {
	component.Authenticator
	component.ConfigurableNotify
	component.Builder
	component.Documented
}
