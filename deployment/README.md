### deployment 롤링 업데이트

```
$ kubectl create -f deployment.yaml --record=true
```

--record=true 옵션을 줘서 히스토리에 백업한다.

```
$ kubectl rollout status deploy http-go
```

rollout 에 관한 정보를 볼 수 있다.

```
$ kubectl rollout history deploy http-go
```

--record=true 옵션으로 히스토리를 볼 수 있다.

```
$  kubectl patch deploy http-go -p '{"spec":{"minReadySeconds":10}}'
```

롤링 업데이트 관찰을 위해
minReadySeconds:10 옵션으로 10초 뒤에 업데이트 하도록 패치한다.

### deploy 모니터링 하기

```
$ kubectl expose deploy http-go
```

deploy를 expose 한다.

```
$ kubectl get svc
```

http-go의 ip를 확인한다.

```
$ kubectl run -it --rm --image busybox -- bash
```

이 명령어로 busybox 이미지의 bash를 실행한다.

```
$ wget -O- -q 10.100.67.131:8080
```

wget 으로 deploy에 접속한다.

```
$ while true; do wget -O- -q 10.100.67.131:8080; sleep 1; done
```

1초마다 확인하도록 loop를 만든다.

### deploy 롤링 업데이트 하기

```
$ kubectl set image deploy http-go http-go=gasbugs/http-go:v2 --record=true
```

http-go deploy 안에 http-go 콘테이너의 버전을 v2로 업데이트한다.

```
$ kubectl get pod -w
```

### 이전 버전으로 돌아가기

```
$ kubectl rollout undo deploy http-go
```

### 특정 버전으로 돌아가기

```
$ kubectl rollout undo deploy http-go --to-revision=1
```

여기서 revision은 hitory 의 번호이다
