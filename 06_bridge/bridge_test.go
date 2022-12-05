package bridge

import "testing"

func TestSevereNotification_Notify(t *testing.T) {
	wms := &WechatMsgSender{}
	sn := &SevereNotification{msgSender: wms}
	sn.Notify("test")

	sn.msgSender = &MobileMsgSender{}
	sn.Notify("test")
}
