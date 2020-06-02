create table if not exists PACKAGE_COUNTS
(
    ID           serial primary key,
    CHANNEL      varchar(50) not null,
    PACKAGE      varchar(64) not null,
    PLATFORM     varchar(10) not null,
    BUILD_STRING varchar(64) not null,
    BUILD_NUMBER int         not null default 0,
    VERSION      varchar(30) not null,
    COUNT        int         not null default 0,
    UPLOAD_DATE  timestamp   not null default NOW()
);

create index IDX_PACKAGE_COUNTS__CHANNEL_PACKAGE on PACKAGE_COUNTS
    (CHANNEL, PACKAGE);

create index IDX_PACKAGE_COUNTS__CHANNEL_PACKAGE_PLATFORM on PACKAGE_COUNTS
    (CHANNEL, PACKAGE, PLATFORM);
