# service

## 포드의 문제점

-   포드는 일시적으로 생성한 컨테이너의 집합
-   때문에 포드가 지속적으로 생겨났을 때 서비스를 하기에 적합하지 않음
-   ip 주소의 지속적인 변동, 로드밸련싱을 관리해줄 또 다른 개체가 필요
-   이 문제를 해결하기 위해 서비스라는 리소스가 존재

## 서비스 생성 방법

1. kubectl의 expose가 가장 쉬운 방법
2. yaml을 통해 버전 관리 기능

## 서비스의 세션 고정하기

-   서비스가 다수의 포드로 구성하면 웹서비스의 세션이 유지되지 않음
-   이를 위해 처음 들어왔던 클라이언트 ip를 그대로 유지해주는 방법이 필요
-   sessionAffinity: ClientIP라는 옵션을 주면 해결(클라이언트가 처음 접속한 pod에만 접속함)