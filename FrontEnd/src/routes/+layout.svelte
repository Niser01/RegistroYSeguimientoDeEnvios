<script lang="ts">
  import '../app.css';
  import Navbar from '../lib/components/Navbar.svelte';
  import Footer from '../lib/components/Footer.svelte';
</script>

<div class="min-h-screen flex flex-col">
  <Navbar />

  <main class="flex-1 pt-20">
    <slot />
  </main>

  <Footer />
</div>
