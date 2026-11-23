package sessions

import "testing"

func TestPublishSkipsNilChannel(t *testing.T) {
	clearParsing()
	publish([]fileEntry{{path: "/nonexistent/path"}})
	t.Log("TestPublishSkipsNilChannel passed: no panic on nil channel")
}

func TestPublishSkipsDuplicatePath(t *testing.T) {
	clearParsing()
	ch := make(chan struct{})
	parsing["/dup/path"] = ch
	publish([]fileEntry{{path: "/dup/path"}, {path: "/dup/path"}})
	t.Log("TestPublishSkipsDuplicatePath passed: no double-close panic")
}

func TestPublishWithValidChannel(t *testing.T) {
	clearParsing()
	ch := make(chan struct{})
	parsing["/valid/path"] = ch
	publish([]fileEntry{{path: "/valid/path"}})
	_, ok := parsing["/valid/path"]
	if ok {
		t.Error("Channel should have been removed from parsing map after close")
	}
	t.Log("TestPublishWithValidChannel passed: channel closed and removed")
}

func clearParsing() {
	parsing = make(map[string]chan struct{})
}
