package ratelimit

func AppKey(route string) string            { return "app:" + route }
func MethodKey(route, method string) string { return "method:" + route + ":" + method }
