// @ts-ignore
const API_URL = import.meta.env.VITE_BACKEND_API_URL;

export async function rastrearPedido(codigoTracking) {
  const response = await fetch(`${API_URL}/rastreopedido?codigoTracking=${codigoTracking}`);
  if (!response.ok) {
    throw new Error("Error consultando el pedido");
  }
  return await response.json();
}

export async function crearPedido(pedido) {
  const response = await fetch(`${API_URL}/crearpedido`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(pedido),
  });

  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText);
  }
  return await response.json();
}

export async function crearCliente(cliente) {
  const response = await fetch(`${API_URL}/crearcliente`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(cliente),
  });

  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText);
  }
  return await response.json();
}

export async function actualizarDireccion(datos) {
  const response = await fetch(`${API_URL}/actualizardireccion`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(datos),
  });

  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText);
  }
}

export async function actualizarEstadoPedido(datos) {
  const response = await fetch(`${API_URL}/actualizarpedido`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(datos),
  });

  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText);
  }
}
