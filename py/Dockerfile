FROM python:3-slim

# Java 설치 추가
RUN apt-get update && \
    apt-get install -y openjdk-17-jre-headless && \
    apt-get clean

WORKDIR /code

COPY ./requirements.txt ./requirements.txt
RUN pip install --no-cache-dir --upgrade -r ./requirements.txt
COPY ./app ./app

ENV ROOT_PATH="/"
ENV JAVA_HOME="/usr/lib/jvm/java-17-openjdk-amd64"
ENV PATH="$JAVA_HOME/bin:$PATH"

# CMD ["sh", "-c", "uvicorn app.main:app --host 0.0.0.0 --port 80 --root-path $ROOT_PATH"]
CMD ["sh", "-c", "uvicorn app.main:app --host 0.0.0.0 --port 80"]