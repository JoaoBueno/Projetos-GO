CREATE OR REPLACE FUNCTION revogar_usuario_serpro()
RETURNS VOID AS $$
DECLARE 
	nome_tabela text;
BEGIN
	FOR nome_tabela IN (SELECT table_name FROM information_schema.tables WHERE table_schema = 'public') 
    LOOP
        EXECUTE 'REVOKE ALL PRIVILEGES ON TABLE public.' || nome_tabela || ' FROM serpro';
    END LOOP;

	REVOKE ALL PRIVILEGES ON DATABASE "fullcontrol" FROM "serpro";
	REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA public FROM serpro;
	REVOKE ALL PRIVILEGES ON SCHEMA public FROM serpro;
	DROP USER serpro;
END;
$$ LANGUAGE plpgsql;

SELECT revogar_usuario_serpro();

CREATE OR REPLACE FUNCTION criar_usuario_serpro()
RETURNS VOID AS $$
BEGIN

	CREATE USER serpro WITH PASSWORD 'Primeira@1linh4';

	GRANT CONNECT on DATABASE "fullcontrol" TO "serpro";

	GRANT USAGE ON SCHEMA public TO serpro;

  GRANT SELECT, INSERT, UPDATE ON TABLE consultas_serpro TO serpro;
  GRANT SELECT, INSERT ON TABLE cnae TO serpro;
  GRANT SELECT ON TABLE cnae_divisao TO serpro;
  GRANT SELECT ON TABLE cnae_grupo TO serpro;
  GRANT SELECT ON TABLE cnae_classe TO serpro;
  GRANT SELECT, INSERT, UPDATE ON TABLE municipios TO serpro;
END;
$$ LANGUAGE plpgsql;

SELECT criar_usuario_serpro();

SELECT * FROM consultas_serpro
