# Namespace

namespace yaml 파일 만들기

```
$ kubectl create ns office --dry-run=client -o yaml > office-ns.yaml
```

만등 namespace에 nginx 이미지 run하기

```
$ kubectl create deploy nginx --image nginx -n office
```

## default namespace 변경

~/.kube/config 파일에서

contexts.context에 namespace: office
추가
