package observerexample

import "testing"

func TestUserController_Register(t *testing.T) {
	uc := &UserController{
		userService: &UserService{Id: 1},
		observers:   make([]IRegisterObserver, 0),
	}

	promotionObserver := &PromotionRegisterObserver{promotionService: &PromotionService{}}
	notificationObserver := &NotificationRegisterObserver{notificationService: &NotificationService{}}
	uc.AddObserver(promotionObserver)
	uc.AddObserver(notificationObserver)

	uc.Register("test", "test")
}
