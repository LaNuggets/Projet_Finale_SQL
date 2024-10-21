INSERT INTO department (Id, Title)
VALUES 
(1, 'Direction'),
(2, 'Human Resources'),
(3, 'Accountability'),
(4, 'Marketing'),
(5, 'Production');


INSERT INTO posts (Id, Title, DepartmentID, PermissionID, SalaryID)
VALUES
(1, 'Directeur Général', 1, 3, 4),        -- Direction
(2, 'Responsable RH', 2, 2, 3),           -- Human Resources
(3, 'Chargé de recrutement', 2, 1, 1),   -- Human Resources
(4, 'Comptable', 3, 2, 2),                -- Accountability
(5, 'Responsable Marketing', 4, 2, 4),    -- Marketing
(6, 'Assistant Marketing', 4, 1, 1),      -- Marketing
(7, 'Responsable de production', 5, 2, 3),-- Production
(8, 'Opérateur de production', 5, 1, 1),  -- Production
(9, 'Commercial', 5, 1, 2);               -- Production


INSERT INTO permission (Id, Title)
VALUES
(1, 'User'),         -- Niveau de permission 1
(2, 'Admin'),        -- Niveau de permission 2
(3, 'Super Admin');  -- Niveau de permission 3

INSERT INTO Salary (Id, Wage)
VALUES
(1, 'Low'),          -- Niveau de salaire faible
(2, 'Medium'),       -- Niveau de salaire moyen
(3, 'High'),         -- Niveau de salaire élevé
(4, 'Very High');    -- Niveau de salaire très élevé

INSERT INTO Employees (Id, LastName, FirstName, BirthDay, Phone, Address, PostId, ReferentId)
VALUES
(1, 'Durand', 'Jean', '1985-05-12', '0601020304', '12 rue de Paris', 1, NULL),   -- Directeur Général
(2, 'Martin', 'Claire', '1990-08-25', '0605060708', '8 avenue des Champs', 2, 1), -- Responsable RH
(3, 'Leclerc', 'Paul', '1982-11-14', '0610101112', '5 rue Victor Hugo', 3, 2), -- Chargé de recrutement
(4, 'Bernard', 'Marie', '1988-04-18', '0617181920', '15 boulevard Voltaire', 4, 1), -- Comptable
(5, 'Richard', 'Julien', '1992-06-30', '0621222324', '7 rue de la Paix', 5, 1),  -- Responsable Marketing
(6, 'Moreau', 'Luc', '1998-07-07', '0625262728', '10 allée des Roses', 7, 1),   -- Responsable de production
(7, 'Garcia', 'Laura', '1993-12-01', '0629303132', '2 impasse des Lilas', 6, 5), -- Assistant Marketing
(8, 'Petit', 'Antoine', '1991-03-15', '0633343536', '20 rue de la Liberté', 6, 5), -- Assistant Marketing
(9, 'Roux', 'Camille', '1989-10-20', '0637383940', '9 rue des Fleurs', 3, 1),  -- Responsable Marketing
(10, 'Blanc', 'Eva', '1994-09-05', '0641424344', '1 avenue des Tilleuls', 6, 5), -- Assistant Marketing
(11, 'Faure', 'Maxime', '1987-02-14', '0645464748', '11 rue du Moulin', 9, 1),   -- Commercial
(12, 'Lambert', 'Nina', '1990-12-31', '0649505152', '13 rue de l’Église', 8, 7), -- Opérateur de production
(13, 'Bonnet', 'Marc', '1996-06-06', '0653545556', '17 rue des Acacias', 8, 7), -- Opérateur de production
(14, 'Lopez', 'Julie', '1997-11-11', '0657585960', '21 boulevard St Michel', 8, 7); -- Opérateur de production


INSERT INTO referent (ReferentId, PostId)
VALUES
(1,2),
(1,3),
(1,4),
(1,5),
(1,7),
(1,9),
(2,3),
(5,6),
(7,8);