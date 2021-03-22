/*

هنا نقطة عرف الوجهات اﻷساسية للـ API
ثم قم بتعريف الوجهات التي تندرج تحتها كلٌ في ملفه الخاص

*/

package router

import (
	"raqin-api/api/router/adminrouter"
	"raqin-api/api/router/clientrouter"

	"github.com/labstack/echo/v4"
)

// InitRouter starts the router
// الدالة التي تقوم بتسجيل وتعريف الوجهات
func InitRouter(e *echo.Echo) {

	// عرف وجهة المستخدم العادي هنا ثم اضف مايندرج تحته في ملفه الخاص كالمثال الموجود
	clientrouter.Init(e.Group("/client"))

	// عرف وجهة المستخدم المتحكم هنا ثم اضف مايندرج تحته في ملفه الخاص كالمثال الموجود
	adminrouter.Init(e.Group("/admin"))

}
