// 用户验证表
DROP TABLE IF EXISTS auths;
CREATE TABLE auths(
id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
username VARCHAR(10) NOT NULL COMMENT '登录账号',
password VARCHAR(10) NOT NULL COMMENT '登录密码',
created_at timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
)
ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

// 新人信息表
DROP TABLE IF EXISTS users;
CREATE TABLE users(
id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
he_name VARCHAR(10) NOT NULL COMMENT '新郎名称',
she_name VARCHAR(10) NOT NULL COMMENT '新娘名称',
wedding_date VARCHAR(10) NOT NULL COMMENT '阳历',
wedding_lunar VARCHAR(10) NOT NULL COMMENT '农历',
share_msg VARCHAR(255) NOT NULL COMMENT '分享消息',
created_at timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
)
ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

// 酒店信息表
DROP TABLE IF EXISTS hotels;
CREATE TABLE hotels(
id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
hotel_name VARCHAR(50) NOT NULL COMMENT '酒店名称',
hotel_address VARCHAR(50) NOT NULL COMMENT '酒店详细地址',
hotel_lat FLOAT(10) NOT NULL COMMENT '酒店经度',
hotel_lng FLOAT(10) NOT NULL COMMENT '酒店纬度',
created_at timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
user_id BIGINT(20) NOT NULL COMMENT '用户id'
)
ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

// 祝福评论表
DROP TABLE IF EXISTS comments;
CREATE TABLE comments(
id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
bless_msg TEXT NOT NULL COMMENT '祝福文字',
open_id varchar(255) NOT NULL COMMENT '微信openid',
created_at timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
user_id BIGINT(20) NOT NULL COMMENT '用户id'
)
ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

// 微信用户信息表
DROP TABLE IF EXISTS wx_users;
CREATE TABLE wx_users  (
id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
open_id varchar(255) NOT NULL COMMENT '微信openid',
union_id varchar(255) NOT NULL COMMENT '微信开放平台union_id',
nick_name TEXT NOT NULL COMMENT '昵称',
wx_face TEXT NULL DEFAULT NULL COMMENT '头像地址',
gender int(1) NULL DEFAULT 0 COMMENT '性别',
province varchar(10) NULL DEFAULT NULL COMMENT '省份',
city varchar(10) NULL DEFAULT NULL COMMENT '城市',
mobile varchar(22) NULL DEFAULT NULL COMMENT '手机号码',
created_at timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
)
ENGINE = InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
