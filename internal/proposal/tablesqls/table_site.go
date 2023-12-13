package tablesqls

func CreateSiteTableSql() (sql string) {

	sql = "CREATE TABLE `site` ("
	sql += "`id` int(10) unsigned NOT NULL AUTO_INCREMENT,"
	sql += "`category_id` int(11) DEFAULT NULL COMMENT '分类id',"
	sql += "`title` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网站标题',"
	sql += "`thumb` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网站logo',"
	sql += "`description` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网站描述',"
	sql += "`url` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网站地址',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',"
	sql += "`is_used` int(1) DEFAULT '-1' COMMENT '是否使用',"
	sql += "PRIMARY KEY (`id`) USING BTREE"
	sql += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='网站信息';"

	return
}

func CreateSiteTableDateSql() (sql string) {

	sql = "INSERT INTO `site` (category_id,title,thumb,description,url,created_at,updated_at,is_used) VALUES"
	sql += "(1,'Dribbble','assets/bootstrap/static/index/images/favicon.png','全球UI设计师作品分享平台。','https://dribbble.com/','2019-01-21 15:23:29','2019-03-12 02:13:08',1),"
	sql += "(2,'磁力搜索','assets/bootstrap/static/index/images/favicon.png','磁力搜索','http://cilisou8.com/','2019-12-22 14:01:00','2019-12-22 14:01:00',1);"


	return
}
