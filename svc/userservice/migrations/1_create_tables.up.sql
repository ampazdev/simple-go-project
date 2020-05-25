-- auto-generated definition
create table USER_CRED
(
    ID          serial  not null primary key,
    EMAIL        varchar not null,
    PASSWORD varchar not null
);

alter table USER_CRED
    owner to admin;

CREATE TABLE USER_INFO(
    ID integer PRIMARY KEY REFERENCES USER_CRED(ID),
    EMAIL varchar not null ,
    FULL_NAME varchar not null,
    PHONE_NUMBER varchar not null);
