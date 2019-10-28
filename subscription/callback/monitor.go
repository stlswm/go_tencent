package callback

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/http"
	"sort"
	"strings"
)

// 接口监听器
type Monitor struct {
	request                      *http.Request
	responseWriter               http.ResponseWriter
	AuthToken                    string       // 授权
	AfterAuthHandler             func([]byte) // 授权成功回调
	TextMsgHandler               func()       // 文本消息回调
	ImageMsgHandler              func()       // 图片消息回调
	VoiceMsgHandler              func()       // 语音消息回调
	VideoMsgHandler              func()       // 视频消息回调
	ShortVideoMsgHandler         func()       // 小视频消息回调
	LocationMsgHandler           func()       // 地理位置消息回调
	LinkMsgHandler               func()       // 链接消息回调
	SubscribeEventHandler        func()       // 关注事件回调
	UnSubscribeEventHandler      func()       // 取消关注事件回调
	ScanEventHandler             func()       // 扫码事件回调
	LocationEventHandler         func()       // 地理位置上报事件回调
	MenuClickEventHandler        func()       // 自定义菜单点击回调
	MenuViewEventHandler         func()       // 点击菜单跳转链接时的事件推送回调
	MenuScanCodePushEventHandler func()       // 扫码推事件的事件推送回调
	MenuScanCodeWaitMsgHandler   func()       // 扫码推事件且弹出“消息接收中”提示框的事件推送回调
	MenuPicSysPhotoHandler       func()       // 弹出系统拍照发图的事件推送回调
	MenuPicPhotoOrAlbumHandler   func()       // 弹出拍照或者相册发图的事件推送回调
	MenuPicWeiXinHandler         func()       // 弹出微信相册发图器的事件推送回调
	MenuLocationSelectHandler    func()       // 弹出地理位置选择器的事件推送回调
	TemplateSendJobFinishHandler func()       // 模板消息发送状态报告回调
}

// 验证授权
func (c *Monitor) checkAuthToken() (err error, isExit bool) {
	mustFiles := []string{
		"signature",
		"timestamp",
		"nonce",
		"signature",
	}
	for _, param := range mustFiles {
		val, ok := c.request.URL.Query()[param]
		if !ok || len(val) == 0 {
			return errors.New("empty param:" + param), false
		}
	}
	// 验证加密内容
	signature, _ := c.request.URL.Query()["signature"]
	timestamp, _ := c.request.URL.Query()["timestamp"]
	nonce, _ := c.request.URL.Query()["nonce"]
	authParam := []string{
		c.AuthToken,
		timestamp[0],
		nonce[0],
	}
	sort.Strings(authParam)
	hasSha1 := sha1.New()
	hasSha1.Write([]byte([]byte(strings.Join(authParam, ""))))
	authStr := hex.EncodeToString(hasSha1.Sum([]byte(nil)))
	if authStr != signature[0] {
		return errors.New("auth fail"), false
	}
	// 输出echo str
	echoStr, ok := c.request.URL.Query()["echostr"]
	if ok && len(echoStr) > 0 {
		_, _ = c.responseWriter.Write([]byte(echoStr[0]))
		return nil, true
	}
	return nil, false
}

// 解析
func (c *Monitor) ParseBody() {}

// 启动
func (c *Monitor) Work() error {
	if c.AuthToken == "" {
		return errors.New("请先设置AuthToken")
	}
	err, exit := c.checkAuthToken()
	if err != nil {
		return err
	}
	if exit {
		return nil
	}
	return nil
}

// 实例化
func NewMonitor(w http.ResponseWriter, r *http.Request) *Monitor {
	return &Monitor{
		request:        r,
		responseWriter: w,
	}
}
