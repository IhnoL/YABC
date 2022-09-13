# YABC

## basic bash ;)
* ls  &emsp;&emsp;&emsp;&emsp;&emsp;&emsp; -> list content
* pwd &emsp;&emsp;&emsp;&emsp;&emsp;&emsp; -> current path
* cd ./directory &emsp;&emsp;-> change directory
* cd ..  &emsp;&emsp;&emsp;&emsp;&emsp;&emsp; -> go one up
* git clone https://github.com/IhnoL/YABC

&emsp;-> use tab-key extensively to autocomplete! - case sensitive!


## go:
* go build ./foo.go
* ./foo
* go build ./bar.go
* ./bar
* curl localhost:6060
* curl localhost:5050

## docker:
* docker build -t <your_docker_id>/foo .
* docker build -t <your_docker_id>/bar .
* docker images &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;  -> outputs existing images
* docker run -p 5050:5050 <your_docker_id>/foo
* docker run -p 6060:6060 <your_docker_id>/bar
* docker ps (--all)  &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp; -> outputs existing container
* docker logs --follow <container_name_or_id>
* docker start <container_name_or_id>
* docker stop <container_name_or_id>

## docker-compose:
* docker-compose up --build
* docker login
* docker-compose push
* docker-compose down
* https://hub.docker.com/u/icidenton

## kubernetes:
* kubectl apply -f .
* kubectl get <pods / services / ingress>
* kubectl describe ingress fancy-ingress
* kubectl get services --all-namespaces
* kubectl logs --follow <pod_name>
* kubectl exec --stdin --tty <pod_name> -- sh
* kubectl delete -f .

## helm:
* snap install helm --classic
* helm template ./foobar
* helm install foobar ./foobar
* helm list
* helm package foobar
* helm uninstall foobar
* helm repo add yabccharts https://ihnol.github.io/YABC-Charts/
* helm search repo
* helm install foobar yabccharts/foobar

#### Charts-Repo: https://github.com/IhnoL/YABC-Charts