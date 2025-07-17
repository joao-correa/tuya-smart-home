package usecases

import (
	"github.com/google/wire"

	"smart-home/modules/tuya_device/api"
)

var AllUsecases = wire.NewSet(
	NewBuildMessageUsecase, wire.Bind(new(api.BuildMessageUsecase), new(*buildMessageUsecase)),
	NewApplySceneUsecase, wire.Bind(new(api.ApplySceneUsecase), new(*applySceneUsecase)),
)
