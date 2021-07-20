# 모니터링 서비스 플랫폼

-   Heapster
-   Metrics Service
-   cAdvisor
-   프로메테우스
-   EFK

# 리소스 모니터링 도구

-   쿠버네티스 클러스터 내의 애플리케이션 성능을 검사
-   쿠버네티스는 각 레벨에서 애플리케이션의 리소스 사용량에 대한 상세 정보를 제공
-   애플리케이션의 성능을 평가하고 병목 현상을 제거하여 전체 성능 향상을 도모

-   리소스 메트릭 파이프라인

    -   kubectl top등의 유틸리티 관련된 메트릭들로 제한된 집합을 제공
    -   단기 메모리 저장소인 metrics-server에 의해 수집
    -   metrics-server는 모든 노드를 발견하고 kubelet에 CPU와 메모리를 질의
    -   kubelet은 kubelet에 통합된 cAdvisor를 통해 레거시 도커와 통합 후 metric-server 리소스 메트릭으로 노출
    -   /metrics/resource/v1beta1 API를 사용

-   완전한 메트릭 파이프라인
    -   보다 풍부한 메트릭에 접근
    -   클러스터의 현재 상태를 기반으로 자동으로 스케일링 하거나 클러스터를 조정
    -   모니터링 파이프라인은 kubelet에서 메트릭을 가져옴
    -   CNCF 프로젝트인 프로메테우스가 대표적
    -   custom.metrics.k8s.io, external.metrics.k8s.io API를 사용

metric-server를 직접 설치해야함
https://blog.naver.com/isc0304/221860790762

```
$ git clone https://github.com/kubernetes-sigs/metrics-server
$ kubectl create -f metrics-server/deploy/kubernetes
```

```
$ kubectl edit deployments.apps -n kube-system metrics-server
```

# 애플리케이션 로그 관리

## kubernetes 애플리케이션 로그 확인

-   로그는 컨테이너 단위로 로그 확인 가능
-   싱글 컨테이너 포드의 경우 포드까지만 지정하여 로그 확인
-   멀티 컨테이너의 경우 포드 뒤에 컨테이너 이름까지 전달하여 로그 확인
-   $ kubectl logs <pod name> <옵션: container name>

## kubeapi가 정산 동작하지 않는 경우

-   쿠버네티스에서 돌아가는 리소스들은 모두 docker 사용
-   따라서 docker의 로깅 기능을 사용
-   docker ps -a 사용하여 조회 가능
-   docker logs <container id>를 사용하여 로그 확인 가능

# 큐브 대시보드

-   kubernetes 클러스터 용 범용 웹 기반 UI
-   사용자는 클러스터에서 실행중인 응용 프로그램을 관리하고 문제를 해결, 클러스터 자체를 관리

## 설치법

-   https://github.com/kubernetes/dashboard
-   다음 명령을 사용하여 git에 올라온 yaml 파일을 바로 적용

```
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/reconneded.yaml
```

```
잘 올라왔는지 확인
$ kubectl get all -n kubernetes-dashboard
```

```
접속
$ kubectl proxy

$ http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
```

```
서비스어카운트 토큰이 있어야 대시보드가 뜬다
$ kubectl get sa -n kubernetes-dashboard

여기서 kubernetes-dashboard 에 대한 정보가 있어야 한다

$ kubectl get secret -n kubernetes-dashboard
에서 kuberntes-dashboard-token-#####가 가지고 있는 토큰을 볼 수 있다

근데 여기서는 base64로 인코딩 된 상태로 보여주기때문에 자세히 볼려면
$ kubectl describe secret -n kubernetes-dashboard kubernetes-dashboard-token-mn55h
로 볼 수 있다

이 토큰을 대시보드 로그인 부분에 복붙하면 접속할 수 있다
```

```
근데 쿠버네티스를 컨트롤할 수 있는 권한이 서비스 어카운트에 없어서 리소스를 볼 수 없다
RBAC를 만들어야 한다

$ kubectl create -f kube-dashboard-role-binding.yaml
```
