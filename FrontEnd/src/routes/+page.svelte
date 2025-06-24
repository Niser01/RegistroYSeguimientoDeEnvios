<script lang="ts">
  const opciones = [
    {
        titulo: 'Rastrea tu pedido',
        descripcion: 'Consulta en tiempo real el estado y ubicación de tu envío con tu código de seguimiento.',
        icono: '/tracking.svg',
        href: '/rastreo',
        boton: 'Rastrear'
    },
    {
        titulo: 'Gestión de envíos',
        descripcion: 'Actualiza el estado de los pedidos en tránsito y confirma entregas desde este panel.',
        icono: '/gestion.svg',
        href: '/gestion_envios',
        boton: 'Actualizar'
    }
  ];
</script>

<div class="min-h-[calc(100vh-27rem)] flex items-center justify-center px-6 my-16">
  <div class="w-full max-w-7xl mx-auto grid grid-cols-1 md:grid-cols-2 gap-8">
    {#each opciones as opcion}
      <div
        class="group relative bg-white text-black p-8 rounded-[20px] border border-[#F0F0F0] flex flex-col items-center justify-center hover:bg-[#f7f7f7] cursor-pointer overflow-hidden h-80 shadow-md"
      >
        <!-- Contenido visible inicialmente -->
        <div
          class="flex flex-col items-center text-center space-y-6 group-hover:hidden"
        >
          <h2 class="text-2xl font-bold">{opcion.titulo}</h2>
          <img src={opcion.icono} alt={opcion.titulo} class="h-30 w-30" />
        </div>

        <!-- Contenido que aparece en hover -->
        <div
          class="absolute inset-0 flex flex-col items-center justify-center text-center space-y-6 hidden group-hover:flex px-6"
        >
          <h2 class="text-2xl font-bold">{opcion.titulo}</h2>
          <p>{opcion.descripcion}</p>
          <a
            href={opcion.href}
            class="bg-[#1E293B] text-white font-semibold px-6 py-2 rounded-full drop-shadow-md hover:bg-[#FF7A00] hover:text-white"
          >
            {opcion.boton}
          </a>
        </div>
      </div>
    {/each}
  </div>
</div>
