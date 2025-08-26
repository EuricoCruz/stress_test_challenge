# stress_test_challenge
Desafio do técnico da pós de golang da fullcycle

## objetivo
O objetivo principal é criar uma aplicação que simule um cenário de stress test, onde a aplicação deve ser capaz de suportar uma alta carga de requisições sem apresentar falhas ou degradação significativa no desempenho.

### 1. Como executar a aplicação

- Clone o repositório: 
```bash
git clone https://github.com/EuricoCruz/stress_test_challenge.git
```
- Navegue até o diretório do projeto:
```bash
cd stress_test_challenge
```
- Execute a aplicação:
```bash 
go run cmd/main.go --url=http://google.com --requests=100 --concurrency=10
```

### 2. Build da aplicação com docker
- Build da imagem:
```bash
docker run --rm stress_test --url=http://google.com --requests=100 --concurrency=10
```
- Executar a imagem:
```bash
docker run --rm stress_test --url=http://google.com --requests=100 --concurrency=10
```