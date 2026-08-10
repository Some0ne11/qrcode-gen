<script setup>
import { ref } from 'vue';
import GenerateTab from './components/GenerateTab.vue';
import ScanTab from './components/ScanTab.vue';

const activeTab = ref('generate'); // 'generate' or 'scan'
</script>

<template>
  <div class="min-h-screen bg-slate-900 text-white font-sans flex flex-col items-center py-12 px-4 selection:bg-indigo-500 selection:text-white">
    
    <header class="mb-12 text-center">
      <div class="inline-block mb-4 p-4 rounded-full bg-indigo-500/20 shadow-[0_0_40px_rgba(99,102,241,0.4)]">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 text-indigo-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
          <rect x="7" y="7" width="3" height="3"></rect>
          <rect x="14" y="7" width="3" height="3"></rect>
          <rect x="7" y="14" width="3" height="3"></rect>
          <rect x="14" y="14" width="3" height="3"></rect>
        </svg>
      </div>
      <h1 class="text-4xl md:text-5xl font-extrabold bg-clip-text text-transparent bg-gradient-to-r from-indigo-400 to-cyan-400">
        QR Studio
      </h1>
      <p class="mt-4 text-slate-400 text-lg">Generate or scan QR codes instantly.</p>
    </header>

    <main class="w-full max-w-xl bg-slate-800/50 backdrop-blur-xl border border-slate-700/50 rounded-3xl shadow-2xl overflow-hidden transition-all duration-300">
      
      <!-- Tabs -->
      <div class="flex border-b border-slate-700/50">
        <button 
          @click="activeTab = 'generate'"
          :class="[
            'flex-1 py-4 px-6 text-center font-semibold text-lg transition-colors duration-200 focus:outline-none',
            activeTab === 'generate' ? 'text-indigo-400 border-b-2 border-indigo-400 bg-indigo-400/5' : 'text-slate-500 hover:text-slate-300 hover:bg-slate-700/30'
          ]"
        >
          Generate
        </button>
        <button 
          @click="activeTab = 'scan'"
          :class="[
            'flex-1 py-4 px-6 text-center font-semibold text-lg transition-colors duration-200 focus:outline-none',
            activeTab === 'scan' ? 'text-indigo-400 border-b-2 border-indigo-400 bg-indigo-400/5' : 'text-slate-500 hover:text-slate-300 hover:bg-slate-700/30'
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

    <footer class="mt-16 text-slate-500 text-sm">
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
