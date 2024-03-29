package cart_items

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s *cartItemHandler) DeleteCartItem() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		cartID, _ := strconv.Atoi(cc.Param("cart_id"))
		itemID, _ := strconv.Atoi(cc.Param("cart_item_id"))

		err := s.cartItemUC.DeleteCartItem(newCtx, uint(cartID), uint(itemID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NotFound()
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok("Upsert CartItem Success")
	}
}
