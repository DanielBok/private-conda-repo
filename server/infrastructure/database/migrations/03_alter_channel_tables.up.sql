-- USERS -> CHANNEL

alter table USERS
    drop constraint USERS_CHANNEL_KEY;

alter table USERS
    rename to CHANNEL;

alter table CHANNEL
    rename column JOIN_DATE to CREATED_ON;

create unique index CHANNEL_CHANNEL_KEY on CHANNEL (lower(CHANNEL));


-- PACKAGE_COUNTS -> PACKAGE_COUNT

drop index IDX_PACKAGE_COUNTS__CHANNEL_PACKAGE;
drop index IDX_PACKAGE_COUNTS__CHANNEL_PACKAGE_PLATFORM;


alter table PACKAGE_COUNTS
    add column CHANNEL_ID serial not null references CHANNEL (ID) on delete cascade;

update PACKAGE_COUNTS
set CHANNEL_ID = Q.ID
from (
         select ID,
                CHANNEL
         from CHANNEL
     ) as Q
where PACKAGE_COUNTS.CHANNEL = Q.CHANNEL;

alter table PACKAGE_COUNTS
    rename to PACKAGE_COUNT;

alter table PACKAGE_COUNT
    drop column CHANNEL;

create index IDX_PACKAGE_COUNT__PACKAGE on PACKAGE_COUNT
    (CHANNEL_ID, PACKAGE);

create index IDX_PACKAGE_COUNT__PACKAGE_PLATFORM on PACKAGE_COUNT
    (CHANNEL_ID, PACKAGE, PLATFORM);
