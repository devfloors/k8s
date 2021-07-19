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
