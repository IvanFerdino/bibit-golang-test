create table tb_user
(
	id serial not null
		constraint user_pk
			primary key,
	user_name varchar,
	parent integer
);

create unique index user_id_uindex
	on tb_user (id);

INSERT INTO tb_user(id, user_name, parent)
VALUES (1,'Ali',2),(2,'Budi',0),(3,'Cecep',1);


SELECT
tbu.id "ID",
tbu.user_name "UserName",
(select u.user_name from tb_user u where u.id = tbu.parent)
FROM tb_user tbu;

