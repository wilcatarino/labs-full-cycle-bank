# Sobre

- Banco fictício com checkout

# Desafio

- Performance entre o banco de dados e a aplicação (checkout)
- gRPC (HTTP2)
- Apache Kafka
- Elasticsearch e Kibana com Kafka Connect
- Istio, Kiali, Prometheus e Grafana
- Docker e Kubernetes

# Arquitetura

Banco (backend) - Golang
Extrato (backend) - 
Extrato (frontend) - 
Checkout (backend) - Nest.js (TypeScript)
Checkout (frontend) - Next.js

Banco publica no tópico do Apache Kafka
Extrato fica lendo o tópico do Apache Kafka
Kibana (frontend) acessa o Elasticsearch
Kafka Connect alimenta o Elasticsearch através do tópico do Apache Kafka 

# Comandos

sudo docker-compose up -d
sudo docker exec -it catarinobank bash
sudo docker-compose down
