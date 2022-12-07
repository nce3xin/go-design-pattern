package observerexample

// IRegisterObserver 用户注册观察者接口
type IRegisterObserver interface {
	HandleRegisterSuccess(userId int)
}

// PromotionRegisterObserver 发放优惠券观察者
type PromotionRegisterObserver struct {
	promotionService *PromotionService
}

// PromotionService 优惠券服务
type PromotionService struct {
}

// PromotionService 优惠券服务的方法：发放新用户体验优惠券
func (p *PromotionService) issueNewUserExperienceCoupon(userId int) {
	//TODO implement me
}

// HandleRegisterSuccess 实现接口IRegisterObserver
func (p *PromotionRegisterObserver) HandleRegisterSuccess(userId int) {
	p.promotionService.issueNewUserExperienceCoupon(userId)
}

type NotificationRegisterObserver struct {
	notificationService *NotificationService
}

// NotificationService 通知服务
type NotificationService struct {
}

// SendEmail 通知服务的方法：发送欢迎邮件
func (n *NotificationService) SendEmail(userId int, msg string) {
	//TODO implement me
}

// HandleRegisterSuccess 实现接口IRegisterObserver
func (n *NotificationRegisterObserver) HandleRegisterSuccess(userId int) {
	n.notificationService.SendEmail(userId, "Welcome")
}

type UserService struct {
	Id int
}

func (u *UserService) register(telephone, password string) int {
	//TODO implement me
	return u.Id
}

type UserController struct {
	userService *UserService
	observers   []IRegisterObserver
}

func (u *UserController) Register(telephone, password string) int {
	userId := u.userService.register(telephone, password)
	// 同步阻塞的观察者模式
	for _, observer := range u.observers {
		observer.HandleRegisterSuccess(userId)
	}
	return userId
}

func (u *UserController) AddObserver(o IRegisterObserver) {
	u.observers = append(u.observers, o)
}
