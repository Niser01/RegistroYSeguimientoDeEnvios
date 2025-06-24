<script lang="ts">
  const year = new Date().getFullYear();
</script>

<footer class="bg-[#1E293B] text-[#FFF] py-16">
  <div class="max-w-7xl mx-auto px-4 grid grid-cols-1 md:grid-cols-2 gap-8 items-center">
    <!-- Columna izquierda: Logo + Contacto -->
    <div class="flex flex-col items-center md:items-start space-y-4">
      <div class="text-sm text-white-400 space-y-1">
        <p>Cll 18A No. 0 – 33 Este, Bloque E.</p>
        <p>Bogotá - Colombia</p>
        <p>Codigo Postal: 111711</p>
        <p>(+57) 60 1 3394949 Ext. 3686</p>
      </div>
    </div>

    <!-- Columna derecha: Info sistema -->
    <div class="flex flex-col items-center md:items-end space-y-2">
      <p class="text-lg font-semibold">Sistema de Verificación de Pagos</p>
      <p class="text-sm text-white-400">&copy; {year} Universidad de los Andes. Todos los derechos reservados.</p>
    </div>

  </div>
</footer>
