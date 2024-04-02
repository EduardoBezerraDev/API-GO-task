# API-GO-task

Esta é uma API RESTful escrita em Go para gerenciar tarefas.

## Como Testar

Siga as instruções abaixo para testar a API localmente.

### Pré-requisitos

- Go instalado (versão 1.13 ou superior)
- MySQL instalado e em execução
- Postman (ou qualquer cliente de API REST) instalado para testar as rotas

### Passo 1: Clone o Repositório

```bash
git clone https://github.com/seu-usuario/API-GO-task.git
```
### Passo 2: Clone o Repositório
Navegue até o diretório do projeto e instale as dependências do Go usando o seguinte comando:

```bash
go mod tidy
```
### Passo 3: Configure o Banco de Dados

Execute o script database.sql localizado em database/schema.sql para criar o banco de dados.

### Passo 4: Inicie o Servidor
Execute o seguinte comando para iniciar o servidor:
```bash
go run main.go
```

### Passo 5: Teste as Rotas
Abra o Postman (ou seu cliente de API REST preferido) e teste as seguintes rotas:

GET /api/tasks: Recupere todas as tarefas.

GET /api/tasks/{id}: Recupere uma tarefa específica pelo ID.

POST /api/tasks: Crie uma nova tarefa.

PUT /api/tasks/{id}: Atualize uma tarefa existente pelo ID.

DELETE /api/tasks/{id}: Exclua uma tarefa pelo ID.

Certifique-se de fornecer os parâmetros necessários nos corpos das solicitações (para rotas que exigem dados) e nos parâmetros de URL conforme necessário.

### Observações
Certifique-se de que o servidor esteja em execução enquanto testa as rotas.

Sempre valide os dados de entrada conforme necessário para evitar erros.

