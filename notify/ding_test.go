package notify

import "testing"

func TestNewDing2(t *testing.T) {
	ding := NewDing(
		DingHook(`https://oapi.dingtalk.com/robot/send`),
		DingToken(``),
		DingSecret(``),
	)

	notify, err := ding.Notify(DingBody{
		MsgType: "markdown",
		Markdown: MarkDownModel{
			Title: " --- test --- ",
			Text:  "test go-core",
		},
	})
	if err != nil {
		t.Fatalf("err: %+v", err)
	}

	t.Logf("body: %s", string(notify.Body()))
}
