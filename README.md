# 간단한 Go([ql](https://github.com/cznic/ql)), Vue(+Electron) Rest-Client
CRUD 연습용  
기존 [rest-client](https://github.com/practice-golang/rest-client)에 일렉트론을 붙여 봄  
[rest-ql-crud](https://github.com/practice-golang/rest-ql-crud)에서 컴파일한 실행파일 필요함  

[Demo](https://www.dropbox.com/s/qxi5hwji6o1h76b/ElectronVue-0.0.1-win.zip?dl=1)  

## Build, Run
```powershell
# 실행 전에 모듈 먼저 설치
npm install

# vue.js 쪽 소스만 브라우저로 띄우기
npm run dev

# vue.js 쪽 소스만 빌드
npm run build:vue

# 아래는 일렉트론 빌드
# 빌드 전에 rest-ql-crud 실행파일을 dist 폴더에 삽입

# vue.js 빌드 + 일렉트론 바로 실행
npm run start

# 윈도우 amd64 빌드
npm run build:win64

# 리눅스 빌드 - 아직 안 됨
npm run build:linux

# 매킨토시 빌드 - 아직 안 됨
npm run build:osx
```
