create table users(
  id        int primary key not null,
  username  varchar(50)        not null,
  job       varchar(50)
);

drop table users;

insert into users (id, username, job) values (1, 'dlsniper', 'Developer Advocate');

select * from users where id = 1;
select * from users where job = 'Developer Advocate'
