<script setup>
defineProps({
  data: {
    type: Object,
    required: true
  }
});
</script>

<template>
  <div>
    <label class="block text-sm font-medium text-zinc-400 mb-2">Plain Text <span class="text-red-500">*</span></label>
    <textarea 
      v-model="data.value" 
      rows="3"
      placeholder="Enter your text here..." 
      class="w-full bg-zinc-950 border border-zinc-700 rounded-md px-4 py-2 text-zinc-100 placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200 resize-none"
    ></textarea>
  </div>
</template>
