package util

import "github.com/golang/glog"

func ErrorHandler(s string, err error) {
	if err != nil {
		glog.Fatalf(s, err)
	}
	return
}
