package post

import "cleanandclean/internal/adapter/interfaces"

func UpdatePost(ctx interfaces.IContext) {
	ctx.JSON(200, map[string]string{"message": "update post"})
}
