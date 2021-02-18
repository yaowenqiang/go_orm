> analyze verbose tablename
> vacuum tablename
 > query buffer
 > psql dbname "select * from table" -H -o /filename.html
  \copy tablename to tables.csv csv
  \copy tablename from  tables.csv csv
j
> mysq> select 1/0 # null
> psql> select now()/0 #error

> create table typesample(
    rowid serial, 
    myname character(50), 
    mylastname character varying(50), 
    myfirstname varchar(50), 
    borndate timestamp check(borndate > '1/1/1982'),
    borndatetz timestamptz check(borndate > '1/1/1982'),
    myprice float, 
    money decimal(10,2),
    luckydate date,
    luckytime time,
    lunchfoot text[],
    myip inet,

)

> insert into typesample (lunchfoot) values('{"rice", "dumpling","noodles"}');
> insert into typesample (luckydate, luckytime) values('today', 'allboalls');
> select lunchfoot[2] from typesample
> select * from typesample where lunchfoot[2] = "rice"
> select * from typesample where  "rice" = any(lunchfoot)
> update typesample set lunchfoot[2] = '';
> insert into typesample (borndate) values('1/1/1988');
> insert into typesample (borndatetz) values('1/1/1988 -6');

> insert into typesample (myname) values ("my name")
> insert into typesample (myname, myprice) values ("my name", "infinity")
> insert into typesample (myname, myprice) values ("my name", "-infinity")
> insert into typesample (myname, myprice) values ("my name", "NaN")

> select currval("typesample_rowid_seq")
> select nextval("typesample_rowid_seq")


# Table inheritance

```sql```
> create table companydrones (
    dronename varchar(100),
    number_of_monitors smallint,
    sickdays smallint
    ) ;

> insert into companydrones(dronename, number_of_monitors, sickdays) values('Rob sullivan', 2,14),('james avery', 1,31);

``````

> create table companydrones (
    supportemail varchar(100),
    keepcool bit,
    ) inherits (companydrones) ;



> create  table supportdrones(supportemail varchar(50), keepcool bit) inherits(companydrones);

insert into supportdrones(dronename, number_of_monitors, sickdays, keepcool, supportemail) values('jack yao', 1, 11, '0' , 'yaowenqiang111@cooco
 m'); 

//only select recoreds in companydrones
> select * from only companydrones

> create  table managementdrones(officenumber int, annoyingparkingspot boolean) inherits(supportdrones);

> insert into managementdrones(dronename, number_of_monitors, sickdays, supportemail, keepcool, officenumber, annoyingparkingspot) values('scoot h
 an', 15, 72, 'scooth@stackover.com', '1', 1, true);

> alter table managementdrones inherit companydrones
> alter table managementdrones no inherit supportdrones

## Concurrency and the MVCC(Multi Version Concurrency Control)

> dirty read
> shared lock
> exclusive lock

> select * from badges (nolock) // sql server query

> create table badges (name varchar(50), badgeid integer, userid integer);

> insert into badges(userid, name, badgeid) values(1, 'old name', 11);


// pg

> begin transaction:
update badges set name = 'new name' where  userid = 1;
select * from badges where userid = 1 // will get the new name in the transaction


> select * from badges where userid = 1 ; // run in another session will get the old name 

> begin transaction:
update badges set name = 'another new name' where  userid = 1;//dead locks

## Performance tuning

> select p.postid, p.answercount, p.viewcount,p.title, p.tags, u.userid, u.displaynae, u.reputation
from posts p inner join users u on p.owneruserid = u.userid
where p.posttypeid = 1
order by p.creationdate desc
limit 20

> create index ix_postfrontpagesearch on posts using btree (posttypeid, creationdate desc)

> show cpu_index_tuple_cost;
> set cpu_index_tuple_cost .0005;
explain analyze select p.postid, p.answercount, p.viewcount,p.title, p.tags, u.userid, u.displaynae, u.reputation
from posts p inner join users u on p.owneruserid = u.userid
where p.posttypeid = 1
order by p.creationdate desc
limit 20

> show random_page_cost;
set random_page_cost = 8;





























