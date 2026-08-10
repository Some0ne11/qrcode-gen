<script setup>
import { ref } from 'vue';

const activeTab = ref('generate'); // 'generate' or 'scan'

// Generate State
const generateText = ref('');
const generateResultUrl = ref(null);
const isGenerating = ref(false);
const generateError = ref(null);

// Scan State
const scanResult = ref(null);
const isScanning = ref(false);
const scanError = ref(null);
const selectedFile = ref(null);
const selectedFilePreview = ref(null);

const handleGenerate = async () => {
  if (!generateText.value) return;
  
  isGenerating.value = true;
  generateError.value = null;
  generateResultUrl.value = null;

  try {
    const response = await fetch('/api/generate', {
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
    const response = await fetch('/api/scan', {
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
        
        <!-- Generate Tab -->
        <div v-show="activeTab === 'generate'" class="animate-in fade-in slide-in-from-bottom-2 duration-300">
          <div class="space-y-6">
            <div>
              <label for="qr-text" class="block text-sm font-medium text-slate-300 mb-2">Enter Text or URL</label>
              <input 
                id="qr-text" 
                v-model="generateText" 
                type="text" 
                placeholder="https://example.com" 
                class="w-full bg-slate-900/50 border border-slate-600 rounded-xl px-4 py-3 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-shadow duration-200"
                @keyup.enter="handleGenerate"
              />
            </div>

            <button 
              @click="handleGenerate" 
              :disabled="!generateText || isGenerating"
              class="w-full bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-3 px-4 rounded-xl transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed shadow-[0_0_15px_rgba(79,70,229,0.3)] hover:shadow-[0_0_25px_rgba(79,70,229,0.5)] transform active:scale-[0.98]"
            >
              <span v-if="isGenerating" class="flex items-center justify-center">
                <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
                Generating...
              </span>
              <span v-else>Generate QR Code</span>
            </button>

            <!-- Error -->
            <div v-if="generateError" class="p-4 bg-red-900/30 border border-red-500/50 rounded-xl text-red-200 text-sm">
              {{ generateError }}
            </div>

            <!-- Result -->
            <div v-if="generateResultUrl" class="mt-8 flex flex-col items-center p-6 bg-slate-900/80 rounded-2xl border border-slate-700/50">
              <img :src="generateResultUrl" alt="Generated QR Code" class="w-48 h-48 rounded-lg shadow-md mb-4 bg-white p-2" />
              <a 
                :href="generateResultUrl" 
                download="qrcode.png"
                class="text-indigo-400 hover:text-indigo-300 font-medium text-sm transition-colors flex items-center"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" /></svg>
                Download PNG
              </a>
            </div>
          </div>
        </div>

        <!-- Scan Tab -->
        <div v-show="activeTab === 'scan'" class="animate-in fade-in slide-in-from-bottom-2 duration-300">
          <div class="space-y-6">
            
            <div class="relative w-full h-48 border-2 border-dashed border-slate-600 rounded-2xl bg-slate-900/30 hover:bg-slate-800/50 transition-colors flex flex-col items-center justify-center overflow-hidden group cursor-pointer">
              <input 
                type="file" 
                accept="image/*" 
                @change="handleFileSelect" 
                class="absolute inset-0 w-full h-full opacity-0 cursor-pointer z-10"
              />
              
              <template v-if="!selectedFilePreview">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 text-slate-400 group-hover:text-indigo-400 transition-colors mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <p class="text-sm text-slate-300 font-medium group-hover:text-white transition-colors">Click or drag image to upload</p>
                <p class="text-xs text-slate-500 mt-1">PNG, JPG up to 10MB</p>
              </template>
              
              <template v-else>
                <img :src="selectedFilePreview" class="absolute inset-0 w-full h-full object-contain bg-slate-900/80 p-2" />
                <div class="absolute inset-0 bg-slate-900/60 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
                  <p class="text-white font-medium">Change Image</p>
                </div>
              </template>
            </div>

            <button 
              @click="handleScan" 
              :disabled="!selectedFile || isScanning"
              class="w-full bg-cyan-600 hover:bg-cyan-500 text-white font-bold py-3 px-4 rounded-xl transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed shadow-[0_0_15px_rgba(8,145,178,0.3)] hover:shadow-[0_0_25px_rgba(8,145,178,0.5)] transform active:scale-[0.98]"
            >
              <span v-if="isScanning" class="flex items-center justify-center">
                <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
                Scanning...
              </span>
              <span v-else>Scan QR Code</span>
            </button>

            <!-- Error -->
            <div v-if="scanError" class="p-4 bg-red-900/30 border border-red-500/50 rounded-xl text-red-200 text-sm">
              {{ scanError }}
            </div>

            <!-- Result -->
            <div v-if="scanResult" class="mt-6 p-5 bg-slate-900 border border-indigo-500/30 rounded-xl">
              <h3 class="text-xs uppercase tracking-wider text-slate-400 font-semibold mb-2">Scanned Result:</h3>
              <div class="relative group">
                <p class="text-white text-lg break-words pr-8">{{ scanResult }}</p>
                <button 
                  @click="navigator.clipboard.writeText(scanResult)" 
                  class="absolute top-0 right-0 p-1 text-slate-500 hover:text-indigo-400 transition-colors opacity-0 group-hover:opacity-100"
                  title="Copy to clipboard"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" /></svg>
                </button>
              </div>
            </div>

          </div>
        </div>

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
