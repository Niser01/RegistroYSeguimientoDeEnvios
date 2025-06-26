<script>
  import { crearCliente } from '$lib/api.js';

  let nombre = '';
  let correo = '';
  let telefono = '';
  let direccion = '';
  let ciudad = '';
  let departamento = '';
  let pais = 'Colombia';
  let codigoPostal = '';
  let mensaje = '';

  async function enviar() {
    try {
      const cliente = {
        nombre,
        correo,
        telefono,
        direccion,
        ciudad,
        departamento,
        pais,
        codigo_postal: codigoPostal
      };
      const res = await crearCliente(cliente);
      mensaje = `✅ Cliente creado con ID: ${res.id}`;
    } catch (error) {
      mensaje = `❌ Error: ${error.message}`;
    }
  }
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-50">
  <section class="max-w-lg w-full p-8 rounded-2xl bg-white shadow-lg border border-gray-200 space-y-6">
    <h1 class="text-3xl font-bold text-center text-gray-800">👤 Crear Cliente</h1>

    <form on:submit|preventDefault={enviar} class="flex flex-col gap-4">
      <input
        bind:value={nombre}
        placeholder="Nombre completo"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        type="email"
        bind:value={correo}
        placeholder="Correo electrónico"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        type="tel"
        bind:value={telefono}
        placeholder="Teléfono"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        bind:value={direccion}
        placeholder="Dirección"
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
        bind:value={pais}
        placeholder="País"
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
        Crear Cliente
      </button>
    </form>

    {#if mensaje}
      <p class="text-center text-lg font-medium {mensaje.startsWith('✅') ? 'text-green-600' : 'text-red-500'}">
        {mensaje}
      </p>
    {/if}
  </section>
</div>
