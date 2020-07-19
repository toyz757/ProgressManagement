DROP TABLE IF EXISTS tickets;

CREATE TABLE tickets (
	id int(8) NOT NULL AUTO_INCREMENT,
	title varchar(30),
	responsible varchar(30),
	deadline date,
	PRIMARY KEY (id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8