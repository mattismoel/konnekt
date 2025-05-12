# ==============================================================================
# BUILD STAGE
# ==============================================================================
#
# We instantiate a node alpine image and copy relevant files into the builder,
# set relevant build environment variables, and clean dev packages.
# After, we build the node SvelteKit project to the output 'build' directory.
#
# Thereafter we instantiate a new node alpine image, copying the 'build'
# directory into the production stage and set runtime environment variables.
# ==============================================================================
ARG PUBLIC_SERVER_API_ORIGIN=${PUBLIC_SERVER_API_ORIGIN:-http://localhost:4000}

FROM node:24-alpine AS build
ARG PUBLIC_SERVER_API_ORIGIN

ENV PUBLIC_SERVER_API_ORIGIN=${PUBLIC_SERVER_API_ORIGIN}

RUN echo "SERVER API ORIGIN: $PUBLIC_SERVER_API_ORIGIN"

WORKDIR /app

COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build
RUN npm prune --omit=dev

# ==============================================================================
# PRODUCTION STAGE =============================================================
# ==============================================================================
FROM node:24-alpine
WORKDIR /app

COPY --from=build /app/build build
COPY --from=build /app/node_modules node_modules
COPY package.json .

ENV NODE_ENV=production
ENV PROTOCOL_HEADER=x-forwarded-proto
ENV HOST_HEADER=x-forwarded-host

EXPOSE 3000

CMD [ "node", "build" ]
