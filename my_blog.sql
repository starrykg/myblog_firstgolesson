-- MySQL dump 10.13  Distrib 5.7.23, for Linux (x86_64)
--
-- Host: localhost    Database: my_blog
-- ------------------------------------------------------
-- Server version	5.7.23-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `article` (
  `uid` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(64) DEFAULT NULL,
  `summary` text,
  `body` text,
  `date` varchar(64) DEFAULT NULL,
  `filename` varchar(64) NOT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article`
--

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;
INSERT INTO `article` VALUES (9,'First post!','This is the summary.','<p>This is the main post!</p>\n\n<h1>Markdown!</h1>\n','2018-08-03 09:39:25','test3'),(10,'First post!','This is the summary.','<p>This is the main post!</p>\n\n<h1>Markdown!</h1>\n\n<p><em>it&rsquo;s</em> <strong>golang</strong>!</p>\n','2018-08-02 16:56:00','test'),(11,'First post!','这是一条简述','<p>This is the main post!</p>\n\n<h1>Markdown!</h1>\n\n<p><em>it&rsquo;s</em> <strong>golang</strong>!</p>\n','2018-08-07 14:36:31','test1'),(12,'First post!','This is the summary.','<p>This is the main post!</p>\n\n<h1>Markdown!</h1>\n','2018-08-02 18:13:29','test2'),(14,'# 江南皮革','This is the summary.','<p>This is the main post!\n浙江温州,浙江温州,最大皮革厂,江南皮革厂倒闭了! 原价都是三百多、二百多、一百多的钱包,通通二十块,通通二十块!</p>\n','2018-08-07 11:38:17','testbyjared'),(15,'# 江南皮革三厂','This is the summary.','<p><p>Golang的测试代码位于某个包的源代码中名称以_test.go结尾的源文件里，测试代码包含测试函数、测试辅助代码和示例函数；</p>\n<h2>测试函数有以Test开头的功能测试函数和以Benchmark开头的性能测试函数两种，测试辅助代码是为测试函数服务的公共函数、初始化函数、测试数据等，\n示例函数则是以Example开头的说明被测试函数用法的函数。</h2></p>\n','2018-08-07 15:07:48','summary');
/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-08-07 15:49:21
