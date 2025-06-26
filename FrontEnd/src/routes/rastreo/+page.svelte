<script lang="ts">
  import { rastrearPedido } from '$lib/api.js';
  import { onMount } from 'svelte';

  let codigoTracking = '';
  let resultado: { estado: string; descripcion: string; direccion_destino: string } | null = null;
  let error = '';
  let recaptchaToken = '';
  const siteKey = import.meta.env.VITE_RECAPTCHA_SITE_KEY;

  const buscarPedido = async () => {
    error = '';
    resultado = null;

    if (!codigoTracking.trim()) {
      error = 'Por favor ingresa un código de rastreo.';
      return;
    }

    if (!recaptchaToken) {
      error = 'Por favor completa el reCAPTCHA antes de continuar.';
      return;
    }

    try {
      const data = await rastrearPedido(codigoTracking);
      resultado = data;
    } catch (err) {
      error = `No se encontró información para el código ${codigoTracking}.`;
    } finally {
      window.grecaptcha.reset(); // Resetea el captcha después de cada consulta
      recaptchaToken = '';
    }
  };

  onMount(() => {
    // Cargar el captcha
    window.grecaptcha.ready(() => {
      window.grecaptcha.render('recaptcha', {
        sitekey: siteKey,
        callback: (token: string) => {
          recaptchaToken = token;
        },
      });
    });
  });
</script>

<div class="min-h-[calc(100vh-19.5rem)] flex items-center justify-center bg-gray-50">
  <section class="max-w-3xl w-full mx-auto py-12 px-4">
    <h1 class="text-3xl font-bold mb-8 text-center">Rastrear Pedido</h1>

    <div class="flex flex-col items-center gap-4 mb-6">
      <input
        type="text"
        bind:value={codigoTracking}
        placeholder="Ingresa tu código de rastreo"
        class="w-full max-w-md p-4 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-[#FF7A00] text-lg"
      />

      <!-- reCAPTCHA widget -->
      <div id="recaptcha" class="mx-auto"></div>

      <button
        on:click={buscarPedido}
        class="bg-[#1E293B] text-white px-6 py-3 rounded-lg hover:bg-[#FF7A00] transition w-full max-w-md text-lg font-semibold"
      >
        Buscar
      </button>
    </div>

    {#if error}
      <p class="text-red-500 mb-4 text-center">{error}</p>
    {/if}

    {#if resultado}
      <div class="p-6 rounded-xl bg-[#F9FAFB] border border-gray-300 shadow-md space-y-3 w-full max-w-md mx-auto">
        <div class="flex items-center space-x-2">
          <span class="text-[#FF7A00] font-semibold">📦 Estado:</span>
          <span class="text-gray-800">{resultado.estado}</span>
        </div>
        <div class="flex items-center space-x-2">
          <span class="text-[#FF7A00] font-semibold whitespace-nowrap">📝 Descripción:</span>
          <span class="text-gray-800 break-words">{resultado.descripcion}</span>
        </div>
        <div class="flex items-center space-x-2">
          <span class="text-[#FF7A00] font-semibold">📍 Dirección de destino:</span>
          <span class="text-gray-800">{resultado.direccion_destino}</span>
        </div>
      </div>
    {/if}
  </section>
</div>
