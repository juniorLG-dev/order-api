# Order API

Uma simples API que usa Arquitetura Orientada a Eventos e RabbitMQ na composição do event bus. 

## Como Executar o Projeto Localmente

Siga os passos abaixo para configurar e rodar o projeto na sua máquina.

### Pré-requisitos

Certifique-se de que você tem o ambiente Go configurado e o Go Modules ativado.

### ⚙️ Configuração

O projeto roda na porta padrão **8080** e requer apenas duas variáveis de ambiente para funcionar corretamente.

#### 1. Variáveis de Ambiente

Crie um arquivo `.env` (ou declare as variáveis diretamente no seu terminal) com as seguintes chaves:

| Variável | Descrição | Exemplo |
| :--- | :--- | :--- |
| **`EMAIL`** | Seu endereço de email para o serviço. | `john@doe.com` |
| **`KEY`** | Sua chave secreta. | `john123` |

**Exemplo de declaração no terminal (Linux/macOS):**

```bash
export EMAIL="john@doe.com"
export KEY="key"
```
#### 2. Rodar o projeto
Baixe as dependências e inicialize o container. 
```bash
go mod tidy
sudo docker-compose up -d
go run main.go
```
