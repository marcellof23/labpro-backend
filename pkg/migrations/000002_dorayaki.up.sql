CREATE TABLE IF NOT EXISTS `dorayaki` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `rasa` longtext NOT NULL,
  `deskripsi` longtext NOT NULL,
  `gambar` longtext NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `dorayaki_store_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dorayaki_store_doryakis` (`dorayaki_store_id`),
  CONSTRAINT `fk_dorayaki_store_doryakis` FOREIGN KEY (`dorayaki_store_id`) REFERENCES `dorayaki_store` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1