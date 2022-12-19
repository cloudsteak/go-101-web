# go-101-web

Web Layer of very first Go lang microservice

## Image Creation

### Default

```bash
docker build --tag cloudsteak/go-101-web .
```

### Linux image

```bash
docker build --tag cloudsteak/go-101-web --platform linux/amd64 .
```

## Run docker image locally

```bash
docker run -d -p 80:3000 --name goweb01 cloudsteak/go-101-web:latest
```

### Force delete docker container

```bash
docker rm goweb01 --force
```

## Push image to DocherHub

### Login to DockerHub

```bash
docker login --username=<your DockerHub username>
```

### Tag image

```bash
docker tag $(docker images cloudsteak/go-101-web:latest -q) <your DockerHub username>/go-101-web:latest
```

### Push image to DockerHub

```bash
docker push <your DockerHub username>/go-101-web:latest
```

## Deploy to Kubernetes

### Create namespace

```bash
kubectl create namespace go-101
```

### Deploy web to new namespace

```bash
kubectl apply -f deployment.yaml --namespace go-101
```

### Check deployment status

```bash
kubectl get deployments -n go-101
```

### Check PODs

```bash
kubectl get pods -n go-101
```

### Expose deployemt

```bash
kubectl expose deployment go-101-web-deployment --type=NodePort --name=go-101-web-svc --target-port=3000 -n go-101
```

### Check nodeport

```bash
kubectl get svc -n go-101
```

## Autoscaling on Kubernetes

### Create

```bash
kubectl autoscale deployment go-101-web-deployment --cpu-percent=50 --min=1 --max=10 -n go-101
```

### Check

```bash
kubectl get hpa -n go-101
```

