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



























