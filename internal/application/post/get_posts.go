package post

import "cleanandclean/internal/adapter/interfaces"

func GetPosts(ctx interfaces.IContext) {
	ctx.JSON(200, map[string]string{"message": "get posts"})
}
