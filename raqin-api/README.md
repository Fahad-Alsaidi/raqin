# بنية البرنامج التحتية
هذا القسم يعطي نظرة أولية حول كيفية تقسيم وهيكلة بنية البرنامج التي سنستخدمها للبناء
> ساعدنا في تحسين الهيكلة 


## الهيكلة الحالية
```bash
├── app
│   ├── cmd `نقطة بداية البرنامج`
│   │   ├── main.go
│   │   
│   ├── api `كل ما يتعلق بالموجه والوجهات وتوابعه`
│   │   ├── middleware ``
│   │   │   ├── middleware.go
│   │   │
│   │   ├── router
│   │   │   ├── router.go
│   │   │   ├── client_routes.go
│   │   │   ├── admin_routes.go
│   │   │
│   │   ├── api.go
│   │   
│   ├── core `هنا تضع كل ما هو اساسي للبرنامج ويرتبط بهدفه`
│   │   ├── controller
│   │   │   ├── controller.go
│   │   │
│   │   ├── service
│   │   │   ├── service.go
│   │   │
│   │   ├── core.go
│   
├── go.mod
├── go.sum
├── README.md
├── Makefile
```