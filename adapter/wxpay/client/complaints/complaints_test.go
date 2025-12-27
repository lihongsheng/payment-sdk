package complaints

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestQueryDetail(t *testing.T) {
	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	resp, apiR, err := c.QueryDetail(ctx, "200000020251110170368514500")
	assert.NoError(t, err)
	t.Log(resp)
	if apiR != nil {
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
	t.Log(fmt.Sprintf("%+v", apiR))
}

func TestSearch(t *testing.T) {
	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	resp, apiR, err := c.Search(ctx, SearchComplaintRequest{
		BeginDate: "2025-11-10",
		EndDate:   "2025-11-11",
		Limit:     10,
		Offset:    0,
	})
	assert.NoError(t, err)
	t.Log(fmt.Sprintf("%+v", resp))
	if apiR != nil {
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
	t.Log(fmt.Sprintf("%+v", apiR))
}

func TestReplay(t *testing.T) {
	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	apiR, err := c.Replay(ctx, "200000020251110170368514500", ReplayComplaintRequest{
		ComplaintedMchid: cccc.MchID,
		ResponseContent:  "测试回复",
	})
	assert.NoError(t, err)
	if apiR != nil {
		t.Log(apiR.Response.StatusCode)
	}
	t.Log(fmt.Sprintf("%+v", apiR))
}

func TestRefund(t *testing.T) {
	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	apiR, err := c.Refund(ctx, "200000020251110170368514500", RefundHandlerRequest{
		Action:          string(RefundType_APPROVE),
		LaunchRefundDay: 1,
		RejectReason:    "测试",
		RejectMediaList: nil,
		Remark:          "测试",
	})
	assert.NoError(t, err)
	if apiR != nil {
		t.Log(apiR.Response.StatusCode)
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
	t.Log(fmt.Sprintf("%+v", apiR))
}

func TestComplete(t *testing.T) {
	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	apiR, err := c.Complete(ctx, "200000020251110170368514500", CompleteRequest{
		ComplaintedMchid: "1713412536",
	})
	assert.NoError(t, err)
	if apiR != nil {
		t.Log(apiR.Response.StatusCode)
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
	t.Log(fmt.Sprintf("%+v", apiR))
}

// proxy.test.jianxindianzi.com

func TestAddNoticeUrl(t *testing.T) {
	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	// proxy.test.jianxindianzi.com
	u := "https://proxy.test.jianxindianzi.com/public/v1/callback/payment/17"
	resp, apiR, err := c.AddNoticeUrl(ctx, AddComplaintNoticeUrl{Url: u})
	assert.NoError(t, err)
	if apiR != nil {
		t.Log(apiR.Response.StatusCode)
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
	t.Log(fmt.Sprintf("%+v", apiR))
	t.Log(fmt.Sprintf("%+v", resp))
}

func TestQueryNoticeUrl(t *testing.T) {
	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	// proxy.test.jianxindianzi.com
	//u := "https://proxy.test.jianxindianzi.com/public/v1/callback/payment/17"
	resp, apiR, err := c.QueryNoticeUrl(ctx)
	assert.NoError(t, err)
	if apiR != nil {
		t.Log(apiR.Response.StatusCode)
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
	t.Log(fmt.Sprintf("%+v", apiR))
	t.Log(fmt.Sprintf("%+v", resp))
	if err != nil {
		t.Log(err.Error())
	}
}

type Cert struct {
	// 微信私钥或者其他平台私钥
	CertPrivateKey string `json:"cert_private_key,omitempty"`
	// 微信证书序列号
	CertificateSerialNumber string `json:"certificate_serial_number,omitempty"`
	// 微信公钥ID
	PublicKeyID string `json:"public_key_id,omitempty"`
	// 微信公钥
	PublicKey string `json:"public_key,omitempty"`
	// 加签盐值
	Secret string `json:"secret,omitempty"`
	// sence id 微信转账需要
	SenseID string `json:"sense_id,omitempty"`
	// 富有协议类型
	ProtocolType string `json:"protocol_type,omitempty"`
}

func TestQueryNoticeUrl2(t *testing.T) {
	var cert = &Cert{}
	certStr := `{"secret": "ZXCBZCVNCVFGHFGHFH14246745646456", "public_key": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2vfo5RuV4aUcv/asYlKl\nErl0hB6r8Tq7ytCpiOurYa4ikFN141hXPoq2rFCvFUAIrr0AbCrwwQPsR5R8Na0U\nhAzpB6smXks1uYiXzdcAGtxlMsVUvmRrWT2gY4dalv/Ia3hFPKeWVW7feelxYW8X\nDyOeVjennYI5ccFeSBqwovCd8Zh8temUQTCGMi4FJvEkhiTXq2Yjr969exzBu6PK\n6Ro6AeE/cGngw4GhYzcdgS98uXlBZnIbqIqMl6l1rACtE0LzoXPgnSTa0bE82dEn\nHxtwKhwwo+QezAbcbYAyk480+h+T0T7a7uzuZdtynnEqdjfPwMsUjhaiNcAhIy9A\nZQIDAQAB\n-----END PUBLIC KEY-----", "complaint_url": "https://huijunlikeji.com", "public_key_id": "PUB_KEY_ID_0117331865082025112400111654001605", "cert_private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC0c2xZmq37+Tpo\nC9PyLFVRK15SumXK1cofwClTOU66nfHHz9IS+aheChk5f2V8LAy0PzCLdes5+UPv\nHZA44dTmR1QsPlTdU1zFSPWwtRbHkCziFjh7k+XbRjejOfokPtULcZzAgB4g/u05\n3qNUwi9YgcsONQ7ILtiv6q7Hb/x953o1yvxRLMGi9CMP96t2V1B3Co+BtGfHNtOl\nZH9hMjmz8wkB2PlJHmkt5qNkE4KWpfVy3oDE/ZZW54aDDt2qkUj5bLYVyGSRdhRX\ns1Eos5PND7MP5w3u1/0b3BRYironL3kTOWmpOBIF+x/Vr5z2HtfE6lnOzAV3uE9P\noAHvDZGdAgMBAAECggEAe+xUeYi3eTakPLX6vPJjORYOdcPaihyd7YYbGzNc9N3p\nIC6Bd1SKouJOhmUN/eOGafaHUQ8PFsYvODRWSioU+nv7u1KnF0PbjwTW7Z6/YReL\nK0zVR1C/ZwHmzarQNToNNwmz+6NAjapkofxasKeWujKQj1Yyq5k4aD9D/mTSwmcg\nXtSMd+e9VZIVD2JLu3bOq/yQBX5W/cr0+kDfqWlaDWhM/hYMXxF9Uza5PI3Br8I5\n20lYrMlAhy3s+l5mmbP6KQ5INx8GQFo9TKvr2pZ0UvC6AIy7gywVykPukpGa918O\nvzDQ8xTCFzkFC85kWa3SzNFaonxihj9S/IsPlWJzoQKBgQDmmQ8v4PZgSCoS+iw0\nJc3HGSe8y7aGA+BnephgUyuceVyl7t0mXCXpIo5mnChx8oIR9eDtZ+qjuYTa6fpm\nxq2UMT91Rpv5F2SBV6E+Ro+UhazCYoB78m0LUmJr8sHqdASfyzzhEDOzWI4gXZmt\nez30mY8GDuaLojCPwT3ZToQx5QKBgQDIVDNb1gVUzh3tyh0mnVPy+iJYv5dNOcpi\nmrz1I2VFc/Rj78CASih3AfzVlmvXDYoWUatmuS4m0pDQ8ZgxbX2J/auAp6KEiq42\nY4HM3HNrVlnxKjm1Z99Drk+XGmxMcByCNTOTWrTm2S7Wf9d1IXB2FMe5/4FWqtHl\nUFx1ZrfFWQKBgQDY8E/YVESU6e8jMVZarOH9l83JkMCjLYx07WHi5PvXVllyBdjp\n2LEVthvPrsNiB6qUaQP1dNjKtKKnLI4VGH5+NyKjGw1rdrVn9V/NbaZwLJ4NtYt4\nj+ZES3oqYhRKlGExT4tzlMF+zSXQjx4flh0AO2LRaT09ShVzeHHOOU6mfQKBgAPN\ndVwc5+UANJk66Oq+ucU28kf0+w5ANpNAlK2iil2TIeRzvJey0KnRo9b6D/n5C9H4\nouRER5f9DzqoG1d02Jkx83txsygI6d2mbeRRmu9CqFpnjsjeu+IxANnBaqTYy/G8\nyMnXQu/O15DCxHg4tkXHhzMEopPekjVkHR0tYfehAoGAbbR2kymXV9V79l6KLhKB\nzCcrdI+uIFijI92jBVtiQWg5gLZg142yrRtTZonv4mAIoVvhis4bcK5fG45lZHW0\nmRNua9JBmQP/atG1ONDsWJRpxsZJEhqckOYv8zR0eI6xj1sFC77wHhuInOd9h9gE\nVToyXvatdz19tub8DRdaMgs=\n-----END PRIVATE KEY-----", "certificate_serial_number": "365DD0DCC0ED83F59DE180D22FCBBF7CD4230A17"}`
	json.Unmarshal([]byte(certStr), &cert)

	cccc := config.Config{
		AppID: "wx2670795c3763951e",
		MchID: "1733186508",
		Cert: config.Cert{
			CertPrivateKey: cert.CertPrivateKey,
			// CertPrivateKey: `-----BEGIN PRIVATE KEY-----
			//			MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC9h7iMvdgwoQgZjspOMMNptXdR6eyPLtN0DA1AU/laYTSKvA6POYDzvv6XA8fOQblGYpM6rV8ebsGNO2rCCTllZ8S79S0JD/WoYfGqCU+A/B5qgFM+IxfH1o1yk7s4kdeKm5e8mXR22fyun60+Aoxsc2rG1BepxbVB44UD+zJyHd+qPnpfcpe+3llPgz9UBWCKRI7bWg/ai33+6DnOtsW8d96zk1zZlhD0tvYoda21Pn6D2kC+PctrbRTORm6Ms8B9Exm519T7Qv4+gTYut5LKODPWA8/yQK2O0usohL0Oa12J5YYyXmFgM/6/xb5eeRBkHKytje7ErjeNh/DKnKNZAgMBAAECggEBAKchx9xUjRBU8I+ZG01YMvpEy7OwVKru4Ai1l/niE0Ff9rVgoHYsf0pyYo9JBikcTAWFZ8+WkwqDIKsqejohaDvEBfi5e71CFZ8mv7TyAOn9adUA1Gc3CwFucc1X+QEpmqjgDC6EI++qyugyZtMH7Ey2erJa1YMglXZE7EdWfGWj7aVDwQJrpBhFMGg9nIRuclZnzqGVuXubJg8PWy7rnDkIi/Z/QM52jZZW50X1W6mLsRNy+eloxg6MgXx+KjAypplQMVWqu21gyudktHAXFSCf3P+NdYb0N1RL/JouUgwzsuGGxRQVi7WOzrP9EJgilFhW07uv6X3hd+KzFJSDDwECgYEA464B3Ucg/j3Oc1KXIN184VQYzzJlPxu2Nx8MwhzekSDU++/E4jvBd65BDgHc+77OwrOO4HNrCGD+iV/Sidqf4bIlGrQlp5Qew+tBb4nhzi+zEsYtwCnJozDACk9PssqP5Hj+pFDvhjy3sMoncPNSm2+jOihcAqEIB7/4jsOFG/MCgYEA1RrrXP8sSmP2WFINRHFhssl5jAV/5MwQ8jRjE6CFX9IwVHPPXjSwwOh16tugM2KPFYuHIqWsR3kyp73Vh4Ce97AJu2G1gO2twl5g0nVwiRr+ejmbJGNasrp62QZggMrHEKTqriosqOj3m6A7n9KvMKd0bsC2uA5WkZeBPx+m0oMCgYEAw9GvDM/WUpR58bnA/aVBeNNJmzru1X5SE8qCwJjv28ZvKFgp76IRXYvjq9ZyZ5rOXartYaIjFkvF4AUoISSFiiobu4HhOOYuJ7c4ymO+cAWacLU+OB44rECLitJ364BIjep6qHxr5fpmyoizr3O3QrSboLOBn0k8jN3RO4hx/X0CgYAA7T4KyH1L0YV3utud6ZRQL7oclsWInC6SrxGjOzZ5RTO6mkpTkY0XOauRmuTmdE5E/LdYujm2kdtbiWLNVQzb7OMN8o3UgrQXvUtUfvg/UGO86lU3Yks5rb/tA68VwEv/UYhHu504GtNA1QCNYGAsqP3DoYjp4f4UYgFI4f1auwKBgCScG5m93LGgJQlE4XI1hGtWCs7qMcdvz8a7X84NXuCxBuw4WZlikeLTQENZzny1u9GKWrczEcKFzMiF7u+e7kfnZzOO1Dph+Ut/8Y0XwVJZDfsZca+dOfDDgJIjZMKqGqJ+67VpVTyguXu7NW7atdnfEMZWEZvzWlXYE6j4DVXF
			//			-----END PRIVATE KEY-----`,
			CertificateSerialNumber: cert.CertificateSerialNumber,
			PublicKeyID:             cert.PublicKeyID,
			PublicKey:               cert.PublicKey,
			// PublicKey: `-----BEGIN PUBLIC KEY-----
			//			MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9CUWaMY8Fd5HfSVNKJtQFjwZ57sruwz4cRK2HCvHUaNsWMdy5S2zfGIwD8Z9YfJESa3AmSV42jb1tFu3tb+Wew6JMLRHolknZuEokDO0zH49fhnVP/6x2FItGy4/lIBaoPg49LHXxCZxVqshykV7m3/bd9MGYgoMHSQKKJhYzNCbL2Gp8L7ChtVctReG6eGjcOOk+a4hztWA7e1Qig8zLsGGnI+rcowMHn9ZTLwYa9taBNotVwa8qqOTCs9s22/6ZCDJHJELMrNrvjRNomjlC0N/dHAbJuftpoXWSxAxJdI0r2+Vy1VqDglJ9KxCiPf9B87eirOnZgn6YgKJ8a35iQIDAQAB
			//			-----END PUBLIC KEY-----`,
		},
		APIKey: cert.Secret,
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	// proxy.test.jianxindianzi.com
	//u := "https://proxy.test.jianxindianzi.com/public/v1/callback/payment/17"
	resp, apiR, err := c.QueryNoticeUrl(ctx)
	assert.NoError(t, err)
	if apiR != nil {
		t.Log(apiR.Response.StatusCode)
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
	t.Log(fmt.Sprintf("%+v", apiR))
	t.Log(fmt.Sprintf("%+v", resp))
	if err != nil {
		t.Log(err.Error())
	}
}

func TestUpdateNoticeUrl(t *testing.T) {
	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	// proxy.test.jianxindianzi.com
	u := "https://api-cabinet.test.jianxindianzi.com/public/v1/callback/complaint/17"
	resp, apiR, err := c.UpdateNoticeUrl(ctx, UpdateComplaintNoticeUrl{Url: u})
	assert.NoError(t, err)
	if apiR != nil {
		t.Log(apiR.Response.StatusCode)
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
	t.Log(fmt.Sprintf("%+v", apiR))
	t.Log(fmt.Sprintf("%+v", resp))
}

func TestDelete(t *testing.T) {
	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := ComplaintApiService{Client: api.Client}
	ctx := context.Background()
	// proxy.test.jianxindianzi.com
	//u := "https://proxy.test.jianxindianzi.com/public/v1/callback/payment/17"
	apiR, err := c.DeleteNoticeUrl(ctx)
	assert.NoError(t, err)
	if apiR != nil {
		t.Log(apiR.Response.StatusCode)
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
	t.Log(fmt.Sprintf("%+v", apiR))
}
