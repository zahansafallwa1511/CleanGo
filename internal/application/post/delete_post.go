package post

import "cleanandclean/internal/adapter/interfaces"

func DeletePost(ctx interfaces.IContext) {
	ctx.JSON(200, map[string]string{"message": "delete post"})
}
