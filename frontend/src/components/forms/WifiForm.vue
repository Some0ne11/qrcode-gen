<script setup>
defineProps({
  data: {
    type: Object,
    required: true
  }
});
</script>

<template>
  <div class="space-y-4">
    <div>
      <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Network Name (SSID) <span class="text-red-500">*</span></label>
      <input 
        v-model="data.ssid" 
        type="text" 
        placeholder="My Network" 
        class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200"
      />
    </div>
    <div class="grid grid-cols-2 gap-4">
      <div>
        <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Security</label>
        <select 
          v-model="data.protocol" 
          class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200"
        >
          <option value="WPA">WPA/WPA2/WPA3</option>
          <option value="WEP">WEP</option>
          <option value="nopass">None</option>
        </select>
      </div>
      <div>
        <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Password <span v-if="data.protocol !== 'nopass'" class="text-red-500">*</span></label>
        <input 
          v-model="data.password" 
          :disabled="data.protocol === 'nopass'"
          type="password" 
          placeholder="Password" 
          class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200 disabled:opacity-50"
        />
      </div>
    </div>
  </div>
</template>
