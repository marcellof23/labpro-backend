CREATE TABLE IF NOT EXISTS `dorayaki_store` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nama` longtext NOT NULL,
  `jalan` longtext NOT NULL,
  `kecamatan` longtext NOT NULL,
  `provinsi` longtext NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1