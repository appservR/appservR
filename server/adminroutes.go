package server

import (
	"net/http"

	"github.com/appservR/appservR/controllers"
	"github.com/gin-gonic/gin"
)

func addAdminRoutes(admin *gin.RouterGroup,
	appsCtl *controllers.AppController, statusCtrl *controllers.StatusController,
	usersCtl *controllers.UserController, groupsCtl *controllers.GroupController) *gin.RouterGroup {

	admin.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin/apps")
	})

	admin.GET("/apps", appsCtl.GetApps())
	admin.GET("/apps/:appname", appsCtl.GetApp())
	admin.POST("/apps/:appname", appsCtl.UpdateApp())
	admin.GET("/apps/:appname/delete", appsCtl.DeleteApp())

	admin.GET("/status/apps", statusCtrl.Apps())
	admin.GET("/status/apps/:appid", statusCtrl.App())

	admin.GET("/users", usersCtl.GetUsers())
	admin.GET("/users/:username", usersCtl.GetUser())
	admin.POST("/users/:username", usersCtl.AdminUpdateUser())
	admin.GET("/users/:username/delete", usersCtl.DeleteUser())

	admin.GET("/groups", groupsCtl.GetGroups())
	admin.GET("/groups/:groupname", groupsCtl.GetGroup())
	admin.POST("/groups/:groupname", groupsCtl.UpdateGroup())
	admin.GET("/groups/:groupname/delete", groupsCtl.DeleteGroup())
	admin.GET("/groups/:groupname/add/:username", groupsCtl.AddGroupMember())
	admin.GET("/groups/:groupname/remove/:username", groupsCtl.RemoveGroupMember())

	return admin
}
