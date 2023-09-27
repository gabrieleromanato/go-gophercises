# SQL Export
# Created by Querious (400047)
# Created: 27 September 2023 at 09:56:36 CEST
# Encoding: Unicode (UTF-8)


SET @ORIG_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;

SET @ORIG_UNIQUE_CHECKS = @@UNIQUE_CHECKS;
SET UNIQUE_CHECKS = 0;

SET @ORIG_TIME_ZONE = @@TIME_ZONE;
SET TIME_ZONE = '+00:00';




DROP TABLE IF EXISTS `numbers`;


CREATE TABLE `numbers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;




LOCK TABLES `numbers` WRITE;
INSERT INTO `numbers` (`id`, `number`) VALUES 
	(1,'(955)562-7937'),
	(2,'+1-003-025-4794x243'),
	(3,'001-801-965-7464x9243'),
	(4,'808.828.3842x03570'),
	(5,'782-153-5870x6709'),
	(6,'(360)121-2686'),
	(7,'587-301-9295'),
	(8,'914-014-3300x126'),
	(9,'456-024-0229x00456'),
	(10,'366-288-3149'),
	(11,'(413)692-1062'),
	(12,'001-507-401-5312x396'),
	(13,'(770)434-5060'),
	(14,'001-128-209-2241x3767'),
	(15,'001-582-285-9283x72011'),
	(16,'670-743-4335x433'),
	(17,'+1-128-940-6451'),
	(18,'0184773164'),
	(19,'001-643-500-1702x296'),
	(20,'+1-205-154-6326x89890'),
	(21,'001-785-588-0010x9567'),
	(22,'865.040.7055'),
	(23,'(746)127-9429x466'),
	(24,'001-439-176-6824x5325'),
	(25,'259.252.3823x08502'),
	(26,'026-203-1650x697'),
	(27,'227-276-3395x3544'),
	(28,'+1-548-453-7547x208'),
	(29,'(760)646-8602'),
	(30,'449.817.3502x651'),
	(31,'(331)191-0059'),
	(32,'445-533-2379'),
	(33,'2247742046'),
	(34,'(360)416-4613'),
	(35,'+1-195-239-8649x6234'),
	(36,'(122)053-4593'),
	(37,'712.005.4172'),
	(38,'+1-959-354-0978'),
	(39,'+1-376-756-6707x4460'),
	(40,'668-876-8270x3311'),
	(41,'(304)219-1439x2837'),
	(42,'001-773-743-7795'),
	(43,'074.418.8778x90182'),
	(44,'+1-095-690-1883x59686'),
	(45,'+1-836-596-0015x426'),
	(46,'(680)102-3399'),
	(47,'(188)523-2451x889'),
	(48,'7221528048'),
	(49,'+1-345-611-6812x7404'),
	(50,'281-555-5626x8411'),
	(51,'698.896.4196x79006'),
	(52,'001-653-060-4477x232'),
	(53,'(453)143-7507'),
	(54,'+1-358-697-1073x9844'),
	(55,'175.861.1593x071'),
	(56,'823.696.3381'),
	(57,'(567)633-1521x751'),
	(58,'001-912-109-4288x816'),
	(59,'(710)142-7188x72039'),
	(60,'001-189-712-8058'),
	(61,'(754)987-5888'),
	(62,'237.375.7196x4419'),
	(63,'+1-061-248-1739'),
	(64,'(468)283-5302'),
	(65,'950-335-9871x5849'),
	(66,'+1-520-308-1789x8947'),
	(67,'001-032-518-2032x9551'),
	(68,'4669922204'),
	(69,'591.422.4942x480'),
	(70,'595-794-9219x20897'),
	(71,'517.229.0896x302'),
	(72,'+1-590-697-8298x72522'),
	(73,'2664730556'),
	(74,'790-233-0946'),
	(75,'279.444.3830x23504'),
	(76,'+1-099-796-8732'),
	(77,'748.240.2497x850'),
	(78,'001-785-704-5084x560'),
	(79,'304-618-0015x1536'),
	(80,'801.240.2444x96970'),
	(81,'081.039.8997x7013'),
	(82,'(131)571-9020'),
	(83,'168-913-2088x14982'),
	(84,'(409)234-5749x12790'),
	(85,'+1-934-114-9761'),
	(86,'060.617.0433x1440'),
	(87,'+1-597-096-9876x430'),
	(88,'(641)137-3487'),
	(89,'208-943-7425'),
	(90,'358.739.7081'),
	(91,'+1-659-067-6181x57283'),
	(92,'001-996-788-5747x849'),
	(93,'(063)869-4797x604'),
	(94,'814.639.2466'),
	(95,'001-723-183-8223x63204'),
	(96,'701-096-8304x16543'),
	(97,'(028)872-0312'),
	(98,'841-052-5952'),
	(99,'+1-695-000-1078x109'),
	(100,'632.319.5085x6701');
UNLOCK TABLES;






SET FOREIGN_KEY_CHECKS = @ORIG_FOREIGN_KEY_CHECKS;

SET UNIQUE_CHECKS = @ORIG_UNIQUE_CHECKS;

SET @ORIG_TIME_ZONE = @@TIME_ZONE;
SET TIME_ZONE = @ORIG_TIME_ZONE;



#  Export Finished: 27 September 2023 at 09:56:36 CEST