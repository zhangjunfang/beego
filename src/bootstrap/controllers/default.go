package controllers

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/beego/i18n"
	"github.com/nfnt/resize"
	"github.com/shirou/gopsutil"
	"html/template"
	"image/jpeg"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type BasicController struct {
	beego.Controller
}
type MainController struct {
	BasicController
	i18n.Locale
}

func (this *MainController) Prepare() {
	lang := this.GetString("lang")
	beego.Info("url 国际化信息===", lang)
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		if len(lang) > 0 {
			this.Lang = lang
			beego.Info("cookie 国际化信息===", this.Lang)
		} else {
			this.Lang = "en-US"
			beego.Info("默认 国际化信息===", this.Lang)
		}

	} else {
		this.Lang = lang
	}
	this.Ctx.SetCookie("lang", this.Lang, 1<<31-1, "/")
}

func (this *MainController) Get() {
	this.Data["hi"] = this.Tr("hi")
	this.Data["bye"] = this.Tr("bye")
	this.Data["Lang"] = this.Lang
	this.TplNames = "index.html"
}
func (this *MainController) FileList() {
	volumeName := this.GetString("volumeName")
	file := make(map[string]interface{}, 20)
	if len(strings.Trim(volumeName, " ")) > 0 { //当访问次数>=2时，执行这个函数
		//获取一个目录下文件个数
		files, err := ioutil.ReadDir(volumeName)
		fmt.Printf("当前目录下有%d个文件\n\r", len(files))
		if err != nil {
			beego.Error("读取文件目录或者文件出错，信息：", err)
			return
		}
		for _, fileStat := range files {
			file[fileStat.Name()] = fileStat
		}
		//time.Sleep(1 * time.Minute)
	} else { //当第一访问功能的时候，系统默认加载系统根目录
		//  go语言中  实现的获取所有的盘符
		disks, _ := gopsutil.DiskPartitions(true)
		for _, disk := range disks {
			file[disk.Device] = disk.Device
		}
	}
	this.Data["file"] = file
	this.ServeJson()

}
func (this *MainController) FileAllList() {
	//  go语言中  实现的获取所有的盘符
	disks, _ := gopsutil.DiskPartitions(true)
	for k, disk := range disks {
		fmt.Println(k, disk.Device)
		filepath.Walk(disk.Device, func(path string, f os.FileInfo, err error) error {
			files, _ := ioutil.ReadDir(path)
			fmt.Printf("当前目录下有%d个文件\n\r", len(files))
			if f == nil {
				fmt.Println("文件目录为空")
				return nil
			}
			if f.IsDir() {
				fmt.Println("文件目录:", path)
				time.Sleep(1 * time.Second)
				return nil
			}
			fmt.Println("文件名称为：", f.Name())
			return nil
		})
	}
	//获取系统所有的盘符【window】
	volume := filepath.VolumeName("d:/aa/bb")
	fmt.Println(volume) //  D:
	//fmt.Println(gopsutil.DiskPartitions(true))
}
func (this *MainController) ExecuteCommand() {
	cmd := this.GetString("cmd")
	var ec *exec.Cmd
	if len(cmd) > 0 {
		//ec.Start()
		ec = exec.Command(cmd)
		//_ := ec.Run()
		b, err := ec.Output()
		if err != nil {
			this.Data["msg"] = err
		} else {
			this.Data["msg"] = string(b)
		}
	}
	this.ServeJson()
}
func (this *MainController) ExecuteCommandWithArgs() {
	cmd := this.GetString("cmd")
	array := strings.Split(cmd, ";")
	if len(array) == 2 {
		var ec *exec.Cmd
		if len(array[0]) > 0 {
			//ec.Start()
			ec = exec.Command(cmd, array[1])
			//_ := ec.Run()
			b, err := ec.Output()
			if err != nil {
				this.Data["msg"] = err
			} else {
				this.Data["msg"] = string(b)
			}
		}
		this.ServeJson()
	} else {
		this.Data["msg"] = "nil"
	}
	return
}
func (this *MainController) FileUpload() {
	fmt.Println("性别：", this.GetString("sex"))
	m := strings.ToLower(this.Ctx.Input.Method())
	if m == "get" {
		this.Data["token"] = template.HTML(TokenFormHtml())
		this.Data["xsrfdata"] = template.HTML(this.XsrfFormHtml())
		this.TplNames = "upload.tpl"
	} else {
		pathstring := path.Join(GetCurrentPath(), "/static/img/")
		//pathstring := strings.Replace(GetCurrentPath(), "\\", "/", -1) + "/static/img/"
		//sign := IsDirExists(pathstring)
		//if !sign {
		//	os.MkdirAll("/aa", 0777)
		//}
		_, fileHeader, _ := this.Ctx.Request.FormFile("imgs")
		fn := fileHeader.Filename
		duplicateName := filepath.Walk(pathstring, func(pathstring string, f os.FileInfo, err error) error {
			if f == nil {
				return errors.New("上传文件目录为空")
			}
			if f.IsDir() {
				return nil
			}
			if f.Name() == fn {
				fmt.Println("Duplicate file name", f.Name())
				fmt.Println("err  content", err)
				return errors.New("Duplicate file name")
			}
			return nil
		})
		if duplicateName != nil {
			this.Data["sign"] = "文件已经存在 ！！！"
			this.TplNames = "sucess.tpl"
			return
		}
		//var temp = pathstring + "/" + fn
		var buf bytes.Buffer
		buf.WriteString(pathstring)
		buf.WriteString("/")
		buf.WriteString(fn)
		os.NewFile(0777, buf.String())
		err2 := this.SaveToFile("imgs", buf.String())
		if err2 != nil {
			fmt.Println("操作失败：", err2)
			this.Data["sign"] = "操作失败"
		} else {
			this.Data["sign"] = "操作成功"
		}
		this.TplNames = "sucess.tpl"
	}
}
func (this MainController) FileDown() {
	//fn实际项目中这个名称可能来自于客户端或者服务器端的映射真正的文件名称
	fn := "bb.jpg"
	//pathString := strings.Replace(GetCurrentPath(), "\\", "/", -1) +
	//pathString:=path.Join(GetCurrentPath(),"/static/img/",fn)
	temp := path.Join(GetCurrentPath(), "/static/img/", fn)
	ext := GetExt(temp)
	//fmt.Println("ext::::::::", ext)
	dn := "aa" + ext
	this.Ctx.Output.Download(temp, dn)
	return
}

