# AGENTS.md — okt-api

Open Korean Text(OKT) 정규화·구문 추출 HTTP API. AI 에이전트·기여자용 규칙입니다.

## 프로젝트 한 줄

Python FastAPI 서버 + Go 클라이언트(`pkg/okt`)로 한국어 normalize / phrases API를 제공합니다.

## 읽기 순서

1. [`README.md`](README.md)
2. [`app/main.py`](app/main.py) — FastAPI 엔드포인트·Bearer 미들웨어
3. [`pkg/okt/okt.go`](pkg/okt/okt.go) — Go 클라이언트
4. [`docker-compose.yml`](docker-compose.yml) / [`Dockerfile`](Dockerfile)

## 레이아웃

| 경로 | 책임 |
|------|------|
| `app/` | FastAPI 서버 (Java/OKT 런타임 필요) |
| `pkg/okt` | Go HTTP 클라이언트 (`OKT_TOKEN`) |
| `examples/` | Go 사용 예제 |
| `requirements.txt` | Python 의존성 |

## 인증·환경

| 변수 | 용도 |
|------|------|
| `OKT_TOKEN` | Bearer 토큰 (**파일 아님**, env 주입) |
| `ROOT_PATH` | FastAPI `root_path` (기본 compose: `/`) |
| `OKT_PORT` | compose 호스트 포트 (기본 `8080`) |

로컬 실행:

```sh
cp .env.example .env   # OKT_TOKEN 설정
docker compose up --build
```

Go 클라이언트도 `OKT_TOKEN` (및 생성자 baseURL)을 사용합니다.

## 코딩 규칙

- 서버는 Python, 클라이언트는 Go. 프로토콜(JSON 스키마·경로)을 양쪽에서 맞출 것.
- 토큰은 반드시 env/`docker-compose`로만; `/secret/token` 파일 경로를 되살리지 말 것.
- Java(OpenJDK)는 Dockerfile이 설치. 로컬 pip-only 실행은 JDK 필요.
- Go 모듈 디렉티브는 **1.27**.

## 하지 말 것

- `token` / `.env` 실비밀 커밋
- 사용자 요청 없이 커밋·푸시

## 커밋

사용자 요청 시에만. 메시지 스타일: `feat|fix|chore: …`
