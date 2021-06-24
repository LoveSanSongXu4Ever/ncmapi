package necmapi

import (
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
	necmcrypto "github.com/benmooo/necm-api/crypto"
)

// 必选参数 :
// phone: 手机号码
// password: 密码
//
// 可选参数 :
// countrycode: 国家码，用于国外手机号登录，例如美国传入：1
// md5_password: md5加密后的密码,传入后 password 将失效
func (api *NeteaseAPI) LoginPhone(phone string, password string) (*APIResponse, error) {
	passwd := necmcrypto.Md5Hex(password)
	payload := defaultLoginPayload().Set("phone", phone).Set("password", passwd)

	opt := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	resp, err := api.Req(http.MethodPost, APIRoutes["login_cellphone"], payload, opt)

	return resp, err
}

func defaultLoginPayload() hm {
	return hm{
		"countrycode":   "86",
		"rememberLogin": true,
	}
}
