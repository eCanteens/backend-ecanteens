package product

import (
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func addFeedbackService(body *FeedbackScheme, userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return err
	}

	feedbacks, err := checkFeedback(userId, uint(id))

	if err != nil {
		return err
	}

	if len(*feedbacks) > 0 {
		return updateFeedback(*(*feedbacks)[0].Id.Id, body)
	} else {
		feedback := &models.ProductFeedback{
			UserId: userId,
			ProductId: uint(id),
			Like: *body.Like,
		}

		return createFeedback(feedback)
	}
}

func removeFeedbackService(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return err
	}

	return deleteFeedback(userId, uint(id))
}