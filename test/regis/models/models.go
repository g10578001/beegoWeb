package 

import {
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/astaxie/beego"
	"gitgub.com/astaxie/beego/orm"
}

type User struct{
		Id int 64
		Name string `orm:"size(256):unique"`
		Eamil String `orm:"size(256):unique"`
	
}

func init() {
	orm.RegisterModel(new(User))
}