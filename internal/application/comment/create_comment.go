package comment

import "cleanandclean/internal/adapter/interfaces"

func CreateComment(ctx interfaces.IContext) {
	ctx.JSON(201, map[string]string{"message": "create comment"})
}
