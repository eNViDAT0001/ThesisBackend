package order

import (
	"context"
	"errors"
	"fmt"
	io2 "github.com/eNViDAT0001/Thesis/Backend/delivery/http/order/order/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/event_background"
	"github.com/eNViDAT0001/Thesis/Backend/external/html_template"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	smtpIO "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/socket"
	io3 "github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var viper = wrap_viper.GetViper()

func (s *orderHandler) UpdateOrderStatus() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io2.UpdateOrderStatusReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		orderID, err := strconv.Atoi(cc.Param("order_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		if input.Status == entities.DeliveredOrder {

			if input.Image == "" {
				cc.BadRequest(errors.New("image is required"))
				return
			}
			err = s.orderUC.UpdateDeliveredOrderStatus(newCtx, uint(orderID), input.Image)
			if err != nil {
				cc.ResponseError(err)
				return
			}

			event_background.AddBackgroundJobs(false, event_background.NewJob(func(ctx context.Context) error {
				order, err := s.orderUC.GetByOrderID(newCtx, uint(orderID))
				if err != nil {
					return err
				}
				user, err := s.useUC.GetUserDetailByID(newCtx, order.UserID)
				if err != nil {
					return err
				}
				email := smtpIO.EmailForm{
					Subject:     "Your order has arrived",
					Content:     html_template.GetOrderSuccessTemplate(fmt.Sprintf(`localhost:3000/user/order/%d`, orderID), order),
					To:          []string{*user.Email},
					Cc:          nil,
					Bcc:         nil,
					AttachFiles: nil,
				}
				seen := false
				newNotification, err := s.notifyUC.CreateNotification(ctx, io.NotificationInput{
					ID:      0,
					UserID:  user.ID,
					Content: "Your order has arrived",
					Seen:    &seen,
					URL:     fmt.Sprintf(`%s:%s/user/order/%d`, viper.GetString("FE.HOST"), viper.GetString("FE.PORT"), orderID),
					Image:   "http://res.cloudinary.com/damzcas3k/image/upload/v1684051785/Product/itl4m7o3jsmtqb2mhtt1.png",
				})

				err = socket.GetManager().EmitNotify(io3.NotificationNew, newNotification, strconv.Itoa(int(user.ID)))
				if err != nil {
					log.Fatal(err)
				}
				return s.smtpUC.SendEmail(ctx, email)
			}))

			cc.Ok("Update Status success")
			return
		}

		err = s.orderUC.UpdateOrderStatus(newCtx, uint(orderID), input.Status)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Status success")
	}
}
