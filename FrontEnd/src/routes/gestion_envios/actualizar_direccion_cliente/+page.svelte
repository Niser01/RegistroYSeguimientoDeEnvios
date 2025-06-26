<script>
  import { actualizarDireccion } from '$lib/api.js';

  let correo = '';
  let direccion = '';
  let ciudad = '';
  let departamento = '';
  let codigoPostal = '';
  let mensaje = '';

  async function enviar() {
    try {
      const datos = {
        correo,
        direccion,
        ciudad,
        departamento,
        codigo_postal: codigoPostal
      };
      await actualizarDireccion(datos);
      mensaje = `✅ Dirección actualizada correctamente.`;
    } catch (error) {
      mensaje = `❌ Error: ${error.message}`;
    }
  }
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-50">
  <section class="max-w-lg w-full p-8 rounded-2xl bg-white shadow-lg border border-gray-200 space-y-6">
    <h1 class="text-3xl font-bold text-center text-gray-800">🏠 Actualizar Dirección</h1>

    <form on:submit|preventDefault={enviar} class="flex flex-col gap-4">
      <input
        type="email"
        bind:value={correo}
        placeholder="Correo electrónico del cliente"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        bind:value={direccion}
        placeholder="Nueva dirección"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        bind:value={ciudad}
        placeholder="Ciudad"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        bind:value={departamento}
        placeholder="Departamento"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        bind:value={codigoPostal}
        placeholder="Código Postal"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />

      <button
        type="submit"
        class="bg-[#1E293B] text-white py-3 rounded-lg hover:bg-[#FF7A00] transition text-lg font-semibold"
      >
        Actualizar Dirección
      </button>
    </form>

    {#if mensaje}
      <p class="text-center text-lg font-medium {mensaje.startsWith('✅') ? 'text-green-600' : 'text-red-500'}">
        {mensaje}
      </p>
    {/if}
  </section>
</div>
