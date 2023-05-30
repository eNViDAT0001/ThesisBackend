package comment

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/grpc/grpc_base"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/comment/convert"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/comment/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	proto "github.com/eNViDAT0001/Thesis/Backend/thesis_proto"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (s commentHandler) CreateComment() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CreateCommentReq
		if err := cc.ShouldBind(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		productID, _ := strconv.Atoi(cc.Param("product_id"))
		userID, _ := strconv.Atoi(cc.Param("user_id"))

		input.UserID = uint(userID)
		input.ProductID = uint(productID)

		inputSto, err := convert.CreateReqToCreateCommentInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		commentID, err := s.commentUC.CreateComment(newCtx, inputSto, input.Media)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		_, err = grpc_base.GetServices().RecommenderService.AddComment(newCtx, &proto.CommentReq{
			UserId:    int32(userID),
			ProductId: int32(productID),
			Rating:    int32(input.Rating),
		})
		if err != nil {
			log.Printf("Error while adding comment to recommender service: %v", err)
		}
		result := map[string]interface{}{
			"CommentID": commentID,
		}
		cc.Ok(result)
	}
}
