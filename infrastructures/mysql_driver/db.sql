-- MySQL dump 10.13  Distrib 8.0.34, for Linux (x86_64)
--
-- Host: localhost    Database: db_xyzmultifinance
-- ------------------------------------------------------
-- Server version	8.0.34-0ubuntu0.22.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `consumers`
--

DROP TABLE IF EXISTS `consumers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `consumers` (
  `id` varchar(255) NOT NULL,
  `nik` varchar(255) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `legal_name` varchar(255) DEFAULT NULL,
  `tempat_lahir` varchar(255) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `gaji` decimal(10,2) DEFAULT NULL,
  `foto_ktp` varchar(255) DEFAULT NULL,
  `foto_selfie` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nik` (`nik`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consumers`
--

LOCK TABLES `consumers` WRITE;
/*!40000 ALTER TABLE `consumers` DISABLE KEYS */;
INSERT INTO `consumers` VALUES ('user-6af79619-c66f-4df7-98a2-b0ee282a3cc1','fc1742154cd0fc06ba378ad35886b478735c8f37fe7d02fdaaa62f0bb4c93c13','Budi','Budi Sutejo','Jakarta','2000-01-01',1500000.00,'foto_ktp_budi.jpg','foto_selfie_budi.jpg'),('user-e17227f6-19d5-4403-a12b-be8a77c74897','9876543210987654','Annisa','Annisa Putri','Bandung','2001-02-02',6700000.00,'foto_ktp_annisa.jpg','foto_selfie_annisa.jpg');
/*!40000 ALTER TABLE `consumers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tenors`
--

DROP TABLE IF EXISTS `tenors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tenors` (
  `id` varchar(255) NOT NULL,
  `consumer_id` varchar(255) NOT NULL,
  `limit_tenor1` int DEFAULT NULL,
  `limit_tenor2` int DEFAULT NULL,
  `limit_tenor3` int DEFAULT NULL,
  `limit_tenor4` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `consumer_id` (`consumer_id`),
  CONSTRAINT `tenors_ibfk_1` FOREIGN KEY (`consumer_id`) REFERENCES `consumers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tenors`
--

LOCK TABLES `tenors` WRITE;
/*!40000 ALTER TABLE `tenors` DISABLE KEYS */;
INSERT INTO `tenors` VALUES ('tenor-0c8215f2-bcb8-4b2b-bb73-6558c6d9a64d','user-6af79619-c66f-4df7-98a2-b0ee282a3cc1',100000,200000,0,700000),('tenor-34d452b0-c4b2-43b5-9db6-45e73592cd34','user-e17227f6-19d5-4403-a12b-be8a77c74897',1000000,1200000,1500000,2000000);
/*!40000 ALTER TABLE `tenors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` varchar(255) NOT NULL,
  `nomor_kontrak` int NOT NULL,
  `consumer_id` varchar(255) NOT NULL,
  `otr` decimal(10,2) DEFAULT NULL,
  `admin_fee` decimal(10,2) DEFAULT NULL,
  `jumlah_cicilan` int DEFAULT NULL,
  `jumlah_bunga` decimal(10,2) DEFAULT NULL,
  `nama_asset` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nomor_kontrak` (`nomor_kontrak`),
  KEY `consumer_id` (`consumer_id`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`consumer_id`) REFERENCES `consumers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES ('transaction-fUPOfqhkZ1NxlDIj',66577011,'user-6af79619-c66f-4df7-98a2-b0ee282a3cc1',1500000.00,5000.00,500000,10.00,'Samsung Galaxy A 01','2023-10-31 05:27:22');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-10-31  5:30:08