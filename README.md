api:存放路由router.go，以及将不同的接口函数按照性质放在不同的文件里
        user.go:与用户信息相关的操作：注册、登录、忘记密码、修改密码
        router.go:存放所有的路由入口

cmd:仅存放main.go文件，作为整个项目的入口

dao:数据访问层,所有数据库相关操作全部封装在这一层。将下层存储以更简单的函数、接口形式暴露给 service 层来使用。
        user.go:与用户信息有关的访问数据库操作
        message.go:与留言板有关的访问数据库操作

model:存放项目中所有的struct，按照文件名归类
        user.go:存放用户信息的有关的struct
        message.go:存放留言板信息的有关的struct

service:将各个服务封装之后供各个接口调用

util:装在出错时的响应的模板


LV3:实现嵌套的评论（我按照贴吧评论的方式写的）
ParentId如果为零，则说明它是Message下面的父级评论，则用ParentUserId来标识不同的评论
则子级评论即回复的ParentId就为对应的ParentUserId的值，不同的回复用ChildId来标识
将评论建成一张表ParentComment，将回复建成一张表ChildComment
评论根据MessageId和ParentUserId来辨别，回复根据MessageId和ParentId和ChildId来辨别