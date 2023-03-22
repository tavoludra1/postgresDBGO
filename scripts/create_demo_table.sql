create table mydemo as select * from pg_tables;

create table x (
id integer,
name varchar(20)
);

insert into x (id, name) values (1, 'Pablo');
insert into x (id, name) values (2, 'Ana');
commit;
