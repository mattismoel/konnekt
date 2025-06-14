FROM oven/bun:1 AS base
WORKDIR /usr/src/app

FROM base AS install
COPY package.json bun.lockb ./

RUN bun install --frozen-lockfile

FROM base AS prerelease
WORKDIR /app
COPY --from=install /usr/src/app/node_modules ./node_modules
COPY --from=install /usr/src/app/package.json .
COPY . .

ENV NODE_ENV=production
RUN bun run build

FROM caddy:2-alpine AS release
WORKDIR /app

COPY --from=prerelease /app/dist/ /usr/share/caddy/
COPY Caddyfile /etc/caddy/Caddyfile
