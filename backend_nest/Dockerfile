# ---------- build stage ----------
FROM node:20-alpine AS builder

WORKDIR /app

# 1️⃣ copy package.json
COPY package*.json ./

# 2️⃣ copy prisma schema SEBELUM npm ci
COPY prisma ./prisma

# 3️⃣ install deps (postinstall aman)
RUN npm ci

# 4️⃣ copy source
COPY . .

# 5️⃣ build
RUN npm run build

# ---------- runtime stage ----------
FROM node:20-alpine

WORKDIR /app
ENV NODE_ENV=production

# copy only needed files
COPY package*.json ./
COPY prisma ./prisma

# install prod deps TANPA postinstall
RUN npm ci --omit=dev --ignore-scripts

COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules/.prisma ./node_modules/.prisma

EXPOSE 3000
CMD ["node", "dist/main.js"]
