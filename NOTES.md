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

### Partitioning

> Partitioning makes large tables or indexes more manageable, because partitioning enables you to manage and access subsets of data quickly and efficiently, while maintaining the integrity of a data collection " - (SQL Server) MSDN"


posttypeid = 1 = questions
posttypeid = 2 = answers

> create table postroot (postid serial, posttypeid int, body text, userid int, lasteditoruserid int);
> create table questions (acceptedanswerid int, viewcount int, title varchar(255), tags varchar(255), closedate timestamptz, score int) inherits (postroot);
> create table answers (accepted bool, questionid int, score int) inherits (postroot);
> create table comments (questionid int) inherits (postroot);


> alter table postroot add constraint pk_postroot primary key (postid);
> alter table questions add constraint pk_questions primary key (postid);
> alter table answers add constraint pk_answers primary key (postid);

create index ix_question_creationdatedesc on questions (creationdate desc)

### Quering Text

select displayname from users where displayname = 'xxx';

create index ix_displayname on users using btree (displayname);

select displayname from users where lower(displayname) = 'xxx';


#### Index functions

create index ix_lowercaseddisplayname on users using btree (loser(displayname::text) collate pg_catelog."default");

### Full text indexing


select count(0) from posts where tags like '%postgresql%';

create index fti_tagsindex
on posts
using gin
(to_tsvector(english'::regconfig, tags::text));


btree  fit for

+ equality(=, <=, >=)
+ range(between in, < >)
+ sorting

rtree fit for

+ lines(line, lseg, interval)
+ spatial(cicle, path, polygon)
+ volumetrix(box)

hash :x=y

gin: full text search

tsvector:
The tsvector type represents a document in a form suited for text search

llexeme:

a levema is an abstract unit of morphological analysis in lingustics, that roughly corresponds to a set of forms taken by a single word


> select count(0) from posts where to_tsvcecor('english', tags) @@ plainto_tsquery('english', 'postgresql')



select to_tsvector('english', 'hello world');
select to_tsquery('english', 'hello');

select to_tsvector('english', 'hello world') @@ to_tsquery('english', 'hello');

tsvector is simply a data type - you can create a column in your table to hold tsvector information rather than convert it on the fly...

tsvector is a first-class data type, which means it can be indexed for performance...


alter table posts add column tagssvector tsvector;

create index ftx_tagsvectorindex
on posts
 using gin
 (tagsvector)



create trigger posts_tags_vector
before insert or update on posts
for each row
execute procedure tsvector_update_trigger('tagsvector', 'pg_catalog.english', 'tags')
update 

select count(0) from posts where tagsvector @@ plainto_tsquery('english', 'postgresql')


create or replace function getpostsbytag(searchy character varying)
    returns error posts as
    $BODY$
    select  * from posts
    where tagsvector @@ plainto_tsquery('english', $1);
    $BODY$
    language SQL
    cost 100
    rows 1000;



select getpostsbytag('postgres')
select * from getpostsbytag('postgres')






















