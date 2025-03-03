# Инструкция по установке и запуску чарта

### 1. Создать namespace `m` и установить `nginx ingress` контроллер:

```bash
kubectl create namespace m && helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/ && helm repo update && helm install nginx ingress-nginx/ingress-nginx --namespace m --create-namespace --set controller.service.externalIPs={$(minikube ip)}
```

### 2. Установить чарт:

```bash
helm upgrade --install user-profile-app ./user-profile-chart -n m 
```

### 3. Для проверки работоспособности запустить `Postman` коллекцию:

```bash
newman run ./postman/hw_06.postman_collection.json
```
