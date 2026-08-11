<script setup>
import { ref, computed } from 'vue';
import { Loader2, Download } from '@lucide/vue';
import UrlForm from './forms/UrlForm.vue';
import TextForm from './forms/TextForm.vue';
import PhoneForm from './forms/PhoneForm.vue';
import WifiForm from './forms/WifiForm.vue';
import ContactForm from './forms/ContactForm.vue';
import ConfirmModal from './ui/ConfirmModal.vue';

const selectedType = ref('url');
const isConfirmModalOpen = ref(false);
const pendingType = ref(null);

const handleTypeChange = (event) => {
  const newType = event.target.value;
  if (generateResultUrl.value) {
    // Revert visually immediately so it doesn't change until confirmed
    event.target.value = selectedType.value;
    pendingType.value = newType;
    isConfirmModalOpen.value = true;
    return;
  }
  selectedType.value = newType;
  generateResultUrl.value = null;
  generateError.value = null;
};

const confirmTypeChange = () => {
  selectedType.value = pendingType.value;
  generateResultUrl.value = null;
  generateError.value = null;
  isConfirmModalOpen.value = false;
  pendingType.value = null;
};

const cancelTypeChange = () => {
  isConfirmModalOpen.value = false;
  pendingType.value = null;
};

const urlData = ref({ value: '' });
const textData = ref({ value: '' });
const phoneData = ref({ number: '' });
const wifiData = ref({ ssid: '', password: '', protocol: 'WPA' });
const contactData = ref({ fullname: '', org: '', address: '', phone: '', email: '', notes: '' });

const generateResultUrl = ref(null);
const isGenerating = ref(false);
const generateError = ref(null);

const formattedText = computed(() => {
  switch (selectedType.value) {
    case 'url':
      return urlData.value.value;
    case 'text':
      return textData.value.value;
    case 'phone':
      return phoneData.value.number ? `tel:${phoneData.value.number}` : '';
    case 'wifi':
      if (!wifiData.value.ssid) return '';
      let str = `WIFI:S:${wifiData.value.ssid};T:${wifiData.value.protocol};`;
      if (wifiData.value.protocol !== 'nopass' && wifiData.value.password) {
        str += `P:${wifiData.value.password};`;
      }
      str += ';';
      return str;
    case 'contact':
      if (!contactData.value.fullname) return '';
      let vcard = `BEGIN:VCARD\nVERSION:3.0\nFN:${contactData.value.fullname}\n`;
      if (contactData.value.org) vcard += `ORG:${contactData.value.org}\n`;
      if (contactData.value.phone) vcard += `TEL:${contactData.value.phone}\n`;
      if (contactData.value.email) vcard += `EMAIL:${contactData.value.email}\n`;
      if (contactData.value.address) vcard += `ADR:;;${contactData.value.address}\n`;
      if (contactData.value.notes) vcard += `NOTE:${contactData.value.notes}\n`;
      vcard += `END:VCARD`;
      return vcard;
    default:
      return '';
  }
});

const isFormValid = computed(() => {
  switch (selectedType.value) {
    case 'url':
      return !!urlData.value.value.trim();
    case 'text':
      return !!textData.value.value.trim();
    case 'phone':
      return !!phoneData.value.number.trim();
    case 'wifi':
      if (!wifiData.value.ssid.trim()) return false;
      if (wifiData.value.protocol !== 'nopass' && !wifiData.value.password.trim()) return false;
      return true;
    case 'contact':
      return !!contactData.value.fullname.trim() && 
             !!contactData.value.org.trim() &&
             !!contactData.value.phone.trim() &&
             !!contactData.value.email.trim() &&
             !!contactData.value.address.trim();
    default:
      return false;
  }
});

const handleGenerate = async () => {
  if (!isFormValid.value) return;
  
  isGenerating.value = true;
  generateError.value = null;
  generateResultUrl.value = null;

  try {
    const response = await fetch(`${import.meta.env.VITE_API_BASE_URL}/api/generate`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ 
        type: selectedType.value,
        text: formattedText.value 
      })
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
      
      <!-- Type Selector Dropdown -->
      <div>
        <label for="qr-type" class="block text-sm font-medium text-zinc-600 dark:text-zinc-300 mb-2">Select Format</label>
        <select 
          id="qr-type" 
          :value="selectedType"
          @change="handleTypeChange"
          class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200"
        >
          <option value="url">Link (URL)</option>
          <option value="text">Plain Text</option>
          <option value="contact">Contact (vCard)</option>
          <option value="phone">Phone Number</option>
          <option value="wifi">WiFi Network</option>
        </select>
      </div>

      <!-- Dynamic Forms -->
      <div class="pt-2 border-t border-zinc-200 dark:border-zinc-800 transition-colors">
        <UrlForm v-if="selectedType === 'url'" :data="urlData" @submit="handleGenerate" />
        <TextForm v-if="selectedType === 'text'" :data="textData" />
        <PhoneForm v-if="selectedType === 'phone'" :data="phoneData" @submit="handleGenerate" />
        <WifiForm v-if="selectedType === 'wifi'" :data="wifiData" />
        <ContactForm v-if="selectedType === 'contact'" :data="contactData" />
      </div>

      <!-- Action Button -->
      <button 
        @click="handleGenerate" 
        :disabled="!isFormValid || isGenerating"
        class="w-full mt-4 bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md transition-colors duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
      >
        <span v-if="isGenerating" class="flex items-center justify-center">
          <Loader2 class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" />
          Generating...
        </span>
        <span v-else>Generate QR Code</span>
      </button>

      <!-- Error -->
      <div v-if="generateError" class="p-3 bg-red-500/10 border border-red-500/20 rounded-md text-red-400 text-sm">
        {{ generateError }}
      </div>

      <!-- Result -->
      <div v-if="generateResultUrl" class="mt-6 flex flex-col items-center p-6 bg-zinc-50 dark:bg-zinc-950/50 rounded-lg border border-zinc-200 dark:border-zinc-800 transition-colors">
        <img :src="generateResultUrl" alt="Generated QR Code" class="w-48 h-48 rounded-md mb-4 bg-white p-2" />
        <a 
          :href="generateResultUrl" 
          download="qrcode.png"
          class="text-blue-500 hover:text-blue-400 font-medium text-sm transition-colors flex items-center"
        >
          <Download class="h-4 w-4 mr-1" />
          Download PNG
        </a>
      </div>
      
      <!-- Custom Confirm Modal -->
      <ConfirmModal 
        :isOpen="isConfirmModalOpen"
        title="Discard QR Code?"
        message="Your currently generated QR code will be lost. Please save it before continuing. Are you sure you want to change the format?"
        confirmText="Discard & Change"
        cancelText="Cancel"
        @confirm="confirmTypeChange"
        @cancel="cancelTypeChange"
      />

    </div>
  </div>
</template>
