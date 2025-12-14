package comment

import "cleanandclean/internal/adapter/interfaces"

func GetComments(ctx interfaces.IContext) {
	ctx.JSON(200, map[string]string{"message": "get comments"})
}
