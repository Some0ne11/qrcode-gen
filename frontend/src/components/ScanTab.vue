<script setup>
import { ref } from 'vue';
import { UploadCloud, Loader2, Check, Copy } from '@lucide/vue';

const scanResult = ref(null);
const isScanning = ref(false);
const scanError = ref(null);
const selectedFile = ref(null);
const selectedFilePreview = ref(null);
const isCopied = ref(false);

const copyToClipboard = async () => {
  try {
    await navigator.clipboard.writeText(scanResult.value);
    isCopied.value = true;
    setTimeout(() => {
      isCopied.value = false;
    }, 2000);
  } catch (err) {
    console.error('Failed to copy: ', err);
    // Fallback if clipboard API fails
    const textArea = document.createElement("textarea");
    textArea.value = scanResult.value;
    document.body.appendChild(textArea);
    textArea.focus();
    textArea.select();
    try {
      document.execCommand('copy');
      isCopied.value = true;
      setTimeout(() => isCopied.value = false, 2000);
    } catch (err) {
      console.error('Fallback copy failed', err);
    }
    document.body.removeChild(textArea);
  }
};

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
      
      <div class="relative w-full h-48 border-2 border-dashed border-zinc-300 dark:border-zinc-700 rounded-md bg-zinc-50 dark:bg-zinc-950 hover:border-blue-500 hover:bg-zinc-100 dark:hover:bg-zinc-900 transition-colors flex flex-col items-center justify-center overflow-hidden group cursor-pointer">
        <input 
          type="file" 
          accept="image/*" 
          @change="handleFileSelect" 
          class="absolute inset-0 w-full h-full opacity-0 cursor-pointer z-10"
        />
        
        <template v-if="!selectedFilePreview">
          <UploadCloud class="h-10 w-10 text-zinc-400 dark:text-zinc-500 group-hover:text-blue-500 transition-colors mb-3" stroke-width="1.5" />
          <p class="text-sm text-zinc-700 dark:text-zinc-300 font-medium transition-colors">Click or drag image to upload</p>
          <p class="text-xs text-zinc-500 mt-1">PNG, JPG up to 10MB</p>
        </template>
        
        <template v-else>
          <img :src="selectedFilePreview" class="absolute inset-0 w-full h-full object-contain bg-white dark:bg-zinc-950 p-2" />
          <div class="absolute inset-0 bg-white/80 dark:bg-zinc-900/60 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
            <p class="text-zinc-900 dark:text-white font-medium">Change Image</p>
          </div>
        </template>
      </div>

      <button 
        @click="handleScan" 
        :disabled="!selectedFile || isScanning"
        class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md transition-colors duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
      >
        <span v-if="isScanning" class="flex items-center justify-center">
          <Loader2 class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" />
          Scanning...
        </span>
        <span v-else>Scan QR Code</span>
      </button>

      <!-- Error -->
      <div v-if="scanError" class="p-3 bg-red-500/10 border border-red-500/20 rounded-md text-red-400 text-sm">
        {{ scanError }}
      </div>

      <!-- Result -->
      <div v-if="scanResult" class="mt-6 p-4 bg-zinc-50 dark:bg-zinc-950 border border-zinc-200 dark:border-zinc-800 rounded-md transition-colors">
        <h3 class="text-xs uppercase tracking-wider text-zinc-500 font-semibold mb-2">Scanned Result:</h3>
        <div class="relative group">
          <p class="text-zinc-900 dark:text-zinc-100 text-base break-words pr-8 transition-colors">{{ scanResult }}</p>
          <button 
            @click="copyToClipboard" 
            class="absolute top-0 right-0 p-1 transition-colors opacity-0 group-hover:opacity-100 focus:opacity-100"
            :class="isCopied ? 'text-green-500' : 'text-zinc-500 hover:text-blue-500'"
            title="Copy to clipboard"
          >
            <Check v-if="isCopied" class="h-4 w-4" />
            <Copy v-else class="h-4 w-4" />
          </button>
        </div>
      </div>

    </div>
  </div>
</template>
