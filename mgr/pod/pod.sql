#
************************************************************
# Sequel Ace SQL dump
# 版本号
： 20048
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主机: 127.0.0.1 (MySQL 8.0.32)
# 数据库: easypaas
# 生成时间: 2023-07-29 08:12:47 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


#
转储表 pod
# ------------------------------------------------------------

DROP TABLE IF EXISTS `pod`;

CREATE TABLE `pod`
(
    `id`              bigint NOT NULL AUTO_INCREMENT,
    `pod_name`        varchar(255) DEFAULT NULL,
    `pod_namespace`   varchar(255) DEFAULT NULL,
    `pod_team_id`     varchar(255) DEFAULT NULL,
    `pod_cpu_min`     double       DEFAULT NULL,
    `pod_cpu_max`     double       DEFAULT NULL,
    `pod_replicas`    int          DEFAULT NULL,
    `pod_memory_min`  double       DEFAULT NULL,
    `pod_memory_max`  double       DEFAULT NULL,
    `pod_pull_policy` varchar(255) DEFAULT NULL,
    `pod_restart`     varchar(255) DEFAULT NULL,
    `pod_type`        varchar(255) DEFAULT NULL,
    `pod_image`       varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_pod_pod_name` (`pod_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



#
转储表 pod_env
# ------------------------------------------------------------

DROP TABLE IF EXISTS `pod_env`;

CREATE TABLE `pod_env`
(
    `id`        bigint NOT NULL AUTO_INCREMENT,
    `pod_id`    bigint       DEFAULT NULL,
    `env_key`   varchar(255) DEFAULT NULL,
    `env_value` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



#
转储表 pod_port
# ------------------------------------------------------------

DROP TABLE IF EXISTS `pod_port`;

CREATE TABLE `pod_port`
(
    `id`             bigint NOT NULL AUTO_INCREMENT,
    `pod_id`         bigint       DEFAULT NULL,
    `container_port` int          DEFAULT NULL,
    `protocol`       varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
