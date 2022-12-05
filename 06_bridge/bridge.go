package bridge

import "k8s.io/klog/v2"

type INotification interface {
	Notify(message string)
}

type IMsgSender interface {
	Send(message string)
}

type SevereNotification struct {
	msgSender IMsgSender
}

func (s *SevereNotification) Notify(message string) {
	s.msgSender.Send(message)
}

type UrgentNotification struct {
	msgSender IMsgSender
}

func (u *UrgentNotification) Notify(message string) {
	u.msgSender.Send(message)
}

type MobileMsgSender struct {
}

func (m *MobileMsgSender) Send(message string) {
	klog.Infof("Send msg via mobile.")
}

type WechatMsgSender struct {
}

func (w *WechatMsgSender) Send(message string) {
	klog.Infof("Send msg via wechat.")
}
