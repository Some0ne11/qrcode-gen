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
    <div class="grid grid-cols-2 gap-4">
      <div>
        <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Full Name <span class="text-red-500">*</span></label>
        <input v-model="data.fullname" type="text" placeholder="John Doe" class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200" />
      </div>
      <div>
        <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Organization <span class="text-red-500">*</span></label>
        <input v-model="data.org" type="text" placeholder="Company Inc" class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200" />
      </div>
    </div>
    <div class="grid grid-cols-2 gap-4">
      <div>
        <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Phone <span class="text-red-500">*</span></label>
        <input v-model="data.phone" type="tel" placeholder="+1234567890" class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200" />
      </div>
      <div>
        <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Email <span class="text-red-500">*</span></label>
        <input v-model="data.email" type="email" placeholder="john@example.com" class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200" />
      </div>
    </div>
    <div>
      <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Address <span class="text-red-500">*</span></label>
      <textarea v-model="data.address" rows="2" placeholder="123 Main St, City, Country" class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200 resize-none"></textarea>
    </div>
    <div>
      <label class="block text-sm font-medium text-zinc-600 dark:text-zinc-400 mb-1">Notes (Optional)</label>
      <textarea v-model="data.notes" rows="2" placeholder="Extra information..." class="w-full bg-white dark:bg-zinc-950 border border-zinc-300 dark:border-zinc-700 rounded-md px-4 py-2 text-zinc-900 dark:text-zinc-100 placeholder-zinc-400 dark:placeholder-zinc-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors duration-200 resize-none"></textarea>
    </div>
  </div>
</template>
