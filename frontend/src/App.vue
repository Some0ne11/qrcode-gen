<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue';
import { QrCode, Sun, Moon } from '@lucide/vue';
import GenerateTab from './components/GenerateTab.vue';
import ScanTab from './components/ScanTab.vue';

const isDarkMode = ref(true);
const activeTab = ref('generate'); // 'generate' or 'scan'

const mouseX = ref(0);
const mouseY = ref(0);

const handleMouseMove = (e) => {
  mouseX.value = e.clientX;
  mouseY.value = e.clientY;
};

onMounted(() => {
  window.addEventListener('mousemove', handleMouseMove);
  
  // Initialize dark mode from localStorage or default to true
  const savedTheme = localStorage.getItem('theme');
  if (savedTheme) {
    isDarkMode.value = savedTheme === 'dark';
  } else {
    isDarkMode.value = true;
  }
  updateTheme();
});

onUnmounted(() => {
  window.removeEventListener('mousemove', handleMouseMove);
});

watch(isDarkMode, () => {
  updateTheme();
  localStorage.setItem('theme', isDarkMode.value ? 'dark' : 'light');
});

const updateTheme = () => {
  if (isDarkMode.value) {
    document.documentElement.classList.add('dark');
  } else {
    document.documentElement.classList.remove('dark');
  }
};
const toggleTheme = () => {
  isDarkMode.value = !isDarkMode.value;
};
</script>

<template>
  <div class="relative min-h-screen bg-gray-50 dark:bg-zinc-950 text-zinc-900 dark:text-zinc-100 font-sans flex flex-col items-center py-12 px-4 selection:bg-blue-600 selection:text-white overflow-hidden transition-colors duration-300">
    
    <!-- Theme Toggle Button -->
    <button 
      @click="toggleTheme" 
      class="absolute top-6 right-6 p-2 rounded-full bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 text-zinc-500 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-zinc-100 hover:bg-zinc-100 dark:hover:bg-zinc-800 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-colors z-20"
      aria-label="Toggle dark mode"
    >
      <Sun v-if="!isDarkMode" class="h-5 w-5" />
      <Moon v-else class="h-5 w-5" />
    </button>

    <!-- Dynamic Spotlight Effect -->
    <div 
      class="pointer-events-none fixed inset-0 z-0 transition-opacity duration-300"
      :style="{
        background: `radial-gradient(600px circle at ${mouseX}px ${mouseY}px, rgba(59, 130, 246, 0.12), transparent 40%)`
      }"
    ></div>

    <div class="relative z-10 w-full flex flex-col items-center">
      <header class="mb-12 text-center">
      <div class="inline-block mb-4 p-4 rounded-lg bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 transition-colors">
        <QrCode class="h-8 w-8 text-blue-500" />
      </div>
      <h1 class="text-3xl font-semibold text-zinc-900 dark:text-zinc-100 tracking-tight transition-colors">
        QR Studio
      </h1>
      <p class="mt-2 text-zinc-500 dark:text-zinc-400 text-sm transition-colors">Generate or scan QR codes instantly.</p>
    </header>

    <main class="w-full max-w-xl bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-lg shadow-sm overflow-hidden transition-all duration-300">
      
      <!-- Tabs -->
      <div class="flex border-b border-zinc-200 dark:border-zinc-800 transition-colors">
        <button 
          @click="activeTab = 'generate'"
          :class="[
            'flex-1 py-3 px-6 text-center font-medium text-sm transition-colors duration-200 focus:outline-none',
            activeTab === 'generate' ? 'text-blue-600 dark:text-blue-500 border-b-2 border-blue-600 dark:border-blue-500 bg-transparent' : 'text-zinc-500 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-zinc-200 hover:bg-zinc-50 dark:hover:bg-zinc-800/50'
          ]"
        >
          Generate
        </button>
        <button 
          @click="activeTab = 'scan'"
          :class="[
            'flex-1 py-3 px-6 text-center font-medium text-sm transition-colors duration-200 focus:outline-none',
            activeTab === 'scan' ? 'text-blue-600 dark:text-blue-500 border-b-2 border-blue-600 dark:border-blue-500 bg-transparent' : 'text-zinc-500 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-zinc-200 hover:bg-zinc-50 dark:hover:bg-zinc-800/50'
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
