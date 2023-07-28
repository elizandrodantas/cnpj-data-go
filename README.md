<br/>

<h1 align="center">
    Cadastro Nacional da Pessoa Jurídica - Ministerio da Fazenda
</h1>

Bem-vindo ao repositório do projeto de Extração e Tratamento de Dados para Bancos de Dados Relacionais! Este projeto foi desenvolvido com o objetivo de facilitar a obtenção, tratamento e inclusão de dados provenientes do Cadastro Nacional da Pessoa Jurídica (CNPJ), um banco de dados gerenciado pela Secretaria Especial da Receita Federal do Brasil (RFB).

A fonte principal dos dados utilizados neste projeto é o Cadastro Nacional da Pessoa Jurídica (CNPJ), um banco de dados que armazena informações cadastrais das pessoas jurídicas e outras entidades de interesse das administrações tributárias da União, dos Estados, do Distrito Federal e dos Municípios. O Ministério da Fazenda, órgão da estrutura administrativa da República Federativa do Brasil, é responsável pela formulação e execução da política econômica e tem papel fundamental na gestão do Cadastro Nacional da Pessoa Jurídica.

## Funcionalidades Principais
O projeto oferece as seguintes funcionalidades principais:

1. Extração de Dados: Capacidade de acessar a fonte de dados do CNPJ e obter informações cadastrais das pessoas jurídicas e entidades correlatas.

2. Tratamento de Dados: Implementação de mecanismos para limpeza, transformação e preparação dos dados extraídos, assegurando a qualidade e consistência das informações.

3. Integração com Bancos de Dados Relacionais: Facilidade para incluir os dados tratados em um banco de dados relacional escolhido, como SQLite, MySQL ou PostgreSQL, tornando-os acessíveis para aplicações e análises.

## Como Utilizar

Para utilizar o projeto, siga os passos a seguir:

1. **Clone o Repositório**: Abra um terminal ou prompt de comando e execute o seguinte comando para clonar o repositório em sua máquina local:

```bash
git clone https://github.com/elizandrodantas/cnpj-data-go.git
```

2. **Navegue até o Diretório do Projeto**: Acesse a pasta do projeto por meio do terminal, navegando para o diretório `cmd/cnpj-gov-go/`:

```bash
cd caminho/do/projeto/cmd/cnpj-data-go/
```

3. **Build com Go**: Com o Go instalado em sua máquina, realize o build do projeto com o seguinte comando:

```bash
go build
```

Após a execução bem-sucedida do comando acima, será gerado um executável com o nome do projeto no diretório atual.

4. **Execução do Projeto**: Agora que o projeto foi compilado, você pode executá-lo para iniciar o processo de extração e tratamento dos dados do CNPJ. Utilize o seguinte comando:

```bash
./cnpj-gov-go -D <driver> -h <host> -u <usuario> -p <senha> -P <porta> -dbname <nome_do_banco> -ssl <modo_ssl>
```

As flags disponíveis são:

- `-D`: Flag para selecionar o driver do banco de dados a ser utilizado. Opções possíveis: `sqlite3`, `mysql`, `postgres`. O padrão é `sqlite3` caso não seja especificado.
- `-h`: Host do banco de dados.
- `-u`: Nome do usuário para autenticação no banco de dados.
- `-p`: Senha do usuário para autenticação no banco de dados.
- `-P`: Porta para conexão com o banco de dados.
- `-dbname`: Nome do banco de dados que será utilizado.
- `-ssl`: Modo de conexão com SSL (caso seja necessário).

Nenhuma das flags é obrigatória. Caso você não forneça nenhuma delas, o projeto utilizará o banco de dados SQLite como padrão.

## Contribuindo

Se você deseja contribuir com melhorias, correções de bugs ou novas funcionalidades, fique à vontade para enviar um pull request.

## Licença

Este projeto é distribuído sob a licença `MIT`. Para mais detalhes, consulte o arquivo `LICENSE`.