-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: bank
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `nasabah`
--

DROP TABLE IF EXISTS `nasabah`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nasabah` (
  `cif` int NOT NULL AUTO_INCREMENT,
  `nik` varchar(45) NOT NULL,
  `nama` varchar(45) NOT NULL,
  `tempat_lahir` varchar(45) NOT NULL,
  `tanggal_lahir` varchar(45) NOT NULL,
  `alamat` varchar(45) NOT NULL,
  `no_telepon` varchar(45) NOT NULL,
  PRIMARY KEY (`cif`),
  UNIQUE KEY `nik_UNIQUE` (`nik`),
  UNIQUE KEY `cif_UNIQUE` (`cif`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000005 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nasabah`
--

LOCK TABLES `nasabah` WRITE;
/*!40000 ALTER TABLE `nasabah` DISABLE KEYS */;
INSERT INTO `nasabah` VALUES (1000000000,'3603172102970002','ryan','bandung','2004-02-21','curug','081122334455'),(1000000001,'3603172102970003','suryo','jakarta','2005-02-21','kulon','082222334455'),(1000000002,'3603172102970005','hadi','tangerang','2006-02-21','cikupa','083322334455'),(1000000003,'3603172102970004','projo','cisauk','2007-02-21','tangerang','084422334455');
/*!40000 ALTER TABLE `nasabah` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nasabah_detail`
--

DROP TABLE IF EXISTS `nasabah_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nasabah_detail` (
  `cif` int NOT NULL,
  `no_rekening` int unsigned NOT NULL,
  `saldo` decimal(15,0) NOT NULL,
  KEY `cif_idx` (`cif`),
  CONSTRAINT `cif` FOREIGN KEY (`cif`) REFERENCES `nasabah` (`cif`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nasabah_detail`
--

LOCK TABLES `nasabah_detail` WRITE;
/*!40000 ALTER TABLE `nasabah_detail` DISABLE KEYS */;
INSERT INTO `nasabah_detail` VALUES (1000000000,1760100001,1513500),(1000000001,1760100002,1527500),(1000000003,1760100003,500000);
/*!40000 ALTER TABLE `nasabah_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaksi`
--

DROP TABLE IF EXISTS `transaksi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaksi` (
  `id_transaksi` int NOT NULL AUTO_INCREMENT,
  `id_user` int NOT NULL,
  `no_rekening` int NOT NULL,
  `tanggal` varchar(45) NOT NULL,
  `jenis_transaksi` varchar(45) NOT NULL,
  `nominal` decimal(15,2) NOT NULL,
  `saldo` decimal(15,2) NOT NULL,
  `berita` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_transaksi`),
  KEY `id_user_idx` (`id_user`),
  CONSTRAINT `id_user` FOREIGN KEY (`id_user`) REFERENCES `user` (`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaksi`
--

LOCK TABLES `transaksi` WRITE;
/*!40000 ALTER TABLE `transaksi` DISABLE KEYS */;
INSERT INTO `transaksi` VALUES (1,1760010,1760100001,'2020-08-10 14:06:56','st',50000.00,1050000.00,'nabung'),(2,1760010,1760100001,'2020-08-10 14:31:11','st',25000.00,1075000.00,'nabung'),(3,1760010,1760100001,'2020-08-10 14:42:11','st',50000.00,1125000.00,'nabung'),(4,1760010,1760100001,'2020-08-10 14:43:37','st',3000.00,1128000.00,'nabung'),(5,1760010,1760100001,'2020-08-10 14:45:29','st',4000.00,1132000.00,'nabung'),(6,1760010,1760100001,'2020-08-10 14:48:16','st',4000.00,1136000.00,'nabung'),(7,1760010,1760100001,'2020-08-10 15:17:04','st',500000.00,1636000.00,'menabung'),(8,1760010,1760100001,'2020-08-10 15:18:47','tt',200000.00,1436000.00,'penarikan'),(9,1760010,1760100001,'2020-08-10 15:35:40','tt',1400000.00,36000.00,'tarik'),(10,1760010,1760100001,'2020-08-10 15:47:19','st',2000000.00,2036000.00,'nabung'),(11,1760010,1760100001,'2020-08-10 19:59:57','st',5000.00,2041000.00,'nabung'),(12,1760010,1760100001,'2020-08-10 22:31:38','pb (d)',500000.00,1541000.00,'transfer'),(13,1760010,1760100002,'2020-08-10 22:31:38','pb (k)',500000.00,1500000.00,'transfer'),(14,1760010,1760100001,'2020-08-10 22:36:45','pb (d)',25000.00,1516000.00,'bayaran'),(15,1760010,1760100002,'2020-08-10 22:36:45','pb (k)',25000.00,1525000.00,'bayaran'),(16,1760010,1760100001,'2020-08-10 23:12:04','pb (d)',2000.00,1514000.00,'trf'),(17,1760010,1760100002,'2020-08-10 23:12:04','pb (k)',2000.00,1527000.00,'trf'),(18,1760010,1760100001,'2020-08-10 23:23:50','pb (d)',500.00,1513500.00,'trf'),(19,1760010,1760100002,'2020-08-10 23:23:50','pb (k)',500.00,1527500.00,'trf');
/*!40000 ALTER TABLE `transaksi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id_user` int NOT NULL AUTO_INCREMENT,
  `password` varchar(45) DEFAULT NULL,
  `nama_user` varchar(45) DEFAULT NULL,
  `role` varchar(45) DEFAULT NULL,
  `cabang` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `id_user_UNIQUE` (`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=1760151 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1760010,'password','ahmad','teller','jakarta'),(1760150,'password','ayu','cs','bandung');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'bank'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-11  4:31:11
