CREATE TABLE IF NOT EXISTS Apply(
    ID  serial primary key,
    CustomerID integer not null,
    WorkID integer not null,
    StartDate integer not null,
    EndDate integer
)