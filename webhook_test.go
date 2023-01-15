package dingtalk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDingTalk_SendTextMessage(t *testing.T) {
	accessToken := os.Getenv("access_token")
	secret := os.Getenv("secret")
	err := InitDingTalkWithSecret(accessToken, secret).SendTextMessage("test dingtalk send text message")
	assert.Nil(t, err)
}

func TestDingTalk_SendTextMessageFailed(t *testing.T) {
	accessToken := os.Getenv("access_token")
	err := InitDingTalkWithSecret(accessToken, "error secret").SendTextMessage("test dingtalk send text message")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "请确认机器人的密钥加密和填写正确")
}
