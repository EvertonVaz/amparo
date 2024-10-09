# Desafio de codigo Amparo

## Pré-requisitos

- **Go**: Certifique-se de ter o Go instalado em sua máquina (versão 1.22 ou superior).
- **Cliente HTTP**: Para testar a API, você pode usar ferramentas como **Postman**, **Insomnia** ou **curl**.

## 1. Configurando o Ambiente

1. **Clone o repositório do projeto** ou copie os arquivos para um diretório local.

   ```bash
   git clone git@github.com:EvertonVaz/amparo.git
   ```

2. **Navegue até o diretório do projeto** no terminal:

   ```bash
   cd caminho/para/o/diretorio/amparo
   ```

3. **Inicialize o módulo Go** (caso ainda não tenha sido feito):

   ```bash
   go mod init amparo
   ```

4. **Baixe as dependências do projeto**:

   ```bash
   go mod tidy
   ```

## 2. Executando a API

Para iniciar o servidor da API, execute:

```bash
go run main.go
```

- O servidor iniciará em `http://localhost:8080`.

## 3. Utilizando a API

### Endpoints Disponíveis

1. **Criar Registro de Datas Importantes**

   - **Método**: `POST`
   - **URL**: `http://localhost:8080/important-dates/{userId}`
   - **Descrição**: Cria um registro de datas importantes para o usuário especificado.
   - **Parâmetros**:
     - `{userId}`: Identificador único do usuário (pode ser numérico ou texto).
   - **Corpo da Requisição** (JSON):

     ```json
     {
       "dateOfDeath": "yyyy-mm-dd"
     }
     ```

   - **Exemplo de Requisição com `curl`**:

     ```bash
     curl -X POST http://localhost:8080/important-dates/1 \
     -H "Content-Type: application/json" \
     -d '{"dateOfDeath": "2023-10-01"}'
     ```

   - **Resposta de Sucesso** (HTTP 200):

     ```json
     {
       "id": 1,
       "seventhDayMass": "2023-10-04",
       "deathRegistration": "2023-10-16",
       "inventoryOpening": "2023-11-30",
       "deathPensionRequest": "2023-12-30",
       "lifeInsuranceClaim": "2024-10-01"
     }
     ```

2. **Consultar Registro de Datas Importantes**

   - **Método**: `GET`
   - **URL**: `http://localhost:8080/important-dates/{userId}`
   - **Descrição**: Retorna o registro de datas importantes do usuário especificado.
   - **Parâmetros**:
     - `{userId}`: Identificador único do usuário.
   - **Exemplo de Requisição com `curl`**:

     ```bash
     curl http://localhost:8080/important-dates/1
     ```

   - **Resposta de Sucesso** (HTTP 200):

     ```json
     {
       "id": 1,
       "seventhDayMass": "2023-10-04",
       "deathRegistration": "2023-10-16",
       "inventoryOpening": "2023-11-30",
       "deathPensionRequest": "2023-12-30",
       "lifeInsuranceClaim": "2024-10-01"
     }
     ```

3. **Deletar Registro de Datas Importantes**

   - **Método**: `DELETE`
   - **URL**: `http://localhost:8080/important-dates/{userId}`
   - **Descrição**: Remove o registro de datas importantes do usuário especificado.
   - **Parâmetros**:
     - `{userId}`: Identificador único do usuário.
   - **Exemplo de Requisição com `curl`**:

     ```bash
     curl -X DELETE http://localhost:8080/important-dates/1
     ```

   - **Resposta de Sucesso** (HTTP 200):

     ```json
     {
       "message": "Registro deletado com sucesso"
     }
     ```

### Usando o Postman ou Insomnia

1. **Criar Registro (POST)**:

   - Selecione o método **POST**.
   - URL: `http://localhost:8080/important-dates/1`
   - No corpo da requisição, escolha o tipo **JSON** e insira:

     ```json
     {
       "dateOfDeath": "2023-10-01"
     }
     ```

   - Envie a requisição e verifique a resposta.

2. **Consultar Registro (GET)**:

   - Selecione o método **GET**.
   - URL: `http://localhost:8080/important-dates/1`
   - Envie a requisição e verifique a resposta.

3. **Deletar Registro (DELETE)**:

   - Selecione o método **DELETE**.
   - URL: `http://localhost:8080/important-dates/1`
   - Envie a requisição e verifique a resposta.

### Mensagens de Erro

- **Registro não encontrado** (HTTP 404):

  ```json
  {
    "error": "Registro não encontrado"
  }
  ```

- **Datas já existem para este usuário** (HTTP 409):

  ```json
  {
    "error": "Datas já existem para este usuário"
  }
  ```

- **Formato de data inválido** (HTTP 400):

  ```json
  {
    "error": "Formato de data inválido. Use yyyy-mm-dd"
  }
  ```

- **Data de óbito no futuro** (HTTP 400):

  ```json
  {
    "error": "Data de óbito não pode ser no futuro"
  }
  ```

## 4. Executando os Testes

Para executar os testes unitários e verificar se todas as funcionalidades estão operando corretamente:

```bash
go test ./tests
```

- Os resultados dos testes serão exibidos no terminal.

## 5. Estrutura do Projeto

O projeto está organizado da seguinte maneira:

```
.
├── main.go
├── handlers
│   └── dateHandler.go
├── models
│   └── date.go
├── tests
│   └── dateHandler_test.go
├── go.mod
└── go.sum
```

- **main.go**: Arquivo principal que inicia o servidor e define as rotas.
- **handlers/**: Contém as funções que lidam com as requisições HTTP.
- **models/**: Define as estruturas de dados utilizadas na aplicação.
- **tests/**: Contém os testes unitários para as funcionalidades implementadas.
- **go.mod** e **go.sum**: Arquivos de gerenciamento de dependências do Go.
