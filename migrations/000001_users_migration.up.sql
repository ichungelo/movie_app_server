CREATE TABLE IF NOT EXISTS `movies` (
    `movie_id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(64) NOT NULL,
    `release_year` year DEFAULT NULL,
    `production` varchar(64) NOT NULL,
    `overview` text,
    PRIMARY KEY (`movie_id`),
    UNIQUE KEY `movie_id` (`movie_id`)
);