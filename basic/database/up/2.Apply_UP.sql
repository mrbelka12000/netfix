CREATE TABLE IF NOT EXISTS Apply(
    ID  serial primary key,
    CustomerID integer not null,
    WorkID integer not null,
    Start integer not null,
    End integer, 
    FOREIGN KEY (CustomerID) REFERENCES customer(id),
    FOREIGN KEY (WorkID) REFERENCES works(id)
)