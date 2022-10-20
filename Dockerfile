FROM denoland/deno:alpine

EXPOSE 8000

WORKDIR /app

ADD . .

RUN deno cache src/main.ts

CMD ["run", "--allow-net", "--allow-read", "--allow-write", "src/main.ts"]
