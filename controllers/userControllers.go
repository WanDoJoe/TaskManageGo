package controllers

import (
	"github.com/astaxie/beego"
	"taskmanage/models"
	"encoding/json"
	"fmt"
)

type UserController struct {
	beego.Controller
}
//http://127.0.0.1:8080/login?jsonParams={"sysinfo":{"drevice":"phone"},"params":{"username":"admin","password":"admin"}}
func (this *UserController) Login(){
          parammsJson:=this.GetString("jsonParams")
	  usernamem :=GetParams(parammsJson,"username")
	 password :=GetParams(parammsJson,"password")

	 user,err :=models.FindByNameAndPwd(usernamem,password)
	user.Logindateid=GetDateAndTime()
	if err!=nil{
		panic(err)
	}  

	use,error :=json.Marshal(user);
	if error!=nil{
		panic(error)
	}

	resultJson :=string(use)
	fmt.Print(resultJson)
	////发送json
	this.Data["json"] = map[string]interface{}{"success":true,"message":"成功","data":resultJson}
	this.ServeJSON(false)
}

//http://127.0.0.1:8080/adduser?jsonParams={"sysinfo":{"drevice":"phone"},"params":{"username":"admin","password":"admin","power":"0","createtime":"2017-07-17pm14:51"}}
var addreturnmessagestring=""
var addreturnsuccessstring=""
func(this *UserController) AddUser(){

	parammsJson:=this.GetString("jsonParams")

	user :=new(models.User)
	user.Username=GetParams(parammsJson,"username")
	user.Password=GetParams(parammsJson,"password")
	user.Power=GetParams(parammsJson,"power")
	user.Createtime=GetParams(parammsJson,"createtime")

	returnid,_ := models.UserAdds(user)
	if returnid > 0 {
		 fmt.Sprintf("%d", returnid)
		addreturnmessagestring="成功"
		addreturnsuccessstring="true"
	}else{
		fmt.Sprintf("%d", "插入未成功")

		addreturnmessagestring=string("插入未成功")
		addreturnsuccessstring="false"
	}

	this.Data["json"] = map[string]interface{}{"success":addreturnsuccessstring,"message":addreturnmessagestring,"data":returnid}
	this.ServeJSON(false)
}