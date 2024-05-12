create table user(
  id int unsigned not null auto_increment,
  username varchar(50) not null comment '用户名',
  password varchar(50) not null comment '密码',
  real_name varchar(50) not null comment '真实姓名',
  phone_number char(11) not null default '18233334444' comment '电话号码',
  department_id  int unsigned not null comment '部门id',
  email varchar(50) not null default 'test@163.com' comment '邮箱',
  is_admin char(1) not null comment '是否是管理员:1-是;2-否',
  jenkins_account varchar(50) not null default 'test@163.com' comment 'Jenkins账号',
  jenkins_password varchar(50) not null default '123456789' comment 'Jenkins密码',
  avatar varchar(100) not null default 'dsfrgsdxcasds' comment '头像md5',
  constraint department foreign key (`department_id`) references department(`id`),
  primary key(`id`),
  index(`username`),
  index(`real_name`)
)engine=INNODB  DEFAULT CHARSET=utf8mb4 COMMENT = '用户表';