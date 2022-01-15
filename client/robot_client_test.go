package client

import (
	"testing"
)

func TestWebHook(t *testing.T) {
	webHook := New("example-access-token")
	payLoad := &PayLoad{}
	err := webHook.sendPayload(payLoad)
	if nil == err {
		t.Error("api request error should be catch!")
	}

	err = webHook.sendPayload(payLoad)
	if nil == err {
		t.Error("api response error should be catch!")
	}

	err = webHook.sendPayload(payLoad)
	if nil == err {
		t.Error("json unmarshal error should be catch!")
	}

	err = webHook.sendPayload(payLoad)
	if nil == err {
		t.Error(err)
	}
	err = webHook.sendPayload(payLoad)
	if nil == err {
		t.Error("response struct error should be catch!")
	}
	payLoad = &PayLoad{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: "test msg",
		},
	}

	// test send text message
	err = webHook.SendTextMsg("Test text message", false, "")
	if nil == err {
		t.Error("token missing error should be catch!")
	}

	// test send link message
	err = webHook.SendLinkMsg("A link message", "Click me to baidu search", "", "https://www.baidu.com")
	if nil == err {
		t.Error("token missing error should be catch!")
	}

	// test send markdown message
	err = webHook.SendMarkdownMsg("A markdown message", "# This is title \n > Hello World", false, "13800138000")
	if nil == err {
		t.Error("token missing error should be catch!")
	}

	// test send action card message
	err = webHook.SendActionCardMsg("A action card message", "This is a action card message", []string{}, []string{}, true, true)
	if nil == err {
		t.Error("links and titles cannot be null error should be catch!")
	}

	err = webHook.SendActionCardMsg("A action card message", "This is a action card message", []string{"Title 1"}, []string{}, true, true)
	if nil == err {
		t.Error("links and titles length not equal error should be catch!")
	}

	err = webHook.SendActionCardMsg("A action card message", "This is a action card message", []string{"Baidu Search"}, []string{"https://www.baidu.com"}, true, true)
	if err == nil {
		t.Error("token missing error should be catch!")
	}

	// test send link card message
	err = webHook.SendLinkCardMsg([]LinkMsg{{Title: "Hello Bob", MessageURL: "https://www.google.com", PicURL: ""}})
	if nil == err {
		t.Error("token missing error should be catch!")
	}

	t.Log("All test had pass ..")

}
