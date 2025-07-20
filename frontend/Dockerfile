FROM node:18
WORKDIR /app
COPY package*.json ./
# Xóa nếu có (an toàn trong môi trường Docker build)
RUN rm -rf node_modules package-lock.json && npm install
COPY . .
EXPOSE 5173
CMD ["npm", "run", "dev"]