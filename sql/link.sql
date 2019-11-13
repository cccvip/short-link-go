CREATE TABLE `ip_statistical` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(32) NOT NULL,
  `short` varchar(11) NOT NULL,
  `total` bigint(11) NOT NULL COMMENT '总访问次数',
  `gmt_create` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

