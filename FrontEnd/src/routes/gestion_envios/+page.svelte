<script>
  const opciones = [
    {
      titulo: 'Crear Cliente',
      descripcion: 'Registra un nuevo cliente en el sistema.',
      enlace: '/gestion_envios/crear_cliente',
      icono: '👤'
    },
    {
      titulo: 'Crear Pedido',
      descripcion: 'Genera un nuevo pedido para un cliente existente.',
      enlace: '/gestion_envios/crear_pedido',
      icono: '📦'
    },
    {
      titulo: 'Actualizar Dirección',
      descripcion: 'Modifica la dirección de entrega de un pedido.',
      enlace: '/gestion_envios/actualizar_direccion_cliente',
      icono: '🏠'
    },
    {
      titulo: 'Actualizar Estado',
      descripcion: 'Cambia el estado actual de un pedido.',
      enlace: '/gestion_envios/actualizar_estado_pedido',
      icono: '🔄'
    }
  ];
</script>

<div class="min-h-[calc(100vh-19.5rem)] flex items-center justify-center">
  <div class="max-w-7xl mx-auto p-6">
    <h1 class="text-4xl font-bold mb-8 text-center">Gestión de Envíos</h1>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      {#each opciones as opcion}
        <a href={opcion.enlace} class="block p-6 bg-white rounded-2xl shadow-lg hover:shadow-xl transition duration-300 border border-gray-200 hover:border-blue-400">
          <div class="text-5xl mb-4 text-center">{opcion.icono}</div>
          <h2 class="text-xl font-semibold mb-2 text-center">{opcion.titulo}</h2>
          <p class="text-gray-600 text-sm text-center">{opcion.descripcion}</p>
        </a>
      {/each}
    </div>
  </div>
</div>
