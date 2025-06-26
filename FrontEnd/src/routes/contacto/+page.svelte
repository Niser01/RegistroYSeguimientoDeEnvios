<script>
  let nombre = '';
  let correo = '';
  let asunto = '';
  let mensaje = '';
  let respuesta = '';

  function enviar() {
    respuesta = `✅ Gracias ${nombre}, hemos recibido tu mensaje y te contactaremos pronto.`;
    nombre = correo = asunto = mensaje = '';
  }
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-50 px-6 py-12">
  <section class="max-w-3xl w-full space-y-8">
    <div class="text-center">
      <h1 class="text-4xl font-extrabold text-gray-800 mb-4">📞 Contáctanos</h1>
      <p class="text-gray-600 text-lg">
        Si tienes dudas, sugerencias o deseas conocer más sobre nuestros servicios, déjanos tu mensaje. 
        En <strong>Pack&Track</strong> estaremos encantados de atenderte.
      </p>
    </div>

    <form on:submit|preventDefault={enviar} class="flex flex-col gap-4 bg-white p-6 rounded-2xl shadow-lg border border-gray-200">
      <input
        bind:value={nombre}
        placeholder="Tu nombre"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        type="email"
        bind:value={correo}
        placeholder="Tu correo electrónico"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <input
        bind:value={asunto}
        placeholder="Asunto"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />
      <textarea
        bind:value={mensaje}
        placeholder="Escribe tu mensaje..."
        rows="4"
        required
        class="p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg resize-none"
      ></textarea>

      <button
        type="submit"
        class="bg-[#1E293B] text-white py-3 rounded-lg hover:bg-[#FF7A00] transition text-lg font-semibold"
      >
        Enviar mensaje
      </button>
    </form>

    {#if respuesta}
      <p class="text-center text-lg font-medium text-green-600">{respuesta}</p>
    {/if}
  </section>
</div>
