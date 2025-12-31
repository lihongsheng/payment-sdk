package refund

import (
	"context"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/stretchr/testify/assert"
	"testing"
)

// {"wxpay": {"secret": "1bde81c041fbaPF2K0d21ccAa0147c8f",
//"public_key": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB\n-----END PUBLIC KEY-----",
//"public_key_id": "PUB_KEY_ID_0117134125362025092400111793000600",
//"cert_private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49\n-----END PRIVATE KEY-----",
//"certificate_serial_number": "59400142A58A56A0AF191C85A2C91CF434AEF32A"}}

func TestJsapi_Pay(t *testing.T) {
	api, err := NewRefund(config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			Private: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			PrivateNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicNumber:  "PUB_KEY_ID_0117134125362025092400111793000600",
			Public: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APISecret: "1bde81c041fbaPF2K0d21ccAa0147c8f",
	})

	assert.NoError(t, err)
	ctx := context.Background()
	req := &dto.RefundRequest{
		RefundNo:  "175879416443300038879",
		TradeNo:   "",
		OrderNo:   "175879416443300038879",
		Reason:    "reason",
		NotifyUrl: "https://api-cabinet.test.jianxindianzi.com/public/v1/callback/refund/5/175879416443300038879",
		Amount: dto.Amount{
			Total:    1,
			Currency: "CNY",
		},
		Goods: nil,
		OrderAmount: dto.Amount{
			Total:    1,
			Currency: "CNY",
		},
	}
	resp, err := api.Refund(ctx, req)
	t.Log(resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestStatus(t *testing.T) {
	assert.Equal(t, enum.Status_Status_UNKNOWN, until.PaymentStatus["UNKNOWN"])
}
