package externals

import (
	"smart-home/modules/tuya_device/api"
)

type ScenesRepo interface {
	LoadScenes() (*api.Scenes, error)
}
