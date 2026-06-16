package user

import (
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"testing"
)

func TestUser_AuthToken(t *testing.T) {

}

func TestUser_AuthToken2(t *testing.T) {
	body1 := []byte(`{"alipay_system_oauth_token_response":{"access_token":"authusrB3379f70b688f4a63a3d5a5fdbea72X83","auth_start":"2025-10-15 13:19:21","expires_in":1296000,"re_expires_in":2592000,"refresh_token":"authusrB79e8d26898f3401b900b666db1aeeX83","open_id":"083luGU2ef8yeC95sroZ9WGyojnYtj3lBoBA6lY_2PlBrU4"},"sign":"GfIz3Q2Rwk2t72xXohXun9bDUrUxUmLr6ATzbaPTGCtBxzRHcN3JAkWJtdtuD/GFxt8lvYVagrFWb50F6n7CYy2dE6hu/7m19md+y6kV60BeORvKCypKC61EegKB4GjeA6uLKdSnjrZI8vaohrHsjw9SRV8lnQVKi/p50Y8XF35nHGIW4t11AvYMQy6+NyscP9bSDodIFXWxN9XT5vnOtd5av4ij8ZR6MFI9v5aWakAOZNWkM4P+GBsZ9ePCwiGD8Se5puvAjljUi5O3oriRaald0e8gPaz1t3BNwWteY6N91B+4EAlzy2oOmpME7ey6U0EDbXB2auQxbF+LKNbQkQ=="}`)
	body2 := []byte(`{"alipay_system_oauth_token_response":{"access_token":"authusrB3379f70b688f4a63a3d5a5fdbea72X83","auth_start":"2025-10-15 13:19:21","expires_in":"1296000","re_expires_in":"2592000","refresh_token":"authusrB79e8d26898f3401b900b666db1aeeX83","open_id":"083luGU2ef8yeC95sroZ9WGyojnYtj3lBoBA6lY_2PlBrU4"},"sign":"GfIz3Q2Rwk2t72xXohXun9bDUrUxUmLr6ATzbaPTGCtBxzRHcN3JAkWJtdtuD/GFxt8lvYVagrFWb50F6n7CYy2dE6hu/7m19md+y6kV60BeORvKCypKC61EegKB4GjeA6uLKdSnjrZI8vaohrHsjw9SRV8lnQVKi/p50Y8XF35nHGIW4t11AvYMQy6+NyscP9bSDodIFXWxN9XT5vnOtd5av4ij8ZR6MFI9v5aWakAOZNWkM4P+GBsZ9ePCwiGD8Se5puvAjljUi5O3oriRaald0e8gPaz1t3BNwWteY6N91B+4EAlzy2oOmpME7ey6U0EDbXB2auQxbF+LKNbQkQ=="}`)
	var model model.AuthTokenResponse
	err := json.Unmarshal(body1, &model)
	if err != nil {
		t.Error(err)
	}
	t.Log(model.AlipayOpenAuthTokenAppResponse.OpenId)
	t.Log(model.AlipayOpenAuthTokenAppResponse.GetExpiresIn())
	err = json.Unmarshal(body2, &model)
	if err != nil {
		t.Error(err)
	}
	t.Log(model.AlipayOpenAuthTokenAppResponse.GetReExpiresIn())

}
