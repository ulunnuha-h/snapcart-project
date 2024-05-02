<script setup>
import { UserCircleIcon } from "@heroicons/vue/24/solid";
</script>

<template>
    <nav class="bg-slate-900 text-sky-400 flex items-center justify-between px-4 border-b-2 border-b-sky-900 sticky top-0 z-10">
        <h3>Snapcart</h3>
        <UserCircleIcon class="text-slate-50 h-8 w-8" />
    </nav>
</template>

<style scoped>
    nav{
        height: 64px;
    }
</style>