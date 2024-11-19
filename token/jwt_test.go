package token

import (
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2/jwt"
	"testing"
)

func TestNewEngine(t *testing.T) {
	openId := uuid.New()
	t.Logf("openId: %+v", openId.String())

	keys := NewKey()
	user := NewUser(
		WithID(333),
		WithOpenId(openId),
		WithEmail(`123@gmail.com`),
		WithClaims(jwt.Claims{
			Issuer:  "test-1",
			Subject: "test-1",
		}),
	)
	token := NewEngine(
		WithKey(keys),
		WithUser(
			user,
		),
	)

	te, re, err := token.Generate()
	if err != nil {
		panic(err)
	}

	t.Logf("token: %s", te)

	verifier, err := token.VerifierToken(te)
	if err != nil {
		panic(err)
	}
	t.Logf("verifier: %+v %+v %+v %+v %+v", verifier.ID(), verifier.OpenID(), verifier.Email(), verifier.IsRefresh(), verifier.Claims().Expiry.Time())

	verifierR, err := token.VerifierRefresh(te)
	if err != nil {
		t.Logf("te VerifierRefresh err: %+v", err)
	}

	t.Logf("refresh: %s", re)
	verifierR, err = token.VerifierRefresh(re)
	if err != nil {
		panic(err)
	}
	t.Logf("verifierR: %+v %+v %+v %+v %+v", verifierR.ID(), verifierR.OpenID(), verifierR.Email(), verifierR.IsRefresh(), verifierR.Claims().Expiry.Time())

	ttoken := `eyJhbGciOiJSU0EtT0FFUC0yNTYiLCJlbmMiOiJBMjU2R0NNIn0.V2tmalWy6UJfwhT53SDOkrWZv2CMOUXCEAm-n071OMrbWwf-b9HWtFMoQw_Agcoz0MyCMt865krNcl8qq2hG8HXRvqrwvH3yFOQ4yqu-2k4A2kt_Oc4DXhb_BtRLTzVGadIVzIJlLD73DYyNJL1IbhOCCT9jjB4y0bR8l8itE-MAt9R2DF9Emuf6BAbYYthbPcu-qofYfZ4oUwrGwaNoz2Rz8yJeE306-YGrafPiGsWYX36jye2jGPxnrYOZjyEkin2WS8czxJR9wZC9y_Wb5UI6KXqFzwtlsvWPSey7DCpz29mTjNodYXh6PLI_k9Sqt1nAFjiYGlh3Ua1fND_cDQ.bUGMlrqrAWSM-i5y.ObLxA2rkjER5yh6j2oSfuVtJOtdm2EWZkZ4wlUjqyXVDh1Uf0VcSAB8_KrYmmyjvk2KucnPVUVZ0xypKexAvgxdkwGjb-0mjUVMkVhgWbuLryDfIt2umOTJHRVzZOP-GZwxQ12mQAFRoJPhGeoWC8-j3TdGvqfugmh5lSRD9moZAxTnblEuBdxOkxPm2K4_4vbJWjMS9LN9eA6Ak0sKTHclnpXTbIN00UOjHkHsnyJJ-K359AFavAIFu2BuIKQr5h5ysU2GiV9knu1tW5kaeDJ356I1XAunLpvLEL8CVch6x_27AvqsVA-AAJq-lpcLl52Kg1ZJRRK93gR0jqp93ZrfpwQIBzJwx-vu8KQinmu5_C1v7yc50voJcO0W-jawxlu6-5IDsn1ev4FyeDW-7cSuTEg-5vwmR4tsiC0gr4OtpUQG5mJOFjAIgzpyBfruDu32LMMfj4V00DwW-5_Jn3Yx7oA3rlnbq41za4gxBB90rRs-7pXMZ9CJoO4ailQRGJGSjfpjyJ70hJlsc5geFsnrQTu0XcO1gxnNH17t8lbHBHdT3nQl3Uv2J6JptH9GulK0qq71z5MQN5OoPjnzNN8a0EKktFOsnS7T7vK5TOdUc8bTX3rCprtV0g9P-0K7DrhdUuQUvO-NREnrvLCwUH4qyUTjcJqs9GCEZO2BblUyQAWqtao-_ZEph9mAKuqEQI2REYkIj8isgnuWiRKIf2ta9DFGR4xTylvtfQ31vmMYrdWdhO3XoLCHP9Mg1smHD__LWc5RaMUqzuc6FdOX7ki4DHFzGrlXLzGxLL3XBlbHGmyAYx7-JMxzPALVsVYnAHmuyuMsyFmv4L_YF29TA8OTl-wHv2cJ9_pdLdhjQZE4HbbVZizZ5dhuUJDrsvSPUdbnzs5dAyTBuLGUM6WLq7tMkP_AkVrU.0sctjiq3om7gZx9P5OSV0g`
	verifier, err = token.VerifierToken(ttoken)
	if err != nil {
		panic(err)
	}
}
