import psycopg2
import random
import json
import hashlib

# Conecta-se ao banco de dados PostgreSQL
conn = psycopg2.connect(
    host="192.168.254.6",
    database="fullcontrol",
    user="postgres",
    password="postgres",
    client_encoding="latin1"
)

# Cria um cursor para executar comandos SQL
cur = conn.cursor()

# Defina o JSONB de exemplo que você deseja inserir em cada registro
exemplo_jsonb = {
    "ni": "123456789",
    "porte": "05",
    "endereco": {
        # Defina os campos de endereço aqui
    },
    "telefones": [
        # Defina os números de telefone aqui
    ],
    "dataAbertura": "2000-05-03",
    "nomeFantasia": "EMPRESA",
    "capitalSocial": 100000000,
    "cnaePrincipal": {
        "codigo": "4642702",
        "descricao": "Comércio atacadista de roupas e acessórios para uso profissional e de segurança do trabalho"
    },
    "cnaeSecundarias": [
        # Defina os CNAEs secundários aqui
    ],
    "nomeEmpresarial": "EMPRESA LTDA",
    "naturezaJuridica": {
        "codigo": "2062",
        "descricao": "Sociedade Empresária Limitada"
    },
    "situacaoEspecial": "",
    "correioEletronico": "asd.COM.BR",
    "situacaoCadastral": {
        "data": "2005-11-03",
        "codigo": "2",
        "motivo": ""
    },
    "municipioJurisdicao": {
        "codigo": "0110100",
        "descricao": "BRASÍLIA"
    },
    "tipoEstabelecimento": "1",
    "dataSituacaoEspecial": ""
}

# Insira registros na tabela com CNPJs variados e CNAEs secundários variados
for i in range(1, 100):
    print(i)
    cnpj = str(i).zfill(14)  # Gere CNPJ com preenchimento de zeros à esquerda

    cnae_secundario = {
        "codigo": f"{random.randint(4000000, 4999999)}",
        "descricao": f"Descrição do CNAE {i}"
    }
    
    # Crie um novo JSONB baseado no exemplo_jsonb e adicione o CNAE secundário
    novo_jsonb = exemplo_jsonb.copy()
    novo_jsonb["cnaeSecundarias"] = [cnae_secundario]
    
    # Colocando todo o JSONB em md5
    md5_hash = hashlib.md5()
    md5 = json.dumps(novo_jsonb).encode('utf-8')
    md5_hash.update(md5)
    md5 = md5_hash.hexdigest()
    print(md5)
    print(type(md5))
    
    data = f"{random.randint(2022, 2023)}-{random.randint(1, 12):02d}-{random.randint(1, 28):02d}"
    print(data)
    print(type(data))
    
    cur.execute(
        "INSERT INTO consultas_serpro (cnpj, data, md5, dados) VALUES (%s, %s, %s, %s)",
        (cnpj, data, md5, json.dumps(novo_jsonb))  # Converte o dicionário para uma string JSON válida
    )

# Confirme as inserções no banco de dados
conn.commit()

# Feche o cursor e a conexão
cur.close()
conn.close()

print("Inserções concluídas.")
