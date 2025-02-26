FROM node:20 AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install --frozen-lockfile

COPY . .

RUN npm run build --configuration=production

FROM nginx:alpine

COPY --from=builder /app/dist/cs-case-tracker-frontend/browser /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]