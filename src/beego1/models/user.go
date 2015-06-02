package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func init() {
	orm.Debug = true
	orm.RegisterModel(new(Person))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/orm?charset=utf8", 30, 30)

}

//模型命名遵循驼峰命名，如果存在驼峰命名【结构体或者字段】，自动生成下划线连键各个单词
type Person struct {
	Id             int32 `PK` //如果主键默认不是id，必须标注为PK
	Name           string
	Sex            string
	Ids            string
	Politicsstatus string
	Nation         string
	Address        string
	Graduation     string
	Date           time.Time
	Education      string
	Specialty      string
	Unit           string
	Department     string
	Workoccupation string
	Job            string
	Professional   string
	Telephone      string `valid:"Mobile"`
	Contactaddress string
}

type User struct {
	Id   int32  `form:"_" PK` //如果主键默认不是id，必须标注为PK
	Name string `form:"xm,text,年龄：" json:"name"`
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (this *Person) Valid(v *validation.Validation) {
	//if strings.Index(u.Name, "admin") != -1 {
	//	// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
	//	v.SetError("Name", "名称里不能含有 admin")
	//}
	fmt.Println("v.HasErrors():", v.HasErrors())
	if v.HasErrors() {
		v.SetError("Telephone", "手机号码不正确")
	}

}

func (this *Person) Add() bool {
	p := orm.NewOrm()
	p.Using("default") // 默认使用 default，你可以指定为其他数据库
	p.Begin()
	id, err := p.Insert(this)
	fmt.Println(id)
	if err == nil {
		p.Commit()
		return true
	} else {
		p.Rollback()
		return false
	}
	return true
}
func (this *Person) Query() bool {
	p := orm.NewOrm()
	p.Using("default") // 默认使用 default，你可以指定为其他数据库
	err := p.Read(this, "Name", "Ids")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return false
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return false
	} else {
		return true
	}
}

func (this *Person) UpdatePerson() bool {
	p := orm.NewOrm()
	p.Using("default")
	p.Begin()
	_, err := p.Update(this, "Name")
	if err != nil {
		fmt.Println("更新操作有误")
		p.Rollback()
		return false
	} else {
		p.Commit()
		return true
	}
	return true
}
func (this *Person) UpdateBatchPerson() bool {
	p := orm.NewOrm()
	p.Using("default")
	p.Begin()
	_, err := p.Raw(" update person set name=?", this.Name).Exec()
	fmt.Println("err:", err)
	if err != nil {
		fmt.Println("更新操作有误")
		p.Rollback()
		return false
	} else {
		p.Commit()
		return true
	}
	return true
}
func (this *Person) DeleteBatchPerson() bool {
	p := orm.NewOrm()
	p.Using("default")
	p.Begin()
	_, err := p.Raw(" delete from  person where  name=?", this.Name).Exec()
	fmt.Println("err:", err)
	if err != nil {
		fmt.Println("更新操作有误")
		p.Rollback()
		return false
	} else {
		p.Commit()
		return true
	}
	return true
}
func (this *Person) DeletePerson() bool {
	p := orm.NewOrm()
	p.Using("default")
	p.Begin()
	_, err := p.Raw(" delete from  person where  id=?", this.Id).Exec()
	fmt.Println("err:", err)
	if err != nil {
		fmt.Println("更新操作有误")
		p.Rollback()
		return false
	} else {
		p.Commit()
		return true
	}
	return true
}

//查询所有的数据
func (this *Person) QueryALL() []*Person {
	p := orm.NewOrm()
	p.Using("default")
	var pers []*Person
	//p.Raw("SELECT * FROM person").QueryRows(pers)
	p.QueryTable(this).All(&pers)

	return pers
}
func (user *Person) String() string {
	return "id:" + strconv.FormatInt(int64(user.Id), 10) + " name:" + user.Name + " ids:" + user.Ids

}
