package controllers

import (
	m "easyui_crud/models"
	"fmt"
)

type UserController struct {
	BaseController
}

func (this *UserController) Index() {
	this.TplName = "main.html"
}

func (this *UserController) UserIndex() {
	this.TplName = "user/usermanage.html"
}

func (this *UserController) UserList() {
	page, _ := this.GetInt64("page")
	pageSize, _ := this.GetInt64("rows")
	nodes, cnt := m.GetUserList(page, pageSize)
	this.Data["json"] = &map[string]interface{}{"total": cnt, "rows": &nodes}
	this.ServeJSON()
}

func (this *UserController) UserDel() {
	idstr := this.GetString("ids")
	err := m.DelUser(idstr)
	if err != nil {
		this.Rsp(false, err.Error())
	} else {
		this.Rsp(true, "删除成功")
	}
}

func (this *UserController) UserAdd() {
	u := m.User{}
	u.Id, _ = this.GetInt("Id")
	if err := this.ParseForm(&u); err != nil {
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "新增成功")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}

func (this *UserController) UserUpdate() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.UpdateUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "修改成功")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *UserController) UserUpload() {

	f, fh, err := this.GetFile("uploadFile")
	defer f.Close()
	if err != nil {
		fmt.Println("get file error ", err)
		this.Data["json"] = &map[string]interface{}{"path": "", "succ": false}
		this.ServeJSON()
	} else {
		fmt.Println(fh.Filename)
		this.SaveToFile("uploadFile", "static/upload/"+fh.Filename)
		this.Data["json"] = &map[string]interface{}{"path": "/static/upload/" + fh.Filename, "succ": true}
		this.ServeJSON()
	}

}