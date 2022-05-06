// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package utils

import (
	"easygoadmin/conf"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/gmd5"
	"easygoadmin/utils/gstr"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

// 调试模式
func AppDebug() bool {
	// 获取配置实例
	return conf.CONFIG.EGAdmin.Debug
}

// 登录用户ID
func Uid(ctx iris.Context) int {
	fmt.Println("全局获取用户ID")
	//sessValues := sessions.Get(ctx).GetAll()
	//fmt.Println(len(sessValues))
	//for k, v := range sessValues {
	//	fmt.Println(k, v)
	//}
	userId := sessions.Get(ctx).GetIntDefault(conf.USER_ID, 0)
	return userId
}

// 判断用户登录状态
func IsLogin(ctx iris.Context) bool {
	// 初始化session对象
	fmt.Println("初始化SESSION")
	userId := Uid(ctx)
	return userId > 0
}

// 判断元素是否在数组中
func InArray(value string, array []interface{}) bool {
	for _, v := range array {
		if gconv.String(v) == value {
			return true
		}
	}
	return false
}

// 获取文件地址
func GetImageUrl(path string) string {
	return conf.CONFIG.EGAdmin.Image + path
}

func Md5(password string) (string, error) {
	// 第一次MD5加密
	password, err := gmd5.Encrypt(password)
	if err != nil {
		return "", err
	}
	// 第二次MD5加密
	password2, err := gmd5.Encrypt(password)
	if err != nil {
		return "", err
	}
	return password2, nil
}

// 图片存放目录
func ImagePath() string {
	return conf.CONFIG.Attachment.FilePath + "/images"
}

func SaveImage(url string, dirname string) (string, error) {
	// 判断文件地址是否为空
	if gstr.Equal(url, "") {
		return "", errors.New("文件地址不能为空")
	}

	// 判断是否本站图片
	if gstr.Contains(url, conf.CONFIG.EGAdmin.Image) {
		// 本站图片

		// 是否临时图片
		if gstr.Contains(url, "temp") {
			// 临时图片

			// 创建目录
			dirPath := ImagePath() + "/" + dirname + "/" + time.Now().Format("20060102")
			if !CreateDir(dirPath) {
				return "", errors.New("文件目录创建失败")
			}
			// 原始图片地址
			oldPath := gstr.Replace(url, conf.CONFIG.EGAdmin.Image, conf.CONFIG.Attachment.FilePath)
			// 目标目录地址
			newPath := ImagePath() + "/" + dirname + gstr.Replace(url, conf.CONFIG.EGAdmin.Image+"/temp", "")
			// 移动文件
			os.Rename(oldPath, newPath)
			return gstr.Replace(newPath, conf.CONFIG.Attachment.FilePath, ""), nil
		} else {
			// 非临时图片
			path := gstr.Replace(url, conf.CONFIG.EGAdmin.Image, "")
			return path, nil
		}
	} else {
		// 远程图片
		// TODO...
	}
	return "", errors.New("保存文件异常")
}

// 处理富文本
func SaveImageContent(content string, title string, dirname string) string {
	str := `<img src="(?s:(.*?))"`
	//解析、编译正则
	ret := regexp.MustCompile(str)
	// 提取图片信息
	alls := ret.FindAllStringSubmatch(content, -1)
	// 遍历图片数据
	for _, v := range alls {
		// 获取图片地址
		item := v[1]
		if item == "" {
			continue
		}
		// 保存图片至正式目录
		image, _ := SaveImage(item, dirname)
		if image != "" {
			content = strings.ReplaceAll(content, item, "[IMG_URL]"+image)
		}
	}
	// 设置ALT标题
	if strings.Contains(content, "alt=\"\"") && title != "" {
		content = strings.ReplaceAll(content, "alt=\"\"", "alt=\""+title+"\"")
	}
	return content
}

// 创建文件夹并设置权限
func CreateDir(path string) bool {
	// 判断文件夹是否存在
	if IsExist(path) {
		return true
	}
	// 创建多层级目录
	err2 := os.MkdirAll(path, os.ModePerm)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// 判断文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		if os.IsExist(err) {
			// 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}

// 数组反转
func Reverse(arr *[]string) {
	length := len(*arr)
	var temp string
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}

func InStringArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
