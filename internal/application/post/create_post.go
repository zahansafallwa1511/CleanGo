package post

import "cleanandclean/internal/adapter/interfaces"

func CreatePost(ctx interfaces.IContext) {
	ctx.JSON(201, map[string]string{"message": "create post"})
}
