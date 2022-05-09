FROM node:16-alpine AS deps

COPY package.json package.json
COPY yarn.lock yarn.lock

RUN yarn install

FROM node:16-alpine AS builder

COPY --from=deps node_modules node_modules
COPY . .
RUN yarn tsc

FROM node:16-alpine AS runner

COPY --from=builder build build
COPY --from=deps node_modules node_modules
COPY package.json package.json
COPY yarn.lock yarn.lock
COPY blacklist.json blacklist.json

RUN yarn install --production

EXPOSE 3000

ENTRYPOINT node build/index.js
