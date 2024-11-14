package notify

import "testing"

func TestNewDing2(t *testing.T) {
	ding := NewDing(
		DingHook(`https://oapi.dingtalk.com/robot/send`),
		DingToken(``),
		DingSecret(``),
		DingBody(DingRequest{
			MsgType: "markdown",
			Markdown: MarkDownModel{
				Title: " --- test --- ",
				Text:  "test go-core",
			},
		}),
	)

	notify, err := ding.Notify()
	if err != nil {
		t.Fatalf("err: %+v", err)
	}

	t.Logf("body: %s", string(notify.Body()))
}
