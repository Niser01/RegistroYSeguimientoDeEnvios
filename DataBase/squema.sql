CREATE TABLE cliente (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    correo VARCHAR(255) UNIQUE NOT NULL,
    telefono VARCHAR(20),
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE direccion (
    id SERIAL PRIMARY KEY,
    cliente_id INTEGER REFERENCES cliente(id) ON DELETE CASCADE,
    direccion TEXT NOT NULL,
    ciudad VARCHAR(100) NOT NULL,
    departamento VARCHAR(100),
    pais VARCHAR(100) DEFAULT 'Colombia',
    codigo_postal VARCHAR(10),
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE envio (
    id SERIAL PRIMARY KEY,
    codigo_tracking VARCHAR(50) UNIQUE NOT NULL,
    cliente_origen_id INTEGER REFERENCES cliente(id) ON DELETE SET NULL,
    cliente_destino_id INTEGER REFERENCES cliente(id) ON DELETE SET NULL,
    descripcion TEXT,
    peso DECIMAL(10,2),
    valor_envio DECIMAL(12,2),
    fecha_envio TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_entrega_estimada TIMESTAMP,
    entregado BOOLEAN DEFAULT FALSE
);

CREATE TABLE estado_envio (
    id SERIAL PRIMARY KEY,
    envio_id INTEGER REFERENCES envio(id) ON DELETE CASCADE,
    estado VARCHAR(100) NOT NULL, 
    descripcion_envio TEXT,
    fecha_estado TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);