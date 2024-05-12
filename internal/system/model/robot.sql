create table robot(
  id int unsigned not null auto_increment,
  name varchar(50) not null comment '机器人名称',
  agent_id varchar(50) not null,
  appkey varchar(50) not null,
  appsecret varchar(50) not null,
  primary key(`id`)
)engine=INNODB  DEFAULT CHARSET=utf8mb4 COMMENT = '机器人表';