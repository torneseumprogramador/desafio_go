-- Criar a base de dados se não existir
CREATE DATABASE IF NOT EXISTS desafio_go;
USE desafio_go;

-- Criar a tabela de Administradores se não existir
CREATE TABLE IF NOT EXISTS Administradores (
    Id VARCHAR(255) NOT NULL,
    Nome VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    Senha VARCHAR(255) NOT NULL,
    PRIMARY KEY (Id)
);

-- Criar a tabela de Donos se não existir
CREATE TABLE IF NOT EXISTS Donos (
    Id VARCHAR(255) NOT NULL,
    Nome VARCHAR(255) NOT NULL,
    Telefone VARCHAR(255) NOT NULL,
    PRIMARY KEY (Id)
);

-- Criar a tabela de Pets se não existir
CREATE TABLE IF NOT EXISTS Pets (
    Id VARCHAR(255) NOT NULL,
    Nome VARCHAR(255) NOT NULL,
    DonoId VARCHAR(255) NOT NULL,
    Tipo int,
    PRIMARY KEY (Id),
    FOREIGN KEY (DonoId) REFERENCES Donos(Id)
);


INSERT INTO Administradores (Id, Nome, Email, Senha) VALUES
('6a4fa193-817c-4a10-9f0f-19c2192991d9', 'Danilo', 'danilo@teste.com', '123456');

INSERT INTO Donos (Id, Nome, Telefone) VALUES
('2872d84c-935e-420d-81fb-c1de636d3295', 'Josevaldo', '(11) 1111-1111'),
('1d81e756-fce2-4bd0-a10c-e7db4bfb224d', 'Marta', '(99) 99999-9999');


INSERT INTO Pets (Id, Nome, DonoId, Tipo) VALUES
('9532833d-a1fb-428c-a6f0-eb6eb332aaf0', 'Lili', '1d81e756-fce2-4bd0-a10c-e7db4bfb224d', 1),
('49b1d5a5-7ed5-4c69-b0f8-a9e092398fa5', 'Max', '2872d84c-935e-420d-81fb-c1de636d3295', 0);
