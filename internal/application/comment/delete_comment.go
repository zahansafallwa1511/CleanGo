package comment

import "cleanandclean/internal/adapter/interfaces"

func DeleteComment(ctx interfaces.IContext) {
	ctx.JSON(200, map[string]string{"message": "delete comment"})
}
