package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	APP_KEY    = "eddycjy"
	APP_SECRET = "go-programming-tour-book"
)

type AccessToken struct {
	Token string `json:"token"`
}

type API struct {
	URL string
}

// 创建一个API对象
func NewAPI(url string) *API {
	return &API{URL: url}
}

// 请求通用方法
func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (a *API) httpPostForm(ctx context.Context, path string,data url.Values) ([]byte, error) {
	resp, err := http.PostForm(fmt.Sprintf("%s/%s", a.URL, path), data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}

// 获取accessToken
func (a *API) getAccessToken(ctx context.Context) (string, error) {
	body, err := a.httpPostForm(ctx, "auth",url.Values{"app_key": {APP_KEY}, "app_secret": {APP_SECRET}})
	if err != nil {
		return "", err
	}
	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

// 获取标签列表
func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s", "api/v1/tags", token))
	if err != nil {
		return nil, err
	}

	return body, nil
}




