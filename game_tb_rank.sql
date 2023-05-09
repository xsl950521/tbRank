CREATE TABLE `game_tb_rank` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `numid` int(11) NOT NULL,
  `score` int(11) NOT NULL DEFAULT '0',
  `nickname` varchar(32) NOT NULL DEFAULT '',
  `ranknum` int(11) NOT NULL DEFAULT '0',
  `logtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `round` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `numid` (`numid`,`round`) USING BTREE,
  KEY `ranktime` (`ranktime`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;