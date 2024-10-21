CREATE TABLE IF NOT EXISTS employees (
    Id INTEGER PRIMARY KEY,
    LastName VARCHAR(25),
    FirstName VARCHAR(25),
    BirthDay DATE,
    Phone VARCHAR(25),
    Address VARCHAR(25),
    PostId INTEGER,
    ReferentId INTEGER,
    FOREIGN KEY (PostId) REFERENCES posts(Id)
    FOREIGN KEY (PostId) REFERENCES posts(Id)
);

CREATE TABLE IF NOT EXISTS posts (
    Id INTEGER PRIMARY KEY,
    Title VARCHAR(25),
    DepartmentId INTEGER,
    PermissionId INTEGER,
    ReferentId INTEGER,
    SalaryId INTEGER,
    FOREIGN KEY (DepartmentId) REFERENCES department(Id),
    FOREIGN KEY (PermissionId) REFERENCES permission(Id),
    FOREIGN KEY (SalaryId) REFERENCES salary(Id),
    FOREIGN KEY (ReferentId) REFERENCES referent(ReferentId)
);

CREATE TABLE IF NOT EXISTS referent (
    ReferentId INTEGER,
    PostId INTEGER,
    FOREIGN KEY (ReferentId) REFERENCES posts(Id)
    FOREIGN KEY (PostId) REFERENCES posts(Id)
);

CREATE TABLE IF NOT EXISTS department (
    Id INTEGER,
    Title VARCHAR(25)
);

CREATE TABLE IF NOT EXISTS permission (
    Id INTEGER,
    Title VARCHAR(25)
);

CREATE TABLE IF NOT EXISTS salary (
    Id INTEGER,
    Wage INTEGER
);

