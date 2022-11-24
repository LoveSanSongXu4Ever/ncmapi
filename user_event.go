package ncmapi

import "net/http"

// 说明 : 登录后调用此接口 , 传入用户 id, 可以获取用户动态

func (api *NeteaseAPI) UserEvent(uid int, opts ...hm) (*APIResponse, error) {
	data := hm{
		"limit":  30,
		"before": 0,
	}.Merge(opts...).Merge(hm{"uid": uid})

	resp, err := api.Req(http.MethodPost, APIRoutes["user_event"], data, nil)
	return resp, err
}

/*func (api *NeteaseAPI) UserEvent(uid int) (*APIResponse, error) {
	u := replaceAllRouteParams(APIRoutes["user_event"], strconv.Itoa(uid))

	resp, err := api.Req(http.MethodPost, u, nil, nil)
	return resp, err
}*/
