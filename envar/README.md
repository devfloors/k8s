# configmap

1234 내용을 넣은 test 파일 만들기

-n 옵션은 1234 뒤에 \n 제거

```
$ echo -n 1234 > test
```

map-name 이름의 configmap 만들기

```
$ kubectl create configmap map-name --from-file=test
```

더 많은 파일 만들기

```
$ kubectl create configmap map-name --from-file=test --from-file=test2
```