//文件重命名
func (this *MainController) RenameFile() {
	path := this.GetString("path")
	info, _ := os.Stat(path)
	if !info.IsDir() {
		pathDir := filepath.Dir(path)
		newName := this.GetString("newName")
		os.Rename(path, filepath.Join(pathDir, newName))
	}
	return
}

//文件移动
func (this *MainController) MoveFile() {
	path := this.GetString("path")
	old, _ := os.Open(path)
	newPath := this.GetString("newPath")
	file, _ := os.Create(newPath)
	io.Copy(file, old)
	//os.OpenFile()
	var filename string = "d:/git.txt"
	var fileInfo []os.FileInfo
HERF:
	fileInfo, _ = ioutil.ReadDir(filename)
	beego.Info(len(fileInfo))
	if len(fileInfo) > 0 {
		for k, v := range fileInfo {
			beego.Info(k, v.IsDir(), v.Name())
			if v.IsDir() {
				filename = filename + "/" + v.Name()
				//append(fileInfo, os.FileInfo{})
			}
		}
		goto HERF
	} else {
		file, _ := os.Open(filename)
		beego.Info(file.Name())

	}
	return
}

//文件删除
func (this *MainController) DeleteFile() {
	//os.Remove("fsfsf")
	os.RemoveAll(this.GetString("path"))
	return
}

//文件压缩
func (this *MainController) Zip() {
	return
}

//文件解压
func (this *MainController) UnZip() {
	return
}

//远程登录服务器
func (this *MainController) Telnet() {

}

//文件编辑并保存
func (this *MainController) EditAndSave() {
	path := this.GetString("path")
	old, _ := os.Open(path)
	defer old.Close()
	if this.Ctx.Request.Method == "get" {
		buf := make([]byte, 1024)
		for {
			n, _ := old.Read(buf)
			if 0 == n {
				break
			}
			this.Ctx.ResponseWriter.Write(buf[:n])
		}
	} else if this.Ctx.Request.Method == "post" {
		old.Write([]byte(this.GetString("content")))
	}
	return
}

// 发送短信
func (this *MainController) SendMC() {
	req := httplib.Post("http://106.ihuyi.cn/webservice/sms.php?method=Submit").Debug(true)
	req.Param("account", "cf_ocean")     //
	req.Param("password", "19860505zjf") //
	req.Param("mobile", "13253663178")
	req.Param("mobile", "18530053050")                                                                                             //
	req.Param("content", "您的验证码是："+strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(99999999))+"。请不要把验证码泄露给其他人。") //
	msg, _ := req.String()
	res, _ := req.Response()
	beego.Info("=====req======", msg)
	//beego.Info("=====res======", res.Write(os.Stdout))
	res.Write(os.Stdout)
}

//支付宝对接
func (this *MainController) Joint() {
	return
}

//远程备份和存储
func (this *MainController) BackupAndStorage() {
	return
}

//调度器
func (this *MainController) Scheduler() {
	time.NewTicker(1 * time.Second).C
	return
}

//获取当前所在的文件路径
func GetCurrentPath() string {
	path, err := os.Getwd()
	if err == nil {
		return path
	}
	return ""
}
func TokenFormHtml() string {
	return "<input type=\"hidden\" name=\"token\" value=\"" +
		GetToken() + "\"/>"
}

//获取文件的扩展名称
func GetExt(filename string) string {
	runtime.Caller(0) //和下面代码实现同样的功能，但是符合操作系统的文件路径符
	return path.Ext(filename)
}

//获取token的值
func GetToken() string {
	cruTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(cruTime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

//图像缩放功能
func ImageResize() {
	// open "test.jpg"
	file, err := os.Open("test.jpg")
	if err != nil {
		print("Open File Error")
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		print("Not image file")
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(300, 0, img, resize.Lanczos3)

	out, err := os.Create("test_go.jpg")
	if err != nil {
		print("Save Image Error!")
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
