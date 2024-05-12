create table `groups`(
  id int unsigned not null auto_increment,
  name varchar(50) not null comment '群组名称',
  conversion_id varchar(50) not null default 'fdefgvrfasdfd',
  primary key(`id`),
  UNIQUE KEY `idx_name` (`name`) USING BTREE
)engine=INNODB  DEFAULT CHARSET=utf8mb4 COMMENT = '群组表';