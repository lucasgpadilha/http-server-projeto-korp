# HTTP Server Projeto Korp

Desafio técnico com Golang, Docker, NGINX, Prometheus, Grafana e Ansible.

## Arquitetura

O serviço `http-server-projeto-korp` roda em container na porta interna `8080`.
O acesso externo acontece somente via NGINX, publicado na porta `80` do host.

Fluxo:

```text
localhost:80 -> nginx -> http-server-projeto-korp:8080
```

O Prometheus coleta métricas em `/metrics` e o Grafana exibe um dashboard provisionado automaticamente.

## Endpoints

| Endpoint            | Descrição                                      |
| ------------------- | ---------------------------------------------- |
| `GET /projeto-korp` | Retorna JSON com nome do projeto e horário UTC |
| `GET /healthz`      | Health check da aplicação                      |
| `GET /metrics`      | Métricas no padrão Prometheus                  |

## Executar com Docker Compose

```bash
docker network create korp_bridge || true
docker compose up -d --build
curl http://localhost/projeto-korp
```

## Executar com Ansible

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml --ask-become-pass
```

## Acessos

| Serviço    | URL                                                            |
| ---------- | -------------------------------------------------------------- |
| Aplicação  | [http://localhost/projeto-korp](http://localhost/projeto-korp) |
| Prometheus | [http://localhost:9090](http://localhost:9090)                 |
| Grafana    | [http://localhost:3000](http://localhost:3000)                 |

Credenciais locais do Grafana:

```text
admin / admin
```

## Métricas

| Métrica                                         | Descrição                                   |
| ----------------------------------------------- | ------------------------------------------- |
| `http_server_projeto_korp_up`                   | Disponibilidade da aplicação                |
| `http_server_projeto_korp_requests_total`       | Volume de requisições no endpoint principal |
| `http_server_projeto_korp_total_requests_total` | Volume total de requisições                 |
| `http_server_projeto_korp_uptime_seconds`       | Tempo de atividade da aplicação             |
