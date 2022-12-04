package proxy

import (
	"k8s.io/klog/v2"
	"time"
)

type IUserController interface {
	Login(username, password string) error
	Register(username, password string) error
}

type UserController struct {
}

func (u UserController) Login(username, password string) error {
	return nil
}

func (u UserController) Register(username, password string) error {
	//TODO implement me
	return nil
}

// 代理类
type UserControllerProxy struct {
	uc               *UserController
	metricsCollector *MetricsCollector
}

// 代理类和原始类实现相同的接口
func (u UserControllerProxy) Login(username, password string) error {
	startTime := time.Now()

	// 委托原始类
	err := u.uc.Login(username, password)
	if err != nil {
		klog.Errorf("Failed to login for %v. Error: %v", username, err)
		return err
	}

	requestInfo := &RequestInfo{
		ApiName:        "login",
		ResponseTime:   time.Since(startTime),
		StartTimeStamp: startTime,
	}
	u.metricsCollector.RecordRequest(requestInfo)
	return nil
}

func (u UserControllerProxy) Register(username, password string) error {
	startTime := time.Now()

	// 委托原始类
	err := u.uc.Register(username, password)
	if err != nil {
		klog.Errorf("Failed to register for %v. Error: %v", username, err)
		return err
	}

	requestInfo := &RequestInfo{
		ApiName:        "register",
		ResponseTime:   time.Since(startTime),
		StartTimeStamp: startTime,
	}
	u.metricsCollector.RecordRequest(requestInfo)
	return nil
}

type MetricsCollector struct {
	requestInfos []*RequestInfo
}

type RequestInfo struct {
	ApiName        string
	ResponseTime   time.Duration
	StartTimeStamp time.Time
}

func (mc *MetricsCollector) RecordRequest(ri *RequestInfo) {
	// TODO
}

func NewUserControllerProxy(uc *UserController) *UserControllerProxy {
	return &UserControllerProxy{
		uc:               uc,
		metricsCollector: &MetricsCollector{},
	}
}
