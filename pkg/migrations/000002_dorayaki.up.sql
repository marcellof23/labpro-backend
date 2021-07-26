CREATE TABLE IF NOT EXISTS `dorayaki` (
  `id` bigint(20) DEFAULT NULL AUTO_INCREMENT,
  `rasa` longtext DEFAULT NULL,
  `deskripsi` longtext DEFAULT NULL,
  `gambar` longtext DEFAULT NULL,
  `jumlah`  bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `dorayaki_store_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dorayaki_store_doryakis` (`dorayaki_store_id`),
  CONSTRAINT `fk_dorayaki_store_doryakis` FOREIGN KEY (`dorayaki_store_id`) REFERENCES `dorayaki_store` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1