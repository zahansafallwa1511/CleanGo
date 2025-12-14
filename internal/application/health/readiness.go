package health

import "cleanandclean/internal/adapter/interfaces"

func Readiness(ctx interfaces.IContext) {
	ctx.JSON(200, map[string]string{"status": "ready"})
}
