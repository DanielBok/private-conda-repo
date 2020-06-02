drop index IDX_PACKAGE_COUNT__PACKAGE_PLATFORM;
drop index IDX_PACKAGE_COUNT__PACKAGE;


alter table PACKAGE_COUNT
    add column CHANNEL varchar(50) not null default '';

alter table PACKAGE_COUNT
    rename to PACKAGE_COUNTS;

update PACKAGE_COUNTS
set CHANNEL = Q.CHANNEL
from (
         select ID, CHANNEL
         from CHANNEL
     ) as Q
where PACKAGE_COUNTS.CHANNEL_ID = Q.ID;

alter table PACKAGE_COUNTS
    alter column CHANNEL drop default;

alter table PACKAGE_COUNTS
    drop column CHANNEL_ID;

create unique index IDX_PACKAGE_COUNTS__CHANNEL_PACKAGE on PACKAGE_COUNTS (CHANNEL, PACKAGE);
create unique index IDX_PACKAGE_COUNTS__CHANNEL_PACKAGE_PLATFORM on PACKAGE_COUNTS (CHANNEL, PACKAGE, PLATFORM);


drop index CHANNEL_CHANNEL_KEY;

alter table CHANNEL
    rename column CREATED_ON to JOIN_DATE;

alter table CHANNEL
    rename to USERS;

alter table USERS
    add constraint USERS_CHANNEL_KEY unique (CHANNEL);
