1.创建数据库和表结构
2.beego过滤器
将根目录定位到static下
3.错误码定义
定义了一些
4.Controler主要负责与前段进行的数据交互,Models负责操作数据库(selece insert update delete)与Conttoler传递数据
(1)/api/v1.0/areas 
获取Area信息,AreaController负责处理该路由Get:GetArea
a.从数据库读取area信息 models.SelectAreaDate(),将读取到的信息以map形式返回给GetArea
b.将返回的map打包成JSON发送的客户端
(2)/api/v1.0/session
有两处该请求
GET 用户注册完之后SetSession,进入index.html的时候GetSession
DELETE 用户退出账号的时候DelSession
(3)/api/v1.0/users
用户注册账号的时候Post注册信息
a.获取前段的JSON数据 json.Unmarshal
b.插入数据库中 InsertUserData
