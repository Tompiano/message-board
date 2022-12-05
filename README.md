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

LV2我实现的：
入参校验：用户名长度，密码长度，文章字数限制，文章内容是否含有敏感词汇
 增加个人信息，点赞，密码加盐，jwt鉴权，预防sql注入

LV3---------------------------------------------------------------------
create table comment(
    id varchar(10) primary key,
    userName varchar(10),
    pId varchar(10),
    content varchar(500)
);

#将要查找的Id作为rootId传入，自定义了一个用于遍历评论的函数
#查询rootId下面的所有子节点
create function getComment (rootId varchar(10))
returns varchar(10)
begin
     declare pId varchar(10);
     declare cId varchar(10);
     set pId=-1;-- 初始化pId
     set cId=rootId;-- 让cId等于要查找的Id
     while cId is not null do
         if (pId=-1)then-- pId=-1则说明这是第一轮的循环
             set pId=cId;-- 将其pId赋值为要查询的根节点的值
             elseif (pId<>-1)then-- 如果不是，则说明这不是第一轮循环
             set pId=concat(pId,',',cId);
         end if;
             select group_concat(id)into cId from comment where find_in_set(pId,cId);  
         end while;
     return pId;
end;

