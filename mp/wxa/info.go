package wxa

import (
	"github.com/charsunny/wechat/mp/core"
)

// 获取帐号基本信息
func GetBaseInfo(clt *core.Client) (info *WxaInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token="

	var result struct {
		core.Error
		WxaInfo
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.WxaInfo
	return
}

// 小程序改名
func SetNickName(clt *core.Client, info *WxaNameRequestInfo) (wording string, audit_id int, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/setnickname?access_token="

	var result struct {
		core.Error
		Wording string `json:"wording"`	// 材料说明
		AuditId int `json:"audit_id"`	// 审核单id 若接口未返回audit_id，说明名称已直接设置成功，无需审核；若返回audit_id则名称正在审核中
	}
	if err = clt.PostJSON(incompleteURL, info, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	wording = result.Wording
	audit_id = result.AuditId
	return
}

// 小程序改名
func QuerySetNickStatus(clt *core.Client,  audit_id int) (info *WxaNameResultInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/api_wxa_querynickname?access_token="

	var request = struct {
		AuditId int `json:"audit_id"`
	}{
		AuditId:audit_id,
	}
	var result struct {
		core.Error
		WxaNameResultInfo
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.WxaNameResultInfo
	return
}

// 微信认证名称检测
func CheckNickName(clt *core.Client,  nickname string) (wording string, hit bool, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxverify/checkwxverifynickname?access_token="

	var request = struct {
		NickName string `json:"nick_name"`
	}{
		NickName:nickname,
	}
	var result struct {
		core.Error
		Wording string `json:"wording"`	// 命中关键字的说明描述（给用户看的）
		HitCondition bool `json:"hit_condition"`	// 是否命中关键字策略。若命中，可以选填关键字材料
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	wording = result.Wording
	hit = result.HitCondition
	return
}

// 修改小程序头像
func ModifyHeadImage(clt *core.Client,  head_img_media_id string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/account/modifyheadimage?access_token="

	var request = struct {
		HeadImgMediaId string `json:"head_img_media_id"`
		X1 float64 `json:"x1"`
		X2 float64 `json:"x2"`
		Y1 float64 `json:"y1"`
		Y2 float64 `json:"y2"`
	}{
		HeadImgMediaId:head_img_media_id,
		X1: 0,
		X2: 1,
		Y1: 0,
		Y2: 0,
	}
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 修改小程序功能介绍
func ModifySignature(clt *core.Client,  signature string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/account/modifysignature?access_token="

	var request = struct {
		Signature string `json:"signature"`
	}{
		Signature:signature,
	}
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// TODO: 类目相关接口 && 换绑小程序管理员接口