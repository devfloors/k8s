# Network

## 한 포드에 있는 다수의 컨테이너끼리 통신

-   pause 명령을 실행해 아무 동작하지 않는 빈 컨테이너 생성
-   인터페이스를 공유
-   포트를 겹치게 구성하지 못함

docker의 기능을 사용해 쿠버네티스 컨테이너 관찰
각 포드마다 하나의 pause 이미지 실행

## 포드끼리 통신

-   포드끼리 통신하기 위해서는 CNI 플러그인이 필요

```
$ sudo netstat -antp | grep weave
```

```
$ sudo docker ps | grep weave
```

## 포드와 서비스 사이의 통신

-   kube-proxy라는 컴포넌트로 서비스 트래픽을 제어
-   iptables는 리눅스 커널 기능인 netfilter를 사용하여 트래픽을 제어

-   ClusterIP를 생성하면 iptables의 설정 적용이 됨
-   iptable에서 목록 확인하는 실습

```
$ kubectl get svc --all-namespaces

$ sudo iptables -S -t nat | grep 10.96.0.1
```

## 외부 클라이언트와 서비스 사이의 통신

-   앞서 학습한대로 netfilter와 kube-proxy 기능을 사용해 원하는 서비스 및 포드로 연결

## DNS 서비스를 이용한 서비스 검색

-   서비스를 생성하면 대응되는 DNS 앤트리가 생성
-   이 앤트리는 <서비스-이름>. <네임스페이스-이름>.svc.cluster.local 형식을 가짐

```
$ kubectl exec -it http-go-rs bash

root@http-go-rs:/# curl http-go-svc.default.svc.cluster.local

root@http-go-rs:/# curl http-go-svc.default
```

## CoreDNS

-   내부에서 DNS 서버 역할을 하는 pod이 존재
-   각 미들웨어를 통해 로깅, 캐시, 쿠버네티스를 질의하는 기능을 가짐
-   해당 DNS에는 configmap 저장소를 사용해 설정 파일을 컨트롤
-   corefile을 통해 현재 클러스터의 NS를 지정
