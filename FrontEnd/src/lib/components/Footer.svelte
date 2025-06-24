<script lang="ts">
  const year = new Date().getFullYear();
</script>

<footer class="bg-[#1E293B] text-[#FFF] py-16">
  <div class="max-w-7xl mx-auto px-4 grid grid-cols-1 md:grid-cols-2 gap-8 items-center">
    
    <!-- Columna izquierda: Información de contacto -->
    <div class="flex flex-col items-center md:items-start space-y-4">
      <div class="text-sm text-white-400 space-y-1">
        <p class="text-lg font-semibold">Pack&Track</p>
        <p>Contacto: soporte@packtrack.app</p>
        <p>Teléfono: (+57) 000 000 0000</p>
        <p>Bogotá, Colombia</p>
      </div>
    </div>

    <!-- Columna derecha: Información de derechos -->
    <div class="flex flex-col items-center md:items-end space-y-2">
      <p class="text-lg font-semibold">Sistema de Gestión y Seguimiento de Envíos</p>
      <p class="text-sm text-white-400">&copy; {year} Pack&Track. Todos los derechos reservados.</p>
    </div>

  </div>
</footer>
