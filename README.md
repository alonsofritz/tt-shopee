# Ticket Processing Queue - Shopee Technical Challenge

## Visão Geral

Este projeto é um sistema esqueleto para processamento de pedidos de tickets para shows, desenvolvido como parte de um teste técnico para a Shopee. O sistema foi projetado de forma que permita a substituição de componentes in-memory por soluções mais robustas como RabbitMQ e MySQL no futuro.

## Problemática

### Contexto
Em plataformas de venda de ingressos, múltiplos usuários podem tentar comprar o mesmo ticket simultaneamente. Isso pode causar:

Race conditions e overselling se não houver controle.

Sobrecarga do sistema com processamento síncrono de muitas requisições.

Lentidão ou travamento durante picos de acesso.

Por isso, é comum enfileirar pedidos a partir de um endpoint e processá-los de forma assíncrona, garantindo consistência e melhor experiência do usuário.

### Solução Proposta
Este sistema implementa um padrão de **processamento assíncrono** utilizando filas de mensageria para:

1. **Desacoplamento**: Separar o recebimento da requisição do processamento do ticket
2. **Escalabilidade**: Permitir processamento paralelo e distribuído
3. **Confiabilidade**: Garantir que nenhum pedido seja perdido
4. **Performance**: Resposta rápida ao usuário enquanto processa em background

### Componentes Principais

#### 1. **Modelos de Domínio**
- **User**: Representa um usuário do sistema
- **Show**: Representa um evento/show disponível
- **Ticket**: Representa um pedido de ticket

#### 2. **Sistema de Mensageria**
- **Interface TicketPublisher**: Define o contrato para publicação de mensagens
- **TicketPublisherMemory**: Implementação in-memory usando Go channels
- **TicketPublisherRabbitMQ**: Implementação futura para RabbitMQ

#### 3. **Repositórios**
- **UserRepository**: Interface para operações com usuários
- **ShowRepository**: Interface para operações com shows
- Implementações in-memory com possibilidade de migração para MySQL

#### 4. **API REST**
- **Health Handler**: Endpoint para verificação de saúde da aplicação
- **Ticket Handler**: Endpoint para criação de pedidos de ticket

## Fluxo de Processamento

1. **Recebimento do Pedido**:
   ```
   POST /tickets
   {
     "show_id": "123",
     "user_id": "456"
   }
   ```

2. **Validação Inicial**:
   - Verificar se o usuário existe
   - Verificar se o show existe

3. **Publicação na Fila**:
   - Criar objeto Ticket
   - Enviar para fila de processamento assíncrono

4. **Processamento Assíncrono**:
   - Consumer processa ticket da fila

5. **Resposta ao Cliente**:
   - Retorno imediato com status "processing"
   - Notificação posterior sobre resultado

## Possíveis Melhorias Futuras

### 1. **Infraestrutura**
- [ ] Migração para RabbitMQ
- [ ] Implementação de banco de dados MySQL/PostgreSQL
- [ ] Sistema de cache com Redis

### 2. **Funcionalidades**
- [ ] Sistema de reserva temporária
- [ ] Notificações em tempo real (WebSockets)
- [ ] Gestão de filas por prioridade

### 3. **Observabilidade**
- [ ] Logging estruturado
- [ ] Métricas com Prometheus
- [ ] Tracing distribuído
- [ ] Health checks avançados

### 4. **Segurança**
- [ ] Autenticação JWT
- [ ] Rate limiting
- [ ] HTTPS obrigatório
- [ ] Configuração de CORS

## Como Executar

```bash
# Compilar a aplicação
make build

# Executar em modo production
make prod

# Ou executar diretamente em modo desenvolvimento
make run
```

## Testando a API

```bash
# Health check
curl http://localhost:3001/health

# Criar pedido de ticket
curl -X POST http://localhost:3001/tickets \
  -H "Content-Type: application/json" \
  -d '{
    "show_id": "230920",
    "user_id": 1
  }'
```
