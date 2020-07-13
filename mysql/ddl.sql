DROP TABLE IF EXISTS 'ticket';

CREATE TABLE 'ticket' (
	'ticket_id' int(8) NOT NULL,
	'title' varchar(30) NOT NULL,
	'responsible' varchar(30) NOT NULL,
	'deadline' date NOT NULL,
	PRIMARY KEY ('ticket_id')
)ENGINE = InnoDB DEFAULT CHARSET = utf8