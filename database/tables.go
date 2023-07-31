package database

var (
	CREATE_TABLES = []string{
		`CREATE TABLE IF NOT EXISTS empresas (
			cnpj_basico VARCHAR(8) NOT NULL,
			razao_social VARCHAR(255),
			natureza VARCHAR(10),
			qualificacao_responsavel VARCHAR(10),
			capital_social VARCHAR(50),
			porte_empresa VARCHAR(3),
			ente_federativo_resp VARCHAR(255)
		)`,
		`CREATE TABLE IF NOT EXISTS estabelecimentos (
			cnpj_basico VARCHAR(8) NOT NULL,
			cnpj_ordem VARCHAR(4),
			cnpj_dv VARCHAR(2),
			identificador_tipo VARCHAR(1),
			nome_fantasia VARCHAR(255),
			situacao_cadastral VARCHAR(11),
			data_situacao_cadastral VARCHAR(8),
			motivo_situacao_cadastral VARCHAR(10),
			nome_cidade_exterior VARCHAR(255),
			pais VARCHAR(11),
			data_inicio_atividade VARCHAR(8),
			cnae_principal VARCHAR(11),
			cnae_secundario VARCHAR(255),
			tipo_logradouro VARCHAR(50),
			logradouro VARCHAR(255),
			numero VARCHAR(50),
			complemento VARCHAR(255),
			bairro VARCHAR(50),
			cep VARCHAR(8),
			uf VARCHAR(2),
			municipio VARCHAR(11),
			ddd1 VARCHAR(2),
			telefone1 VARCHAR(9),
			ddd2 VARCHAR(2),
			telefone2 VARCHAR(9),
			ddd_fax VARCHAR(2),
			fax VARCHAR(10),
			correio_eletronico VARCHAR(255),
			situacao_especial VARCHAR(255),
			data_situacao_especial VARCHAR(8)
		);`,
		`CREATE TABLE IF NOT EXISTS simples (
			cnpj_basico VARCHAR(8) NOT NULL,
			opcao_simples VARCHAR(11),
			data_opcao_simples VARCHAR(8),
			data_exclusao_simples VARCHAR(8),
			opcao_mei VARCHAR(11),
			data_opcao_mei VARCHAR(8),
			data_exclusao_mei VARCHAR(8)
		);`,
		`CREATE TABLE IF NOT EXISTS socios (
			cnpj_basico VARCHAR(8) NOT NULL,
			identificador_socio VARCHAR(1),
			nome VARCHAR(255),
			documento VARCHAR(18),
			qualificacao VARCHAR(5),
			data_entrada VARCHAR(8),
			pais VARCHAR(11),
			representante_legal VARCHAR(11),
			nome_representante_legal VARCHAR(255),
			qualificacao_representante_legal VARCHAR(5),
			faixa_etaria VARCHAR(5)
		);`,
		`CREATE TABLE IF NOT EXISTS paises (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE IF NOT EXISTS municipios (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE IF NOT EXISTS qualificacoes (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE IF NOT EXISTS naturezas (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE IF NOT EXISTS cnaes (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE IF NOT EXISTS motivos (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
	}

	TABLES_NAME_LIST = []string{
		"empresas",
		"estabelecimentos",
		"simples",
		"socios",
		"paises",
		"municipios",
		"qualificacoes",
		"naturezas",
		"cnaes",
		"motivos",
	}

	ALTER_CHARSET = "ALTER TABLE `%s` CONVERT TO CHARACTER SET %s COLLATE %s;"
	ALTER_ENGINE  = "ALTER TABLE `%s` ENGINE=%s;"
	DROP_TABLES   = "DROP TABLE IF EXISTS %s;"
)
