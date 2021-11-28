package emsg

const (
	Success = 10000
	Error   = 20000
	//用户错误
	UsernameExist   = 10001
	UsernameNoExist = 10002
	PasswordWrong   = 10003
	NoLogin         = 10004
	//文章错误
	GetArticleFailed    = 20001
	PostArticleFailed   = 20002
	ArticleNoExist      = 20003
	DeleteArticleFailed = 20004
	LikeArticleFailed   = 20005
)

var ErrorMsg map[int]string

//初始化状态码与错误信息对应的map
func init() {
	ErrorMsg = make(map[int]string)
	ErrorMsg[Success] = "操作成功"
	ErrorMsg[Error] = "操作失败"
	ErrorMsg[UsernameExist] = "用户名已存在"
	ErrorMsg[UsernameNoExist] = "用户名不存在"
	ErrorMsg[NoLogin] = "用户未登录"
	ErrorMsg[PasswordWrong] = "密码错误"
	ErrorMsg[GetArticleFailed] = "获取文章失败"
	ErrorMsg[PostArticleFailed] = "发布文章失败"
	ErrorMsg[ArticleNoExist] = "文章不存在"
	ErrorMsg[DeleteArticleFailed] = "删除文章失败"
	ErrorMsg[LikeArticleFailed] = "点赞文章失败"
}

// @title 获取错误信息
// @description 接受状态码返回对应的错误信息
// @param code int "状态码"
// @return msg string "状态码对应的错误信息"
func GetErrorMsg(code int) string {
	return ErrorMsg[code]
}
