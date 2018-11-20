create keyspace goland with replication = {'class':'SimpleStrategy','replication_factor' : 2};

use goland;

create table users
(
  id        int primary key,
  firstname varchar,
  lastname  varchar,
  job       varchar
);

create index idx_job on users(job);

insert into users (id, firstname, lastname, job) values (1, 'Florin', 'Patan', 'Developer Advocate');

select * from users where id = 1;
select * from users where job = 'Developer Advocate'