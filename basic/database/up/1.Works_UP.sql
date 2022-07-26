CREATE TABLE IF NOT EXISTS works (
    ID  serial primary key,
    Name text  not null,
    WorkField text  not null,
    Description text not null,
    Price float8 not null,
    Date DATE not null,
    CompanyID integer not null
)