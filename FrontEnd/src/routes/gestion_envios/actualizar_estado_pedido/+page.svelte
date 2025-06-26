<script>
  import { actualizarEstadoPedido } from '$lib/api.js';

  let codigoTracking = '';
  let estado = 'Pendiente';
  let descripcionEnvio = '';
  let mensaje = '';

  async function enviar() {
    try {
      const datos = {
        codigo_tracking: codigoTracking,
        estado,
        descripcion_envio: descripcionEnvio
      };
      await actualizarEstadoPedido(datos);
      mensaje = `✅ Estado de pedido actualizado correctamente.`;
    } catch (error) {
      mensaje = `❌ Error: ${error.message}`;
    }
  }
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-50">
  <section class="max-w-lg w-full p-8 rounded-2xl bg-white shadow-lg border border-gray-200 space-y-6">
    <h1 class="text-3xl font-bold text-center text-gray-800">🚚 Actualizar Estado de Pedido</h1>

    <form on:submit|preventDefault={enviar} class="flex flex-col gap-4">
      <input
        bind:value={codigoTracking}
        placeholder="Código de tracking"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />

      <select
        bind:value={estado}
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      >
        <option>Pendiente</option>
        <option>En Camino</option>
        <option>Entregado</option>
        <option>Cancelado</option>
      </select>

      <textarea
        bind:value={descripcionEnvio}
        placeholder="Descripción del estado de envío"
        rows="3"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg resize-none"
      ></textarea>

      <button
        type="submit"
        class="bg-[#1E293B] text-white py-3 rounded-lg hover:bg-[#FF7A00] transition text-lg font-semibold"
      >
        Actualizar Estado
      </button>
    </form>

    {#if mensaje}
      <p class="text-center text-lg font-medium {mensaje.startsWith('✅') ? 'text-green-600' : 'text-red-500'}">
        {mensaje}
      </p>
    {/if}
  </section>
</div>
