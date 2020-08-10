CREATE TABLE `datasource` (
  `id` int NOT NULL,
  `addr` varchar(255) DEFAULT NULL,
  `cluster` varchar(255) DEFAULT NULL,
  `annotations` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `prometheus`.`datasource`(`id`, `addr`, `cluster`, `annotations`) VALUES (1, 'http://localhost:9090', 'p8s-test', '测试 Prometheus');
INSERT INTO `prometheus`.`datasource`(`id`, `addr`, `cluster`, `annotations`) VALUES (2, 'http://localhost:9090', 'p8s-local', '本地 Prometheus');

CREATE TABLE `rule` (
  `id` int NOT NULL AUTO_INCREMENT,
  `scene` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `interval` int DEFAULT NULL,
  `expr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `for` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `labels` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `annotations` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `group_id` int DEFAULT NULL,
  `datasource` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `maintainer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`scene`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `prometheus`.`rule`(`id`, `scene`, `name`, `interval`, `expr`, `for`, `labels`, `annotations`, `group_id`, `datasource`, `maintainer`) VALUES (1, 'alert', 'test up 1', 5, 'up==1', '5s', 'severity=page', 'summary={{ $labels.instance }} of {{ $labels.job }} has been down for more than 5 seconds.', 1, 'p8s-test', 'chenming');
INSERT INTO `prometheus`.`rule`(`id`, `scene`, `name`, `interval`, `expr`, `for`, `labels`, `annotations`, `group_id`, `datasource`, `maintainer`) VALUES (2, 'alert', 'trst up 2', 5, 'up==1', '5s', 'severity=page', 'summary={{ $labels.instance }} of {{ $labels.job }} has been down for more than 5 seconds.', 2, 'p8s-local', 'heyang');
