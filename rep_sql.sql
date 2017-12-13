/*
SQLyog Ultimate v11.11 (64 bit)
MySQL - 5.7.14 : Database - localstudy
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`localstudy` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `localstudy`;

/*Table structure for table `score` */

DROP TABLE IF EXISTS `score`;

CREATE TABLE `score` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL DEFAULT '0',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '姓名',
  `type_position` varchar(25) NOT NULL COMMENT '位置固定标志',
  `course_id` smallint(5) NOT NULL DEFAULT '0' COMMENT '科目id',
  `score` tinyint(3) NOT NULL DEFAULT '0' COMMENT '分数',
  PRIMARY KEY (`id`),
  KEY `c_id` (`course_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4;

/*Data for the table `score` */

insert  into `score`(`id`,`user_id`,`name`,`type_position`,`course_id`,`score`) values (1,1,'张三','',1,89),(2,0,'例子','',2,78),(3,0,'刘德华','',1,89),(4,0,'成龙','',2,32),(5,0,'郭富城','',3,89),(6,0,'周润发','',2,39),(7,0,'刘翔','',1,77),(8,1,'张三','',2,98),(9,0,'例子','',1,99),(10,0,'刘德华','',2,8),(11,0,'刘翔','',2,68),(12,0,'周润发','',1,96),(13,1,'张三','',3,67),(14,0,'郭富城','',1,88),(15,0,'成龙','',1,99),(16,0,'刘翔','',3,79),(17,0,'郭富城','',2,89),(18,0,'例子','',3,77),(19,0,'周润发','',3,56),(20,0,'成龙','',3,88),(21,0,'刘德华','',3,100);

/*Table structure for table `student` */


DROP TABLE IF EXISTS `student`;

CREATE TABLE `student` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `birthday` varchar(10) NOT NULL DEFAULT '',
  `sex` tinyint(1) NOT NULL DEFAULT '0',
  `email` varchar(32) NOT NULL DEFAULT '',
  `address` varchar(255) NOT NULL DEFAULT '',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `student` */

insert  into `student`(`id`,`name`,`birthday`,`sex`,`email`,`address`,`update_time`,`create_time`) values (1,'王刚','14211122',1,'1044070679@qq.com','zhenjinaghangz','2017-12-13 19:44:45','2017-12-13 19:44:45');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
