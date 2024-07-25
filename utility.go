package jsontime

import "time"

func SecResNow() SecRes {
	return SecRes{Time: time.Now()}
}

func MsResNow() MsRes {
	return MsRes{Time: time.Now()}
}

func UsResNow() UsRes {
	return UsRes{Time: time.Now()}
}

func NsResNow() NsRes {
	return NsRes{Time: time.Now()}
}
