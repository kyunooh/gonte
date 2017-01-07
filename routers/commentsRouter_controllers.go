package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:NoteController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:NoteController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:nid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:NoteController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:nid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:NoteController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:nid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"] = append(beego.GlobalControllerRouter["kyunooh/gonte/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
