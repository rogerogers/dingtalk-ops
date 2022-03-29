package dingtalk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/dingtalk/contact_1_0"
	"github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	v1 "github.com/rogerogers/dingtalk-ops/api/helloworld/v1"
)

var (
	client        *oauth2_1_0.Client
	once          = &sync.Once{}
	contactClient *contact_1_0.Client
)

func init() {
	var _err error
	once.Do(func() {
		config := &openapi.Config{}
		config.Protocol = tea.String("https")
		config.RegionId = tea.String("central")
		client, _err = oauth2_1_0.NewClient(config)
		if _err != nil {
			log.Fatalln(_err)
		}
		contactClient, _err = contact_1_0.NewClient(config)
		if _err != nil {
			log.Fatalln(_err)
		}
	})
}

//GetAccessToken 获取企业access_token
func GetAccessToken() (token string) {
	getAccessTokenRequest := &oauth2_1_0.GetAccessTokenRequest{
		AppKey:    tea.String(os.Getenv("DING_KEY")),
		AppSecret: tea.String(os.Getenv("DING_SECRET")),
	}
	token, tryErr := func() (token string, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		res, _err := client.GetAccessToken(getAccessTokenRequest)
		if _err != nil {
			return "", _err
		}
		return *res.Body.AccessToken, nil
	}()

	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		if !tea.BoolValue(util.Empty(err.Code)) && !tea.BoolValue(util.Empty(err.Message)) {
			// err 中含有 code 和 message 属性，可帮助开发定位问题
			log.Println(err)
		}
	}
	return token
}

//GetUserToken 获取用户token
func GetUserToken(code string) (token string, err error) {
	getUserTokenRequest := oauth2_1_0.GetUserTokenRequest{
		ClientId:     tea.String(os.Getenv("DING_KEY")),
		ClientSecret: tea.String(os.Getenv("DING_SECRET")),
		Code:         tea.String(code),
		GrantType:    tea.String("authorization_code"),
	}
	result, err := client.GetUserToken(&getUserTokenRequest)
	if err != nil {
		return "", err
	} else {
		return *result.Body.AccessToken, nil
	}
}

func GetUserInfoByToken(token string) *string {
	header := contact_1_0.GetUserHeaders{
		XAcsDingtalkAccessToken: &token,
	}
	unionId := "me"
	res, err := contactClient.GetUserWithOptions(&unionId, &header, &util.RuntimeOptions{})
	if err != nil {
		log.Println(err)
	}
	return res.Body.UnionId
}

func GetUserIdByUnionId(unionid string) (userId string, err error) {
	requestBody, err := json.Marshal(map[string]string{"unionid": unionid})
	if err != nil {
		return "", err

	}
	res, err := http.Post(fmt.Sprintf("https://oapi.dingtalk.com/topapi/user/getbyunionid?access_token=%s", GetAccessToken()), "application/json", strings.NewReader(string(requestBody)))
	if err != nil {
		return "", err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil

}

//GetUserInfoByUserId 根据用户id获取用户信息
func GetUserInfoByUserId(userId string) (userInfoRes *UserInfoResult, err error) {
	requestBody, err := json.Marshal(map[string]string{"userid": userId})
	if err != nil {
		return &UserInfoResult{}, err
	}
	res, err := http.Post(fmt.Sprintf("https://oapi.dingtalk.com/topapi/v2/user/get?access_token=%s", GetAccessToken()), "application/json", strings.NewReader(string(requestBody)))

	if err != nil {
		return &UserInfoResult{}, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &UserInfoResult{}, err
	}
	err = json.Unmarshal(resBody, &userInfoRes)
	if err != nil {
		log.Println(err)
	}
	return
}

type UserInfoResult struct {
	Result v1.GetUserInfoByUserIdReply `json:"result"`
}
