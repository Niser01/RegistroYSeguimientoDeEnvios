<nav class="bg-[#FF7A00] text-white py-4 fixed top-0 left-0 w-full z-50 shadow">
  <div class="w-full max-w-7xl mx-auto px-6 flex justify-between items-center">

    <!-- Logo -->
    <div class="flex items-center space-x-3">
      <a href="/">
        <img src="/P&T_Blanco.svg" alt="Logo Empresa" class="h-12 w-auto" />
      </a>
    </div>

    <!-- Botón menú hamburguesa (solo en móvil) -->
    <input type="checkbox" id="menu-toggle" class="peer hidden">
    <label for="menu-toggle" class="cursor-pointer md:hidden block">
      <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
          d="M4 6h16M4 12h16M4 18h16" />
      </svg>
    </label>

    <!-- Enlaces navegación (en desktop) -->
    <div
      class="hidden peer-checked:block md:flex md:space-x-8 md:items-center absolute md:static top-full left-0 w-full md:w-auto bg-[#FF7A00] md:bg-transparent shadow md:shadow-none mt-2 md:mt-0 py-4 md:py-0 px-6 md:px-0 z-40">

      <div class="flex flex-col space-y-4 md:flex-row md:space-y-0 md:space-x-8 md:items-center w-full text-center">

        <a href="/rastreo" class="block py-3 border-t border-white md:border-none md:py-0 hover:underline">Rastrea tu pedido</a>
        <a href="/gestion_envios" class="block py-3 border-t border-white md:border-none md:py-0 hover:underline">Gestión de envíos</a>
        <a href="/sobre_nosostros" class="block py-3 border-t border-white md:border-none md:py-0 hover:underline">Sobre nosotros</a>

        <!-- Botón Contacto (solo móvil) -->
        <a href="/contacto"
          class="block md:hidden font-semibold bg-white text-[#1E293B] hover:bg-[#1E293B] hover:text-white px-5 py-3 rounded-full transition text-center">
          Contactanos
        </a>

      </div>
    </div>

    <!-- Botón Contacto (solo desktop) -->
    <div class="hidden md:block">
      <a href="/contacto"
        class="font-semibold bg-white text-[#1E293B] hover:bg-[#1E293B] hover:text-white px-5 py-2 rounded-full transition">
        Contactanos
      </a>
    </div>

  </div>
</nav>
