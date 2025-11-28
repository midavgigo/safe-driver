CREATE TABLE Services(
    Id                  SERIAL PRIMARY KEY UNIQUE,
    Name                VARCHAR(20)
);

INSERT INTO Services (Name) VALUES ('pet');
INSERT INTO Services (Name) VALUES ('luggage_rack');
INSERT INTO Services (Name) VALUES ('english');


CREATE TABLE Tarifs(
    Id                  SERIAL PRIMARY KEY UNIQUE,
    Name                VARCHAR(20)
);

INSERT INTO Tarifs (Name) VALUES ('econom');
INSERT INTO Tarifs (Name) VALUES ('comfort');
INSERT INTO Tarifs (Name) VALUES ('buisiness');


CREATE TABLE OrderStatuses(
    Id                  SERIAL PRIMARY KEY UNIQUE,
    Name                VARCHAR(30)
);

INSERT INTO OrderStatuses VALUES (0, 'pending');
INSERT INTO OrderStatuses VALUES (1, 'searching');
INSERT INTO OrderStatuses VALUES (2, 'driver_assigned');
INSERT INTO OrderStatuses VALUES (3, 'waiting_for_confirmation');
INSERT INTO OrderStatuses VALUES (4, 'in_progress');
INSERT INTO OrderStatuses VALUES (5, 'completed');

CREATE TABLE Drivers(
    Id                  SERIAL PRIMARY KEY UNIQUE,
    Available           BOOLEAN,
    Latitude            FLOAT,
    Longitude           FLOAT
);

INSERT INTO Drivers (Id) VALUES (0);

CREATE TABLE Orders(
    Id                  SERIAL PRIMARY KEY UNIQUE,
    AddressFrom         TEXT,
    AddressTo           TEXT,
    Tarif               INTEGER REFERENCES Tarifs (Id),
    PassengerId         INTEGER,
    SelectedServices    INTEGER[],
    Comment             TEXT,
    CurrentStatus       INTEGER REFERENCES OrderStatuses (Id),
    Driver              INTEGER REFERENCES Drivers (Id)
);

CREATE OR REPLACE FUNCTION MakeOrder(
    pAddressFrom         TEXT,
    pAddressTo           TEXT,
    pTarif               VARCHAR(20),
    pPassengerId         INTEGER,
    pSelectedServices    INTEGER[] DEFAULT '{}',
    pComment             TEXT DEFAULT ''
)
RETURNS INTEGER
AS $$
DECLARE
    tarif_id INTEGER;
    order_id INTEGER;
BEGIN
    SELECT Id 
    INTO tarif_id
    FROM Tarifs
    WHERE Name = pTarif;

    IF tarif_id IS NULL THEN
        RAISE EXCEPTION 'Tarif "%" not found', pTarif;
    END IF;
   
    INSERT INTO Orders (AddressFrom, AddressTo, Tarif, PassengerId, SelectedServices, Comment, CurrentStatus)
    VALUES (pAddressFrom, pAddressTo, tarif_id, pPassengerId, pSelectedServices, pComment, 0)
    RETURNING Id
    INTO order_id;
    RETURN order_id;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION GetOrderStatusName(
    pId INTEGER
)
RETURNS VARCHAR(30)
AS $$
DECLARE 
    name VARCHAR(30);
BEGIN
    SELECT os.Name 
    INTO name
    FROM Orders as o
    INNER JOIN OrderStatuses AS os
    ON o.CurrentStatus = os.Id;
    RETURN name;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE CancelOrder(
    pId INTEGER
)
AS $$
DECLARE
    status INTEGER;
BEGIN
    SELECT CurrentStatus
    INTO status
    FROM Orders
    WHERE Id = pId;
    IF status = 0 OR status = 1 THEN
        DELETE 
        FROM Orders 
        WHERE Id = pId;
    ELSE
        RAISE EXCEPTION 'Order is in progress or has already been completed';
    END IF;
END;
$$
LANGUAGE plpgsql;