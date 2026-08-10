<script setup>
import { ref } from 'vue';

const generateText = ref('');
const generateResultUrl = ref(null);
const isGenerating = ref(false);
const generateError = ref(null);

const handleGenerate = async () => {
  if (!generateText.value) return;
  
  isGenerating.value = true;
  generateError.value = null;
  generateResultUrl.value = null;

  try {
    const response = await fetch(`${import.meta.env.VITE_API_BASE_URL}/api/generate`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ text: generateText.value })
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || 'Failed to generate QR Code');
    }

    const blob = await response.blob();
    generateResultUrl.value = URL.createObjectURL(blob);
  } catch (error) {
    generateError.value = error.message;
  } finally {
    isGenerating.value = false;
  }
};
</script>

<template>
  <div class="animate-in fade-in slide-in-from-bottom-2 duration-300">
    <div class="space-y-6">
      <div>
        <label for="qr-text" class="block text-sm font-medium text-zinc-300 mb-2">Enter Text or URL</label>
        <input 
          id="qr-text" 
          v-model="generateText" 
          type="text" 
          placeholder="https://example.com" 
          class="w-full bg-zinc-950 border border-zinc-700 rounded-md px-4 py-2 text-zinc-100 placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200"
          @keyup.enter="handleGenerate"
        />
      </div>

      <button 
        @click="handleGenerate" 
        :disabled="!generateText || isGenerating"
        class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md transition-colors duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
      >
        <span v-if="isGenerating" class="flex items-center justify-center">
          <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
          Generating...
        </span>
        <span v-else>Generate QR Code</span>
      </button>

      <!-- Error -->
      <!-- Error -->
      <div v-if="generateError" class="p-3 bg-red-500/10 border border-red-500/20 rounded-md text-red-400 text-sm">
        {{ generateError }}
      </div>

      <!-- Result -->
      <div v-if="generateResultUrl" class="mt-6 flex flex-col items-center p-6 bg-zinc-950/50 rounded-lg border border-zinc-800">
        <img :src="generateResultUrl" alt="Generated QR Code" class="w-48 h-48 rounded-md mb-4 bg-white p-2" />
        <a 
          :href="generateResultUrl" 
          download="qrcode.png"
          class="text-blue-500 hover:text-blue-400 font-medium text-sm transition-colors flex items-center"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" /></svg>
          Download PNG
        </a>
      </div>
    </div>
  </div>
</template>
