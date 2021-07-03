# eustoma
Blog golang Gin

###调用顺序：
main->router->application/controllers->logic->dao/daomysql  
main 入口
router 路由  
controllers 控制器  
logic 逻辑层 使用tools （验证码 随机字符串 日志 雪花id response md5）  
dao 数据库操作  

###模型：
参数模型-models params.go  
库实例模型：models：user.go  

###返回码
settings code_setting.go  
