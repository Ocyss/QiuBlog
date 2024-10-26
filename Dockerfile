ARG USE_CN_PROXY=false

# 第一阶段: 编译Go程序
FROM golang:1.22.3-alpine AS go-builder
ARG USE_CN_PROXY

WORKDIR /app

COPY app .

RUN if [ "$USE_CN_PROXY" = "true" ]; then \
  go env -w GOPROXY=https://goproxy.cn,direct; \
  fi


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# 第二阶段: 准备Node.js环境
FROM node:22-slim AS node-builder
ARG USE_CN_PROXY

RUN if [ "$USE_CN_PROXY" = "true" ]; then \
  npm config set registry https://registry.npmmirror.com; \
  fi

COPY web /app
WORKDIR /app

RUN npm install -g pnpm
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --no-frozen-lockfile

RUN pnpm run build

# 第三阶段: 准备运行环境
FROM node:22-slim AS runner
ARG USE_CN_PROXY

WORKDIR /app

COPY --from=go-builder /app/main .
COPY --from=node-builder /app/dist web

COPY web/server.prod.js web/index.mjs
COPY web/package.prod.json web/package.json

WORKDIR /app/web

RUN if [ "$USE_CN_PROXY" = "true" ]; then \
  npm config set registry https://registry.npmmirror.com; \
  fi

RUN --mount=type=cache,id=npm,target=/root/.npm \
  npm install --frozen-lockfile

WORKDIR /app

EXPOSE 16879
EXPOSE 3000
EXPOSE 9001

COPY --from=ochinchina/supervisord:latest /usr/local/bin/supervisord /usr/local/bin/supervisord

COPY supervisord.conf /app/supervisord.conf

CMD ["/usr/local/bin/supervisord", "-c", "/app/supervisord.conf"]