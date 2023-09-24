package database

var (
	CREATE_TABLES = []string{
		`CREATE TABLE empresas (
			cnpj_basico VARCHAR(8) NOT NULL,
			razao_social VARCHAR(255),
			natureza VARCHAR(10),
			qualificacao_responsavel VARCHAR(10),
			capital_social VARCHAR(50),
			porte_empresa VARCHAR(3),
			ente_federativo_resp VARCHAR(255)
		);`,
		`CREATE TABLE estabelecimentos (
			cnpj_basico VARCHAR(8) NOT NULL,
			cnpj_ordem VARCHAR(4),
			cnpj_dv VARCHAR(4),
			identificador_tipo VARCHAR(1),
			nome_fantasia VARCHAR(255),
			situacao_cadastral VARCHAR(11),
			data_situacao_cadastral VARCHAR(8),
			motivo_situacao_cadastral VARCHAR(10),
			nome_cidade_exterior VARCHAR(255),
			pais VARCHAR(11),
			data_inicio_atividade VARCHAR(8),
			cnae_principal VARCHAR(11),
			cnae_secundario TEXT,
			tipo_logradouro VARCHAR(50),
			logradouro TEXT,
			numero VARCHAR(50),
			complemento VARCHAR(255),
			bairro VARCHAR(50),
			cep VARCHAR(8),
			uf VARCHAR(4),
			municipio VARCHAR(11),
			ddd1 VARCHAR(4),
			telefone1 VARCHAR(9),
			ddd2 VARCHAR(4),
			telefone2 VARCHAR(9),
			ddd_fax VARCHAR(4),
			fax VARCHAR(10),
			correio_eletronico VARCHAR(255),
			situacao_especial VARCHAR(255),
			data_situacao_especial VARCHAR(8)
		);`,
		`CREATE TABLE simples (
			cnpj_basico VARCHAR(8) NOT NULL,
			opcao_simples VARCHAR(11),
			data_opcao_simples VARCHAR(8),
			data_exclusao_simples VARCHAR(8),
			opcao_mei VARCHAR(11),
			data_opcao_mei VARCHAR(8),
			data_exclusao_mei VARCHAR(8)
		);`,
		`CREATE TABLE socios (
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
		`CREATE TABLE paises (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE municipios (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE qualificacoes (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE naturezas (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE cnaes (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,
		`CREATE TABLE motivos (
			codigo VARCHAR(11),
			descricao VARCHAR(255)
		);`,

		`CREATE INDEX idx_empresas_cnpj_basico ON empresas (cnpj_basico);`,
		`CREATE INDEX idx_estabelecimentos_cnpj_basico ON estabelecimentos (cnpj_basico);`,
		`CREATE INDEX idx_simples_cnpj_basico ON simples (cnpj_basico);`,
		`CREATE INDEX idx_socios_cnpj_basico ON socios (cnpj_basico);`,

		`CREATE INDEX idx_paises_codigo ON paises (codigo);`,
		`CREATE INDEX idx_municipios_codigo ON municipios (codigo);`,
		`CREATE INDEX idx_qualificacoes_codigo ON qualificacoes (codigo);`,
		`CREATE INDEX idx_naturezas_codigo ON naturezas (codigo);`,
		`CREATE INDEX idx_cnaes_codigo ON cnaes (codigo);`,
		`CREATE INDEX idx_motivos_codigo ON motivos (codigo);`,
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
