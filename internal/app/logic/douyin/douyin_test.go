package douyin

import (
	"spider/internal/config"
	"testing"
)

func TestDouyinClientUserInfo(t *testing.T) {
	client := NewDouyinClient("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	resp, err := client.DouyinUserInfo(nil, "MS4wLjABAAAAz-Nssy-G6nNshJODTK3VpEpjWsH1pMHODDPexGS5K-D6EAo5iASK_qCGRb7M5Rbe")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp, status)
}
func TestDouyinUserVideo(t *testing.T) {
	initd()
	client := NewDouyinClient("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	resp, status, err := client.DouyinUserVideo(nil, "MS4wLjABAAAAz-Nssy-G6nNshJODTK3VpEpjWsH1pMHODDPexGS5K-D6EAo5iASK_qCGRb7M5Rbe")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp, status)
}

func TestDouyinClientUserSearch(t *testing.T) {
	initd()
	client := NewDouyinClient("")
	search, err := client.DouyinUserSearch(nil, "dyu741ej0t5u")
	if err != nil {
		t.Error(err)
	}
	t.Log(search)
}

func TestDouyinAwemeDetail(t *testing.T) {
	initd()
	client := NewDouyinClient("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	search, err := client.DouyinAwemeDetail(nil, "7372484719365098803")
	if err != nil {
		t.Error(err)
	}
	t.Log(search)
}
func TestDouyinCommentList(t *testing.T) {
	initd()
	var (
		cursor = "0"
		limit  = "10"
	)
	client := NewDouyinClient("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	search, err := client.DouyinComment(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(search)

	//item_id
	//7561795410549853481
	//comment_id
	//7565069968577479451
}
func TestDouyinCommentReplyList(t *testing.T) {
	initd()
	var (
		cursor = "0"
		limit  = "10"
	)
	client := NewDouyinClient("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	search, err := client.DouyinCommentReply("7561795410549853481", "7565069968577479451", cursor, limit)
	if err != nil {
		t.Error(err)
	}
	t.Log(search)

	//item_id
	//7561795410549853481
	//comment_id
	//7565069968577479451
}

func initd() {
	config.InitConfig("E:\\workspace\\src\\spider\\config.yaml")
}
