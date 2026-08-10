<script setup>
import { ref } from 'vue';

const scanResult = ref(null);
const isScanning = ref(false);
const scanError = ref(null);
const selectedFile = ref(null);
const selectedFilePreview = ref(null);

const handleFileSelect = (event) => {
  const file = event.target.files[0];
  if (file) {
    selectedFile.value = file;
    selectedFilePreview.value = URL.createObjectURL(file);
    scanResult.value = null;
    scanError.value = null;
  }
};

const handleScan = async () => {
  if (!selectedFile.value) return;
  
  isScanning.value = true;
  scanError.value = null;
  scanResult.value = null;

  const formData = new FormData();
  formData.append('image', selectedFile.value);

  try {
    const response = await fetch(`${import.meta.env.VITE_API_BASE_URL}/api/scan`, {
      method: 'POST',
      body: formData
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || 'Failed to scan QR Code');
    }

    const data = await response.json();
    scanResult.value = data.result;
  } catch (error) {
    scanError.value = error.message;
  } finally {
    isScanning.value = false;
  }
};
</script>

<template>
  <div class="animate-in fade-in slide-in-from-bottom-2 duration-300">
    <div class="space-y-6">
      
      <div class="relative w-full h-48 border-2 border-dashed border-zinc-700 rounded-md bg-zinc-950 hover:border-blue-500 hover:bg-zinc-900 transition-colors flex flex-col items-center justify-center overflow-hidden group cursor-pointer">
        <input 
          type="file" 
          accept="image/*" 
          @change="handleFileSelect" 
          class="absolute inset-0 w-full h-full opacity-0 cursor-pointer z-10"
        />
        
        <template v-if="!selectedFilePreview">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 text-zinc-500 group-hover:text-blue-500 transition-colors mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          <p class="text-sm text-zinc-300 font-medium transition-colors">Click or drag image to upload</p>
          <p class="text-xs text-zinc-500 mt-1">PNG, JPG up to 10MB</p>
        </template>
        
        <template v-else>
          <img :src="selectedFilePreview" class="absolute inset-0 w-full h-full object-contain bg-zinc-950 p-2" />
          <div class="absolute inset-0 bg-zinc-900/60 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
            <p class="text-white font-medium">Change Image</p>
          </div>
        </template>
      </div>

      <button 
        @click="handleScan" 
        :disabled="!selectedFile || isScanning"
        class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md transition-colors duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
      >
        <span v-if="isScanning" class="flex items-center justify-center">
          <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
          Scanning...
        </span>
        <span v-else>Scan QR Code</span>
      </button>

      <!-- Error -->
      <div v-if="scanError" class="p-3 bg-red-500/10 border border-red-500/20 rounded-md text-red-400 text-sm">
        {{ scanError }}
      </div>

      <!-- Result -->
      <div v-if="scanResult" class="mt-6 p-4 bg-zinc-950 border border-zinc-800 rounded-md">
        <h3 class="text-xs uppercase tracking-wider text-zinc-500 font-semibold mb-2">Scanned Result:</h3>
        <div class="relative group">
          <p class="text-zinc-100 text-base break-words pr-8">{{ scanResult }}</p>
          <button 
            @click="navigator.clipboard.writeText(scanResult)" 
            class="absolute top-0 right-0 p-1 text-zinc-500 hover:text-blue-500 transition-colors opacity-0 group-hover:opacity-100 focus:opacity-100"
            title="Copy to clipboard"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" /></svg>
          </button>
        </div>
      </div>

    </div>
  </div>
</template>
