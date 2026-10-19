# Demo App - Полный CI/CD цикл и мониторинг

Этот репозиторий содержит исходный код приложения, Dockerfile, Helm-чарт и пайплайны CI/CD для его автоматической сборки и развертывания в Kubernetes.

## 🏗 Об устройстве репозитория

- `main.go`, `go.mod` — исходный код приложения на Go (отдает JSON, имеет эндпоинты `/healthz` и `/metrics`).
- `Dockerfile` — многоступенчатая сборка Docker-образа (на базе Alpine).
- `helm/demo-app` — Helm-чарт для развертывания приложения в Kubernetes (Deployment, Service, ServiceMonitor).
- `.github/workflows/` — пайплайны GitHub Actions:
  - `ci.yml` — тестирование и сборка при каждом пуше в `main`.
  - `cd.yml` — сборка Docker-образа и публикация в GitHub Container Registry (GHCR) при создании тега.

## 🚀 Развертывание инфраструктуры

Инфраструктура описана с помощью Terraform в отдельном репозитории `infra-repo`.
1. Настройте доступ к Yandex Cloud: `yc init`.
2. Перейдите в папку `terraform` и выполните `terraform init` и `terraform apply`.
3. Terraform создаст виртуальную машину в Yandex Cloud и автоматически установит на нее легковесный Kubernetes-кластер **K3s**.
4. Состояние Terraform (state) хранится в S3-бакете (Yandex Object Storage).

## 📦 Развертывание приложения

Приложение разворачивается с помощью Helm:
```bash
helm install my-demo helm/demo-app
