CREATE TABLE cnae_divisao (
  id VARCHAR(2) PRIMARY KEY,
  descricao VARCHAR NOT NULL
);

SELECT * FROM cnae_divisao;

CREATE TABLE cnae_grupo (
	id VARCHAR(3) PRIMARY KEY,
	descricao VARCHAR NOT NULL
)

SELECT * FROM cnae_grupo

CREATE TABLE cnae_classe (
	id VARCHAR(5) PRIMARY KEY, 
	descricao VARCHAR NOT NULL
)

SELECT * FROM cnae_classe ORDER BY id

CREATE TABLE cnae (
	id VARCHAR(7) PRIMARY KEY,
	descricao VARCHAR NOT NULL
)

SELECT * FROM cnae

-- Atualiza a descrição das tabelas tirando do uppercase
update cnae_classe
set descricao = upper(substring(descricao from 1 for 1)) ||
         lower(substring(descricao from 2 for length(descricao)))

