# 오토 스케일링

## 포드 스케일링의 두가지 방법

-   HPA: 포드 자체를 복제하여 처리할 수 있는 포드의 개수를 늘리는 방법
-   VPA: 리소스를 증가시켜 포드의 사용 가능한 리소르를 늘리는 방법
-   CA: 번외로 클러스터 자체를 늘리는 방법(노드 추가)

### HPA(Horizontal Pod Autoscaler)

-   쿠버네티스에는 기본 오토스케일링 기능 내장
-   CPU 사용률을 모니터링하여 실행된 포드의 개수를 늘리거나 줄임

## 설정 방법

```
$ kubectl autoscale deployment my-app --max 6 --min 4 --cpu-percent 50
```
