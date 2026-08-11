<script setup>
import { ref, onMounted, watch, computed } from 'vue';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';

// Fix Leaflet's default icon paths issue with module bundlers
delete L.Icon.Default.prototype._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.9.4/images/marker-icon-2x.png',
  iconUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.9.4/images/marker-icon.png',
  shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.9.4/images/marker-shadow.png',
});

const props = defineProps({
  data: {
    type: Object,
    required: true
  }
});

const emit = defineEmits(['submit']);

const mapContainer = ref(null);
let map = null;
let marker = null;

const mapUrl = computed(() => {
  if (props.data.lat !== null && props.data.lng !== null) {
    return `https://www.google.com/maps/place/@${props.data.lat},${props.data.lng},15z`;
  }
  return '';
});

onMounted(() => {
  // Default coordinates (e.g., center of the world or specific default location)
  const defaultLat = props.data.lat !== null ? props.data.lat : 51.505;
  const defaultLng = props.data.lng !== null ? props.data.lng : -0.09;

  map = L.map(mapContainer.value).setView([defaultLat, defaultLng], 2);

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap contributors'
  }).addTo(map);

  if (props.data.lat !== null && props.data.lng !== null) {
    marker = L.marker([props.data.lat, props.data.lng]).addTo(map);
    map.setView([props.data.lat, props.data.lng], 15);
  }

  map.on('click', (e) => {
    const { lat, lng } = e.latlng;
    props.data.lat = parseFloat(lat.toFixed(6));
    props.data.lng = parseFloat(lng.toFixed(6));
    
    updateMarker(props.data.lat, props.data.lng);
  });
});

const updateMarker = (lat, lng) => {
  if (!map) return;
  if (!marker) {
    marker = L.marker([lat, lng]).addTo(map);
  } else {
    marker.setLatLng([lat, lng]);
  }
  map.setView([lat, lng], 15);
};

// Watch for manual input changes
watch(() => props.data.lat, (newLat) => {
  if (newLat !== null && props.data.lng !== null && !isNaN(newLat)) {
    updateMarker(newLat, props.data.lng);
  }
});

watch(() => props.data.lng, (newLng) => {
  if (newLng !== null && props.data.lat !== null && !isNaN(newLng)) {
    updateMarker(props.data.lat, newLng);
  }
});
</script>

<template>
  <div class="space-y-4">
    <!-- Map Container -->
    <div>
      <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Select Location <span class="text-red-500">*</span></label>
      <div class="w-full h-64 rounded-md border border-zinc-300 dark:border-zinc-700 overflow-hidden relative z-0">
        <div ref="mapContainer" class="w-full h-full"></div>
      </div>
    </div>

    <!-- Manual Coordinate Inputs -->
    <div class="grid grid-cols-2 gap-4">
      <div>
        <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Latitude</label>
        <input 
          v-model.number="data.lat" 
          type="number" 
          step="any"
          placeholder="e.g. 51.505" 
          class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200" 
        />
      </div>
      <div>
        <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Longitude</label>
        <input 
          v-model.number="data.lng" 
          type="number" 
          step="any"
          placeholder="e.g. -0.09" 
          class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200" 
        />
      </div>
    </div>

    <!-- Auto-generated Google Maps URL -->
    <div>
      <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Google Maps URL</label>
      <input 
        :value="mapUrl" 
        type="text" 
        readonly
        placeholder="Auto-generated URL..." 
        class="w-full bg-zinc-50 dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-md px-4 py-2 text-zinc-500 dark:text-zinc-400 cursor-not-allowed transition-colors duration-200" 
      />
      <p class="text-xs text-zinc-500 mt-1">This URL will be encoded into the QR code.</p>
    </div>
  </div>
</template>

<style scoped>
/* Ensure leaflet controls don't overlap awkwardly with modals or dropdowns */
:deep(.leaflet-control-container) {
  z-index: 10;
}
</style>
