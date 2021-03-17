/*

هنا نقطة بدأ تشغيل الـ API


*/

package api

import (
	"raqin-api/api/middleware"
	"raqin-api/api/router"

	"github.com/labstack/echo/v4"
)

// Init starts the REST server
// الدالة التي تقوم بتشغيل الموجه
func Init() {
	// Echo instance
	// نقوم بإنشاء معرف من نوع ايكو لنستخدمه في بقية البرنامج
	e := echo.New()

	// Middleware
	// قم بتشغيل ماتحتاجه كالمسجل مثلا بداخل الدالة التالية في ملفها الخاص
	middleware.InitMiddleware(e)

	// router
	// هنا البرنامج يبدأ بتشغيل الموجه وتسجيل الوجهات المعرفة بداخله
	router.InitRouter(e)

	// Start echo server
	// هنا يتم ربط الموجه على المنفذ المحدد ويتم التشغيل
	e.Logger.Fatal(e.Start(":1323"))
}
