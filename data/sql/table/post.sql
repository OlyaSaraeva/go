CREATE TABLE post
(
   `post_id`      INT NOT NULL AUTO_INCREMENT,
   `Title`        VARCHAR(255) NOT NULL,
   `Subtitle`     VARCHAR(255) NOT NULL,
   `Emblema`       VARCHAR(255),
   `EmblemaTitle`   VARCHAR(255),
   `Outt`            VARCHAR(255),
   `BlockDirection`    VARCHAR(255),
   `Author`    VARCHAR(255),
   `AuthorImg`    VARCHAR(255),
   `Background`    VARCHAR(255),
   `SizeSmall`    VARCHAR(255),
   `PublishDate`    VARCHAR(255) NOT NULL,
   `featured`     TINYINT(1),
   PRIMARY KEY (`post_id`)
) ENGINE = InnoDB
CHARACTER SET = utf8mb4
COLLATE utf8mb4_unicode_ci
;