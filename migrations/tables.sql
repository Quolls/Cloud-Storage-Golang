CREATE TABLE `file_metadata` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file_sha1` char(40) NOT NULL DEFAULT '',
  `file_name` varchar(256) NOT NULL DEFAULT '',
  `file_size` bigint(20) DEFAULT '0',
  `file_path` varchar(1024) NOT NULL DEFAULT '',
  `create_at` datetime default NOW(),
  `update_at` datetime default NOW() on update current_timestamp(),
  `status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_file_hash` (`file_sha1`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;