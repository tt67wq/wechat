package oauth2

import (
	mpoauth2 "github.com/charsunny/wechat/mp/oauth2"
	"net/http"
	"net/url"
	"strconv"
)

type AuthFuncInfo struct {
	Category struct{
		Id int `json:"id"`
	} `json:"funcscope_category"`
}

type AuthorizationInfo struct {
	AppId string `json:"authorizer_appid"`
	AccessToken string `json:"authorizer_access_token"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"authorizer_refresh_token"`
	FuncInfo [] *AuthFuncInfo `json:"func_info"`

}

// AuthWebURL 生成网页授权地址.
//  appId:       开放平台
// 	preAuthCode: 预授权code， 从平台获取
// authType : 1则商户点击链接后，手机端仅展示公众号、2表示仅展示小程序，3表示公众号和小程序都展示。如果为未指定，则默认小程序和公众号都展示
//  redirectURI: 授权后重定向的回调链接地址
func AuthWebURL(appId, redirectURI, preAuthCode string, authType int) string {
	return "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&pre_auth_code" + preAuthCode +
		"&auth_type" + strconv.Itoa(authType)
}

// AuthH5Link 生成微信内点击的授权地址.
//  appId:       开放平台
// 	preAuthCode: 预授权code， 从平台获取
// 	authType : 1则商户点击链接后，手机端仅展示公众号、2表示仅展示小程序，3表示公众号和小程序都展示。如果为未指定，则默认小程序和公众号都展示
//  redirectURI: 授权后重定向的回调链接地址
func AuthWechatLink(appId, redirectURI, preAuthCode string, authType int) string {
	return "https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&component_appid=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&pre_auth_code" + preAuthCode +
		"&auth_type" + strconv.Itoa(authType) +
		"#wechat_redirect"
}

// Auth 检验授权凭证 access_token 是否有效.
//  accessToken: 网页授权接口调用凭证
//  openId:      用户的唯一标识
//  httpClient:  如果不指定则默认为 util.DefaultHttpClient
func Auth(accessToken, openId string, httpClient *http.Client) (valid bool, err error) {
	return mpoauth2.Auth(accessToken, openId, httpClient)
}


