# Etapa de build
FROM node:18 AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

# Etapa de producción (usando nginx o http-server)
FROM node:18 AS production

WORKDIR /app

# Solo copiamos los archivos de producción (dist)
COPY --from=build /app/dist ./dist

# Instalamos un servidor estático (puede ser http-server)
RUN npm install -g http-server

# Puerto que se expone
EXPOSE 4173

# Comando para servir el contenido estático
CMD ["http-server", "dist"]
