package post

import "cleanandclean/internal/adapter/interfaces"

func GetPost(ctx interfaces.IContext) {
	ctx.JSON(200, map[string]string{"message": "get post"})
}
