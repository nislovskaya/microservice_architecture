# Домашнее задание №3: Основы работы с Kubernetes (Часть 2)

---

## Цель:

В этом ДЗ вы научитесь создавать минимальный сервис.

## Описание/Пошаговая инструкция выполнения домашнего задания:

**Шаг 1.** Создать минимальный сервис, который

1. отвечает на порту 8000
2. имеет http-метод:

- GET /health/
- RESPONSE: {"status": "OK"}

**Шаг 2.** Собрать локально образ приложения в докер.  
Запушить образ в Dockerhub.

**Шаг 3.** Написать манифесты для деплоя в k8s для этого сервиса.  
Манифесты должны описывать сущности: Deployment, Service, Ingress.

В Deployment могут быть указаны Liveness, Readiness пробы.

Количество реплик должно быть не меньше 2. Image контейнера должен быть указан с Dockerhub.

Хост в ингрессе должен быть `arch.homework`.
В итоге после применения манифестов GET запрос на http://arch.homework/health должен отдавать `{“status”: “OK”}`.

**Шаг 4.**  На выходе необходимо предоставить:

- ссылку на github c манифестами. Манифесты должны лежать в одной директории, так чтобы можно было их все применить
  одной командой `kubectl apply -f .`
- url, по которому можно будет получить ответ от сервиса (либо тест в postman).

#### Задание со звездой

В Ingress-е должно быть правило, которое форвардит все запросы с `/otusapp/{student name}/*` на сервис с rewrite-ом
пути.
Где `{student name}` - это имя студента.

Например: `curl arch.homework/otusapp/aeugene/health` -> рерайт пути на `arch.homework/health`

### Рекомендации по форме сдачи дз:

- использовать `nginx ingress` контроллер, установленный через хелм, а не встроенный в миникубик:
  `kubectl create namespace m && helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/ && helm repo update && 
  helm install nginx ingress-nginx/ingress-nginx --namespace m --create-namespace --set controller.service.externalIPs={$(minikube ip)}`
- https://kubernetes.github.io/ingress-nginx/user-guide/basic-usage/
  необходимо в новых версиях nginx добавлять класс ингресса
  `ingressClassName: nginx`
- прикладывать к 2 дз урл для проверки: curl http://arch.homework/health или как указано в дз со *.
  К 3 дз и далее дз прикладывать коллекцию postman и проверять ее работу через `newman run имя_коллекции`
  прикладывать кроме команд разворачивания приложения, команду удаления)
- прописать у себя в `/etc/hosts` хост `arch.homework` с адресом своего миникубика (minikube ip), чтобы обращение было
  по
  имени хоста в запросах, а не айпи
- Обратите внимание, что при сборке на m1 при запуске вашего контейнера на стандартных платформах будет ошибка такого
  вида:
  `standard_init_linux.go:228: exec user process caued: exec format error`

Для сборки рекомендую указать тип платформы linux/amd64:
`docker build --platform linux/amd64 -t tag .`

Более подробно можно прочитать в статье:
https://programmerah.com/how-to-solve-docker-run-error-standard_init_linux-go219-exec-user-process-caused-exec-format-error-39221/

## Критерии оценки:

"Принято" - задание выполнено полностью   
"Возвращено на доработку" - задание не выполнено полностью   