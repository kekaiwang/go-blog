package model

type ArticleRelation struct {
	Id         int64  `gorm:"primary_key"`
	Name       string `gorm:"name" json:"name"` // 分类名字
	RouterLink string `gorm:"router_link"`
}

// CREATE TABLE `wkk_category` (
// 	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
// 	`name` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '分类名字',
// 	`router_link` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '路由地址',
// 	`link_article` int(11) unsigned DEFAULT '0' COMMENT '关联文章数量',
// 	`status` tinyint(1) DEFAULT NULL COMMENT '1:正常 2:禁用 3:已删除',
// 	`created` datetime DEFAULT NULL COMMENT '创建时间',
// 	`updated` datetime DEFAULT NULL COMMENT '更新时间',
// 	PRIMARY KEY (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='文章分类表';
