# Full Cycle Rate Limiter
Rate limiter criando em GoLang para a pós graduação Full Cycle

### Como configurar
Defina os limites de requests no arquivo `.env` \
Também é possível determinar qual tipo de rate limiter está ativado.

### Rate Limit By IP
Possibilita X requests por Y tempo do mesmo IP

### Rate Limit By API Key
Possibilita X requests por Y tempo com a mesma API Key informada no Header `API_KEY`

### Como iniciar o sistema

1. Na pasta raíz do sistema execute `make up` para iniciar as dependências do sistema.
2. Em outro terminal, execute `make run` para iniciar o servidor.
3. Com o servidor inicializado, é possível testar as requests:
```http request
GET http://localhost:8080/
API_KEY: "abc123"
```