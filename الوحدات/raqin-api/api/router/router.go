/*

هنا نقطة عرف الوجهات اﻷساسية للـ API
ثم قم بتعريف الوجهات التي تندرج تحتها كلٌ في ملفه الخاص

*/

package router

import (
	"github.com/labstack/echo/v4"
)

// InitRouter starts the router
// الدالة التي تقوم بتسجيل وتعريف الوجهات
func InitRouter(e *echo.Echo) {

	// this group will have all client routes
	// عرف وجهة المستخدم العادي هنا ثم اضف مايندرج تحته في ملفه الخاص كالمثال الموجود
	// انظر الى ملف client_routes.go
	clientRouter := e.Group("/client")
	InitClientRouter(clientRouter)

	// this group will have all client routes
	// عرف وجهة المستخدم المتحكم هنا ثم اضف مايندرج تحته في ملفه الخاص كالمثال الموجود
	// انظر الى ملف admin_routes.go
	adminRouter := e.Group("/admin")
	InitAdminRouter(adminRouter)

}
