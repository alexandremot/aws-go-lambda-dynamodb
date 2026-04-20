# AWS Go Lambda DynamoDB (PoC)

Este projeto é uma prova de conceito (PoC) criada para demonstrar a integração entre AWS Lambda (usando Go) e DynamoDB. Foi desenvolvido como base para uma possível implementação em um projeto real.

## Estrutura do Projeto

- `cmd/main.go`: Ponto de entrada da aplicação Lambda.
- `cmd/lambda_deploy.sh`: Script para empacotamento e deploy da Lambda.
- `myaws/dynamoHandler.go`: Implementação das funções de manipulação do DynamoDB.
- `go.mod`: Gerenciamento de dependências Go.
- `.gitignore`: Arquivos e pastas ignorados pelo Git.

## Como executar

1. Instale as dependências:
   ```bash
   go mod tidy
   ```
2. Compile o projeto:
   ```bash
   go build -o main ./cmd/main.go
   ```
3. (Opcional) Use o script `lambda_deploy.sh` para empacotar e fazer o deploy na AWS Lambda.

## Funcionalidades
- Exemplo de integração entre AWS Lambda (Go) e DynamoDB
- Estrutura modular para facilitar manutenção e expansão

## Observações
- Este projeto é uma PoC e pode ser expandido conforme as necessidades do projeto final.
- Adaptável para outros cenários de integração com AWS e bancos NoSQL.

---

**Autor:** Alexandre Mota
