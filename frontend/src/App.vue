<script setup>
import { ref } from 'vue';
import GenerateTab from './components/GenerateTab.vue';
import ScanTab from './components/ScanTab.vue';

const activeTab = ref('generate'); // 'generate' or 'scan'
</script>

<template>
  <div class="min-h-screen bg-zinc-950 text-zinc-100 font-sans flex flex-col items-center py-12 px-4 selection:bg-blue-600 selection:text-white">
    
    <header class="mb-12 text-center">
      <div class="inline-block mb-4 p-4 rounded-lg bg-zinc-900 border border-zinc-800">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-blue-500" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
          <rect x="7" y="7" width="3" height="3"></rect>
          <rect x="14" y="7" width="3" height="3"></rect>
          <rect x="7" y="14" width="3" height="3"></rect>
          <rect x="14" y="14" width="3" height="3"></rect>
        </svg>
      </div>
      <h1 class="text-3xl font-semibold text-zinc-100 tracking-tight">
        QR Studio
      </h1>
      <p class="mt-2 text-zinc-400 text-sm">Generate or scan QR codes instantly.</p>
    </header>

    <main class="w-full max-w-xl bg-zinc-900 border border-zinc-800 rounded-lg shadow-sm overflow-hidden transition-all duration-300">
      
      <!-- Tabs -->
      <div class="flex border-b border-zinc-800">
        <button 
          @click="activeTab = 'generate'"
          :class="[
            'flex-1 py-3 px-6 text-center font-medium text-sm transition-colors duration-200 focus:outline-none',
            activeTab === 'generate' ? 'text-blue-500 border-b-2 border-blue-500 bg-transparent' : 'text-zinc-400 hover:text-zinc-200 hover:bg-zinc-800/50'
          ]"
        >
          Generate
        </button>
        <button 
          @click="activeTab = 'scan'"
          :class="[
            'flex-1 py-3 px-6 text-center font-medium text-sm transition-colors duration-200 focus:outline-none',
            activeTab === 'scan' ? 'text-blue-500 border-b-2 border-blue-500 bg-transparent' : 'text-zinc-400 hover:text-zinc-200 hover:bg-zinc-800/50'
          ]"
        >
          Scan
        </button>
      </div>

      <!-- Tab Content -->
      <div class="p-8">
        <GenerateTab v-show="activeTab === 'generate'" />
        <ScanTab v-show="activeTab === 'scan'" />
      </div>
    </main>

    <footer class="mt-16 text-zinc-500 text-xs tracking-wide uppercase">
      Built with Vue & Go
    </footer>

  </div>
</template>

<style>
/* Basic fade in animations from bottom */
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes slideInFromBottom {
  from { transform: translateY(0.5rem); }
  to { transform: translateY(0); }
}
.animate-in {
  animation-fill-mode: both;
}
.fade-in {
  animation-name: fadeIn;
}
.slide-in-from-bottom-2 {
  animation-name: slideInFromBottom;
}
.duration-300 {
  animation-duration: 300ms;
}
</style>
