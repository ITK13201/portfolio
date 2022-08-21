ARG node_version=18-buster

# Development environment
FROM node:${node_version} AS development

WORKDIR /frontend

RUN apt-get update

COPY ./frontend/package.json /frontend
COPY ./frontend/yarn.lock /frontend
RUN yarn install --frozen-lockfile

USER node

COPY --chown=node:node ./frontend /frontend

CMD ["yarn", "dev"]


# Builder
FROM node:${node_version} AS builder

WORKDIR /frontend

RUN apt-get update

COPY ./frontend /frontend
RUN yarn build && yarn install --production --ignore-scripts --prefer-offline


# Production environment
FROM node:${node_version} AS production

WORKDIR /frontend

RUN apt-get update

COPY --from=builder /frontend/package.json ./package.json
COPY --from=builder /frontend/public ./public
COPY --from=builder /frontend/node_modules ./node_modules
COPY --from=builder --chown=node:node /frontend/.next ./.next
COPY --from=builder /frontend/next.config.js ./next.config.js

USER node

CMD ["yarn", "start"]
