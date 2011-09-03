package selenium

import (
	"testing"
)

var caps = &Capabilities {
	"browserName": "firefox",
}

func TestStatus(t *testing.T) {
	wd, err := NewRemote(nil, "", nil)
	if err != nil {
		t.Error(err)
	}

	status, err := wd.Status()
	if err != nil {
		t.Error(err)
	}

	if len(status.OS.Name) == 0 {
		t.Error("No OS")
	}
}

func TestNewSession(t *testing.T) {
	wd := &remoteWD{Capabilities: caps, Executor: DEFAULT_EXECUTOR}
	sid, err := wd.NewSession()
	defer wd.Quit()

	if err != nil {
		t.Errorf("error in new session - %s", err)
	}

	if len(sid) == 0 {
		t.Error("Empty session id")
	}

	if wd.SessionId != sid {
		t.Error("Session id mismatch")
	}
}

func TestCurrentWindowHandle(t *testing.T) {
	wd, _ := NewRemote(caps, "", nil)
	defer wd.Quit()

	handle, err := wd.CurrentWindowHandle()

	if err != nil {
		t.Error(err)
	}

	if len(handle) == 0 {
		t.Error("Empty handle")
	}
}

func TestWindowHandles(t *testing.T) {
	wd, _ := NewRemote(caps, "", nil)
	defer wd.Quit()

	handles, err := wd.CurrentWindowHandle()
	if err != nil {
		t.Error(err)
	}

	if len(handles) == 0 {
		t.Error("No handles")
	}
}

func TestGet(t *testing.T) {
	wd, _ := NewRemote(caps, "", nil)
	defer wd.Quit()

	url := "http://www.google.com/"
	err := wd.Get(url)
	if err != nil {
		t.Error(err)
	}

	newUrl, err := wd.CurrentURL()
	if err != nil {
		t.Error(err)
	}

	if newUrl != url {
		t.Error("%s != %s", newUrl, url)
	}
}

func TestNavigation(t *testing.T) {
	wd, _ := NewRemote(caps, "", nil)
	defer wd.Quit()

	url1 := "http://www.google.com/"
	err := wd.Get(url1)
	if err != nil {
		t.Error(err)
	}

	url2 := "http://golang.org/"
	err = wd.Get(url2)
	if err != nil {
		t.Error(err)
	}

	err = wd.Back()
	if err != nil {
		t.Error(err)
	}
	url, _ := wd.CurrentURL()
	if url != url1 {
		t.Error("back go me to %s (expected %s)", url, url1)
	}
	err = wd.Forward()
	if err != nil {
		t.Error(err)
	}
	url, _ = wd.CurrentURL()
	if url != url2 {
		t.Error("back go me to %s (expected %s)", url, url2)
	}

	err = wd.Refresh()
	if err != nil {
		t.Error(err)
	}
	url, _ = wd.CurrentURL()
	if url != url2 {
		t.Error("back go me to %s (expected %s)", url, url2)
	}
}
