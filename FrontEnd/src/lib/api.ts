const API_URL = import.meta.env.VITE_BACKEND_API_URL;

export async function rastrearPedido(codigoTracking: string): Promise<any> {
  const response = await fetch(`${API_URL}/rastreopedido?codigoTracking=${codigoTracking}`);
  if (!response.ok) {
    throw new Error("Error consultando el pedido");
  }
  return await response.json();
}

export async function crearPedido(pedido: any): Promise<any> {
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

export async function crearCliente(cliente: any): Promise<any> {
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

export async function actualizarDireccion(datos: any): Promise<void> {
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

export async function actualizarEstadoPedido(datos: any): Promise<void> {
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
