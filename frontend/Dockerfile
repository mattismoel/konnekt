# ==============================================================================
# BUILD STAGE
# ==============================================================================
#
# We instantiate a node alpine image and copy relevant files into the builder,
# whereafter we build the node SvelteKit project to the output 'build' 
# directory.
# ==============================================================================

FROM oven/bun:1 AS base

# Change work directory to 'app'.
WORKDIR /app

FROM base AS install
RUN mkdir -p /temp/dev
COPY package.json bun.lockb /temp/dev/
RUN cd /temp/dev/ && bun install --frozen-lockfile

RUN mkdir -p /temp/prod
COPY package.json bun.lockb /temp/prod/
RUN cd /temp/prod && bun install --frozen-lockfile --production

FROM base AS prerelease
COPY --from=install /temp/dev/node_modules node_modules
COPY . .

ENV NODE_ENV=production
RUN bun test

ENV PUBLIC_BACKEND_URL=http://backend:8080

RUN bun run build

RUN ls build

FROM base AS release
COPY --from=install /temp/prod/node_modules node_modules
COPY --from=prerelease /app/build build
COPY --from=prerelease /app/package.json .

USER bun
EXPOSE 3000
CMD [ "bun", "./build/index.js" ]
