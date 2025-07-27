FROM node:18
WORKDIR /app
# Chỉ copy file cần thiết trước để cache hiệu quả
COPY package.json package-lock.json ./
# Cài đặt các package một cách sạch sẽ
RUN rm -rf node_modules \
 && npm ci --ignore-scripts
# Copy toàn bộ mã nguồn sau (để không làm mất cache npm install nếu code thay đổi)
COPY . .
EXPOSE 5173
CMD ["npm", "run", "dev"]