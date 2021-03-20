create table tb_search_log
(
    id serial not null
        constraint tb_search_log_pk
            primary key,
    http_request varchar,
    timestamp timestamp with time zone default timezone('Asia/Jakarta'::text, now())
);

create unique index tb_search_log_id_uindex
	on tb_search_log (id);

