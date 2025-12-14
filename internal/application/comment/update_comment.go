package comment

import "cleanandclean/internal/adapter/interfaces"

func UpdateComment(ctx interfaces.IContext) {
	ctx.JSON(200, map[string]string{"message": "update comment"})
}
