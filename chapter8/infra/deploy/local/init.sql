CREATE DATABASE IF NOT EXISTS `loon-on-boarding`;
USE `loon-on-boarding`;
CREATE TABLE IF NOT EXISTS employees(
    ID          INT AUTO_INCREMENT PRIMARY KEY,
    FullName    VARCHAR(255),
    Position    INT,
    Salary      FLOAT,
    Joined      DATE,
    OnProbation BOOLEAN,
    CreatedAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO employees ( FullName, Position, Salary, Joined, OnProbation )
        VALUES
            ("Oscar Contreras Palacios",2,500.0,"2022-12-31",false),
            ("Hector Contreras Palacios",3,400.0,"2022-12-31",true);