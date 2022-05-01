CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS `usuarios`;

CREATE TABLE `usuarios` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(50) NOT NULL,
  `nick` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `senha` varchar(100) NOT NULL,
  `criadoEm` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nick` (`nick`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

DROP TABLE IF EXISTS `seguidores`;

CREATE TABLE `seguidores` (
  `usuario_id` int NOT NULL,
  `seguidor_id` int NOT NULL,
  FOREIGN KEY (`usuario_id`) REFERENCES usuarios(id) ON DELETE CASCADE,
  FOREIGN KEY (`seguidor_id`) REFERENCES usuarios(id) ON DELETE CASCADE,
  PRIMARY KEY(`usuario_id`, `seguidor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;