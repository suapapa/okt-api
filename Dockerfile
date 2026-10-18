FROM python:3.12-slim-bookworm

# Java 설치 추가 (bookworm: openjdk-17 제공)
RUN apt-get update && \
    apt-get install -y openjdk-17-jre-headless && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /code

COPY ./requirements.txt ./requirements.txt
RUN pip install --no-cache-dir --upgrade -r ./requirements.txt
COPY ./app ./app

ENV ROOT_PATH="/"
ENV JAVA_HOME="/usr/lib/jvm/java-17-openjdk-$(dpkg --print-architecture)"
ENV PATH="$JAVA_HOME/bin:$PATH"
# OKT_APIKEY is injected at runtime via environment / docker-compose

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "80"]
