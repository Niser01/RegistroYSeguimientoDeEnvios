<script>
  import { crearPedido } from '$lib/api.js';

  let correoClienteOrigen = '';
  let correoClienteDestino = '';
  let descripcion = '';
  let peso = '';
  let valorEnvio = '';
  let moneda = 'COP';
  let fechaEntregaEstimada = '';
  let mensaje = '';

  async function enviar() {
    try {
      const pedido = {
        correo_cliente_origen: correoClienteOrigen,
        correo_cliente_destino: correoClienteDestino,
        descripcion,
        peso: parseFloat(peso),
        valor_envio: parseFloat(valorEnvio),
        moneda,
        fecha_entrega_estimada: fechaEntregaEstimada
      };
      const res = await crearPedido(pedido);
      mensaje = `✅ Pedido creado con ID: ${res.id}`;
    } catch (error) {
      mensaje = `❌ Error: ${error.message}`;
    }
  }
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-50">
  <section class="max-w-lg w-full p-8 rounded-2xl bg-white shadow-lg border border-gray-200 space-y-6">
    <h1 class="text-3xl font-bold text-center text-gray-800">📦 Crear Pedido</h1>

    <form on:submit|preventDefault={enviar} class="flex flex-col gap-4">
      <input
        type="email"
        bind:value={correoClienteOrigen}
        placeholder="Correo Cliente Origen"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        type="email"
        bind:value={correoClienteDestino}
        placeholder="Correo Cliente Destino"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        bind:value={descripcion}
        placeholder="Descripción del producto"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        type="number"
        step="0.01"
        bind:value={peso}
        placeholder="Peso (kg)"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        type="number"
        step="0.01"
        bind:value={valorEnvio}
        placeholder="Valor del envío"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <select
        bind:value={moneda}
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      >
        <option value="COP">COP</option>
        <option value="USD">USD</option>
        <option value="EUR">EUR</option>
      </select>
      <input
        type="datetime-local"
        bind:value={fechaEntregaEstimada}
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />

      <button
        type="submit"
        class="bg-[#1E293B] text-white py-3 rounded-lg hover:bg-[#FF7A00] transition text-lg font-semibold"
      >
        Crear Pedido
      </button>
    </form>

    {#if mensaje}
      <p class="text-center text-lg font-medium {mensaje.startsWith('✅') ? 'text-green-600' : 'text-red-500'}">
        {mensaje}
      </p>
    {/if}
  </section>
</div>
