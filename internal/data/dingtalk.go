package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/dingtalk/contact_1_0"
	"github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	v1 "github.com/rogerogers/dingtalk-ops/api/helloworld/v1"
	"github.com/rogerogers/dingtalk-ops/internal/biz"
	"github.com/rogerogers/dingtalk-ops/internal/conf"
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
			log.Fatal(_err)
		}
		contactClient, _err = contact_1_0.NewClient(config)
		if _err != nil {
			log.Fatal(_err)
		}
	})
}

type dingtalkRepo struct {
	data *Data
	log  *log.Helper
}

type UserInfoResult struct {
	Result v1.GetUserInfoByUserIdReply `json:"result"`
}

// NewDingtalkRepo .
func NewDingtalkRepo(data *Data, logger log.Logger) biz.DingtalkRepo {
	return &dingtalkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

var cacheCtx = context.Background()

//GetDingtalkToken 获取钉钉企业token .
func (r *dingtalkRepo) GetDingtalkToken() (string, error) {
	key := fmt.Sprintf("%sdingtalk-token", r.data.CacheConfig["prefix"])
	getValue, err := r.data.Cache.Get(cacheCtx, key).Result()
	if err == redis.Nil {
		token := r.GetAccessToken()
		r.data.Cache.Set(cacheCtx, key, token, 3600*time.Second)
		return token, nil
	} else if err != nil {
		return "", errors.New(500, "SERVER_ERROR", "redis connect error")
	} else {
		return getValue, nil
	}
}

//GetUserToken 获取用户token
func GetUserToken(code string, ding *conf.Dingtalk) (token string, err error) {
	getUserTokenRequest := oauth2_1_0.GetUserTokenRequest{
		ClientId:     tea.String(ding.Key),
		ClientSecret: tea.String(ding.Secret),
		Code:         tea.String(code),
		GrantType:    tea.String("authorization_code"),
	}
	result, err := client.GetUserToken(&getUserTokenRequest)
	if err != nil {
		return "", errors.New(400, "PARAM_ERROR", "param error")
	} else {
		return *result.Body.AccessToken, nil
	}
}

func (r *dingtalkRepo) GetUserToken(ctx context.Context, d *v1.GetUserTokenRequest) (*v1.GetUserTokenReply, error) {
	result, err := GetUserToken(d.AuthCode, r.data.DingtalkConfig)
	if err != nil {
		return &v1.GetUserTokenReply{}, err
	} else {
		return &v1.GetUserTokenReply{
			AccessToken: result,
		}, nil
	}
}

//GetAccessToken 获取企业access_token
func (r *dingtalkRepo) GetAccessToken() (token string) {
	getAccessTokenRequest := &oauth2_1_0.GetAccessTokenRequest{
		AppKey:    tea.String(r.data.DingtalkConfig.Key),
		AppSecret: tea.String(r.data.DingtalkConfig.Secret),
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
			log.Fatal(err)
		}
	}
	return token
}

func GetUserInfoByToken(token string) (*string, error) {
	header := contact_1_0.GetUserHeaders{
		XAcsDingtalkAccessToken: &token,
	}
	unionId := "me"
	res, err := contactClient.GetUserWithOptions(&unionId, &header, &util.RuntimeOptions{})
	if err != nil {
		return &unionId, errors.New(400, "PARAM_ERROR", "param error")
	}
	return res.Body.UnionId, nil
}

func (r *dingtalkRepo) GetUserInfoByToken(ctx context.Context, d *v1.GetUserInfoByTokenRequest) (*v1.GetUserInfoByTokenReply, error) {
	unionId, err := GetUserInfoByToken(d.AccessToken)
	if err != nil {
		return &v1.GetUserInfoByTokenReply{}, err
	}

	return &v1.GetUserInfoByTokenReply{
		UnionId: *unionId,
	}, nil
}

func GetUserIdByUnionId(accessToken string, unionid string) (userId string, err error) {
	requestBody, err := json.Marshal(map[string]string{"unionid": unionid})
	if err != nil {
		return "", errors.New(400, "PARAM_ERROR", "param error")

	}
	res, err := http.Post(fmt.Sprintf("https://oapi.dingtalk.com/topapi/user/getbyunionid?access_token=%s", accessToken), "application/json", strings.NewReader(string(requestBody)))
	if err != nil {
		return "", errors.New(400, "PARAM_ERROR", "param error")
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New(400, "PARAM_ERROR", "param error")
	}
	return string(resBody), nil
}

func (r *dingtalkRepo) GetUserIdByUnionId(ctx context.Context, d *v1.GetUserIdByUnionIdRequest) (*v1.GetUserIdByUnionIdReply, error) {
	token, err := r.GetDingtalkToken()
	if err != nil {
		return &v1.GetUserIdByUnionIdReply{}, err
	}
	userId, err := GetUserIdByUnionId(token, d.UnionId)

	if err != nil {
		return &v1.GetUserIdByUnionIdReply{}, err
	}
	return &v1.GetUserIdByUnionIdReply{
		UserId: userId,
	}, nil
}

func GetUserInfoByUserId(accessToken string, userId string) (userInfoRes *UserInfoResult, err error) {
	requestBody, err := json.Marshal(map[string]string{"userid": userId})
	if err != nil {
		return &UserInfoResult{}, err
	}
	res, err := http.Post(fmt.Sprintf("https://oapi.dingtalk.com/topapi/v2/user/get?access_token=%s", accessToken), "application/json", strings.NewReader(string(requestBody)))

	if err != nil {
		return &UserInfoResult{}, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		errors.New(500, "SERVER_ERROR", "server error")
	}
	err = json.Unmarshal(resBody, &userInfoRes)
	if userInfoRes.Result.Unionid == "" || err != nil {
		return &UserInfoResult{}, errors.New(400, "PARAM_ERROR", "param error")
	}
	return
}

//GetUserInfoByUserId 根据用户id获取用户信息
func (r *dingtalkRepo) GetUserInfoByUserId(ctx context.Context, d *v1.GetUserInfoByUserIdRequest) (*v1.GetUserInfoByUserIdReply, error) {
	token, err := r.GetDingtalkToken()
	if err != nil {
		return &v1.GetUserInfoByUserIdReply{}, err
	}
	result, err := GetUserInfoByUserId(token, d.UserId)
	if err != nil {
		return &v1.GetUserInfoByUserIdReply{}, err
	} else {
		return &result.Result, nil
	}
}
